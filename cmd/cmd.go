package cmd

import (
	"bookkeeping-server/database"
	"bookkeeping-server/internal/router"
	"github.com/spf13/cobra"
)

func Run() {
	var rootCmd = &cobra.Command{
		Use:   "run",
		Short: "启动程序",
		Long:  `启动程序`,
	}
	var dbCmd = &cobra.Command{
		Use:   "db",
		Short: "启动数据库",
		Long:  `启动数据库`,
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
		},
	}
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "启动服务器",
		Long:  `启动服务器`,
		Run: func(cmd *cobra.Command, args []string) {
			router.RunServer()
		},
	}
	var allCmd = &cobra.Command{
		Use:   "all",
		Short: "启动数据库和服务器",
		Long:  `启动数据库和服务器`,
		Run: func(cmd *cobra.Command, args []string) {
			database.Connect()
			router.RunServer()
		},
	}
	rootCmd.AddCommand(dbCmd, serverCmd, allCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
