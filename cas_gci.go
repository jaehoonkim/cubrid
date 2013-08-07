package cubrid

/*
#cgo CFLAGS: -I./CUBRID/include
#cgo LDFLAGS: -L./CUBRID/lib -lcascci -lnsl
#include <stdio.h>
#include <stdlib.h>
#include "cas_cci.h"
int ex_cci_connect(char *ip, int port, char *db_name, char *db_user, char *db_password) {
	int con = cci_connect(ip, port, db_name, db_user, db_password);
	return con;
}
*/
import "C"
import (
	"unsafe"
)

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

func gci_connect(ip string, port int, db_name string, db_user string, db_password string) int {
	serverAddress := C.CString(ip)
	serverPort := C.int(port)
	dbName := C.CString(db_name)
	dbUser := C.CString(db_user)
	dbPassword := C.CString(db_password)

	defer C.free(unsafe.Pointer(serverAddress))
	defer C.free(unsafe.Pointer(dbName))
	defer C.free(unsafe.Pointer(dbUser))
	defer C.free(unsafe.Pointer(dbPassword))

	con := C.ex_cci_connect(serverAddress, serverPort, dbName, dbUser, dbPassword)
	return int(con)

}
