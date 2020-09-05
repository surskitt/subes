package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func initParams(cmd *cobra.Command, params interface{}) error {
	v := viper.New()

	if err := v.BindPFlags(cmd.Flags()); err != nil {
		return fmt.Errorf("couldn't process command flags, %v", err)
	}

	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("couldn't find config directory, %v", err)
	}

	v.AddConfigPath(configDir)
	v.SetConfigName("subes")

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return fmt.Errorf("couldn't read config file, %v", err)
		}
	}

	if err := v.Unmarshal(params); err != nil {
		return fmt.Errorf("couldn't parse config, %v", err)
	}

	return nil
}
