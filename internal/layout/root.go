package layout

import (
	"github.com/not-for-prod/speedrun/internal/pkg/logger"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "layout",
	Short: "creates common folder structure",
	Long:  `creates common folder structure`,
	Run: func(cmd *cobra.Command, args []string) {
		serviceName := cmd.Flag("svc").Value.String()

		if serviceName == "" {
			logger.Fatalf("svc flag is required: go run main.go layout --svc <service_name>")
		}

		command := layoutCommand{
			serviceName: serviceName,
		}
		command.execute()
	},
}

func init() {
	Cmd.Flags().String("svc", "", "service name")
}
