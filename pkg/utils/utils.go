package utils

import (
	"fmt"
	"time"
)

func DisplayInfo(start, finish *time.Time) {
	fmt.Println("------------------------------------------------------------")
	fmt.Printf("Start\t\t\t: %s\n", (*start).Format("2006-01-02 15:04:05 Z0700 MST"))
	fmt.Printf("Finish\t\t\t: %s\n", (*finish).Format("2006-01-02 15:04:05 Z0700 MST"))
	fmt.Printf("Response time\t\t: %dms \n", (*finish).Sub(*start).Abs().Milliseconds())
}
