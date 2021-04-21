package main

import (
	"context"
	"os"

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
		Run: func(_ *cobra.Command, args []string) {
			osutil.ExitIfError(
				logic(osutil.CancelOnInterruptOrTerminate(nil)))
		},
	})

	osutil.ExitIfError(app.Execute())
}

func logic(ctx context.Context) error {
	return nil
}
