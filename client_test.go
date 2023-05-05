package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHash(t *testing.T) {
	tests := []struct {
		name    string
		handler http.Handler
		want    string
		wantErr bool
	}{
		{
			name: "should return correct hash",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("foo"))
			}),
			want:    "acbd18db4cc2f85cedef654fccc4a4d8",
			wantErr: false,
		},
		{
			name: "should return error when get wrong status code",
			handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusNotFound)
			}),
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newSimpleClient()
			srv := httptest.NewServer(tt.handler)

			got, err := c.getHash(srv.URL)

			if (err != nil) != tt.wantErr {
				t.Errorf("simpleClient.getHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("simpleClient.getHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
