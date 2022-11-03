package cmd

import (
	// "log"
	// "net/url"

	// qrcode "github.com/skip2/go-qrcode"
	"github.com/spf13/cobra"
)

// 环境选项
// var environment int8
var content string
var size int


var QRCodeCmd = cobra.Command{
	Use:   "qrcode",
	Short: "qr",
	Run: func(cmd *cobra.Command, args []string) {
		
	},
}

func init() {
	QRCodeCmd.Flags().StringVarP(&content, "url", "u", "https://www.baidu.com", "请输入二维码地址")
	QRCodeCmd.Flags().IntVarP(&size, "size", "s", 400, "请输入二维码大小")
	// QRCodeCmd.Flags().StringVarP(&content, "size", "u", "https://www.baidu.com", "请输入二维码地址")
}
