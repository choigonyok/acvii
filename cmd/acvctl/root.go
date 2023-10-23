package cmd

import (
	"flag"
	"path/filepath"
	"strings"

	"github.com/choigonyok/goopt/pkg/env"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	// acvctlConfigPath = "$HOME/.acvctl/config.yml"
	acvctlConfigPath = "../../../.acvctl/config.yml"
)

var rootCmd = &cobra.Command{
	Use:   "acvctl",
	Short: "help",
	Long:  "to see help, please use flag --help / -h",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

var installCMD = &cobra.Command{
	Use:   "install",
	Short: "install acvii pod",
	Long:  "install acvii pod in Kubernetes cluster",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Print("INSTALL COMMAND RUNNING")
	},
}

// var acvctl = env.NewStringVar("CONFIG", "for initial configuration of acvctl", AcvctlConfigPath)

func init() {
	env.ForceNewStringVar("CONFIG", "for initial configuration of acvctl", acvctlConfigPath)
	rootCmd.AddCommand(installCMD)
}

func Execute() error {
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	return rootCmd.Execute()
}

func ConfigAcvctl() error {
	acvctlPath := env.Get("CONFIG").DefaultValue
	// if !ok {
	// 	return errors.New("acvctl config file is not registered as env variables")
	// }
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
