package strings2

import (
	"testing"
)

func TestEqual(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{ // []string{}
			name: "both nil",
			args: args{
				a: nil,
				b: nil,
			},
			want: true,
		},
		{
			name: "both one element",
			args: args{
				a: []string{"a"},
				b: []string{"a"},
			},
			want: true,
		},
		{
			name: "both one element unequal",
			args: args{
				a: []string{"a"},
				b: []string{"b"},
			},
			want: false,
		},
		{
			name: "one nil one non-nil 1",
			args: args{
				a: nil,
				b: []string{},
			},
			want: true,
		},
		{
			name: "one nil one non-nil 2",
			args: args{
				a: []string{},
				b: nil,
			},
			want: true,
		},
		{
			name: "one element one nil 1",
			args: args{
				a: nil,
				b: []string{"a"},
			},
			want: false,
		},
		{
			name: "one element one nil 2",
			args: args{
				a: []string{"a"},
				b: nil,
			},
			want: false,
		},
		{
			name: "one element one none 1",
			args: args{
				a: []string{},
				b: []string{"a"},
			},
			want: false,
		},
		{
			name: "one element one none 1",
			args: args{
				a: []string{"a"},
				b: []string{},
			},
			want: false,
		},
		{
			name: "two elements wrong sort order",
			args: args{
				a: []string{"a", "b"},
				b: []string{"b", "a"},
			},
			want: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortedEqual(t *testing.T) {
	type args struct {
		a []string
		b []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "two elements wrong sort order",
			args: args{
				a: []string{"a", "b"},
				b: []string{"b", "a"},
			},
			want: true,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := SortedEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SortedEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
