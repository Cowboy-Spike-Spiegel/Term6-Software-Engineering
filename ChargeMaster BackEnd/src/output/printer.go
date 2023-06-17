package output

import (
	"demo/src/consts"
	"fmt"
	"os"
	"strings"
	"time"
)

func Printer(layer string, info string) {
	cur := strings.Split(time.Now().Format("2006-01-02 15:04:05"), " ")
	for len(layer) < consts.LayerLength {
		layer = layer + " "
	}
	info = "[" + cur[0] + " " + cur[1] + "] ---- " + layer + "\t> " + info

	if os.Args[1] == consts.ModeTest {
		Log_insert(info)
		fmt.Println("\t" + info)
	} else if os.Args[1] == consts.ModeRelease {
		if strings.Contains(layer, "Dao") != true {
			fmt.Println("\t" + info)
		}
	}
}
