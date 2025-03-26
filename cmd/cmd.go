package cmd

import (
	"fuck-the-world/database"
	"fuck-the-world/internal/pkg/email"
	"fuck-the-world/internal/router"
	"fuck-the-world/unit"
	"github.com/spf13/cobra"
	"log"
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
	var emailCmd = &cobra.Command{
		Use:   "email",
		Short: "发送邮件",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				log.Println("请输入目标邮箱")
			} else {
				err := email.SendCode(args[0], "123456")
				if err != nil {
					unit.HandleError("sendEmail接口发送邮件失败", err)
				}
			}
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
	database.Connect()
	rootCmd.AddCommand(serverCmd)
	rootCmd.AddCommand(dbCmd)
	rootCmd.AddCommand(emailCmd)
	dbCmd.AddCommand(migrateCmd)
	err := rootCmd.Execute()
	if err != nil {
		panic(err)
	}
}
