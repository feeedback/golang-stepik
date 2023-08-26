package main // https://stepik.org/lesson/359395/step/8

import (
	"fmt"
	"time"
)

func main() {
	var mins, secs int
	fmt.Scanf("%d мин. %d сек.", &mins, &secs)

	periodDuration := time.Duration(mins)*time.Minute + time.Duration(secs)*time.Second

	const inputUnixTime int64 = 1589570165
	outputTimeWithAddedPeriod := time.Unix(inputUnixTime, 0).UTC().Add(periodDuration)

	fmt.Println(outputTimeWithAddedPeriod.Format(time.UnixDate))
}
