/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Gocyber-world/navigator-demo/core"
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	"github.com/Gocyber-world/navigator-demo/middleware"
	"github.com/Gocyber-world/navigator-demo/service/system"
	"github.com/Gocyber-world/navigator-demo/utils/obfuse"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("serve command called", zap.String("stage", global.STAGE))
		initServeEnv()
		core.RunWindowsServer()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initServeEnv() {
	if global.STAGE == "" {
		// 本地环境敏感字段从配置文件中读取
		global.JWT_AUTH = middleware.NewJWT([]byte(global.GVA_VP.GetString("jwt.signing-key")))
		global.OBFUSE = obfuse.NewObfuscator(global.GVA_VP.GetString("hashids.salt"), global.GVA_VP.GetInt("hashids.minlength"))
	} else {
		// scf 环境敏感字段来源于环境变量
		global.JWT_AUTH = middleware.NewJWT([]byte(global.GVA_VP.GetString("JWT_SIGNING_KEY")))
		global.OBFUSE = obfuse.NewObfuscator(global.GVA_VP.GetString("HASHIDS_SALT"), global.GVA_VP.GetInt("hashids.minlength"))
	}
	global.JWT_EXPIRE_TIME = global.GVA_VP.GetInt("jwt.expiretime")
	global.JWT_COOKIES_DOMAIN = global.GVA_VP.GetString("jwt.cookies-domain")
	global.HOST = global.GVA_VP.GetString("host")

	initDBService := system.InitDBService{}
	if err := initDBService.ConnectGlobalDB(global.MYSQL_CONFIG); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

}
