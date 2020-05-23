package helper

import (
	"fmt"
	"testing"

	"golang.org/x/xerrors"
)

func TestContainsDoNotPrintHelp(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "normal error",
			args: args{
				err: fmt.Errorf("abc"),
			},
			want: false,
		},
		{
			name: "xerror",
			args: args{
				err: xerrors.Errorf("abc"),
			},
			want: false,
		},
		{
			name: "directly doNotHelp",
			args: args{
				err: DoNotPrintHelp(xerrors.Errorf("abc")),
			},
			want: true,
		},
		{
			name: "doNotHelp wrapped again",
			args: args{
				err: xerrors.Errorf("an abc error after: %w", DoNotPrintHelp(xerrors.Errorf("abc"))),
			},
			want: true,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsDoNotPrintHelp(tt.args.err); got != tt.want {
				t.Errorf("ContainsDoNotPrintHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_doNotPrintHelp_Unwrap(t *testing.T) {
	abcErr := fmt.Errorf("abc")

	type fields struct {
		err error
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr error
	}{
		{
			name: "stay nil",
			fields: fields{
				err: nil,
			},
			wantErr: nil,
		},
		{
			name: "unwrap correctly",
			fields: fields{
				err: abcErr,
			},
			wantErr: abcErr,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			doNot := doNotPrintHelp{
				err: tt.fields.err,
			}
			if err := doNot.Unwrap(); err != tt.wantErr {
				t.Errorf("doNotPrintHelp.Unwrap() error = %#v, wantErr %#v", err, tt.wantErr)
			}
		})
	}
}
