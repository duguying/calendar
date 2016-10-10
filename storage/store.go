package storage

import (
	"encoding/json"
	"fmt"
	"github.com/duguying/calendar/sms"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/gogather/lunar"
	"path/filepath"
	"sort"
	"time"
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

func (r *Reminder) CheckBirth() {

	now := time.Now()
	nowSolor := lunarsolar.TimeToSolar(now)
	nowLunar := lunarsolar.SolarToLunar(*nowSolor)
	if r.RemindTimeIsLunar {
		lunarDate := lunarsolar.Lunar{
			IsLeap:     r.RemindTimeLeap,
			LunarDay:   r.RemindTimeDay,
			LunarMonth: r.RemindTimeMonth,
			LunarYear:  r.RemindTimeYear,
		}

		// compare
		if lunarDate.LunarMonth == nowLunar.LunarMonth {
			if lunarDate.LunarDay == nowLunar.LunarDay {
				if now.Hour() == r.RemindTimeHour {
					if now.Minute() == r.RemindTimeMinute {
						if now.Second() == r.RemindTimeSecond {
							// shoot
							log.Blueln("==== Birthday Remind ====")
							sms.SendSMS(r.Phone, fmt.Sprintf("%d-%d-%d", r.RemindTimeYear, r.RemindTimeMonth, r.RemindTimeDay), r.Event)
						}
					}
				}
			}
		}
	}
}

type Calendar struct {
	Count int              `json:"count"`
	Data  map[int]Reminder `json:"data"`
}

func (c *Calendar) ListAll() {
	count := c.Count
	data := c.Data
	log.Bluef("共%d条提醒\n", count)

	sorted_keys := make([]int, 0)
	for k, _ := range data {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Ints(sorted_keys)

	for _, idx := range sorted_keys {
		r := data[idx]
		log.Pinkf("[%d]\t", idx)
		log.Greenf("%d-%d-%d %d:%d:%d\t%s\n", r.RemindTimeYear, r.RemindTimeMonth, r.RemindTimeDay, r.RemindTimeHour, r.RemindTimeMinute, r.RemindTimeSecond, r.Event)
	}
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

func AddItem(phone string, lunar bool, year int, month int, day int, h int, m int, leap bool, event string) {
	reminder := NewReminder(phone, lunar, year, month, day, h, m, 0, leap, event)
	count := calendar.Count + 1
	calendar.Data[count] = reminder
	calendar.Count = count
	log.Pinkf("add [%d] ", count)
	log.Bluef("%d-%d-%d %d:%d %s\n", reminder.RemindTimeYear, reminder.RemindTimeMonth, reminder.RemindTimeDay, reminder.RemindTimeHour, reminder.RemindTimeMinute, reminder.Event)
	SaveData()
}

func RemoveItem(id int) {
	count := calendar.Count
	if _, ok := calendar.Data[id]; ok {
		data := calendar.Data[id]
		log.Pinkf("remove [%d] ", id)
		log.Bluef("%d-%d-%d %d:%d %s\n", data.RemindTimeYear, data.RemindTimeMonth, data.RemindTimeDay, data.RemindTimeHour, data.RemindTimeMinute, data.Event)
		delete(calendar.Data, id)
		calendar.Count = count - 1

		SaveData()
	} else {
		log.Yellowf("reminder [%d] not exist\n", id)
	}
}

func LoadData() {
	path := getConfigPath()
	jsonBytes, err := com.ReadFileByte(path)
	if err != nil {
		log.Redln(err)
	}
	err = json.Unmarshal(jsonBytes, &calendar)
	if err != nil {
		log.Redln("load config failed")
	}
}

func SaveData() {
	if b, err := json.Marshal(calendar); err == nil {
		fullpath := getConfigPath()
		com.WriteFile(fullpath, string(b))
	}
}

func GetCalendar() *Calendar {
	return &calendar
}
