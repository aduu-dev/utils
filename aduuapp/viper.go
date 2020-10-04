package aduuapp

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"k8s.io/klog/v2"
)

func prefixFlagset(head *pflag.FlagSet, prefix string) {
	head.SetNormalizeFunc(func(f *pflag.FlagSet, name string) pflag.NormalizedName {
		return pflag.NormalizedName(prefix + "." + name)
	})
}

// ViperSetup stores config requires to fully setup an aduu app.
type ViperSetup struct {
	defaultConfigPath string
	defaultConfig     string
	writeConfigFlag   *bool
	readConfigErr     error

	config *SetupViperConfig

	v *viper.Viper
}

// SetupViperConfig options ot pass to SetupViper.
type SetupViperConfig struct {
}

// SetupViper sets up viper for immediate use.
func SetupViper(appName string, req *SetupViperConfig) (setup *ViperSetup, v *viper.Viper) {
	defaultConfigPath := filepath.Join(os.Getenv("HOME"), "."+appName)
	defaultConfig := filepath.Join(defaultConfigPath, appName+".yaml")
	v = viper.New()
	v.SetEnvPrefix(appName)
	v.AutomaticEnv()
	v.SetConfigName(appName)
	v.AddConfigPath("/etc/" + appName)
	v.AddConfigPath(defaultConfigPath)
	v.AddConfigPath(".")

	return &ViperSetup{
		defaultConfig:     defaultConfig,
		defaultConfigPath: defaultConfigPath,
		readConfigErr:     v.ReadInConfig(),
		config:            req,
		v:                 v,
	}, v
}

// SetupFlagsConfig contains options to enable or disable certain flags from being added.
type SetupFlagsConfig struct {
	DisableWriteConfigFlag bool
}

// SetupFlags adds klog flags to the given command together with a -w flag.
// The -w flag enables saving the current viper state to the default config file.
func (setup *ViperSetup) SetupFlags(
	cmd *cobra.Command,
	flagsConfig *SetupFlagsConfig,
) {
	fs := flag.NewFlagSet("log", flag.ExitOnError)
	klog.InitFlags(fs)

	cmd.PersistentFlags().AddGoFlag(fs.Lookup("v"))
	cmd.PersistentFlags().AddGoFlag(fs.Lookup("logtostderr"))

	if !flagsConfig.DisableWriteConfigFlag {
		setup.writeConfigFlag = cmd.Flags().BoolP(
			"write-given-settings", "w", false,
			"writes all given persistent flags to the config file and exits the program",
		)
	}
}

// IsWriteConfigSet return whether the config write flag was set.
func (setup *ViperSetup) IsWriteConfigSet() bool {
	return setup.writeConfigFlag != nil && *setup.writeConfigFlag
}

func (setup *ViperSetup) WriteConfig() (err error) {
	klog.InfoS("Writing config", "config", setup.defaultConfig)

	if err = os.MkdirAll(setup.defaultConfigPath, 0700); err != nil {
		return
	}

	// If we read config successfully then write there.
	if setup.readConfigErr == nil {
		return setup.v.WriteConfig()
	}

	// Write to default config location.
	return setup.v.WriteConfigAs(setup.defaultConfig)
}
