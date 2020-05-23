package errors2

import (
	"os"
	"testing"
)

type myError struct {
	value int
}

func (myError) Error() string {
	panic("implement me")
}

type subError struct {
	myError
}

func TestIsType(t *testing.T) {
	type args struct {
		err    error
		target any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one nil one not 1",
			args: args{
				err:    nil,
				target: myError{},
			},
			want: false,
		},
		{
			name: "one nil one not 2",
			args: args{
				err:    myError{},
				target: nil,
			},
			want: false,
		},
		{
			name: "same type same value",
			args: args{
				err: myError{
					value: 0,
				},
				target: myError{
					value: 0,
				},
			},
			want: true,
		},
		{
			name: "same type different values 1",
			args: args{
				err: myError{
					value: 0,
				},
				target: myError{
					value: 1,
				},
			},
			want: true,
		},
		{
			name: "same type different values 1",
			args: args{
				err: myError{
					value: 1,
				},
				target: myError{
					value: 0,
				},
			},
			want: true,
		},
		{
			name: "subtype and type itself 1",
			args: args{
				err:    subError{},
				target: myError{},
			},
			want: false,
		},
		{
			name: "subtype and type itself 2",
			args: args{
				err:    myError{},
				target: subError{},
			},
			want: false,
		},
		{
			name: "subtype and type itself - different values 1",
			args: args{
				err: myError{
					value: 0,
				},
				target: subError{
					myError: myError{
						value: 1,
					},
				},
			},
			want: false,
		},
		{
			name: "subtype and type itself - different values 2",
			args: args{
				err: subError{
					myError: myError{
						value: 1,
					},
				},
				target: myError{
					value: 0,
				},
			},
			want: false,
		},
		{
			name: "wholly different type 1",
			args: args{
				err:    myError{},
				target: os.ErrClosed,
			},
			want: false,
		},
		{
			name: "wholly different type 2",
			args: args{
				err:    os.ErrClosed,
				target: myError{},
			},
			want: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := IsType(tt.args.err, tt.args.target); got != tt.want {
				t.Errorf("IsType() = %v, want %v", got, tt.want)
			}
		})
	}
}
