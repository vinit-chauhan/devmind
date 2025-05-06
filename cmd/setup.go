package cmd

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	ui "github.com/vinit-chauhan/devmind/cmd/ui/setup"
	"github.com/vinit-chauhan/devmind/internal/constants"
)

var setupCmd = &cobra.Command{
	Use:     "setup",
	Args:    cobra.NoArgs,
	Short:   "Setup devmind",
	Long:    `Setup devmind. This command will set up the devmind environment and install all the required dependencies. It will also create a config file if it does not exist. It is recommended to run this command before using devmind for the first time.`,
	Example: `devmind setup`,
	RunE:    runSetup,
}

func runSetup(cmd *cobra.Command, args []string) error {
	// TUI to create config, setup log path.
	config_path := constants.SETUP_PATH

	fmt.Println("Config path:", config_path)

	p := tea.NewProgram(ui.InitialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		return err
	}

	return nil
}

func init() {
	rootCmd.AddCommand(setupCmd)
}
