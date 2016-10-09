package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/duguying/calendar/storage"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"strconv"
	"strings"
)

func AddRemind(c *cli.Context) {
	log.Pinkln("add")
	phone, lunar, year, month, day, hours, minute, leap, event := parseArgs(c)
	storage.AddItem(phone, lunar, year, month, day, hours, minute, leap, event)
}

func parseArgs(c *cli.Context) (phone string, lunar bool, year int, month int, day int, hours int, minute int, leap bool, event string) {
	args := c.Args()
	dateType := args.First()

	if dateType == "L" {
		lunar = true
	} else if dateType == "S" {
		lunar = false
	}

	tail := args.Tail()
	for i := 0; i < len(tail); i++ {
		arg := tail[i]
		head := strings.ToUpper(com.SubString(arg, 0, 1))
		left := com.SubString(arg, 1, len(arg)-1)
		switch head {
		case "R":
			{
				if left == "1" {
					leap = true
				} else {
					leap = false
				}
				break
			}
		case "D":
			{
				var err error = nil
				date := strings.Split(left, ".")
				if len(date) > 2 {
					day, err = strconv.Atoi(date[2])
				}
				if len(date) > 1 {
					month, err = strconv.Atoi(date[1])
				}
				if len(date) > 0 {
					year, err = strconv.Atoi(date[0])
				} else {
					log.Fatalln("illeage date")
				}
				if err != nil {
					log.Fatalln("illeage date")
				}
				break
			}
		case "T":
			{
				var err error = nil
				t := strings.Split(left, ":")
				if len(t) > 1 {
					minute, err = strconv.Atoi(t[1])
				}
				if len(t) > 0 {
					hours, err = strconv.Atoi(t[0])
				} else {
					minute = 0
					hours = 8
				}
				if err != nil {
					log.Fatalln("illeage time")
				}
				break
			}
		case "P":
			{
				phone = left
				break
			}
		case "E":
			{
				event = left
				break
			}
		}
	}
	return
}

func RemoveRemind(c *cli.Context) {
	log.Pinkln("remove")
}
