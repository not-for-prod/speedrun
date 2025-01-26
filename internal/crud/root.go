package crud

import "C"
import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "crud",
	Short: "writes simple crud 4 u",
	Long:  `writes simple crud 4 u`,
	Run: func(cmd *cobra.Command, args []string) {
		src := cmd.Flag("src").Value.String()
		dst := cmd.Flag("dst").Value.String()

		command := generationCommandFromString(src, dst)
		command.execute()
	},
}

func init() {
	Cmd.Flags().String("src", "", "path in format <path to file>::<struct name>::<id field>")
	Cmd.Flags().String("dst", "", "path where to generate code")
}
