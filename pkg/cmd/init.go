package cmd

import (
	"github.com/axis-hineno/shifttimes-backend/pkg/api/core/tool/config"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init database",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		confPath, err := cmd.Flags().GetString("config")
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		if config.GetConfig(confPath) != nil {
			log.Fatalf("error config process |%v", err)
		}

		db, err := gorm.Open("mysql", config.Conf.DB.User+":"+config.Conf.DB.Pass+"@"+
			"tcp("+config.Conf.DB.IP+":"+strconv.Itoa(config.Conf.DB.Port)+")"+"/"+config.Conf.DB.DBName+"?parseTime=true")
		if err != nil {
			panic(err)
		}
		result := db.AutoMigrate()
		log.Println(result.Error)
		log.Println("end")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
