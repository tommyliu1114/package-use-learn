package local_utils

import (
	"fmt"
	"time"
)

func Logf(vals ...interface{}) {
	outs := make([]interface{}, 0, len(vals)+2)
	outs = append(outs, fmt.Sprintf("[%s] [%s] ", time.Now().String(), GCFG.Server.Id))
	outs = append(outs, vals...)
	fmt.Println(outs...)
}
