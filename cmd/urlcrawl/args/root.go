package args

import (
	"context"
	"errors"
	"fmt"
	"rytsh/urlcrawl/internal/process"
	"rytsh/urlcrawl/internal/resource"
	"rytsh/urlcrawl/internal/storage/file"

	"github.com/rakunlabs/logi"
	"github.com/spf13/cobra"
)

var Config = struct {
	LogLevel string
	Storage  struct {
		File struct {
			Destionation string
		}
	}
}{
	LogLevel: "info",
}

var BuildVars = struct {
	Version string
	Commit  string
	Date    string
}{
	Version: "v0.0.0",
	Commit:  "-",
	Date:    "-",
}

var rootCmd = &cobra.Command{
	Use:   "urlcrawl",
	Short: "url crawl to download path recursively",
	PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
		if err := logi.SetLogLevel(Config.LogLevel); err != nil {
			return err //nolint:wrapcheck // no need
		}

		return nil
	},
	Example:       `  urlcrawl -d /tmp https://example.com`,
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if Config.Storage.File.Destionation == "" {
			return errors.New("destination is required")
		}

		if len(args) == 0 {
			return errors.New("url is required")
		}

		logi.Log(fmt.Sprintf(
			"urlcrawl [%s] buildCommit=[%s] buildDate=[%s]",
			BuildVars.Version, BuildVars.Commit, BuildVars.Date,
		))

		if err := runRoot(cmd.Context(), args[0]); err != nil {
			return err
		}

		return nil
	},
}

// setFlags sets flags for root CLI.
func setFlags() {
	rootCmd.PersistentFlags().StringVarP(&Config.LogLevel, "log-level", "l", Config.LogLevel, "log level")
	rootCmd.PersistentFlags().StringVarP(&Config.Storage.File.Destionation, "destination", "d", Config.Storage.File.Destionation, "destination")
}

func Execute(ctx context.Context) error {
	setFlags()

	rootCmd.Version = BuildVars.Version
	rootCmd.Long = fmt.Sprintf(
		"%s\nversion:[%s] commit:[%s] buildDate:[%s]",
		rootCmd.Long, BuildVars.Version, BuildVars.Commit, BuildVars.Date,
	)

	return rootCmd.ExecuteContext(ctx) //nolint:wrapcheck // no need
}

func runRoot(ctx context.Context, url string) error {
	resourcePort, err := resource.New(url)
	if err != nil {
		return fmt.Errorf("failed to create resource: %w", err)
	}

	storagePort, err := file.New(Config.Storage.File.Destionation)
	if err != nil {
		return fmt.Errorf("failed to create storage: %w", err)
	}

	processor := process.New(storagePort, resourcePort)

	// //////////////////////////////

	if err := processor.Process(ctx, url); err != nil {
		return fmt.Errorf("failed to process: %w", err)
	}

	return nil
}
