package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/duguying/calendar/storage"
	"github.com/gogather/com/log"
)

func AddRemind(c *cli.Context) {
	log.Pinkln("add")
	storage.AddItem("15989553147", true, 2016, 10, 10, 10, 10, 10, true, "hello event")
}

func RemoveRemind(c *cli.Context) {
	log.Pinkln("remove")
}
