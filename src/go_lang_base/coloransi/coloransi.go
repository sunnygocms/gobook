package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {

	beeStrUp := `______
| ___ \
| |_/ /  ___   ___
| ___ \ / _ \ / _ \
| |_/ /|  __/|  __/
\____/  \___| \___| v1.5.0
`
	beeStrDown := `
├── Beego     : 1.7.0
├── GoVersion : go1.6.2
├── GOOS      : windows
├── GOARCH    : amd64
├── NumCPU    : 8
├── GOPATH    : x:\Eagle\go
├── GOROOT    : y:\Go\
├── Compiler  : gc
└── Date      : Friday, 19 Aug 2016`
	color.Set(color.FgMagenta, color.Bold)
	defer color.Unset()
	fmt.Println(beeStrUp)
	color.Set(color.FgGreen, color.Bold)
	fmt.Println(beeStrDown)
}
