package cubrid

/*
#cgo CFLAGS: -I./CUBRID/include
#cgo LDFLAGS: -L./CUBRID/lib -lcascci -lnsl
#include <stdio.h>
#include <stdlib.h>
#include "cas_cci.h"
#include "cas_error.h"
int ex_cci_connect(char *ip, int port, char *db_name, char *db_user, char *db_password) {
	int con = cci_connect(ip, port, db_name, db_user, db_password);
	return con;
}

char* ex_cci_get_result_info_name(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_NAME(res_info, index);
}

T_CCI_U_TYPE ex_cci_get_result_info_type(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_TYPE(res_info, index);
}

int ex_cci_is_set_type(T_CCI_U_TYPE type) {
	return CCI_IS_SET_TYPE(type);
}

int ex_cci_is_collection_type(T_CCI_U_TYPE type) {
	return CCI_IS_COLLECTION_TYPE(type);
}
*/
import "C"
import (
	"unsafe"
	"strconv"
	"fmt"
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

func gci_prepare(conn_handle int, sql_stmt string, flag byte) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cQuery *C.char = C.CString(sql_stmt)
	var cci_error C.T_CCI_ERROR
	var req C.int
	var err GCI_ERROR

	defer C.free(unsafe.Pointer(cQuery))

	req = C.cci_prepare(cHandle, cQuery, 0, &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])

	return int(req), err
}

func gci_disconnect(conn_handle int) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_disconnect(cHandle, &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])

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

