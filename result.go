package cubrid

/*
#include "cas_cci.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
	"strconv"
)

type cubridResult struct {
	c *cubridConn
}


func (result *cubridResult) LastInsertId() (int64, error) {
	var err C.int
	var conn_handle C.int
	conn_handle = result.c.con
	var value *C.char
	var cci_error C.T_CCI_ERROR
	err = C.cci_get_last_insert_id(conn_handle, unsafe.Pointer(value), &cci_error)
	if err < 0 {
		return 0, fmt.Errorf("cci_get_last_insert_id err: %d", err)
	}
	id := C.GoString(value)
	res, _ := strconv.ParseInt(id, 0, 64)
	return res, nil
}

func (result *cubridResult) RowsAffected() (int64, error) {
	var err C.int
	var conn_handle C.int
	conn_handle = result.c.con
	var row_count C.int
	var cci_error C.T_CCI_ERROR
	err = C.cci_row_count(conn_handle, &row_count, &cci_error)
	if err < 0 {
		return 0, fmt.Errorf("cci_row_count err: %d", cci_error.err_code)
	}
	return int64(row_count), nil
}

