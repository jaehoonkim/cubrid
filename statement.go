package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
	"log"
)

type cubridStmt struct {
	c *cubridConn
	req C.int
}

func (s *cubridStmt) Close() error {
	log.Println("cubridStmt:Close")
	var err C.int
	err = C.cci_close_req_handle(s.req)
	if err == 0 {
		return nil
	}
	return fmt.Errorf("cci_close_req_handle err : %d", err)
}

func (s *cubridStmt) NumInput() int {
	var param_cnt C.int
	param_cnt = C.cci_get_bind_num(s.req)
	if param_cnt < 0 {
		fmt.Errorf("cci_get_bind_num err : %d", param_cnt)
	}
	return int(param_cnt)
}

func (s *cubridStmt) Exec(args []driver.Value) (driver.Result, error) {
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	result := &cubridResult{ c: s.c }
	return result, nil
}

func (s *cubridStmt) Query(args []driver.Value) (driver.Rows, error) {
	log.Println("cubridStmt : Query")
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	rows := &cubridRows{ s: s }
	return rows, nil
}

func (s *cubridStmt) execute(args []driver.Value) (error) {
	var err C.int
	var cci_error C.T_CCI_ERROR
	if args != nil {
		err := s.bindParam(args)
		if err != nil {
			return err
		}
	}
	err = C.cci_execute(s.req, 0, 0, &cci_error)
	if int(err) < 0 {
		return fmt.Errorf("cci_execute err: %d, %s", cci_error.err_code, cci_error.err_msg)
	}

	return nil
}

func (s *cubridStmt) bindParam(args []driver.Value) error {
	return nil
}
