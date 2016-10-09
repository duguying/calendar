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

	app.Commands = []cli.Command{
		cmd.CmdServe,
		cmd.CmdAdd,
		cmd.CmdRemove,
	}

	storage.InitStorage()

	app.Run(os.Args)
}
