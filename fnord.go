package DDate
import (
	"time"
)

const (
	YOLD = 1166
	SL = 73
)

type DTime struct{
	Year int
	Season DSeason
	Day int
	DayN DDay
	Hour int
	Minute int
	Second int
}

type DSeason int
const (
	Chaos DSeason = iota
	Discord
	Confusion
	Bureaucracy
	Aftermath
)
var seasons = [...]string{
	"Chaos",
	"Discord",
	"Confusion",
	"Bureaucracy",
	"The Aftermath"}

func(s DSeason) String() string{
	return seasons[s]
}

type DDay int
const (
	Sweetmorn DDay = iota
	BoomTime
	Pungenday
	Prickle
	SettingOrange
)
var days = [...]string{
	"Sweetmorn",
	"BoomTime",
	"Pungenday",
	"Prickle-Prickle",
	"Setting Orange"}

func (d DDay) String() string{
	return days[d]
}

func TimeToDTime(t time.Time) (f DTime){
	f.Year = t.Year() + YOLD
	d := t.YearDay()
	f.Season = DSeason(d/SL)
	f.Day = d%SL
	f.DayN = DDay( (f.Day%5)-1)
	f.Hour= t.Hour()
	f.Minute= t.Minute()
	f.Second= t.Second()
	return
}
