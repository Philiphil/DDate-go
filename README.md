# DDate-go

Discordian Dates for Go(lang)

## FNORD
```go get github.com/Philiphil/Ddate-go```


## FNORD?
DDate comes with 5 structs (actually 3 but i'm still working on the missing ones)

DTime : Represents Discordian Time  (with standard POEE holydays)
DSeason : for discordian seasons (0-4)  
DDay : Discordian Day (0-4)  

function timeToDTime(t time) returns t converted in DTime
function UnixToDTime(i int64) returns i converted in DTime
function DTimeToTime(d Dtime) returns d converted in Time
