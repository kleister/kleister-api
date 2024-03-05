package command

import (
	"github.com/kleister/kleister-api/pkg/config"
	"github.com/kleister/kleister-api/pkg/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:           "kleister-api",
		Short:         "Manage mod packs for minecraft",
		Version:       version.String,
		SilenceErrors: false,
		SilenceUsage:  true,

		PersistentPreRunE: func(ccmd *cobra.Command, args []string) error {
			return setupLogger()
		},

		CompletionOptions: cobra.CompletionOptions{
			DisableDefaultCmd: true,
		},
	}

	cfg *config.Config
)

func init() {
	cfg = config.Load()
	cobra.OnInitialize(setupConfig)

	rootCmd.PersistentFlags().BoolP("help", "h", false, "Show the help")
	rootCmd.PersistentFlags().BoolP("version", "v", false, "Print the version")

	rootCmd.PersistentFlags().String("config-file", "", "Path to config file")
	_ = viper.BindPFlag("config.file", rootCmd.PersistentFlags().Lookup("config-file"))

	rootCmd.PersistentFlags().String("log-level", "info", "Set logging level")
	viper.SetDefault("log.level", "info")
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	rootCmd.PersistentFlags().Bool("log-pretty", true, "Enable pretty logging")
	viper.SetDefault("log.pretty", true)
	_ = viper.BindPFlag("log.pretty", rootCmd.PersistentFlags().Lookup("log-pretty"))

	rootCmd.PersistentFlags().Bool("log-color", true, "Enable colored logging")
	viper.SetDefault("log.color", true)
	_ = viper.BindPFlag("log.color", rootCmd.PersistentFlags().Lookup("log-color"))
}

// Run parses the command line arguments and executes the program.
func Run() error {
	return rootCmd.Execute()
}
