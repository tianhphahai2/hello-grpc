package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	RootCmd = cobra.Command{}
)

func Excute()  {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init()  {
	// viper
}
