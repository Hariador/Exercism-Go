package meetup

import (
	"time"
)

type WeekSchedule struct {
	val int
}

var (
	First  = WeekSchedule{0}
	Second = WeekSchedule{1}
	Third  = WeekSchedule{2}
	Fourth = WeekSchedule{3}
	Teenth = WeekSchedule{6}
	Last   = WeekSchedule{7}
)

func Day(wSched WeekSchedule, wDay time.Weekday, month time.Month, year int) int {
	t := time.Date(year, month, 1, 0, 0, 0, 0, time.UTC)

	for t.Weekday() != wDay {
		t = t.AddDate(0, 0, 1)
	}
	switch wSched.val {
	case First.val, Second.val, Third.val, Fourth.val:

		t = t.AddDate(0, 0, wSched.val*7)

	case Teenth.val:

		for t.Day() < 13 {
			t = t.AddDate(0, 0, 7)
		}

	case Last.val:
		{
			t = t.AddDate(0, 0, 28)
			if t.Day() < 21 {
				t = t.AddDate(0, 0, -7)
			}
		}
	}
	return t.Day()

}
