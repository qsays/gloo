package cmd

import (
	"context"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/remove"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/route"
	version2 "github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/version"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/add"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/create"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/del"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/edit"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/get"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/install"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/upgrade"
	"github.com/solo-io/go-utils/cliutils"

	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/gateway"
	"github.com/solo-io/gloo/projects/gloo/cli/pkg/cmd/options"
	"github.com/spf13/cobra"
)

var versionTemplate = `{{with .Name}}{{printf "%s community edition " .}}{{end}}{{printf "version %s" .Version}}
`

func App(optionsFunc ...cliutils.OptionsFunc) *cobra.Command {

	app := &cobra.Command{
		Use:   "glooctl",
		Short: "CLI for Gloo",
		Long: `glooctl is the unified CLI for Gloo.
	Find more information at https://solo.io`,
	}

	// Complete additional passed in setup
	cliutils.ApplyOptions(app, optionsFunc)

	app.SetVersionTemplate(versionTemplate)

	return app
}

func GlooCli() *cobra.Command {
	opts := &options.Options{
		Top: options.Top{
			Ctx: context.Background(),
		},
	}

	optionsFunc := func(app *cobra.Command) {
		pflags := app.PersistentFlags()
		pflags.BoolVarP(&opts.Top.Interactive, "interactive", "i", false, "use interactive mode")

		app.SuggestionsMinimumDistance = 1
		app.AddCommand(
			get.RootCmd(opts),
			del.RootCmd(opts),
			install.InstallCmd(opts),
			install.UninstallCmd(opts),
			add.RootCmd(opts),
			remove.RootCmd(opts),
			route.RootCmd(opts),
			create.RootCmd(opts),
			edit.RootCmd(opts),
			upgrade.RootCmd(opts),
			gateway.RootCmd(opts),
			version2.VersionCmd(opts),
		)
	}
	cmd := App(optionsFunc)
	return cmd
}
