package gci

/*
#cgo CFLAGS: -I../CUBRID/include
#cgo LDFLAGS: -L../CUBRID/lib -lcascci -lnsl
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
	"reflect"
	"log"
)

func Gci_init() {
	C.cci_init()
}

func Gci_end() {
	C.cci_end()
}
/*
func gci_get_version_string(string &str) int {
}
*/

func Gci_connect(ip string, port int, db_name string, db_user string, db_password string) int {
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

func Gci_connect_with_url(url string, user string, password string) int {
	serverUrl := C.CString(url)
	serverUser := C.CString(user)
	serverPassword := C.CString(password)

	defer C.free(unsafe.Pointer(serverUrl))
	defer C.free(unsafe.Pointer(serverUser))
	defer C.free(unsafe.Pointer(serverPassword))

	con := C.cci_connect_with_url(serverUrl, serverUser, serverPassword)

	return int(con)
}

func Gci_disconnect(conn_handle int) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_disconnect(cHandle, &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Gci_prepare(conn_handle int, sql_stmt string, flag byte) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cQuery *C.char = C.CString(sql_stmt)
	var cci_error C.T_CCI_ERROR
	var req C.int
	var err GCI_ERROR

	defer C.free(unsafe.Pointer(cQuery))

	req = C.cci_prepare(cHandle, cQuery, 0, &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(req), err
}

func Gci_close_req_handle(req_handle int) int {
	var err C.int
	var handle C.int = C.int(req_handle)

	err = C.cci_close_req_handle(handle)

	return int(err)
}

func Gci_get_bind_num(req_handle int) int {
	var param_cnt C.int
	var handle C.int = C.int(req_handle)

	param_cnt = C.cci_get_bind_num(handle)

	return int(param_cnt)
}

func Gci_execute(req_handle int, flag int, max_col_size int) (int, GCI_ERROR) {
	var res C.int
	var cci_error C.T_CCI_ERROR
	var handle C.int = C.int(req_handle)
	var err GCI_ERROR

	res = C.cci_execute(handle, C.char(flag), C.int(max_col_size), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Gci_set_autocommit(conn_handle int autocommit_mode AUTOCOMMIT_MODE) int {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var mode C.int = C.int(autocommit_mode)

	res = C.cci_set_autocommit(handle, C.CCI_AUTOCOMMIT_MODE(mode))

	return int(res)
}

func Gci_end_tran(conn_handle int, tran_type int) (int, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_end_tran(handle, C.char(tran_type), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Gci_get_last_insert_id(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var value *C.char
	var nid int64

	res = C.cci_get_last_insert_id(handle, unsafe.Pointer(value), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])
	if res < 0 {
		return int64(res), err
	}

	id := C.GoString(value)
	nid, _ = strconv.ParseInt(id, 0, 64)

	return nid, err
}

func Gci_row_count(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var row_count C.int
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_row_count(handle, &row_count, &cci_error)
	if res < 0 {
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	return int64(row_count), err
}

func Gci_bind_param_int(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.int(value.(int64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_INT, 
				unsafe.Pointer(&c_param), C.CCI_U_TYPE_INT, C.char(flag))

	return int(res)
}

func Gci_bind_param_string(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	ss := fmt.Sprint(value)
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_STR, 
				unsafe.Pointer(C.CString(ss)), C.CCI_U_TYPE_STRING, C.char(flag))

	return int(res)
}

func Gci_bind_param_float(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.float(value.(float64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_FLOAT, 
				unsafe.Pointer(&c_param), C.CCI_U_TYPE_FLOAT, C.char(flag))

	return int(res)
}

func Gci_get_result_info(req_handle int) ([]GCI_COL_INFO, GCI_CUBRID_STMT, int) {
	var handle C.int = C.int(req_handle)
	var c_col_info *C.T_CCI_COL_INFO
	var go_col_info []C.T_CCI_COL_INFO
	var cubrid_stmt C.T_CCI_CUBRID_STMT
	var col_count C.int
	var gci_col_info []GCI_COL_INFO
	var gci_cubrid_stmt GCI_CUBRID_STMT

	c_col_info = C.cci_get_result_info(handle, &cubrid_stmt, &col_count)
	gci_cubrid_stmt = GCI_CUBRID_STMT(cubrid_stmt)

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&go_col_info)))
	sliceHeader.Cap = int(col_count)
	sliceHeader.Len = int(col_count)
	sliceHeader.Data = uintptr(unsafe.Pointer(c_col_info))
	gci_col_info = make([]GCI_COL_INFO, int(col_count))
	for i := 0; i < int(col_count); i++ {
		gci_col_info[i].u_type = GCI_U_TYPE(go_col_info[C.int(i)]._type)
		gci_col_info[i].is_non_null = C.GoString(&go_col_info[C.int(i)].is_non_null)
		gci_col_info[i].scale = int16(go_col_info[C.int(i)].scale)
		gci_col_info[i].precision = int(go_col_info[C.int(i)].precision)
	}

	return gci_col_info, gci_cubrid_stmt, int(col_count)
}

func Gci_get_result_info_name(col_info []GCI_COL_INFO, idx int) string {
	var result string
	result = col_info[idx - 1].col_name
	return result
}

func Gci_get_result_info_type(col_info []GCI_COL_INFO, idx int) GCI_U_TYPE {
	var result GCI_U_TYPE
	result = col_info[idx - 1].u_type
	return result
}

func Gci_is_collection_type(u_type GCI_U_TYPE) int {
	var result int
	// 이게 맞는건가????
	res := (u_type) & GCI_CODE_COLLECTION
	if( (res != 0) || ((u_type) == U_TYPE_SET) || ((u_type) == U_TYPE_MULTISET) || ((u_type) == U_TYPE_SEQUENCE) ) {
		result = 1
	} else {
		result =  0
	}
	return result
}

func Gci_cursor(req_handle int, offset int, origin GCI_CURSOR_POS) (int, GCI_ERROR) {
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
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Gci_fetch(req_handle int) (int, GCI_ERROR) {
	var handle C.int = C.int(req_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_fetch(handle, &cci_error)
	if res < C.int(0) {
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	return int(res), err
}

func Gci_get_data_string(req_handle int, idx int) (int, string, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf *C.char
	var res C.int
	var indicator C.int
	var data string

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &indicator)
	data = C.GoString(buf)

	return int(res), data, int(indicator)
}

func Gci_get_data_int(req_handle int, idx int) (int, int, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.int
	var res C.int
	var indicator C.int
	var data int

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_INT, unsafe.Pointer(&buf), &indicator)
	data = int(buf)

	return int(res), data, int(indicator)
}

func Gci_get_data_float(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.float
	var res C.int
	var indicator C.int
	var data float64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func Gci_get_data_double(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.double
	var res C.int
	var indicator C.int
	var data float64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func Gci_get_data_bit(req_handle int, idx int) (int, GCI_BIT, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.T_CCI_BIT
	var res C.int
	var indicator C.int
	var data GCI_BIT

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIT, unsafe.Pointer(&buf), &indicator)
	data.size = int(buf.size)
	data.buf = C.GoBytes(unsafe.Pointer(buf.buf), buf.size)

	return int(res), data, int(indicator)
}

func Gci_get_data_set(req_handle int, idx int) (int, GCI_SET, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.T_CCI_SET
	var res C.int
	var indicator C.int
	var data GCI_SET

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_SET, unsafe.Pointer(&buf), &indicator)
	data = GCI_SET(buf)

	return int(res), data, int(indicator)

}

func Gci_set_size(set GCI_SET) int {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var res C.int

	res = C.cci_set_size(data);

	return int(res)
}

/*
	set 버퍼 안에서의 index와 a_type
	ex) {'a', 'b', 'c'}
*/
func Gci_set_get(set GCI_SET, index int, a_type GCI_A_TYPE) (int, interface{}, int) {
	var indicator int
	var res int
	var data interface{}

	switch a_type {
	case A_TYPE_STR:
		res, data, indicator = gci_set_get_str(set, index)
	case A_TYPE_INT:
		res, data, indicator = gci_set_get_int(set, index)
	case A_TYPE_FLOAT:
		res, data, indicator = gci_set_get_float(set, index)
	case A_TYPE_DOUBLE:
		res, data, indicator = gci_set_get_float(set, index)
	case A_TYPE_BIT:
		res, data, indicator = gci_set_get_bit(set, index)
	case A_TYPE_DATE:
		res, data, indicator = gci_set_get_date(set, index)
	case A_TYPE_BIGINT:
		res, data, indicator = gci_set_get_bigint(set, index)
	// todo
	//case A_TYPE_BLOB:
	//	res, data, indicator = gci_set_get_blob(set, index)
	//case A_TYPE_CLOB:
	//	res, data, indicator = gci_set_get_clob(set, index)
	}

	return res, data, indicator
}

func gci_set_get_str(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value *C.char
	var indicator C.int
	var res C.int

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_STR, unsafe.Pointer(&value), &indicator)

	rv := C.GoString(value)
	return int(res), rv, int(indicator)
}

func gci_set_get_int(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.int
	var indicator C.int
	var res C.int

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_INT, unsafe.Pointer(&value), &indicator)

	rv := int(value)
	return int(res), rv, int(indicator)
}

func gci_set_get_float(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.float
	var indicator C.int
	var res C.int

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&value), &indicator)

	rv := float64(value)
	return int(res), rv, int(indicator)
}

func gci_set_get_double(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.double
	var indicator C.int
	var res C.int

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&value), &indicator)

	rv := float64(value)
	return int(res), rv, int(indicator)
}

