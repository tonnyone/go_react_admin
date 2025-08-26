package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tonnyone/go_react_admin/internal/config"
	"github.com/tonnyone/go_react_admin/internal/database"
	"github.com/tonnyone/go_react_admin/internal/logger"
	"github.com/tonnyone/go_react_admin/router"
)

var rootCmd = &cobra.Command{
	Use:   "go_react_admin",
	Short: "go_react_admin is a backend service",
}

var configPath string

var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Start the web server",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.LoadConfigWithPath(configPath)
		if err != nil {
			logger.Errorf("加载配置失败: %v", err)
			os.Exit(1)
		}
		logger.Init(cfg.Log.Level, cfg.Log.Format)
		// 初始化数据库
		_, err = database.InitDB(cfg.DB.DSN, cfg.Log.Level)
		if err != nil {
			logger.Errorf("数据库初始化失败: %v", err)
			os.Exit(1)
		}
		logger.Infof("配置加载成功，服务端口: %d", cfg.App.Port)
		// 初始化日志
		gin.SetMode(cfg.App.Mode)
		r := gin.New()
		r.Use(gin.LoggerWithWriter(logger.UnderlyingLogger().Out))
		r.Use(gin.RecoveryWithWriter(logger.UnderlyingLogger().Out))
		router.RegisterRoutes(r)
		r.Run(fmt.Sprintf(":%d", cfg.App.Port))
	},
}

func main() {
	webCmd.Flags().StringVarP(&configPath, "config", "c", "", "配置文件目录路径")
	rootCmd.AddCommand(webCmd)
	// 这里可以继续添加其他子命令，如 migrateCmd、seedCmd 等
	if err := rootCmd.Execute(); err != nil {
		logger.Errorf("Error: %v", err)
		os.Exit(1)
	}
}
