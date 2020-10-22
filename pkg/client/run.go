package client

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"

	"github.com/spf13/cobra"

	"github.com/datawire/telepresence2/pkg/common"
	"github.com/datawire/telepresence2/pkg/rpc"
)

// runner contains all parameters needed in order to run an intercepted command.
type runner struct {
	rpc.ConnectRequest
	rpc.InterceptRequest
	DNS      string
	Fallback string
	NoWait   bool
	Quit     bool
	Status   bool
	Version  bool
}

// run will ensure that an intercept is in place and then execute the command given by args[0]
// and the arguments starting at args[1:].
func (p *runner) run(cmd *cobra.Command, args []string) error {
	switch {
	case p.Quit:
		return Quit(cmd, args)
	case p.Status:
		return status(cmd, args)
	case p.Version:
		return version(cmd, args)
	case p.NoWait:
		if p.InterceptRequest.Deployment != "" {
			return p.addIntercept(cmd, args)
		}
		return p.connect(cmd, args)
	}

	out := cmd.OutOrStdout()
	var exe string
	subShellMsg := ""
	if len(args) == 0 {
		exe = os.Getenv("SHELL")
		subShellMsg = fmt.Sprintf("Starting a %s subshell", exe)
	} else {
		exe = args[0]
		args = args[1:]
	}

	if p.InterceptRequest.Deployment != "" {
		err := prepareIntercept(cmd, &p.InterceptRequest)
		if err != nil {
			return err
		}
		return p.runWithIntercept(out, func(_ *interceptState) error {
			if subShellMsg != "" {
				fmt.Fprintln(out, subShellMsg)
			}
			return start(exe, args, true, cmd.InOrStdin(), out, cmd.OutOrStderr())
		})
	}
	return p.runWithConnector(out, func(cs *connectorState) error {
		if subShellMsg != "" {
			fmt.Fprintln(out, subShellMsg)
		}
		return start(exe, args, true, cmd.InOrStdin(), out, cmd.OutOrStderr())
	})
}

func (p *runner) runWithDaemon(out io.Writer, f func(ds *daemonState) error) error {
	ds, err := newDaemonState(out, p.DNS, p.Fallback)
	if err != nil && err != daemonIsNotRunning {
		return err
	}
	return common.WithEnsuredState(ds, func() error { return f(ds) })
}

func (p *runner) runWithConnector(out io.Writer, f func(cs *connectorState) error) error {
	return p.runWithDaemon(out, func(ds *daemonState) error {
		p.InterceptEnabled = true
		cs, err := newConnectorState(ds.grpc, &p.ConnectRequest, out)
		if err != nil && err != connectorIsNotRunning {
			return err
		}
		return common.WithEnsuredState(cs, func() error { return f(cs) })
	})
}

func (p *runner) runWithIntercept(out io.Writer, f func(is *interceptState) error) error {
	return p.runWithConnector(out, func(cs *connectorState) error {
		is := newInterceptState(cs.grpc, &p.InterceptRequest, out)
		return common.WithEnsuredState(is, func() error { return f(is) })
	})
}

func runAsRoot(exe string, args []string) error {
	if os.Geteuid() != 0 {
		args = append([]string{"-n", "-E", exe}, args...)
		exe = "sudo"
	}
	return start(exe, args, false, nil, nil, nil)
}

func start(exe string, args []string, wait bool, stdin io.Reader, stdout, stderr io.Writer) error {
	cmd := exec.Command(exe, args...)
	cmd.Stdout = stdout
	cmd.Stderr = stderr
	cmd.Stdin = stdin
	if !wait {
		// Process must live in a process group of its own to prevent
		// getting affected by <ctrl-c> in the terminal
		cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	}

	var err error
	if err = cmd.Start(); err != nil {
		return fmt.Errorf("%s %s: %v\n", exe, strings.Join(args, " "), err)
	}
	if !wait {
		_ = cmd.Process.Release()
		return nil
	}

	// Ensure that SIGINT and SIGTERM are propagated to the child process
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigCh
		if sig == nil {
			return
		}
		_ = cmd.Process.Signal(sig)
	}()
	s, err := cmd.Process.Wait()
	if err != nil {
		return fmt.Errorf("%s %s: %v\n", exe, strings.Join(args, " "), err)
	}

	sigCh <- nil
	exitCode := s.ExitCode()
	if exitCode != 0 {
		return fmt.Errorf("%s %s: exited with %d\n", exe, strings.Join(args, " "), exitCode)
	}
	return nil
}
