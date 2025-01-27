package gateway

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "gateway",
	Short: "creates all u need for grpc-gateway",
	Long:  `creates all u need for grpc-gateway, buf.yaml, proto, gateway, swagger`,
	Run: func(cmd *cobra.Command, args []string) {
		src := cmd.Flag("src").Value.String()
		dst := cmd.Flag("dst").Value.String()

		command := gatewayCommand{
			src: src,
			dst: dst,
		}
		command.execute()
	},
}

func init() {
	Cmd.Flags().String("src", "", "path to proto")
	Cmd.Flags().String("dst", "", "path to generated files")
}
