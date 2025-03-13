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
	var serverCmd = &cobra.Command{
		Use:   "server",
		Short: "启动服务器",
		Long:  `启动服务器`,
		Run: func(cmd *cobra.Command, args []string) {
			router.RunServer()
		},
	}
	var dbCmd = &cobra.Command{
		Use: "db",
	}
	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "迁移数据库",
		Long:  `迁移数据库`,
		Run: func(cmd *cobra.Command, args []string) {
			database.Migrate()
		},
	}
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(dbCmd)
	dbCmd.AddCommand(migrateCmd)
	database.Connect()
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
