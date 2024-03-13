package cmd

import (
	"os"
	"fmt"
	"os/exec"
	"strings"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)


var rootCmd = &cobra.Command{
	Use: "ouroboros",
	Short: "Ouroboros automates needed tasks for running Ubuntu on a MacBook",
	Long: "To keep the Ubuntu powered MacBooks functioning at work, Ouroboros is used to run commands that are often needed",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/ouroboros/")
	err := viper.ReadInConfig()
	if err != nil { 
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
  
func Execute() {
	if err := rootCmd.Execute(); err != nil {
	  fmt.Fprintln(os.Stderr, err)
	  os.Exit(1)
	}
  }


func ExecuteCommand(command []string, successMsg string, errorMsg string) error {
	execCommand := exec.Command("sudo", command...)
	output , err := execCommand.Output()
	if err != nil {
		if len(string(output)) > 0 {
			fmt.Println(errorMsg)
			return err
		}
	} else {
		stringOutput :=  strings.TrimSpace(string(output))
		printString := strings.ToLower(string(stringOutput))
		fmt.Println(successMsg, printString)
	}
	return nil
}
