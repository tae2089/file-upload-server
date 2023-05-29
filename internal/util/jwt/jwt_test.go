//go:build test
// +build test

package util

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tae2089/file-upload-server/internal/config"
)

func TestCreateAccessToken(t *testing.T) {
	config.LoadEnv()
	gotAccessToken, err := CreateAccessToken([]byte("test"))
	if err != nil {
		t.Error(err)
	}
	fmt.Println(gotAccessToken)
}

func TestExtractIDFromToken(t *testing.T) {
	config.LoadEnv()
	type args struct {
		requestToken string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "test1",
			args:    args{requestToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYmMtbGFicy10b2tlbiIsImlkIjoiZThmOTFhMzYtNDE0Zi00ODhiLThmODYtZmZkZWIyZWM2ZWI5IiwiaXNzIjoiZGV2b3BzLXRlYW0iLCJzdWIiOiJ0ZXN0IiwiYXVkIjpbImJjLWxhYnMtcHJvamVjdCJdLCJleHAiOjIwMDA1MTQwNDIsIm5iZiI6MTY4NTE1NDA0MiwiaWF0IjoxNjg1MTU0MDQyfQ.RreUhDnQGjtr9493Z9Y_zgA3xi-OJf7-LN1tuccJwoSS9mDH6nA_lIACy1JwzfOpzYPDZyymVQjZVci-3tsfPS_ijn7klqyKJnzfiPmiq-9v4-ljYuv6FeuIx5zwnJAnhhFzn1ehB-iV7pNiSWIxOQ5kliJGaHfGxpz_vTvWopa5-seUjSYtJb2otMJZn134Zf0uiNNsxlx3KgL6kT0aqnQxQnGVwxNzOOzddiDpLoQISsh9R3pL9P5tEURber_aguDh5FaO3lOyuoEBRH3H-pq1U3GPrS89xX4hX2pS2i8OLj3HdrMwTC4ysi7gyL3oF9k8IG8X48R-9-_2GGThZhTcT_UvaRgZAkCECKSTv1MUOy8Kv-zqtkAObVqJ6KTI5-dAICtzG75i3inQjKvOxDePF7Zp0C2nGTIPxT_dkjTykxtuCPVkuBQPDW3BqKNuqA50YRwny6oQzrtTZXZ5pFU9RuhrRyQvn8rM5P8RhqLiK1fk95Pz4bTHmggSMemjZ-aRnSVPXgbPRShNd0G7mV01cpQ5kzdcidKiqQpvt61f_r99fENPXZyb-Ey_Z_QSF5wtThjxZyO-AQj9m-MyuYnZTuGvLK7xU3GA-jHU0llqlKLWxrQDxfk3t6yUWaSizy7QeYk_pf11yb1iwvpuSOb77V3ofuAxrpvE-U_4dAg"},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := ExtractIDFromToken(tt.args.requestToken)
			fmt.Println(got, err)
			assert.Nil(t, err)
			assert.NotEmpty(t, got)
		})
	}
}
