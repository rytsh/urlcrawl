package file

import (
	"os"
	"testing"
)

func TestFile_Has(t *testing.T) {
	type fields struct {
		destionation   string
		filePermission os.FileMode
		dirPermission  os.FileMode
	}
	type args struct {
		path string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "has file",
			fields: fields{
				destionation: "testdata",
			},
			args: args{
				path: "exist",
			},
			want: true,
		},
		{
			name: "non exist",
			fields: fields{
				destionation: "testdata",
			},
			args: args{
				path: "non-exist",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := New(tt.fields.destionation)
			if err != nil {
				t.Fatalf("failed to create file: %v", err)
			}

			if got := f.Has(tt.args.path); got != tt.want {
				t.Errorf("File.Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_Save(t *testing.T) {
	type fields struct {
		destionation   string
		filePermission os.FileMode
		dirPermission  os.FileMode
	}
	type args struct {
		data []byte
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "save file",
			fields: fields{
				destionation: "testdata",
			},
			args: args{
				data: []byte("test"),
				path: "example.com/test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := New(tt.fields.destionation)
			if err != nil {
				t.Fatalf("failed to create file: %v", err)
			}

			if err := f.Set(tt.args.path, tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("File.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
