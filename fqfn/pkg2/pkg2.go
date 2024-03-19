package pkg2

import (
	"runtime"
)

func Func() string {
	pc, _, _, ok := runtime.Caller(0)
	if ok {
		fn := runtime.FuncForPC(pc)

		return fn.Name()
	} else {
		return "unknown"
	}
}
