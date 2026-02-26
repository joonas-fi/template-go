package main

import (
	"context"

	"github.com/function61/gokit/app/cli"
	"github.com/spf13/cobra"
)

func main() {
	app := &cobra.Command{
		// Short:   "Description for app",
	}

	app.AddCommand(&cobra.Command{
		Use: "reticulate-splines",
		// Short: "Reticulates splines",
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			<-cmd.Context().Done()
			return nil

		},
	})

	cli.Execute(app)
}

func logic(ctx context.Context) error {
	return nil
}
