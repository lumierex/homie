package cmd

import (
	"log"
	"strings"

	"github.com/miltian/homie/internal/env"
	"github.com/spf13/cobra"
)

const (
	EnvMiniConda = iota + 1 // MiniConda环境安装
	EnvNode                 // Node环境安卓
	EnvTest                 // 测试
)

// 环境选项
var environment int8

var LongDesc = strings.Join([]string{
	"1: MiniConda env install",
	"2: Node env install",
	"3: Test env install",
}, "\n")

var EnvCmd = cobra.Command{
	Use:   "env",
	Short: "常用环境安装配置",
	Long:  LongDesc,
	Run: func(cmd *cobra.Command, args []string) {
		switch environment {
		case EnvMiniConda:
			{
				// 执行EnvMiniConda shell环境安装
				err := env.MiniCondaInstall()
				if err != nil {
					log.Println(err)
				}
			}
		default:
			{
				log.Println("未输入任何安装任务")
			}
		}

	},
}

func init() {
	EnvCmd.Flags().Int8VarP(&environment, "env", "e", -1, "请输入需要安装的环境类型")
}
