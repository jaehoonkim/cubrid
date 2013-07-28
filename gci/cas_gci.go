package gci

/*
#cgo CFLAGS: -I../CUBRID/include
#cgo LDFLAGS: -L../CUBRID/lib -lcascci -lnsl
#include "cas_cci.h"
*/
import "C"

func gci_init() {
	C.cci_init()
}

func gci_end() {
	C.cci_end()
}
/*
func gci_get_version_string(string &str) int {
}
*/

