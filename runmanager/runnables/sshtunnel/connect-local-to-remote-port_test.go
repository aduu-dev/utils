package sshtunnel

import (
	"context"
	"testing"

	"aduu.dev/utils/dash"
	"github.com/stretchr/testify/assert"
)

func Test_connectLocalPortToRemotePort(t *testing.T) {
	type args struct {
		ctx context.Context
		d   dash.Runner
		con *localConnection
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				con: &localConnection{
					LocalPort:    "8080",
					RemoteAdress: "127.0.0.1",
					RemotePort:   "9090",
					SSHName:      "dev1",
				},
			},
			want: []string{
				"ssh",
				"-L",
				"8080:127.0.0.1:9090",
				"dev1",
			},
			wantErr: false,
		},
		{
			name: "split does not Args[1] into multiple parts",
			args: args{
				ctx: context.Background(),
				con: &localConnection{
					LocalPort:    "8080 def",
					RemoteAdress: "127.0.0.1",
					RemotePort:   "9090",
					SSHName:      "dev1",
				},
			},
			want: []string{
				"ssh",
				"-L",
				"8080 def:127.0.0.1:9090",
				"dev1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the test shell to check what was called.
			td := dash.NewTestRunner()
			tt.args.d = td

			if err := connectLocalPortToRemotePort(tt.args.ctx, tt.args.d,
				tt.args.con); (err != nil) != tt.wantErr {
				t.Errorf("connectLocalPortToRemotePort() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if !assert.Equal(t, 1, len(td.Commands())) {
				t.Fail()
			}

			got := (td.Commands()[0]).Command

			if !assert.Equal(t, tt.want, got) {
				t.Fail()
			}
		})
	}
}

func Test_connectRemotePortToLocalPort(t *testing.T) {
	type args struct {
		ctx context.Context
		d   dash.Runner
		con *localConnection
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				ctx: context.Background(),
				con: &localConnection{
					LocalPort:    "8080",
					RemoteAdress: "127.0.0.1",
					RemotePort:   "9090",
					SSHName:      "dev1",
				},
			},
			want: []string{
				"ssh",
				"-R",
				"9090:127.0.0.1:8080",
				"dev1",
			},
			wantErr: false,
		},
		{
			name: "split does not Args[1] into multiple parts",
			args: args{
				ctx: context.Background(),
				con: &localConnection{
					LocalPort:    "8080 def",
					RemoteAdress: "127.0.0.1",
					RemotePort:   "9090",
					SSHName:      "dev1",
				},
			},
			want: []string{
				"ssh",
				"-R",
				"9090:127.0.0.1:8080 def",
				"dev1",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set the test shell to check what was called.
			td := dash.NewTestRunner()
			tt.args.d = td

			if err := connectRemotePortToLocalPort(tt.args.ctx, tt.args.d,
				tt.args.con); (err != nil) != tt.wantErr {
				t.Errorf("connectLocalPortToRemotePort() error = %v, wantErr %v", err, tt.wantErr)
			}

			if tt.wantErr {
				return
			}

			if !assert.Equal(t, 1, len(td.Commands())) {
				t.Fail()
			}

			got := (td.Commands()[0]).Command

			if !assert.Equal(t, tt.want, got) {
				t.Fail()
			}
		})
	}
}
