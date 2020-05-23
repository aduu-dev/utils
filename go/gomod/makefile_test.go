package gomod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMake(t *testing.T) {
	type args struct {
		packageLines []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "no packages",
			args: args{
				packageLines: nil,
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "two packages",
			args: args{
				packageLines: []string{"hello.de", "world2"},
			},
			want: `module abc.test

go 1.12

require (
	hello.de
	world2
)`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt2 := tt
		t.Run(tt2.name, func(t *testing.T) {

			got, err := Make(tt2.args.packageLines)
			if (err != nil) != tt2.wantErr {
				t.Errorf("Make() error = %v, wantErr %v", err, tt2.wantErr)
				return
			}

			assert.Equal(t, tt2.want, got)
		})
	}
}
