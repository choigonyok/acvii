package cmd

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/choigonyok/acvii/pkg/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// acvctlConfigPath = "$HOME/.acvctl/config.yml"
	acvctlConfigPath = "../../../.acvctl/config.yml"
	binaryFilePath   = "../../acvii/acvii"
)

var rootCmd = &cobra.Command{
	Use:   "acvctl",
	Short: "help",
	Long:  "to see help, please use flag --help / -h",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func New() *cobra.Command {
	return rootCmd
}

var startCMD = &cobra.Command{
	Use:   "start",
	Short: "start",
	Long:  "start server",
	Run: func(cmd *cobra.Command, args []string) {
		command := exec.Command(binaryFilePath)

		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err := command.Run()
		if err != nil {
			fmt.Println("실행 중 오류 발생:", err)
		}
	},
}

var installCMD = &cobra.Command{
	Use:   "install",
	Short: "deploy acvii pod",
	Long:  "deploy acvii pod in Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("INSTALL COMMAND RUNNING")
	},
}

var planCMD = &cobra.Command{
	Use:   "plan",
	Short: "check how AuthorizationPolicy gonna change",
	Long:  "check how AuthorizationPolicy gonna change",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("PLAN COMMAND RUNNING")
	},
}

var applyCMD = &cobra.Command{
	Use:   "apply",
	Short: "apply current AuthroizationPolicy settings",
	Long:  "apply current AuthroizationPolicy settings",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("APPLY COMMAND RUNNING")
	},
}

var dashboardCMD = &cobra.Command{
	Use:   "dashboard",
	Short: "show dashboard web UI in browser",
	Long:  "show dashboard web UI in browser",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("DASHBOARD COMMAND RUNNING")
	},
}

var uninstallCMD = &cobra.Command{
	Use:   "uninstall",
	Short: "delete acvii pod",
	Long:  "delete acvii pod in Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("UNINSTALL COMMAND RUNNING")
	},
}

func init() {
	env.ForceNewStringVar("CONFIG", "for initial configuration of acvctl", acvctlConfigPath)
	rootCmd.AddCommand(installCMD)
	rootCmd.AddCommand(planCMD)
	rootCmd.AddCommand(applyCMD)
	rootCmd.AddCommand(dashboardCMD)
	rootCmd.AddCommand(uninstallCMD)
	rootCmd.AddCommand(startCMD)
}

func Execute() error {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	return rootCmd.Execute()
}

func ConfigAcvctl() error {
	acvctlPath := env.Get("CONFIG").DefaultValue

	configPath := filepath.Dir(acvctlPath)
	baseName := filepath.Base(acvctlPath)
	configType := filepath.Ext(acvctlPath)
	configName := baseName[0 : len(baseName)-len(configType)]
	if configType != "" {
		configType = configType[1:]
	}

	viper.SetEnvPrefix("ACVCTL")
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	err := viper.ReadInConfig()

	return err
}