func gci_execute(req_handle int, flag int, max_col_size int) (int, GCI_ERROR) {
	var res C.int
	var cci_error C.T_CCI_ERROR
	var handle C.int = C.int(req_handle)
	var err GCI_ERROR

	res = C.cci_execute(handle, C.char(flag), C.int(max_col_size), &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func gci_set_autocommit(conn_handle int autocommit_mode AUTOCOMMIT_MODE) int {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var mode C.int = C.int(autocommit_mode)

	res = C.cci_set_autocommit(handle, C.CCI_AUTOCOMMIT_MODE(mode))
	return int(res)
}

func gci_end_tran(conn_handle int, tran_type int) (int, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_end_tran(handle, C.char(tran_type), &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func gci_get_last_insert_id(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var value *C.char
	var nid int64

	res = C.cci_get_last_insert_id(handle, unsafe.Pointer(value), &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])
	if res < 0 {
		return int64(res), err
	}

	id := C.GoString(value)
	nid, _ = strconv.ParseInt(id, 0, 64)
	return nid, err
}

func gci_row_count(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var row_count C.int
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_row_count(handle, &row_count, &cci_error)
	if res < 0 {
		err.Err_code = int(cci_error.err_code)
		err.Err_msg = C.GoString(&cci_error.err_msg[0])
	}

	return int64(row_count), err
}

/*
현재는 prototyping만 해 놓자,,,
void *타입에 대한 처리를 어떻게 할지 고민이 필요,,,
*/
func gci_bind_param_int(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.int(value.(int64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_INT, unsafe.Pointer(&c_param), C.CCI_U_TYPE_INT, C.char(flag))
	
	return int(res)
}

func gci_bind_param_string(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	ss := fmt.Sprint(value)
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_STR, unsafe.Pointer(C.CString(ss)), C.CCI_U_TYPE_STRING, C.char(flag))

	return int(res)
}

func gci_bind_param_float(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.float(value.(float64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&c_param), C.CCI_U_TYPE_FLOAT, C.char(flag))

	return int(res)
}

func gci_get_result_info(req_handle int) (*GCI_COL_INFO, GCI_CUBRID_STMT, int) {
	var handle C.int = C.int(req_handle)
	var col_info *C.T_CCI_COL_INFO
	var cubrid_stmt C.T_CCI_CUBRID_STMT
	var col_count C.int
	var gci_col_info *GCI_COL_INFO
	var gci_cubrid_stmt GCI_CUBRID_STMT

	col_info = C.cci_get_result_info(handle, &cubrid_stmt, &col_count)
	gci_cubrid_stmt = GCI_CUBRID_STMT(cubrid_stmt)

	gci_col_info = new(GCI_COL_INFO)
	gci_col_info.u_type = GCI_U_TYPE(col_info._type)
	gci_col_info.is_non_null = C.GoString(&col_info.is_non_null)
	gci_col_info.scale = int16(col_info.scale)
	gci_col_info.precision = int(col_info.precision)

	return gci_col_info, gci_cubrid_stmt, int(col_count)
}

func gci_get_result_info_name(col_info *GCI_COL_INFO, idx int) string {
	var c_name *C.char
	var result string
	
	// todo : GCI_COL_INFO -> T_CCI_COL_INFO
	var cci_col_info *C.T_CCI_COL_INFO
	cci_col_info = col_info.To()
	c_name = C.ex_cci_get_result_info_name(cci_col_info, C.int(idx))
	result = C.GoString(c_name)
	return result
}

func gci_get_result_info_type(col_info *GCI_COL_INFO, idx int) GCI_U_TYPE {
	var result GCI_U_TYPE
	var cci_u_type C.T_CCI_U_TYPE
	
	var cci_col_info *C.T_CCI_COL_INFO
	cci_col_info = col_info.To()
	cci_u_type = C.ex_cci_get_result_info_type(cci_col_info, C.int(idx))

	result = (GCI_U_TYPE)(cci_u_type)
	return result

}

func gci_cursor(req_handle int, offset int, origin GCI_CURSOR_POS) (int, GCI_ERROR) {
	var handle C.int
	var c_offset C.int
	var c_origin C.T_CCI_CURSOR_POS
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var res C.int

	handle = C.int(req_handle)
	c_offset = C.int(offset)
	c_origin = C.T_CCI_CURSOR_POS(origin)

	res = C.cci_cursor(handle, c_offset, c_origin, &cci_error)
	err.Err_code = int(cci_error.err_code)
	err.Err_msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func gci_fetch(req_handle int) (int, GCI_ERROR) {
	var handle C.int = C.int(req_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_fetch(handle, &cci_error)
	
	if res < C.int(0) {
		err.Err_code = int(cci_error.err_code)
		err.Err_msg = C.GoString(&cci_error.err_msg[0])
	}

	return int(res), err
}

func gci_get_data_string(req_handle int, idx int) (int, string, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf *C.char
	var res C.int
	var indicator C.int
	var data string

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &indicator)
	
	data = C.GoString(buf)

	return int(res), data, int(indicator)
}

func gci_get_data_int(req_handle int, idx int) (int, int, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.int
	var res C.int
	var indicator C.int
	var data int

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_INT, unsafe.Pointer(&buf), &indicator)
	data = int(buf)

	return int(res), data, int(indicator)
}

func gci_get_data_float(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.float
	var res C.int
	var indicator C.int
	var data float64
	
	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func gci_get_data_double(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.double
	var res C.int
	var indicator C.int
	var data float64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func gci_get_data_bit(req_handle int, idx int) (int, GCI_BIT, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_BIT
	var res C.int
	var indicator C.int
	var data GCI_BIT

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIT, unsafe.Pointer(&buf), &indicator)
	data.size = int(buf.size)
	//data.buf = make([]byte, data.size)
	data.buf = C.GoBytes(unsafe.Pointer(buf.buf), buf.size)

	return int(res), data, int(indicator)
}

func gci_get_data_date(req_handle int, idx int) (int, GCI_DATE, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_DATE
	var res C.int
	var indicator C.int
	var data GCI_DATE

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DATE, unsafe.Pointer(&buf), &indicator)
	data.yr = int(buf.yr)
	data.mon = int(buf.mon)
	data.day = int(buf.day)
	data.hh = int(buf.hh)
	data.mm = int(buf.mm)
	data.ss = int(buf.ss)
	data.ms = int(buf.ms)

	return int(res), data, int(indicator)
}

func gci_get_data_bigint(req_handle int, idx int) (int, int64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.int64_t
	var res C.int
	var indicator C.int
	var data int64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIGINT, unsafe.Pointer(&buf), &indicator)
	data = int64(buf)

	return int(res), data, int(indicator)
}

func gci_get_data_blob(req_handle int, idx int) (int, GCI_BLOB, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_BLOB
	var res C.int
	var indicator C.int
	var data GCI_BLOB

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BLOB, unsafe.Pointer(&buf), &indicator)
	data = GCI_BLOB(buf)
	
	return int(res), data, int(indicator)
}

func gci_blob_size(blob GCI_BLOB) int64 {
	var size C.longlong
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)
	size = C.cci_blob_size(data)

	return int64(size)
}

func gci_blob_read(con_handle int, blob GCI_BLOB, start_pos int64, length int64) (GCI_BLOB, GCI_ERROR) {
	var handle C.int = C.int(con_handle)
	var res C.int
	var c_start_pos C.longlong = C.longlong(start_pos)
	var c_length C.int = C.int(length)
	var c_blob string
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)
	var res_blob GCI_BLOB

	c_buf := C.CString(c_blob)
	defer C.free(unsafe.Pointer(c_buf))
	res = C.cci_blob_read(handle, data, c_start_pos, c_length, c_buf, &cci_error)
	if res < C.int(0) {
		err.Err_code = int(cci_error.err_code)
		err.Err_msg = C.GoString(&cci_error.err_msg[0])
	}

	res_blob = GCI_BLOB(c_buf)


	return res_blob, err 
}

func gci_blob_free(blob GCI_BLOB) {
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob) 
	C.cci_blob_free(data)
}
