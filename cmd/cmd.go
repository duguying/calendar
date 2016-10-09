package cmd

import (
	"github.com/codegangsta/cli"
)

var CmdServe = cli.Command{
	Name:        "serve",
	Usage:       "This command should will start as a service",
	Description: `Serv provide access auth for repositories`,
	Action:      Serve,
}

var CmdAdd = cli.Command{
	Name:        "add",
	Usage:       "This command should will start as a service",
	Description: `Serv provide access auth for repositories`,
	Action:      AddRemind,
}

var CmdRemove = cli.Command{
	Name:        "remove",
	Usage:       "This command should will start as a service",
	Description: `Serv provide access auth for repositories`,
	Action:      RemoveRemind,
}
