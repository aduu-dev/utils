package regex

import (
	"reflect"
	"testing"
)

func TestGroupMatch(t *testing.T) {
	type args struct {
		data        string
		regexString string
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
		wantErr   bool
		wantM     map[string]string
	}{
		{
			name: "invalid regex",
			args: args{
				data:        "",
				regexString: "[abc",
			},
			wantMatch: false,
			wantM:     nil,
			wantErr:   true,
		},
		{
			name: "found match",
			args: args{
				data:        "aabb",
				regexString: "(?P<abc>aa)(?P<def>bb)",
			},
			wantMatch: true,
			wantM:     map[string]string{"": "aabb", "abc": "aa", "def": "bb"},
			wantErr:   false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotM, err := GroupMatch(tt.args.data, tt.args.regexString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMatch != tt.wantMatch {
				t.Errorf("GroupMatch() gotMatch = %v, want %v", gotMatch, tt.wantMatch)
			}
			if !reflect.DeepEqual(gotM, tt.wantM) {
				t.Errorf("GroupMatch() gotM = %v, want %v", gotM, tt.wantM)
			}
		})
	}
}

func TestGroupMatchAll(t *testing.T) {
	type args struct {
		data        string
		regexString string
	}
	tests := []struct {
		name      string
		args      args
		wantMatch bool
		wantErr   bool
		wantMaps  []map[string]string
	}{
		{
			name: "invalid regex",
			args: args{
				data:        "",
				regexString: "[abc",
			},
			wantMatch: false,
			wantMaps:  nil,
			wantErr:   true,
		},
		{
			name: "found match",
			args: args{
				data:        "aabb aabb",
				regexString: "(?P<abc>aa)(?P<def>bb)",
			},
			wantMatch: true,
			wantMaps: []map[string]string{
				{"": "aabb", "abc": "aa", "def": "bb"},
				{"": "aabb", "abc": "aa", "def": "bb"},
			},
			wantErr: false,
		},
	}
	for _, tt2 := range tests {
		tt := tt2
		t.Run(tt.name, func(t *testing.T) {
			gotMatch, gotMaps, err := GroupMatchAll(tt.args.data, tt.args.regexString)
			if (err != nil) != tt.wantErr {
				t.Errorf("GroupMatch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMatch != tt.wantMatch {
				t.Errorf("GroupMatch() gotMatch = %#v, want %#v", gotMatch, tt.wantMatch)
			}
			if !reflect.DeepEqual(gotMaps, tt.wantMaps) {
				t.Errorf("GroupMatch() \ngotMaps = %#v, \n   want = %#v", gotMaps, tt.wantMaps)
			}
		})
	}
}