func gci_set_get_bit(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.T_CCI_BIT
	var indicator C.int
	var res C.int
	var rv GCI_BIT

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_BIT, unsafe.Pointer(&value), &indicator)

	rv.size = int(value.size)
	rv.buf = C.GoBytes(unsafe.Pointer(value.buf), value.size)

	return int(res), rv, int(indicator)
}

func gci_set_get_date(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.T_CCI_DATE
	var indicator C.int
	var res C.int
	var rv GCI_DATE

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_DATE, unsafe.Pointer(&value), &indicator)

	rv.yr = int(value.yr)
	rv.mon = int(value.mon)
	rv.day = int(value.day)
	rv.hh = int(value.hh)
	rv.mm = int(value.mm)
	rv.ss = int(value.ss)
	rv.ms = int(value.ms)
	
	return int(res), rv, int(indicator)
}

func gci_set_get_bigint(set GCI_SET, index int) (int, interface{}, int) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var value C.int64_t
	var indicator C.int
	var res C.int

	res = C.cci_set_get(data, C.int(index), C.CCI_A_TYPE_BIGINT, unsafe.Pointer(&value), &indicator)

	rv := int64(value)
	return int(res), rv, int(indicator)
}


func Gci_set_free(set GCI_SET) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)

	C.cci_set_free(data)
}

