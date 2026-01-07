// Code written by github.com/ryznxx
package target

import (
	"bootstrap-open-academia/target/bsd"
	"bootstrap-open-academia/target/linux"
	"bootstrap-open-academia/target/windows"
	"bootstrap-open-academia/utils"
	"fmt"
)

// supported soon
func Windows() {
	windows.BuildInstaller()
}

func Linux() bool {
	fmt.Print("\n========== Install On Linux ==========")
	fmt.Print("\n\n[Checking]")
	checkRuntime := utils.RequiredRuntime()
	checkTools := utils.RequiredTools()

	if checkRuntime && checkTools {
		linux.BuildInstaller()
		return true
	}
	return false
}

// supported soon
func Bsd() {
	bsd.BuildInstaller()
}
