package domain

import "time"

type Apod struct {
	Date           time.Time
	Explanation    string
	Image          string
	ServiceVersion string
	Title          string
}
