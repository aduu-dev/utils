package sshtunnel_test

import "testing"

func TestConnectRemotePorts(t *testing.T) {
	args := struct {
		fromSSHName string
		startPort   string
		toSSHName   string
		endPort     string
	}{
		fromSSHName: "dev1",
		startPort:   "7000",
		toSSHName:   "dev2",
		endPort:     "9000",
	}

	want := [][]string{
		{ // Aufbau: remote:7000 zu host:8000
			"ssh",
			"-R",
			"127.0.0.1:7000:127.0.0.1:8000",
			"dev2",
			"-N",
		},
		{
			"ssh",
			"-L",
			"127.0.0.1:7000:127.0.0.1:8000",
			"dev1",
			"-N",
		},
	}
}
