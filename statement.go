package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
	"log"
	"unsafe"
	"time"
)

type cubridStmt struct {
	c *cubridConn
	req C.int
}

func (s *cubridStmt) Close() error {
	//log.Println("cubridStmt:Close")
	var err C.int
	err = C.cci_close_req_handle(s.req)
	if err == 0 {
		return nil
	}
	return fmt.Errorf("cci_close_req_handle err : %d", err)
}

func (s *cubridStmt) NumInput() int {
	//log.Println("cubridStmt:NumInput")
	var param_cnt C.int
	param_cnt = C.cci_get_bind_num(s.req)
	if param_cnt < 0 {
		fmt.Errorf("cci_get_bind_num err : %d", param_cnt)
	}
	//log.Printf("numInput:%d\n", int(param_cnt))
	return int(param_cnt)
}

func (s *cubridStmt) Exec(args []driver.Value) (driver.Result, error) {
	//log.Println("cubridStmt:Exec")
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	result := &cubridResult{ c: s.c }
	return result, nil
}

func (s *cubridStmt) Query(args []driver.Value) (driver.Rows, error) {
	//log.Println("cubridStmt : Query")
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	rows := &cubridRows{ s: s }
	return rows, nil
}

func (s *cubridStmt) execute(args []driver.Value) (error) {
	//log.Println("cubridStmt:execute")
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
	//log.Println("cbubridStmt:bindParam")
	var ss string
	var err C.int
	for i, arg := range args {
		//var c_arg *C.char
		//c_arg = C.CString(arg)
		switch arg.(type) {
		case int64:
			c_param := C.int(arg.(int64))
			err = C.cci_bind_param(s.req, C.int(i + 1), C.CCI_A_TYPE_INT, unsafe.Pointer(&c_param), C.CCI_U_TYPE_INT, C.CCI_BIND_PTR)
			if int(err) < 0 {
				return fmt.Errorf("cci_bind_param : %d", int(err))
			}
		case string:
			ss = fmt.Sprint(arg)
			//log.Printf("cubridStmt:bindParam:%s\n", ss)
			err = C.cci_bind_param(s.req, C.int(i+1), C.CCI_A_TYPE_STR, unsafe.Pointer(C.CString(ss)), C.CCI_U_TYPE_STRING, C.CCI_BIND_PTR)
			if err < 0 {
				return fmt.Errorf("cci_bind_param : %d", err)
			}
		case time.Time:
			log.Println("statement:bindParam:time.Tile")
		case float64:
			log.Println("statement:bindParam:float64")
		case bool:
			log.Println("statement:bindParam:bool")
		case []byte:
			log.Println("statement:bindParam:[]byte")
		}
		//log.Printf("cubridStmt:bindParam:idx : %d, arg:%s", i, arg.(string))
	}
	return nil
}
