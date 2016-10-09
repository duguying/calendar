package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/duguying/calendar/loop"
)

func Serve(c *cli.Context) {
	loop.CalendarLoop()
}
