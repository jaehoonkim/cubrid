package cubrid

/*
#cgo CFLAGS: -ICUBRID/include
#cgo LDFLAGS: -LCUBRID/lib -lcascci -lnsl
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
