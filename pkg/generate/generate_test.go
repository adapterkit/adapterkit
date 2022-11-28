package generate

import (
	"testing"
)

func TestGetTemplateDirName(t *testing.T) {
	tests := []struct {
		path   string
		want   string
		hasErr bool
	}{
		{
			path:   "/tmp/adapterkit-template-foo",
			want:   "adapterkit-template-foo",
			hasErr: false,
		},
		{
			path:   "/tmp/adapterkit-template-foo/bar",
			want:   "bar",
			hasErr: false,
		},
		{
			path:   "",
			want:   "",
			hasErr: true,
		},
		{
			path:   "/tmp/template/",
			want:   "template",
			hasErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, err := getTemplateDirName(tt.path)
			switch {
			case err != nil && !tt.hasErr:
				t.Errorf("getTemplateDirName() error = %v, wantErr %v", err, tt.hasErr)
			case err == nil && tt.hasErr:
				t.Errorf("getTemplateDirName() error = %v, wantErr %v", err, tt.hasErr)
			case got != tt.want:
				t.Errorf("getTemplateDirName() = %v, want %v", got, tt.want)
			}
		})
	}
}
