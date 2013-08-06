/*
CUBRID Database driver for Go
*/
package cubrid

/*
#cgo CFLAGS: -I./CUBRID/include
#cgo LDFLAGS: -L./CUBRID/lib -lcascci -lnsl
#include "cas_cci.h"
*/
import "C"
import (
	"database/sql"
	"database/sql/driver"
	"unsafe"
	"strings"
	"errors"
	"strconv"
	"fmt"
)

type cubridDriver struct {}


/*
	name : ip/port/db_name/db_user/db_password
*/
func (d *cubridDriver) Open(name string) (driver.Conn, error) {
	opt := strings.SplitN(name, "/", 5)
	if(len(opt) != 5) {
		return nil, errors.New("error options")
	}

	port, _ := strconv.Atoi(opt[1])

	serverAddress := C.CString(opt[0])
	serverPort := C.int(port)
	dbName := C.CString(opt[2])
	dbUser := C.CString(opt[3])
	dbPassword := C.CString(opt[4])

	defer C.free(unsafe.Pointer(serverAddress))
	defer C.free(unsafe.Pointer(dbName))
	defer C.free(unsafe.Pointer(dbUser))
	defer C.free(unsafe.Pointer(dbPassword))
	con := C.CCI_CONNECT_INTERNAL_FUNC_NAME(serverAddress, serverPort, dbName, dbUser, dbPassword)
	if int(con) < 0 {
		fmt.Printf("connect error code : %d", int(con))
		return nil, errors.New("cannot connect to database")
	}
	conn := &cubridConn{ con: con }
	return conn, nil
}
func init() {
	sql.Register("cubrid", &cubridDriver{})
}

