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
	Holyday string
}
var holydays = [...]DTime{
		DTime{0,Chaos,5,0,0,0,0,"Mungday"},
		DTime{0,Chaos,50,0,0,0,0,"Chaoflux"},
		DTime{0,Chaos,60,0,0,0,0,"St. Tib's Day"},
		DTime{0,Discord,5,0,0,0,0,"Mojoday"},
		DTime{0,Discord,50,0,0,0,0,"Discoflux"},
		DTime{0,Confusion,5,0,0,0,0,"Syaday"},
		DTime{0,Confusion,50,0,0,0,0,"Confuflux"},
		DTime{0,Bureaucracy,5,0,0,0,0,"Zaraday"},
		DTime{0,Bureaucracy,50,0,0,0,0,"Bureflux"},
		DTime{0,Aftermath,5,0,0,0,0,"Maladay"},
		DTime{0,Aftermath,50,0,0,0,0,"Afflux"}	}

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
	if f.Day % 5 > 0 {
		f.DayN = DDay( (f.Day%5)-1)
	}else{
		f.DayN = DDay( (f.Day%5))
	}
	f.Hour= t.Hour()
	f.Minute= t.Minute()
	f.Second= t.Second()
	for _,v := range holydays{
		if v.Season == f.Season && f.Day == v.Day{
			if v.Day == 60 {
				if t.Month() != 2{
						continue
				}
			}
			f.Holyday = v.Holyday
		}
	}
	return
}
