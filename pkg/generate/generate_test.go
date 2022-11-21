package generate

import (
	"testing"
)

func TestGetTemplateDirName(t *testing.T) {
	tests := []struct {
		path string
		want string
		err  bool
	}{
		{
			path: "/tmp/adapterkit-template-foo",
			want: "adapterkit-template-foo",
			err:  false,
		},
		{
			path: "/tmp/adapterkit-template-foo/bar",
			want: "bar",
			err:  false,
		},
		{
			path: "",
			want: "",
			err:  true,
		},
		{
			path: "/tmp/template/",
			want: "template",
			err:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			got, err := getTemplateDirName(tt.path)
			if err != nil {
				if !tt.err {
					t.Errorf("error = %v, but didn't expect any errors", err)
				}
				return
			} else if tt.err {
				t.Errorf("didn't get any errors but expected one")
				return
			}

			if got != tt.want {
				t.Errorf("getTemplateDirName() = %v, want %v", got, tt.want)
			}
		})
	}
}
