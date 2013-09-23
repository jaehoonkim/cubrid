package cubrid

import (
	"database/sql/driver"
	"fmt"
)

type cubridConn struct {
	con int
}


func (c *cubridConn) Prepare(query string) (driver.Stmt, error) {
	var req int
	var err GCI_ERROR
	stmt := &cubridStmt { c: c }
	req, err = Gci_prepare(c.con, query, 0)
	if req  < 0 {
		c.Close()
		return nil, fmt.Errorf("error : %d, %s", err.Code, err.Msg)
	}
	stmt.req = req
	return stmt, nil
}

func (c *cubridConn) Close() error {
	var err GCI_ERROR
	var err_no int
	err_no, err = Gci_disconnect(c.con)
	if err_no < 0 {
		return fmt.Errorf("error: %d, %s", err.Code, err.Msg)
	}
	return nil
}

/*
	cubrid는 기본으로 auto commit 모드가 켜져 있음
	이를 off하면 transaction을 사용하는걸로,,
*/
func (c *cubridConn) Begin() (driver.Tx, error) {
	var err int
	err = Gci_set_autocommit(c.con, AUTOCOMMIT_FALSE)
	if err == 0 {
		tx := &cubridTx{ c: c }
		return tx, nil
	}
	return nil, fmt.Errorf("cci_set_autocommit err : %d", err)
}

func (c *cubridConn) Exec(query string, args []driver.Value) (driver.Result, error) {
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

