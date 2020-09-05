package cmd

import (
	playlists "github.com/shanedabes/subes/internal/playlists"
	"github.com/shanedabes/subes/internal/subscriptions"
	"github.com/spf13/cobra"
)

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subes",
		Short: "A simple cli for managing YouTube resources",
	}

	cmd.AddCommand(
		getCmd(),
	)

	return cmd
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get resources",
	}

	cmd.AddCommand(
		getPlaylistsCmd(),
		getSubsCommand(),
	)

	return cmd
}

func getPlaylistsCmd() *cobra.Command {
	c := playlists.NewCollector()

	cmd := &cobra.Command{
		Use:     "playlists",
		Aliases: []string{"pl"},
		Short:   "Get playlists",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initParams(cmd, c.Params)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Playlists()
		},
	}

	return cmd
}

func getSubsCommand() *cobra.Command {
	c := subscriptions.NewCollector()

	cmd := &cobra.Command{
		Use:     "subscriptions",
		Aliases: []string{"subs"},
		Short:   "Get subscriptions",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return initParams(cmd, c.Params)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return c.Subscriptions()
		},
	}

	return cmd
}
