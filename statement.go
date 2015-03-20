package cubrid

import (
	"database/sql/driver"
	"fmt"
	"log"
	"time"

	"github.com/jaehoonkim/cubrid/gci"
)

type cubridStmt struct {
	c   *cubridConn
	req int
}

func (s *cubridStmt) Close() error {
	err := gci.Close_req_handle(s.req)
	if err == 0 {
		return nil
	}
	return fmt.Errorf("cci_close_req_handle err : %d", err)
}

func (s *cubridStmt) NumInput() int {
	param_cnt := gci.Get_bind_num(s.req)
	if param_cnt < 0 {
		fmt.Errorf("cci_get_bind_num err : %d", param_cnt)
	}
	return param_cnt
}

func (s *cubridStmt) Exec(args []driver.Value) (driver.Result, error) {
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	result := &cubridResult{c: s.c}
	return result, nil
}

func (s *cubridStmt) Query(args []driver.Value) (driver.Rows, error) {
	err := s.execute(args)
	if err != nil {
		return nil, err
	}
	rows := &cubridRows{s: s}
	return rows, nil
}

func (s *cubridStmt) execute(args []driver.Value) error {
	var gciError gci.GCI_ERROR
	if args != nil {
		err := s.bindParam(args)
		if err != nil {
			return err
		}
	}
	_, gciError = gci.Execute(s.req, 0, 0)
	if gciError.Code < 0 {
		return fmt.Errorf("cci_execute err: %d, %s", gciError.Code, gciError.Msg)
	}

	return nil
}

/*
- cubrid cci 문서의 내용
  prepared statement에서 bind변수에 데이터를 바인딩하기 위하여 사용되는 함수이다.
  이때, 주어진 a_type의 value의 값을 실제 바인딩되어야 하는 타입으로 변환하여 저장한다.
  이후, cci_execute()가 호출될 때 저장된 데이터가 서버로 전송된다.:
*/
func (s *cubridStmt) bindParam(args []driver.Value) error {
	var res int
	for i, arg := range args {
		switch arg.(type) {
		case int64:
			res = gci.Bind_param_int(s.req, i+1, arg, gci.GCI_BIND_PTR)
			if res < 0 {
				return fmt.Errorf("gci_bind_param : %d", res)
			}
		case string:
			res = gci.Bind_param_string(s.req, i+1, arg, gci.GCI_BIND_PTR)
			if res < 0 {
				return fmt.Errorf("gci_bind_param : %d", res)
			}
		case time.Time:
			log.Println("statement:bindParam:time.Tile")
		case float64:
			res = gci.Bind_param_float(s.req, i+1, arg, gci.GCI_BIND_PTR)
			if res < 0 {
				return fmt.Errorf("gci_bind_param : %d", res)
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
