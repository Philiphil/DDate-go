// Package DDate provides Discordian Dates for go
package DDate
import (
	"time"
)
//YOLD stands for Year of Our Lady of Discord, which is when greyface's cuses began.
//SL Stands for (Discordian) Season length, which is 73 days fnord
const (
	YOLD = 1166
	SL = 73
)

//DTime is Discordian time
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
//Discordian Holydays
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

//Discordian actually have 5 seasons of 73 days
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
// A discordian week length 5 discordian days
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

//Converts time.Time to DDate.DTime
// input : time.time
// Ouput : DDate.DTime
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

//Converts unix timestamp to DDate.DTime
// input : int64
// Ouput : DDate.DTime
func UnixToDTime(i int64) DTime{
	return TimeToDTime( time.Unix(i,0) )
}

//Converts DDate.DTime to time.Time
// input : DDate.Dtime
// Ouput : time.time
func DTimeToTime(d DTime) (t time.Time){
	tmp,_ := time.LoadLocation("UTC")
	y := time.Date( d.Year - YOLD, 1,0,0,0,0,0, tmp)

	n := int( (int(d.Season * SL)) + int(d.Day) )
	h := d.Hour
	m := d.Minute
	s := d.Second
	var u int64 
	u = int64((n*24*60*60*1000*1000*1000) + (h*60*60*1000*1000*1000) + (m*60*1000*1000*1000)+ (s*1000*1000*1000))
	return y.Add(time.Duration(u))
}
