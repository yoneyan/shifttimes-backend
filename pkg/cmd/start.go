package cmd

import (
	"github.com/axis-hineno/shifttimes-backend/pkg/api"
	"github.com/axis-hineno/shifttimes-backend/pkg/api/core/tool/config"
	"github.com/spf13/cobra"
	logging "github.com/vmmgr/controller/pkg/api/core/tool/log"
	"log"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start server",
	Long:  `start server`,
	Run: func(cmd *cobra.Command, args []string) {
		confPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if config.GetConfig(confPath) != nil {
			log.Fatalf("error config process |%v", err)
		}

		logging.WriteLog("------Application Start(User)------")

		api.ControllerAPI()
		log.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.PersistentFlags().StringP("config", "c", "", "config path")
}
