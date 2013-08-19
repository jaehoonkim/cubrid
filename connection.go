package cubrid

import (
	"database/sql/driver"
	"fmt"
	"log"
)

type cubridConn struct {
	con int
}


func (c *cubridConn) Prepare(query string) (driver.Stmt, error) {
	var req int
	var cci_error CCI_ERROR
	stmt := &cubridStmt { c: c }
	req, cci_error = gci_prepare(c.con, query, 0)
	if req  < 0 {
		c.Close()
		return nil, fmt.Errorf("error : %d, %s", cci_error.err_code, cci_error.err_msg)
	}
	stmt.req = req
	return stmt, nil
}

func (c *cubridConn) Close() error {
	//log.Println("cubridConn:Close")
	var cci_error CCI_ERROR
	var err_no int
	err_no, cci_error = gci_disconnect(c.con)
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
	var err int
	err = gci_set_autocommit(c.con, AUTOCOMMIT_FALSE)
	if err == 0 {
		tx := &cubridTx{ c: c }
		return tx, nil
	}
	return nil, fmt.Errorf("cci_set_autocommit err : %d", err)
}

func (c *cubridConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	log.Println("cubridConn:Exec")
	stmt, err := c.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	result, err := stmt.Exec(args)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *cubridConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	log.Println("cubridConn:Query")
	stmt, err := c.Prepare(query)
	defer stmt.Close()
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query(args)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	
	return rows, nil
}
