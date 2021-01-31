package main

import (
	"github.com/tubone24/redump/pkg/cmd"
	"github.com/tubone24/redump/pkg/config"
	"github.com/docopt/docopt-go"
)


func main() {
	usage := `Redump.

Usage:
  redump migrate
  redump list
  redump -h | --help
  redump --version

Options:
  -h --help     Show this screen.
  --version     Show version.`

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	arguments, _ := docopt.ParseDoc(usage)
	//fmt.Println(arguments)
	flag, err := arguments.Bool("migrate")
	if flag {
		cmd.Migrate(cfg.ServerConfig.ProjectId)
	}
	flag, err = arguments.Bool("list")
	if flag {
		cmd.ListAll(cfg.ServerConfig.ProjectId)
	}

}
