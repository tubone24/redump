package main

import (
	"github.com/tubone24/redump/pkg/cmd"
	"github.com/tubone24/redump/pkg/config"
	"github.com/docopt/docopt-go"
)


func main() {
	usage := `REDUMP
A tool to migrate data in your Redmine without admin accounts.

Usage:
  redump migrate
  redump list
  redump dump [-c|--concurrency]
  redump restore
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
	flag, err := arguments.Bool("migrate")
	if flag {
		err = cmd.Migrate(cfg.ServerConfig.ProjectId)
		if err != nil {
			panic(err)
		}
	}
	flag, err = arguments.Bool("list")
	if flag {
		err = cmd.ListAll(cfg.ServerConfig.ProjectId)
		if err != nil {
			panic(err)
		}
	}
	flag, err = arguments.Bool("dump")
	concurrency, err := arguments.Bool("--concurrency")
	if flag {
		cmd.Dump(cfg.ServerConfig.ProjectId, concurrency)
	}
	flag, err = arguments.Bool("restore")
	if flag {
		err = cmd.RestoreDataFromLocal(cfg.ServerConfig.ProjectId)
		if err != nil {
			panic(err)
		}
	}

}
