package cmd 

import "github.com/spf13/cobra"

var EnvCmd = cobra.Command{
	Use: "env",
	Short: "常用环境安装配置",
	Long: "环境安装，Miniconda, Node, Python 等环境安装" ,
	Run: func(cmd *cobra.Command, args []string) {},
}

func init() {

}
