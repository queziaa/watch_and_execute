package execute

import (
	"fmt"
)

func InitExecute(command string, timeChan chan uint) {
	var otv uint = 0
	for {
		tv := <-timeChan
		if (tv - otv) < 400000 {
			fmt.Print(2)
		}
		otv = tv

		fmt.Print(command, tv)
	}

}
