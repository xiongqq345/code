package main

import "fmt"
import "time"

type LogMessage struct {
	Date    int64
	Source  string
	Message string
}

type SyslogMessage struct {
	Server string
	*LogMessage
}

//<start id="declaration"/>
func (m *LogMessage) TimeAgoInSeconds() int64 {
	//<end id="declaration"/>
	var seconds int64
	seconds = time.Seconds() - m.Date
	return seconds
}

func main() {

	lm := &LogMessage{Date: time.Seconds(),
		Source:  "GoCron",
		Message: "GoCron failed to create a task"}
	sm := &SyslogMessage{Server: "192.168.2.6",
		LogMessage: lm}
	time.Sleep(5000e6)
	fmt.Println("It happended ",
		sm.LogMessage.TimeAgoInSeconds(),
		"seconds ago")

}
