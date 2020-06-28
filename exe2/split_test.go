package exe2

import (
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestSplitCommand(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"kubectl get external ips", args{`kubectl get svc --all-namespaces -o` +
			` jsonpath='{range .items[?(@.spec.type=="LoadBalancer")]}` +
			`{.metadata.name}:{.status.loadBalancer.ingress[0].ip}{"\n"}{end}'`},
			[]string{"kubectl", "get", "svc", "--all-namespaces", "-o",
				`jsonpath={range .items[?(@.spec.type=="LoadBalancer")]}` +
					`{.metadata.name}:{.status.loadBalancer.ingress[0].ip}{"\n"}{end}`}},
		{
			name: "go list -f {{.Deps}}",
			args: args{
				s: "go list -f {{.Deps}}",
			},
			want: []string{"go", "list", "-f", "{{.Deps}}"},
		},
		{
			name: "start ",
			args: args{
				s: `cmd.exe /C start "" "C:\Program Files (x86)\Mozilla Thunderbird\thunderbird.exe"` +
					` -compose "to = 'CloudCoin@Protonmail.com',` +
					`subject = 'Subject1',body = 'Hello'" -offline`,
			},
			want: []string{
				"cmd.exe",
				"/C",
				"start",
				"",
				`C:\Program Files (x86)\Mozilla Thunderbird\thunderbird.exe`,
				"-compose",
				"to = 'CloudCoin@Protonmail.com',subject = 'Subject1',body = 'Hello'",
				"-offline",
			},
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := splitCommandOld(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				assert.Equal(t, got, tt.want)
				t.Errorf("SplitCommand()\n got: %#v\nwant: %#v", got, tt.want)
			}
		})
	}
}
