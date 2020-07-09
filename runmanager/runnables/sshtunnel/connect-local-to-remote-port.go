package sshtunnel

import (
	"context"
	"fmt"
	"strconv"

	"aduu.dev/utils/dash"
	"aduu.dev/utils/runmanager"
	"k8s.io/klog/v2"
)

type localConnection struct {
	LocalPort    string
	RemoteAdress string
	RemotePort   string
	SSHName      string
}

func newLocalConnection(
	localPort string,
	remoteAdress string,
	remotePort string,
	sshName string,
) (con *localConnection, err error) {
	if _, err = strconv.ParseInt(localPort, 10, 64); err != nil {
		return nil, fmt.Errorf("local port is invalid: %w", err)
	}

	if _, err = strconv.ParseInt(remotePort, 10, 64); err != nil {
		return nil, fmt.Errorf("remote port is invalid: %w", err)
	}

	return &localConnection{LocalPort: localPort,
		RemoteAdress: remoteAdress,
		RemotePort:   remotePort,
		SSHName:      sshName,
	}, nil
}

func connectLocalPortToRemotePort(
	ctx context.Context,
	d dash.Runner,
	con *localConnection) (err error,
) {
	return d.RunE(ctx,
		dash.SplitTemplateExpand(`ssh -L {{.LocalPort}}:{{.RemoteAdress}}:{{.RemotePort}} {{.SSHName}}`, con))
}

func connectRemotePortToLocalPort(
	ctx context.Context,
	d dash.Runner,
	con *localConnection) (err error,
) {
	return d.RunE(ctx,
		dash.SplitTemplateExpand(`ssh -R {{.RemotePort}}:{{.RemoteAdress}}:{{.LocalPort}} {{.SSHName}}`, con))
}

// ConnectLocalPortToRemotePort starts a service which connects
// the local host with the remote host on the given ports.
func ConnectLocalPortToRemotePort(m *runmanager.RunManager,
	localPort string,
	remoteAdress string,
	remotePort string,
	sshName string,
) {
	con, err := newLocalConnection(localPort,
		remoteAdress, remotePort,
		sshName)

	if err != nil {
		m.ErrChan <- err
	}

	m.Run(&runmanager.Runner{
		Run: func() error {
			return connectLocalPortToRemotePort(m.Context, dash.NewRunner(), con)
		},
		Shutdown: func() error {
			klog.InfoS("Shutting down local-to-remote-port connection")
			return nil
		},
	})
}

// ConnectLocalPortToRemotePort starts a service which connects
// the remote host with the local host on the given ports.
func ConnectRemotePortToLocalPort(m *runmanager.RunManager,
	remoteAdress string,
	remotePort string,
	localPort string,
	sshName string,
) {
	con, err := newLocalConnection(localPort,
		remoteAdress, remotePort,
		sshName)

	if err != nil {
		m.ErrChan <- err
	}

	m.Run(&runmanager.Runner{
		Run: func() error {
			return connectRemotePortToLocalPort(m.Context, dash.NewRunner(), con)
		},
		Shutdown: func() error {
			klog.InfoS("Shutting down local-to-remote-port connection")
			return nil
		},
	})
}
