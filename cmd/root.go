package cmd

import (
	"fmt"
	"github.com/rakib-09/golang-clean-architecture/utils/logger"
	"os"

	"github.com/rakib-09/golang-clean-architecture/config"
	"github.com/rakib-09/golang-clean-architecture/conn"

	"github.com/spf13/cobra"
)

var (
	RootCmd = &cobra.Command{
		Use: "golang-clean-architecture",
	}
)

func init() {
	RootCmd.AddCommand(serveCmd)
	RootCmd.AddCommand(workerCmd)
}

// Execute executes the root command
func Execute() {
	logger.InitLogger()
	config.LoadConfig()
	conn.ConnectRedis()
	conn.InitAsynqClient()
	conn.InitAsyncInspector()

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
