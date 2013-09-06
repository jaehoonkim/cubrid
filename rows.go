package cubrid

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
	var col_info *GCI_COL_INFO
	var stmt_type GCI_CUBRID_STMT
	var col_count int
	var idx int
	//log.Println("cbubridStmt:bindParam")
	col_info, stmt_type, col_count = gci_get_result_info(rows.s.req)
	if col_info == nil {
		return nil
	}

	col_name  := make([]string, int(col_count))
	for idx = 1; idx <= col_count; idx++ {
		col_name[idx] = gci_get_result_info_name(col_info, idx)
	}
	return col_name
}

func (rows *cubridRows) Close() error {
	var res int
	res = gci_close_req_handle(rows.s.req)
	if res < 0 {
		return fmt.Errorf("close_req_handle err : %d", res)
	}
	return nil
}

func (rows *cubridRows) Next(dest []driver.Value) error {
	var res int
	var err GCI_ERROR
	var col_info *GCI_COL_INFO
	var stmt_type GCI_CUBRID_STMT
	var col_count int

	res, err = gci_cursor(rows.s.req, 1, GCI_CURSOR_CURRENT)
	if res == int(GCI_ER_NO_MORE_DATA) {
		return io.EOF
	}
	if res < 0 {
		return fmt.Errorf("cursor err: %d, %s", err.Err_code, err.Err_msg)
	}

	res, err = gci_fetch(rows.s.req)
	if res < 0 {
		return fmt.Errorf("fetch err: %d, %s", err.Err_code, err.Err_msg)
	}

	col_info, stmt_type, col_count = gci_get_result_info(rows.s.req)
	var columnType GCI_U_TYPE
	var i int
	var ind int
	for i = 1; i <= col_count; i++ {
		columnType = gci_get_result_info_type(col_info, i)
		
		switch columnType {
		case GCI_U_TYPE_CHAR, GCI_U_TYPE_STRING, GCI_U_TYPE_NCHAR, GCI_U_TYPE_VARNCHAR:
			//log.Println("cci_a_type_str")
			var buf *C.char
			err = C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &ind)
			if int(err) < 0 {
				return fmt.Errorf("get_data err : %d, %d\n", err, int(i))
			}
			dest[int(i - 1)] = C.GoString(buf)
		case GCI_U_TYPE_INT, GCI_U_TYPE_NUMERIC, GCI_U_TYPE_SHORT:
			//log.Println("cci_a_type_int")
			var buf C.int
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_INT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = int(buf)
 		case GCI_U_TYPE_FLOAT:
			//log.Println("cci_a_type_float")
			var buf C.float
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = float64(buf)
		case GCI_U_TYPE_DOUBLE:
			//log.Println("cci_a_type_double")
			var buf C.double
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = float64(buf)
		case GCI_U_TYPE_BIT, GCI_U_TYPE_VARBIT:
			//log.Println("cci_a_type_bit")
			var buf C.T_CCI_BIT
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BIT, unsafe.Pointer(&buf), &ind)
			_bit := CCI_BIT{ buf }
			dest[int(i - 1)] = _bit
		case GCI_U_TYPE_DATE, GCI_U_TYPE_TIME, GCI_U_TYPE_TIMESTAMP:
			//log.Println("cci_a_type_date")
			var buf C.T_CCI_DATE
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_DATE, unsafe.Pointer(&buf), &ind)

			_date := CCI_DATE{ buf }
			dest[int(i - 1)] = _date
		//case C.CCI_U_TYPE_OBJECT, C.CCI_U_TYPE_RESULTSET:
			//log.Println("cci_a_type_set")
		case GCI_U_TYPE_BIGINT:
			//log.Println("cci_a_type_bigint")
			var buf C.int64_t
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BIGINT, unsafe.Pointer(&buf), &ind)
			dest[int(i - 1)] = int64(buf)
		case GCI_U_TYPE_BLOB:
			//log.Println("cci_a_type_blob")
			var blob C.T_CCI_BLOB
			var size C.longlong
			var buf string
			cBuf := C.CString(buf)
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_BLOB, unsafe.Pointer(&blob), &ind)
			size = C.cci_blob_size(blob)
			C.cci_blob_read(rows.s.c.con, blob, 0, C.int(size), cBuf, &cci_error)
			_blob := CCI_BLOB { _BLOB : C.GoBytes(unsafe.Pointer(cBuf), C.int(size)) }
			C.cci_blob_free(blob)
			C.free(unsafe.Pointer(cBuf))
			dest[int(i - 1)] = _blob
		case GCI_U_TYPE_CLOB:
			//log.Println("cci_u_type_clob")
			var clob C.T_CCI_CLOB
			var size C.longlong
			var buf string
			cBuf := C.CString(buf)
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_CLOB, unsafe.Pointer(&clob), &ind)
			size = C.cci_clob_size(clob)
			C.cci_clob_read(rows.s.c.con, clob, 0, C.int(size), cBuf, &cci_error)
			_clob := CCI_CLOB { _CLOB : C.GoString(cBuf) }
			C.free(unsafe.Pointer(cBuf))
			C.cci_clob_free(clob)
			dest[int(i - 1)] = _clob
		default:
			if int(C.ex_cci_is_collection_type(columnType)) == 1 {
				//log.Println("ex_cci_is_set_type")
				var set C.T_CCI_SET
				C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_SET, unsafe.Pointer(&set), &ind)
				if(int(ind) == -1) {
					log.Println("set data is nil")
					return nil
				}
				var set_size C.int
				var buf *C.char
				set_size = C.cci_set_size(set)

				var _set CCI_SET
				_set.MakeBuf(int(set_size))
				for j := C.int(0); j < set_size; j++ {
					res := C.cci_set_get(set, j+1, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &ind)
					if res < 0 {
						C.cci_set_free(set)
						return nil
					}
					_set.SetBuf(int(j), C.GoString(buf))
				}
				dest[int(i - 1)] = _set
				C.cci_set_free(set)
			}
		}
	}
	return nil
}

