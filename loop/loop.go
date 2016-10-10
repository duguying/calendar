package loop

import (
	"github.com/duguying/calendar/storage"
	"time"
)

func CalendarLoop() {
	timerSecond := 0
	timerMinute := 0
	timerHour := 0
	timerDay := 0
	for {
		time.Sleep(1 * time.Second)
		perSecond(timerSecond)

		timerSecond++
		if timerSecond > 59 {
			timerSecond = 0
			perMinute(timerMinute)

			timerMinute++
			if timerMinute > 59 {
				timerMinute = 0
				perHour(timerHour)

				timerHour++
				if timerHour > 23 {
					timerHour = 0
					perDay(timerDay)

					timerDay++
				}
			}
		}
	}
}

func perMinute(tick int) {

}

func perHour(tick int) {

}

func perDay(tick int) {

}

func perSecond(tick int) {
	storage.LoadData()
	calendar := storage.GetCalendar()
	data := calendar.Data
	for _, reminder := range data {
		reminder.CheckBirth()
	}
}
