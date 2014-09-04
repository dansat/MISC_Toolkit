/*
@author : dana satriya
*/
package xutil

import (
	"fmt"
	"runtime"
)

var x_mode int
var X_LOGMODE int

const (
	X_INFO = 0x01
	X_DBG  = 0x02
	X_ERR  = 0x03
	X_CRIT = 0x04
)

const (
	NORMAL        = 0x01
	DEBUG         = 0X02
	EXTREME_DEBUG = 0X03
)

/*
simple leveled logger
X_INFO : fmt.printf wrapper
X_DBG  : print function caller LINE NUMBER only
X_ERR  : print function caller FILENAME only
X_CRIT : print function caller FILENAME + LINENUMBER


logmode : if NORMAL, display only X_INFO. this is used for shipping logmode
     : if DEBUG, X_CRIT won't be displayed, display only X_INFO, X_DBG, X_ERR
     : if EXTREME_DEBUG, display ALL!

the logmode is NORMALLY initialized during init,
however a flexiblity of the caller to change the logmode on run time is also allowed

sample usage


LOG_XROUTER(X_INFO, "checking parser cisco")
LOG_XROUTER(X_CRIT, "checking parser juniper with router id %s ", 220)

*/
func LOG_XROUTER(log_type uint8, text string, param ...interface{}) {

	x_mode = X_LOGMODE
	text_info := "\033[1;32m" + text + "\033[0m"
	text_dbg := "\033[1;34m" + text + "\033[0m"
	text_err := "\033[1;35m" + text + "\033[0m"
	text_crit := "\033[1;31m" + text + "\033[0m"

	if x_mode < 0 || x_mode > 3 {
		fmt.Printf("XROUTER_FAIL: MODE IS NOT HANDLED! > %08x \n", x_mode)
	}

	_, filename, line, _ := runtime.Caller(1)

	switch log_type {
	case X_INFO:
		fmt.Printf("\033[1;32mXROUTER_INFO: \033[0m"+text_info, param...)
	case X_DBG:
		if x_mode == DEBUG || x_mode == EXTREME_DEBUG {
			dbg := fmt.Sprintf("\033[1;34m\033[0m"+text_dbg, param...)
			fmt.Printf("\033[1;34mXROUTER_DBG : IN LINE : [%d] \033[0m"+dbg, line)
		}
	case X_ERR:
		if x_mode == DEBUG || x_mode == EXTREME_DEBUG {
			err := fmt.Sprintf("\033[1;35m \033[0m"+"\t      "+text_err, param...)
			fmt.Printf("\033[1;35mXROUTER_ERR : IN FILE : [%s]\n\033[0m"+"  "+err, filename)
		}
	case X_CRIT:
		if x_mode == EXTREME_DEBUG {
			err := fmt.Sprintf("\033[1;31m \033[0m"+"\t      "+text_crit, param...)
			fmt.Printf("\033[1;31mXROUTER_CRIT: IN FILE : [%s] IN LINE : [%d] \033[0m\n"+"  "+err, filename, line)
		}
	default:
		fmt.Printf("XROUTER_FAIL: log type not handled! > %d \n", log_type)
		return
	}

}