package main

import (
	"github.com/docopt/docopt-go"
	"github.com/tubone24/redump/pkg/cmd"
	"github.com/tubone24/redump/pkg/config"
)

func main() {
	usage := `REDUMP
A tool to migrate data in your Redmine without admin accounts.

Usage:
  redump migrate [-i|--issue <number>]
  redump list
  redump dump [-c|--concurrency] [-i|--issue <number>]
  redump restore [-i|--issue <number>]
  redump -h|--help
  redump --version

Options:
  -h --help                  Show this screen.
  -c --concurrency           Concurrency Request Danger!
  -i --issue                 Specify Issues
  --version                  Show version.`

	cfg, err := config.GetConfig()
	if err != nil {
		panic(err)
	}
	arguments, _ := docopt.ParseDoc(usage)
	err = arguments.Bind(&cmd.DocOptConf)
	if err != nil {
		panic(err)
	}
	if cmd.DocOptConf.Migrate {
		if cmd.DocOptConf.Issue && cmd.DocOptConf.Number != 0 {
			//
		} else {
			err = cmd.Migrate(cfg.ServerConfig.ProjectId)
			if err != nil {
				panic(err)
			}
		}
	}
	if cmd.DocOptConf.List {
		err = cmd.ListAll(cfg.ServerConfig.ProjectId)
		if err != nil {
			panic(err)
		}
	}
	if cmd.DocOptConf.Dump {
		if cmd.DocOptConf.Issue && cmd.DocOptConf.Number != 0 {
			//
		} else {
			cmd.Dump(cfg.ServerConfig.ProjectId, cmd.DocOptConf.Concurrency)
		}
	}
	if cmd.DocOptConf.Restore {
		if cmd.DocOptConf.Issue && cmd.DocOptConf.Number != 0 {
			//
		} else {
			err = cmd.RestoreDataFromLocal(cfg.ServerConfig.ProjectId)
			if err != nil {
				panic(err)
			}
		}
	}

}
