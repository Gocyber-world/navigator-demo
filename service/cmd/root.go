/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/Gocyber-world/navigator-demo/config"
	"github.com/Gocyber-world/navigator-demo/global"
	"github.com/Gocyber-world/navigator-demo/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gocyber-service",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initGlobalEnv)
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "configs/config.local.yaml", "config file (default is configs/config.local.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gocyber-service" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gocyber-service")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Info("using config file:" + viper.ConfigFileUsed())
	}
}

// 全局变量初始化
func initGlobalEnv() {
	global.GVA_VP = viper.GetViper()
	global.STAGE = global.GVA_VP.GetString("STAGE")

	// zap 初始化
	if err := logger.InitLogger(); err != nil {
		panic("failed to init zap:" + err.Error())
	}

	// mysql
	var mysqlConfig config.Mysql
	if err := viper.UnmarshalKey("mysql", &mysqlConfig); err != nil {
		zap.S().Fatal("failed to read mysql config", err.Error())
	}

	// STAGE == "" 时，说明程序在本地运行
	// STAGE 指明时，说明程序在scf或actions中运行，此时数据库密码来源于环境变量而非配置文件
	if global.STAGE != "" {
		mysqlConfig.Password = global.GVA_VP.GetString("DB_PASSWORD")
	}

	global.MYSQL_CONFIG = mysqlConfig

}
