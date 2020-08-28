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

	v *viper.Viper
}

// SetupViper sets up viper for immediate use.
func SetupViper(appName string) (setup *ViperSetup, v *viper.Viper) {
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
		v:                 v,
	}, v
}

// SetupKlogFlags adds klog flags to the given command together with a -w flag.
// The -w flag enables saving the current viper state to the default config file.
func (setup *ViperSetup) SetupKlogFlags(cmd *cobra.Command) {
	fs := flag.NewFlagSet("klog", flag.ExitOnError)
	klog.InitFlags(fs)
	flag.Parse()

	cmd.PersistentFlags().AddGoFlag(fs.Lookup("v"))
	cmd.PersistentFlags().AddGoFlag(fs.Lookup("logtostderr"))

	setup.writeConfigFlag = cmd.Flags().BoolP(
		"write-given-settings", "w", false,
		"writes all given persistent flags to the config file and exits the program",
	)
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
