package strings2

import (
	"testing"
)

func TestReplaceSpaces(t *testing.T) {
	type args struct {
		str             string
		replacementRune rune
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "replace1",
			args: args{
				str:             "hello world\na\t",
				replacementRune: '_',
			},
			want: "hello_world_a_",
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpaces(tt.args.str, tt.args.replacementRune); got != tt.want {
				t.Errorf("ReplaceSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
