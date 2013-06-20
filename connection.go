package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
	"unsafe"
	"log"
)

type cubridConn struct {
	con C.int
}


func (c *cubridConn) Prepare(query string) (driver.Stmt, error) {
	//log.Println("cubridConn:Prepare")
	var cQuery *C.char = C.CString(query)
	defer C.free(unsafe.Pointer(cQuery))
	var req C.int
	var cci_error C.T_CCI_ERROR
	stmt := &cubridStmt { c: c }
	req = C.cci_prepare(c.con, cQuery, 0, &cci_error)
	if req  < 0 {
		c.Close()
		return nil, fmt.Errorf("error : %d, %s", cci_error.err_code, cci_error.err_msg)
	}
	stmt.req = req
	return stmt, nil
}

func (c *cubridConn) Close() error {
	//log.Println("cubridConn:Close")
	var cci_error C.T_CCI_ERROR
	var err_no C.int
	err_no = C.cci_disconnect(c.con, &cci_error)
	if err_no < 0 {
		return fmt.Errorf("error: %d, %s", cci_error.err_code, cci_error.err_msg)
	}
	return nil
}

/*
	cubrid는 기본으로 auto commit 모드가 켜져 있음
	이를 off하면 transaction을 사용하는걸로,,
*/
func (c *cubridConn) Begin() (driver.Tx, error) {
	//log.Println("cubridConn:Begin")
	var con C.int
	var err C.int
	con = c.con
	err = C.cci_set_autocommit(con, C.CCI_AUTOCOMMIT_FALSE)
	if err == 0 {
		tx := &cubridTx{ c: c }
		return tx, nil
	}
	return nil, fmt.Errorf("cci_set_autocommit err : %d", err)
}

func (c *cubridConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	log.Println("cubridConn:Exec")
	return nil, nil
}

func (c *cubridConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	log.Println("cubridConn:Query")
	return nil, nil
}
