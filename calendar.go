package calendar

import (
	"fmt"
	"strings"
	"time"
)

type Calendar struct {
	WeekLabels  []string
	MonthLabels []string

	DateAt time.Time
	DateTo time.Time

	CellNameFunc func(t time.Time, n int) string
	CellDataFunc func(t time.Time, n int) string
}

var DefaultWeeekLabels = []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

func DefaultIsDayOffFunc(t time.Time) bool {
	return t.Weekday() == time.Sunday || t.Weekday() == time.Saturday
}

func NewCalendar() *Calendar {
	return &Calendar{
		WeekLabels: DefaultWeeekLabels,
		// IsDayOffFunc: DefaultIsDayOffFunc,
		DateAt: time.Now(),
		DateTo: time.Now().AddDate(0, 1, 0),
	}
}

func (c *Calendar) Html() string {
	wc := len(c.WeekLabels)
	wd := int(c.DateAt.Weekday())

	s := "<table class='calendar'>\n"
	s += " <tr><th>" + strings.Join(c.WeekLabels, "</th><th>") + "</th></tr>\n"
	s += " <tr>" + strings.Repeat("<td></td>", wd)

	last := int(c.DateTo.Sub(c.DateAt).Hours()/24) + 1
	for d := 0; d <= last; d++ {
		t := c.DateAt.AddDate(0, 0, d)
		name := fmt.Sprint(t.Day())
		data := ""

		if c.CellNameFunc != nil {
			name = c.CellNameFunc(t, d)
		}

		if c.CellDataFunc != nil {
			data = c.CellDataFunc(t, d)
		}

		s += fmt.Sprintf("<td><span class='calendar-cell-name'>%s</span><span class='calendar-cell-data'>%s</span></td>", name, data)

		wd = (wd + 1) % wc
		if wd == 0 {
			s += "</tr>\n"
			if d != last {
				s += " <tr>"
			}
		}
	}
	s += strings.Repeat("<td></td>", wc-wd) + "<tr>\n"

	s += "</table>\n"
	return s
}
