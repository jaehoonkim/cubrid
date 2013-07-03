package cubrid

/*
#include <stdio.h>
#include <stdlib.h>
#include "cas_cci.h"
char* ex_cci_get_result_info_name(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_NAME(res_info, index);
//	return  res_info[index - 1].col_name;
}

T_CCI_U_TYPE ex_cci_get_result_info_type(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_TYPE(res_info, index);
//	return res_info[index - 1].type;
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
	"database/sql/driver"
	"fmt"
	"io"
	"unsafe"
	"log"
)

type cubridRows struct {
	s *cubridStmt
}

func (rows *cubridRows) Columns() []string {
	//log.Println("cubridRow:Columns")
	var col_info *C.T_CCI_COL_INFO
	var stmt_type C.T_CCI_CUBRID_STMT
	var col_count, idx  C.int
	col_info = C.cci_get_result_info(rows.s.req, &stmt_type, &col_count)
	if col_info == nil {
		return nil
	}

	var c_name *C.char
	col_name  := make([]string, int(col_count))
	for idx = C.int(1); idx <= col_count; idx++ {
		c_name = C.ex_cci_get_result_info_name(col_info, idx)
		col_name[int(idx-1)] = C.GoString(c_name);
		//log.Println(col_name[int(idx-1)])
	}
	return col_name
}

func (rows *cubridRows) Close() error {
	//log.Println("cubridRows:Close")
	var err C.int
	err = C.cci_close_req_handle(rows.s.req)
	if int(err) < 0 {
		return fmt.Errorf("close_req_handle err : %d", int(err))
	}
	return nil
}

func (rows *cubridRows) Next(dest []driver.Value) error {
	//log.Println("cubridRows:Next")
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
	var ind C.int
	for i = C.int(1); i <= col_count; i++ {
		columnType = C.ex_cci_get_result_info_type(col_info, i)
		log.Printf("columnType : %d", int(columnType))
		switch columnType {
		case C.CCI_U_TYPE_CHAR, C.CCI_U_TYPE_STRING, C.CCI_U_TYPE_NCHAR, C.CCI_U_TYPE_VARNCHAR:
			log.Println("cci_a_type_str")
			var buf *C.char
			err = C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &ind)
			if int(err) < 0 {
				return fmt.Errorf("get_data err : %d, %d\n", err, int(i))
			}
			//log.Printf("cci_a_type_str: %s", C.GoString(buf))
			dest[int(i - 1)] = C.GoString(buf)
		case C.CCI_U_TYPE_INT, C.CCI_U_TYPE_NUMERIC, C.CCI_U_TYPE_SHORT:
			log.Println("cci_a_type_int")
			var buf C.int
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_INT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = int(buf)
 		case C.CCI_U_TYPE_FLOAT:
			log.Println("cci_a_type_float")
			var buf C.float
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = float64(buf)
		case C.CCI_U_TYPE_DOUBLE:
			log.Println("cci_a_type_double")
			var buf C.double
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = float64(buf)
		case C.CCI_U_TYPE_BIT, C.CCI_U_TYPE_VARBIT:
			log.Println("cci_a_type_bit")
			var buf C.T_CCI_BIT
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BIT, unsafe.Pointer(&buf), &ind)
			_bit := CCI_BIT{ buf }
			//log.Printf("cci_bit : %x, %d", C.GoString(buf.buf), int(buf.size))
			dest[int(i - 1)] = _bit
		case C.CCI_U_TYPE_DATE, C.CCI_U_TYPE_TIME, C.CCI_U_TYPE_TIMESTAMP:
			log.Println("cci_a_type_date")
			var buf C.T_CCI_DATE
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_DATE, unsafe.Pointer(&buf), &ind)

			_date := CCI_DATE{ buf }
			dest[int(i - 1)] = _date
			//log.Printf("cci_a_type_date:%d,%d,%d", int(_date._DATE.yr), _date._DATE.mon, _date._DATE.day)
		case /*C.CCI_U_TYPE_SET, C.CCI_U_TYPE_MULTISET, C.CCI_U_TYPE_SEQUENCE,*/ C.CCI_U_TYPE_OBJECT, C.CCI_U_TYPE_RESULTSET:
			log.Println("cci_a_type_set")
			//var buf C.T_CCI_SET
			//C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_SET, unsafe.Pointer(&buf), &ind)
			//_set := CCI_SET { buf }
			//dest[int(i - 1)] = _set
		case C.CCI_U_TYPE_BIGINT:
			log.Println("cci_a_type_bigint")
			var buf C.int64_t
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BIGINT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = int64(buf)
		case C.CCI_U_TYPE_BLOB:
			log.Println("cci_a_type_blob")
			var buf C.T_CCI_BLOB
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BLOB, unsafe.Pointer(&buf), &ind)
			_blob := CCI_BLOB { buf }
			dest[int(i - 1)] = _blob
		case C.CCI_U_TYPE_CLOB:
			log.Println("cci_u_type_clob")
			var buf C.T_CCI_CLOB
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_CLOB, unsafe.Pointer(&buf), &ind)
			_clob := CCI_CLOB { buf }
			dest[int(i - 1)] = _clob
		default:
			if int(C.ex_cci_is_collection_type(columnType)) == 1 {
				log.Println("ex_cci_is_set_type")
				var set C.T_CCI_SET
				C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_SET, unsafe.Pointer(&set), &ind)
				var set_size C.int
				var buf *C.char
				set_size = C.cci_set_size(set)
				var _set CCI_SET
				_set.makeBuf(int(set_size))
				for j := C.int(0); j < set_size; j++ {
					res := C.cci_set_get(set, j+1, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &ind)
					if res < 0 {
						C.cci_set_free(set)
						return nil
					}
					_set.setBuf(int(j), C.GoString(buf))
				}
				dest[int(i - 1)] = _set
				C.cci_set_free(set)
			}
		}
	}
	return nil
}

