package DDate
import (
	"time"
	"fmt"
)
var (
	YOLD = 1166
	SL = 73
)
type DTime struct{
	Year int
	Season DSeason
	Day DDay
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
type DDay int
const (
	Sweetmorn DDay = iota
	BoomTime
	Pugenday
	Prickle
	SettingOrange
)



func TimeToDTime(t time.Time) (f DTime){
	f.Year = t.Year() + YOLD
	d := t.YearDay()
	f.Season = DSeason(d/SL)
	f.Day = DDay(d%SL)
	f.Hour= t.Hour()
	f.Minute= t.Minute()
	f.Second= t.Second()
	return
}

func main(){
	fmt.Println(TimeToDTime(time.Now()))
}