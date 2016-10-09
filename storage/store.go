package storage

import (
	"encoding/json"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"path/filepath"
)

func InitStorage() {
	initProfileDir()
}

func initProfileDir() {
	home := getHome()
	path := filepath.Join(home, ".calendar")

	if !com.FileExist(path) {
		err := com.Mkdir(path)
		if err != nil {
			log.Fatalln("Create profile directory failed!")
		}
	}
}

// getHome - get current user path
func getHome() string {
	home, err := com.Home()
	if err != nil {
		log.Fatalln("Can NOT find user path!")
	}
	return home
}

func getConfigPath() string {
	home := getHome()
	path := filepath.Join(home, ".calendar", "config.json")
	return path
}

type Reminder struct {
	Phone string `json:"phone"`

	RemindTimeIsLunar bool `json:"lunar"`
	RemindTimeYear    int  `json:"year"`
	RemindTimeMonth   int  `json:"month"`
	RemindTimeDay     int  `json:"day"`

	RemindTimeHour   int  `json:"hour"`
	RemindTimeMinute int  `json:"minute"`
	RemindTimeSecond int  `json:"second"`
	RemindTimeLeap   bool `json:"leap"`

	Event string `json:"event"`
}

type Calendar struct {
	Count int              `json:"count"`
	Data  map[int]Reminder `json:"data"`
}

var calendar Calendar

func init() {
	calendar = Calendar{Count: 0, Data: make(map[int]Reminder)}
	path := getConfigPath()
	if com.FileExist(path) {
		LoadData()
	} else {
		SaveData()
	}
}

func NewReminder(phone string, lunar bool, year int, month int, day int, h int, m int, s int, leap bool, event string) Reminder {
	return Reminder{
		Phone:             phone,
		RemindTimeIsLunar: lunar,
		RemindTimeYear:    year,
		RemindTimeMonth:   month,
		RemindTimeDay:     day,
		RemindTimeHour:    h,
		RemindTimeMinute:  m,
		RemindTimeSecond:  s,
		RemindTimeLeap:    leap,
		Event:             event,
	}
}

func AddItem(phone string, lunar bool, year int, month int, day int, h int, m int, s int, leap bool, event string) {
	reminder := NewReminder(phone, lunar, year, month, day, h, m, s, leap, event)
	count := calendar.Count + 1
	calendar.Data[count] = reminder
	SaveData()
}

func LoadData() {
	path := getConfigPath()
	jsonBytes, err := com.ReadFileByte(path)
	if err != nil {
		log.Redln(err)
	}
	if err := json.Unmarshal(jsonBytes, &calendar); err == nil {
		log.Greenln(calendar)
	}
}

func SaveData() {
	if b, err := json.Marshal(calendar); err == nil {
		fullpath := getConfigPath()
		com.WriteFile(fullpath, string(b))
	}
}
