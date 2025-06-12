package wrap

import (
	"fmt"
	"runtime"
	"strings"
)

func Wrap(err error) error {
	pc, _, _, ok := runtime.Caller(1)

	funcName := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			full := fn.Name()
			parts := strings.Split(full, ".")
			funcName = parts[len(parts)-1]
		}
	}

	return fmt.Errorf("%s: %w", funcName, err)
}