func Gci_get_data_date(req_handle int, idx int) (int, GCI_DATE, int) {
	log.Println("gci_get_data_date_start")
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_DATE
	var res C.int
	var indicator C.int
	var data GCI_DATE

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DATE, 
				unsafe.Pointer(&buf), &indicator)
	data.yr = int(buf.yr)
	data.mon = int(buf.mon)
	data.day = int(buf.day)
	data.hh = int(buf.hh)
	data.mm = int(buf.mm)
	data.ss = int(buf.ss)
	data.ms = int(buf.ms)
	
	log.Println("gci_get_data_date_end")
	return int(res), data, int(indicator)
}

func Gci_get_data_bigint(req_handle int, idx int) (int, int64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.int64_t
	var res C.int
	var indicator C.int
	var data int64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIGINT, 
				unsafe.Pointer(&buf), &indicator)
	data = int64(buf)

	return int(res), data, int(indicator)
}

func Gci_get_data_blob(req_handle int, idx int) (int, GCI_BLOB, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_BLOB
	var res C.int
	var indicator C.int
	var data GCI_BLOB

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BLOB, 
				unsafe.Pointer(&buf), &indicator)
	data = GCI_BLOB(buf)

	return int(res), data, int(indicator)
}

func Gci_blob_size(blob GCI_BLOB) int64 {
	var size C.longlong
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)

	size = C.cci_blob_size(data)

	return int64(size)
}

func Gci_blob_read(con_handle int, blob GCI_BLOB, start_pos int64, length int64) (GCI_BLOB, GCI_ERROR) {
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
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	res_blob = GCI_BLOB(c_buf)

	return res_blob, err
}

func Gci_blob_free(blob GCI_BLOB) {
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)
	C.cci_blob_free(data)
}

