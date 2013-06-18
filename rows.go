package cubrid

/*
#include "cas_cci.h"
char* ex_cci_get_result_info_name(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_NAME(res_info, index);
}
*/
import "C"
import (
	"database/sql/driver"
	"fmt"
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
	return nil
}

