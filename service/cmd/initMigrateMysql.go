/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	"github.com/Gocyber-world/navigator-demo/service/system"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// initMigrateMysqlCmd represents the initMigrateMysql command
var initMigrateMysqlCmd = &cobra.Command{
	Use:   "initMigrateMysql",
	Short: "Init or migrate MySQL table schemas",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("initMigrateMysql called")
		// 在actions中连接腾讯云数据库并进行初始化时, 没有STAGE的概念 远程数据库 地址 端口 密码 都来自于环境变量
		MYSQL_ADDR := viper.GetViper().GetString("DB_ADDR")
		MYSQL_PORT := viper.GetViper().GetString("DB_PORT")

		// 如果两个环境变量都不为空，那么认为运行环境为actions, 需要通过环境变量来设置(配置文件中的设置只用于scf vpc内部)
		if MYSQL_ADDR != "" && MYSQL_PORT != "" {
			logger.Info("init migrate mysql in actions")
			global.MYSQL_CONFIG.Path = MYSQL_ADDR
			global.MYSQL_CONFIG.Port = MYSQL_PORT
		}

		initDBService := system.InitDBService{}
		if err := initDBService.InitMsqlDB(global.MYSQL_CONFIG); err != nil {
			logger.Error(err.Error())
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initMigrateMysqlCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initMigrateMysqlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initMigrateMysqlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
