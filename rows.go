package cubrid

/*
#include "cas_cci.h"
char* ex_cci_get_result_info_name(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_NAME(res_info, index);
}

T_CCI_U_TYPE ex_cci_get_result_info_type(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_TYPE(res_info, index);
}
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
	"io"
	"unsafe"
)

type cubridRows struct {
	s *cubridStmt
}

func (rows *cubridRows) Columns() []string {
	var col_info *C.T_CCI_COL_INFO
	var stmt_type C.T_CCI_CUBRID_STMT
	var col_count, idx  C.int
	col_info = C.cci_get_result_info(rows.s.req, &stmt_type, &col_count)
	if col_info == nil {
		return nil
	}

	col_name  := make([]string, col_count)
	for idx = C.int(0); idx < col_count; idx++ {
		col_name[idx] = C.GoString(C.ex_cci_get_result_info_name(col_info, idx))
	}
	return col_name
}

func (rows *cubridRows) Close() error {
	var err C.int
	err = C.cci_close_req_handle(rows.s.req)
	if int(err) < 0 {
		return fmt.Errorf("close_req_handle err : %d", int(err))
	}
	return nil
}

func (rows *cubridRows) Next(dest []driver.Value) error {
	var err C.int
	var cci_error C.T_CCI_ERROR
	var col_info *C.T_CCI_COL_INFO
	var stmt_type C.T_CCI_CUBRID_STMT
	var col_count C.int

	err = C.cci_cursor(rows.s.req, 1, C.CCI_CURSOR_CURRENT, &cci_error)
	if err == C.CCI_ER_NO_MORE_DATA {
		return io.EOF
	}
	if int(err) < 0 {
		return fmt.Errorf("cursor err: %d, %s", cci_error.err_code, cci_error.err_msg)
	}
	
	err = C.cci_fetch(rows.s.req, &cci_error)
	if int(err) < 0 {
		return fmt.Errorf("fetch err: %d, %s", cci_error.err_code, cci_error.err_msg)
	}
	
	col_info = C.cci_get_result_info(rows.s.req, &stmt_type, &col_count)
	var columnType C.T_CCI_U_TYPE
	var i C.int
	var value C.void
	var ind C.int
	for i = C.int(0); i < col_count; i++ {
		columnType = C.ex_cci_get_result_info_type(col_info, i)
		switch columnType {
		case C.CCI_U_TYPE_CHAR:
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_STR, unsafe.Pointer(&value), &ind)
		case C.CCI_U_TYPE_INT:
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_INT, unsafe.Pointer(&value), &ind)
 
		}
	}
	return nil
}
