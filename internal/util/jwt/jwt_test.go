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
			args:    args{requestToken: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiYmMtbGFicy10b2tlbiIsImlkIjoiNzcxNmU3YTEtZjlmYS00NDk4LWIyMjItYTRkM2I1MzM3NGRjIiwiaXNzIjoiZGV2b3BzLXRlYW0iLCJzdWIiOiJ0ZXN0IiwiYXVkIjpbImJjLWxhYnMtcHJvamVjdCJdLCJleHAiOjIwMDA3Nzk3NzAsIm5iZiI6MTY4NTQxOTc3MCwiaWF0IjoxNjg1NDE5NzcwfQ.DZ8cGI_-wqgynTTcRnQ3826BP_wnu8Tnx11zExaf7Ui67dXApho1cnYRRxF0zFEMjFKimgYp10HwpGBBcgHvf8sJJYt1dJFH7G98tK5IzwlkxhOdWLNI0_HgVdX2iW0uL-fpvppA-qESHRTZF0MlR7C1hEktqBeL6esjjNERi_ybNHlVDS971v49YT3wYlyVDCJkSXeOZev2k7pUYLCThpzLH8ZcbiSGlkErz5sOf3Q3JvSbhaUCLgnsiYKn6pPNeirPyDzyZoHSbWCQ3KREJYTo2bm_O0nA4tKSfJAKbFgkNu46O9oRVORGC9OwGt4rUvj3vzzY3le8xr9pSVlgUZY3PvuMxfUP4chW_Ims91RoBy2_urORKY0z4pYQ62YwYyDep3nswCBLEcbKC9qBjkqtL0D0RVRvnynYzGUV22hI3IRhZRpBNiVHdnD8uD7f54mYw8aMfoOPrO_tCa8Spu23sRqsaekAkmzEB_ksPLNRM_jukBboj3X4bcp3C7xKLWHap3948dNkT84M_1Lt0Mhji59nCiq_PQQLn3x8frV4oh79cASZ-jht2mPr0paEGuH6EmMUShK1faEq7erFA1ky-eVqVFYLUQKKO9kmcZuDrGEwjuPmHMd0ND0Xp2cKhDTh2rgXEmEOCy7lmt_jxeIxFaKLCfVv82upYnF7ajI"},
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
