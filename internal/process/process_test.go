package process

import (
	"context"
	"testing"

	gomock "go.uber.org/mock/gomock"
)

func TestProcess_Process(t *testing.T) {
	type fields struct {
		storage    func(*MockStorage) *MockStorage
		resource   func(*MockResource) *MockResource
		maxProcess int
	}
	type args struct {
		ctx    context.Context
		urlStr string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "process",
			fields: fields{
				storage: func(m *MockStorage) *MockStorage {
					m.EXPECT().Has("test").Return(false)
					m.EXPECT().Set("test", []byte("test")).Return(nil)

					return m
				},
				resource: func(m *MockResource) *MockResource {
					m.EXPECT().Fetch(gomock.Any(), "test").Return([]byte("test"), nil)
					m.EXPECT().Links([]byte("test")).Return([]string{})

					return m
				},
			},
			args: args{
				ctx:    context.Background(),
				urlStr: "test",
			},
			wantErr: false,
		},
		{
			name: "process",
			fields: fields{
				storage: func(m *MockStorage) *MockStorage {
					m.EXPECT().Has("test.com").Return(false)
					m.EXPECT().Set("test.com", []byte(`lorem ipsum dolor sit amet <a title="abc" href="http://test.com/abc">test</a>`)).Return(nil)

					m.EXPECT().Has("test.com/abc").Return(true)
					m.EXPECT().Get("test.com/abc").Return([]byte("abc-data"), nil)

					return m
				},
				resource: func(m *MockResource) *MockResource {
					m.EXPECT().Fetch(gomock.Any(), "http://test.com").Return([]byte(`lorem ipsum dolor sit amet <a title="abc" href="http://test.com/abc">test</a>`), nil)
					m.EXPECT().Links([]byte(`lorem ipsum dolor sit amet <a title="abc" href="http://test.com/abc">test</a>`)).Return([]string{"http://test.com/abc"})

					m.EXPECT().Links([]byte("abc-data")).Return([]string{})
					return m
				},
			},
			args: args{
				ctx:    context.Background(),
				urlStr: `http://test.com`,
			},
			wantErr: false,
		},
		{
			name: "process",
			fields: fields{
				storage: func(m *MockStorage) *MockStorage {
					m.EXPECT().Has("start.url/abc").Return(false)
					m.EXPECT().Set("start.url/abc", []byte(`<a title="foo" href="https://start.url/abc/foo">FOO</a> <a href="https://start.url/baz">BAZ</a> <a href="https://another.domain">ANOTHER</a> `)).Return(nil)

					m.EXPECT().Has("start.url/abc/foo").Return(true)
					m.EXPECT().Get("start.url/abc/foo").Return([]byte("fooo"), nil)

					return m
				},
				resource: func(m *MockResource) *MockResource {
					m.EXPECT().Fetch(gomock.Any(), "https://start.url/abc").Return([]byte(`<a title="foo" href="https://start.url/abc/foo">FOO</a> <a href="https://start.url/baz">BAZ</a> <a href="https://another.domain">ANOTHER</a> `), nil)
					m.EXPECT().Links([]byte(`<a title="foo" href="https://start.url/abc/foo">FOO</a> <a href="https://start.url/baz">BAZ</a> <a href="https://another.domain">ANOTHER</a> `)).Return([]string{"https://start.url/abc/foo"})

					m.EXPECT().Links([]byte("fooo")).Return(nil)
					return m
				},
			},
			args: args{
				ctx:    context.Background(),
				urlStr: `https://start.url/abc`,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			p := New(tt.fields.storage(NewMockStorage(ctrl)), tt.fields.resource(NewMockResource(ctrl)))

			if err := p.Process(tt.args.ctx, tt.args.urlStr); (err != nil) != tt.wantErr {
				t.Errorf("Process.Process() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
