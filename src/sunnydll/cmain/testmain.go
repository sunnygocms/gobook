package main

/*
#include <stdio.h>
#include "libsunny.h"
#cgo linux CFLAGS: -L./ -I./
#cgo linux LDFLAGS: -L./ -I./ -lsunny
*/
import "C"

import (
	"fmt"
)

func main() {

	str := C.Hello()
	C.Test()
	fmt.Println(str)
}
