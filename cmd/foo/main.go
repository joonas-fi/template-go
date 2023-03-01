package main

import (
	"context"
	"log"
	"os"

	"github.com/function61/gokit/app/cli"
	"github.com/function61/gokit/app/dynversion"
	"github.com/function61/gokit/os/osutil"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		Use: os.Args[0],
		// Short:   "Description for app",
		Version: dynversion.Version,
	}

	app.AddCommand(&cobra.Command{
		Use: "reticulate-splines",
		// Short: "Reticulates splines",
		Args: cobra.NoArgs,
		Run: cli.RunnerNoArgs(func(ctx context.Context, _ *log.Logger) error {
			<-ctx.Done()
			return nil
		}),
	})

	osutil.ExitIfError(app.Execute())
}

func logic(ctx context.Context) error {
	return nil
}
