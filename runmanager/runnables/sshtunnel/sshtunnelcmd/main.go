// Package runs ssh tunnels.
package main

import (
	"context"
	"os"
	"strings"

	"aduu.dev/utils/runmanager"
	"aduu.dev/utils/runmanager/runnables/sshtunnel"
	"github.com/spf13/cobra"
	"k8s.io/klog/v2"
)

// RootCMD creates a root command of a program.
func RootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sshtunnel",
		Short: "runs ssh tunnels",
	}

	localToRemote := cmd.Flags().StringArrayP("local", "l", []string{}, "sets local to remote forwarding: local-port:remote-port:ssh-name")

	remoteToLocal := cmd.Flags().StringArrayP("remote", "r", []string{}, "sets remote to local forwarding: remote-port:local-port:ssh-name")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var services []runmanager.Service

		for _, loc := range *localToRemote {
			split := strings.SplitN(loc, ":", 3)

			localPort := split[0]
			remotePort := split[1]
			sshName := split[2]

			services = append(services, sshtunnel.ConnectLocalPortToRemotePort(localPort, "127.0.0.1", remotePort, sshName))
		}

		for _, loc := range *remoteToLocal {
			split := strings.SplitN(loc, ":", 3)

			remotePort := split[0]
			localPort := split[1]
			sshName := split[2]

			services = append(services, sshtunnel.ConnectRemotePortToLocalPort("127.0.0.1", remotePort, localPort, sshName))
		}

		klog.InfoS("number of services", "service-count", len(services))

		runmanager.RunServices(
			context.Background(),
			append(
				[]runmanager.Service{
					runmanager.RunInterruptHandler,
				},
				services...),
		)

		return nil
	}
	cmd.SetOut(os.Stdout)
	cmd.SetErr(os.Stderr)
	// cmd.AddCommand(completion.NewCompletionCMD())
	cmd.AddCommand()
	return cmd
}

func main() {
	errorExitCode := 1

	if err := RootCMD().Execute(); err != nil {
		// if !helper.ContainsDoNotPrintHelp(err) {
		//    fmt.Println(err)
		// }

		os.Exit(errorExitCode)
	}
}
