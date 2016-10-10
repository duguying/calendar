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
	Name:  "add",
	Usage: "This command should will add a reminder to calendar config",
	Description: `lunar - L R1 D2016.10.10 T10:23 P15911113147 Eevents    D:date    T:time    P:Phone    E:Events    R1:Leap
   solar - S D2016.10.10 T12:23 P15911113147 Eevents       D:date    T:time    P:Phone    E:Events`,
	Action: AddRemind,
}

var CmdRemove = cli.Command{
	Name:        "remove",
	Usage:       "This command should will remove a reminder from calendar config",
	Description: `id as arguments`,
	Action:      RemoveRemind,
}

var CmdList = cli.Command{
	Name:        "list",
	Usage:       "This command should will list all reminder from calendar config",
	Description: `list all reminder`,
	Action:      ListRemind,
}
