package yuml2svg

import "testing"

func Test_svgPath(t *testing.T) {
	type args struct {
		yumlPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "works",
			args: args{yumlPath: "doc/test.yuml"},
			want: "doc/test.svg",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := svgPath(tt.args.yumlPath); got != tt.want {
				t.Errorf("svgPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
