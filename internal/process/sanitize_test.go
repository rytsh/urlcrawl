package process

import "testing"

func TestURLToPath(t *testing.T) {
	type args struct {
		urlStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "url to path",
			args: args{
				urlStr: "http://test.com/test",
			},
			want:    "test.com/test",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := URLToPath(tt.args.urlStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLToPath() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("URLToPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
