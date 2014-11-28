package tools

import (
	"fmt"
)

func E(desc string, err interface{}) {
	switch fault := err.(type) {
	case error:
		{
			if fault != nil {
				panic(fmt.Errorf(desc+": %v", err))
			}
		}
	case bool:
		{
			if !fault {
				panic(fmt.Errorf(desc))
			}
		}
	}

}
