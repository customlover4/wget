package utils

import (
	"fmt"
	"strconv"
	"time"
)

type Len interface {
	Len() int
}

func DownloadProgress(buf Len, cLen int) {
	percent := cLen / 100
	for v := buf.Len(); v != cLen; v = buf.Len() {
		s := "\rprocess: ["
		l := v / percent
		tmp := ""
		for i := 0; i < l/5; i++ {
			tmp += "#"
		}
		tmp += "=>"
		for i := 0; i < 20-(l/5+2); i++ {
			tmp += "."
		}
		s += tmp + "] " + strconv.Itoa(l) + "%"
		fmt.Print(s)
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Print("\rprocess: [####################] 100%")
	fmt.Print("\n\n")
}
