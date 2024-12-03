package resource

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"golang.org/x/net/context"
)

func TestResource_Fetch(t *testing.T) {
	type args struct {
		ctx context.Context
		url string
	}
	tests := []struct {
		name    string
		args    args
		handler http.Handler
		want    []byte
		wantErr bool
	}{
		{
			name: "fetch data",
			args: args{
				ctx: context.Background(),
				url: "/test",
			},
			handler: func() http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("test"))
				})
			}(),
			want: []byte("test"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(tt.handler)
			defer server.Close()

			r, err := New(server.URL)
			if err != nil {
				t.Fatalf("failed to create resource: %v", err)
			}

			got, err := r.Fetch(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("Resource.Fetch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resource.Fetch() = %v, want %v", got, tt.want)
			}
		})
	}
}
