package clone

import "testing"

func TestGitClone(t *testing.T) {
	type args struct {
		url string
	}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Should clone the repository 'https://github.com/halimath/mini-httpd.git'",
			args: args{
				url: "https://github.com/halimath/mini-httpd.git",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GitClone(tt.args.url)
		})
	}
}
