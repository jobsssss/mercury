package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"mercury/app/cmd"
	"mercury/app/cmd/make"
	"mercury/bootstrap"
	btsConfig "mercury/config"
	"mercury/pkg/config"
	"mercury/pkg/console"
	"os"
)

func init() {
	btsConfig.Initialize()
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "mercury",
		Short: "A web application skeleton",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommand`,

		PersistentPreRun: func(command *cobra.Command, args []string) {
			config.InitConfig(cmd.Env)
			bootstrap.SetupLogger()
			bootstrap.SetupDB()
			bootstrap.SetupRedis()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,
		cmd.CmdKey,
		cmd.CmdPlay,
		make.CmdMake,
	)
	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)
	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}
