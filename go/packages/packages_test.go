package packages

import (
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

func TestListPackages(t *testing.T) {
	type args struct {
		fromPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "works",
			args: args{
				fromPath: "",
			},
			want:    []string{"my.test/a", "my.test/b", "my.test"},
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			tt.args.fromPath = filepath.Join("testdata/data")
			sort.Strings(tt.want)

			got, err := ListPackages(tt.args.fromPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListPackages() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListPackages() = %v, want %v", got, tt.want)
			}
		})
	}
}
