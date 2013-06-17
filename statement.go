package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
)

type cubridStmt struct {
	c *cubridConn
	req C.int
}

func (s *cubridStmt) Close() error {
	var err C.int
	err = C.cci_close_req_handle(s.req)
	if err == 0 {
		return nil
	}
	return fmt.Errorf("cci_close_req_handle err : %d", err)
}

func (s *cubridStmt) NumInput() int {
	var param_cnt C.int
	param_cnt = C.cci_get_bind_num(s.c.con)
	if param_cnt < 0 {
		fmt.Errorf("cci_get_bind_num err : %d", param_cnt)
	}
	return int(param_cnt)
}

func (s *cubridStmt) Exec(args []driver.Value) (driver.Result, error) {
	err := s.exec(args)
	if err != nil {
		return nil, err
	}
	result := &cubridResult{ c: s.c }
	return result, nil
}

func (s *cubridStmt) Query(args []driver.Value) (driver.Rows, error) {
	err := s.exec(args)
	if err != nil {
		return nil, err
	}
	rows := &cubridRows{ s: s }
	return rows, nil
}

func (s *cubridStmt) exec(args []driver.Value) (error) {
	return nil
}

