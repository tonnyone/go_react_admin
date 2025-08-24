package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"github.com/tonnyone/go_react_admin/internal/config"
	"github.com/tonnyone/go_react_admin/internal/database"
	"github.com/tonnyone/go_react_admin/router"
	"os"
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
			fmt.Fprintf(os.Stderr, "加载配置失败: %v\n", err)
			os.Exit(1)
		}
		// 初始化数据库
		_, err = database.InitDB(cfg.DB.DSN)
		if err != nil {
			fmt.Fprintf(os.Stderr, "数据库初始化失败: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("配置加载成功，服务端口: %d\n", cfg.App.Port)
		gin.SetMode(gin.DebugMode)
		r := gin.New()
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
		router.RegisterRoutes(r)
		r.Run(fmt.Sprintf(":%d", cfg.App.Port))
	},
}

func main() {
	webCmd.Flags().StringVarP(&configPath, "config", "c", "", "配置文件目录路径")
	rootCmd.AddCommand(webCmd)
	// 这里可以继续添加其他子命令，如 migrateCmd、seedCmd 等
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
