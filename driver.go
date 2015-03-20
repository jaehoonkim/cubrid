/*
CUBRID Database driver for Go
*/
package cubrid

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"

	"github.com/jaehoonkim/cubrid/gci"
)

type cubridDriver struct{}

/*
	name : ip/port/db_name/db_user/db_password
*/
func (d *cubridDriver) Open(name string) (driver.Conn, error) {
	opt := strings.SplitN(name, "/", 5)
	if len(opt) != 5 {
		return nil, fmt.Errorf("error options : %d", len(opt))
	}
	port, _ := strconv.Atoi(opt[1])
	con := gci.Connect(opt[0], port, opt[2], opt[3], opt[4])

	if con < 0 {
		fmt.Printf("connect error code : %d", con)
		return nil, fmt.Errorf("cannot connect to database : %d", con)
	}
	conn := &cubridConn{con: con}
	return conn, nil
}

func init() {
	sql.Register("cubrid", &cubridDriver{})
}
