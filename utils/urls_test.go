package utils

import (
	"reflect"
	"testing"
)

func TestParseUrl(t *testing.T) {
	type args struct {
		u      string
		host   string
		scheme string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				u:      "//test.com",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://test.com/",
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				u:      "/path",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/path/",
			wantErr: false,
		},
		{
			name: "3",
			args: args{
				u:      "/path/to_file/?arg1=val",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/path/to_file/?arg1=val",
			wantErr: false,
		},
		{
			name: "4",
			args: args{
				u:      "/",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/",
			wantErr: false,
		},
		{
			name: "5",
			args: args{
				u:      "/?testarg=testval&testarg2=testval2",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/?testarg=testval&testarg2=testval2",
			wantErr: false,
		},
		{
			name: "6",
			args: args{
				u:      "/path",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/path/",
			wantErr: false,
		},
		{
			name: "7",
			args: args{
				u:      "/path/to_file?arg1=val",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://our.site/path/to_file/?arg1=val",
			wantErr: false,
		},
		{
			name: "1",
			args: args{
				u:      "//test.com",
				host:   "our.site",
				scheme: "https",
			},
			want:    "https://test.com/",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUrl(tt.args.u, tt.args.host, tt.args.scheme)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got.String(), tt.want) {
				t.Errorf("ParseUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
