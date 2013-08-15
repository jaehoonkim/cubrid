package cubrid

import (
	"database/sql/driver"
	"fmt"
	"log"
	"unsafe"
	"time"
)

type cubridStmt struct {
	c *cubridConn
	req int
}

func (s *cubridStmt) Close() error {
	//log.Println("cubridStmt:Close")
	err := gci_close_req_handle(s.req)
	if err == 0 {
		return nil
	}
	return fmt.Errorf("cci_close_req_handle err : %d", err)
}

func (s *cubridStmt) NumInput() int {
	//log.Println("cubridStmt:NumInput")
	param_cnt := gci_get_bind_num(s.req)
	if param_cnt < 0 {
		fmt.Errorf("cci_get_bind_num err : %d", param_cnt)
	}
	//log.Printf("numInput:%d\n", int(param_cnt))
	return param_cnt
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
	var res int
	var cci_error CCI_ERROR
	if args != nil {
		err := s.bindParam(args)
		if err != nil {
			return err
		}
	}
	res, cci_error = gci_execute(s.req, 0, 0)
	if err < 0 {
		return fmt.Errorf("cci_execute err: %d, %s", cci_error.err_code, cci_error.err_msg)
	}

	return nil
}

func (s *cubridStmt) bindParam(args []driver.Value) error {
	//log.Println("cbubridStmt:bindParam")
	var ss string
	var err C.int
	for i, arg := range args {
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
			c_param := C.float(arg.(float64))
			err = C.cci_bind_param(s.req, C.int(i + 1), C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&c_param), C.CCI_U_TYPE_FLOAT, C.CCI_BIND_PTR)
			if int(err) < 0 {
				return fmt.Errorf("cci_bind_param : %d", int(err))
			}
		case bool:
			log.Println("statement:bindParam:bool")
			return fmt.Errorf("not supported type")
		case []byte:
			log.Println("statement:bindParam:[]byte")
		}
		//log.Printf("cubridStmt:bindParam:idx : %d, arg:%s", i, arg.(string))
	}
	return nil
}
