package main

import (
	"github.com/codegangsta/cli"
	"github.com/duguying/calendar/cmd"
	"github.com/duguying/calendar/storage"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Calendar"
	app.Usage = "My Calendar!"
	app.Version = "1.0.1010"
	app.Author = "独孤影"
	app.Email = "root@duguying.net"

	app.Commands = []cli.Command{
		cmd.CmdServe,
		cmd.CmdAdd,
		cmd.CmdRemove,
		cmd.CmdList,
	}

	storage.InitStorage()

	app.Run(os.Args)
}
