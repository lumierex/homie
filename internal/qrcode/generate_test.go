package qrcode

import (
	"testing"
)

func TestQRCodeBatchGenerate(t *testing.T) {
	type args struct {
		contents []QRCodeMetaInfo
		size     int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "批量测试",
			args: args{
				contents: []QRCodeMetaInfo{
					{
						name: "百度",
						url:  "https://www.baidu.com",
					},
				},
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QRCodeBatchGenerate(tt.args.contents, tt.args.size, "./qr/")
		})
	}
}

func TestQRCodeGenerateWithAvatar(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				url: "https://www.baidu.com",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := QRCodeGenerateWithAvatar(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("QRCodeGenerateWithAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
