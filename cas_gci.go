package cubrid

/*
#cgo CFLAGS: -I./CUBRID/include
#cgo LDFLAGS: -L./CUBRID/lib -lcascci -lnsl
#include <stdio.h>
#include <stdlib.h>
#include "cas_cci.h"
int ex_cci_connect(char *ip, int port, char *db_name, char *db_user, char *db_password) {
	int con = cci_connect(ip, port, db_name, db_user, db_password);
	return con;
}
*/
import "C"
import (
	"unsafe"
	"strconv"
)

func gci_init() {
	C.cci_init()
}

func gci_end() {
	C.cci_end()
}
/*
func gci_get_version_string(string &str) int {
}
*/

func gci_connect(ip string, port int, db_name string, db_user string, db_password string) int {
	serverAddress := C.CString(ip)
	serverPort := C.int(port)
	dbName := C.CString(db_name)
	dbUser := C.CString(db_user)
	dbPassword := C.CString(db_password)

	defer C.free(unsafe.Pointer(serverAddress))
	defer C.free(unsafe.Pointer(dbName))
	defer C.free(unsafe.Pointer(dbUser))
	defer C.free(unsafe.Pointer(dbPassword))

	con := C.ex_cci_connect(serverAddress, serverPort, dbName, dbUser, dbPassword)
	return int(con)

}

func gci_prepare(conn_handle int, sql_stmt string, flag byte) (int, CCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cQuery *C.char = C.CString(sql_stmt)
	var cci_error C.T_CCI_ERROR
	var req C.int
	var err CCI_ERROR

	defer C.free(unsafe.Pointer(cQuery))

	req = C.cci_prepare(cHandle, cQuery, 0, &cci_error)
	err.err_code = int(cci_error.err_code)
	err.err_msg = C.GoString(cci_error.err_msg)

	return int(req), err
}

func gci_disconnect(conn_handle int) (int, CCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err CCI_ERROR

	res = C.cci_disconnect(cHandle, &cci_error)
	err.err_code = int(cci_error.err_code)
	err.err_msg = C.GoString(cci_error.err_msg)

	return int(res), err
}

func gci_close_req_handle(req_handle int) int {
	var err C.int
	var handle C.int = C.int(req_handle)
	err = C.cci_close_req_handle(handle)

	return int(err)
}

func gci_get_bind_num(req_handle int) int {
	var param_cnt C.int
	var handle C.int = C.int(req_handle)
	param_cnt = C.cci_get_bind_num(handle)

	return int(param_cnt)
}

func gci_execute(req_handle int, flag int, max_col_size int) (int, CCI_ERROR) {
	var res C.int
	var cci_error C.T_CCI_ERROR
	var handle C.int = C.int(req_handle)
	var err CCI_ERROR

	res = C.cci_execute(handle, flag, max_col_size, &cci_error)
	err.err_code = int(cci_error.err_code)
	err.err_msg = C.GoString(cci_error.err_msg)

	return int(res), err
}

func gci_set_autocommit(conn_handle int autocommit_mode AUTOCOMMIT_MODE) int {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var mode C.int = C.int(autocommit_mode)

	res = C.cci_set_autocommit(handle, mode)
	return int(res)
}

func gci_end_tran(conn_handle int, tran_type int) (int, CCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err CCI_ERROR

	res = C.cci_end_tran(handle, tran_type, &cci_error)
	err.err_code = int(cci_error.err_code)
	err.err_msg = C.GoString(cci_error.err_msg)

	return int(res), err
}

func gci_get_last_insert_id(conn_handle int) (int64, CCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err CCI_ERROR
	var value *C.char
	var nid int64

	res = C.cci_get_last_insert_id(handle, unsafe.Pointer(value), &cci_error)
	err.err_code = int(cci_error.err_code)
	err.err_msg = C.GoString(cci_error.err_msg)
	if res < 0 {
		return int64(res), err
	}

	id := C.GoString(value)
	nid, _ = strconv.ParseInt(id, 0, 64)
	return nid, err
}

func gci_row_count(conn_handle int) (int64, CCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var row_count C.int
	var cci_error C.T_CCI_ERROR
	var err CCI_ERROR

	res = C.cci_row_count(handle, &row_count, &cci_error)
	if res < 0 {
		err.err_code = int(cci_error.err_code)
		err.err_msg = C.GoString(cci_error.err_msg)
	}

	return int64(row_count), err
}
