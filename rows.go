package cubrid

import (
	"database/sql/driver"
	"fmt"
	"io"
)

type cubridRows struct {
	s *cubridStmt
}

func (rows *cubridRows) Columns() []string {
	//log.Println("cubridRow:Columns")
	var col_info *GCI_COL_INFO
	//var stmt_type GCI_CUBRID_STMT
	var col_count int
	var idx int
	//log.Println("cbubridStmt:bindParam")
	col_info, _, col_count = gci_get_result_info(rows.s.req)
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
	//var stmt_type GCI_CUBRID_STMT
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

	col_info, _, col_count = gci_get_result_info(rows.s.req)
	var columnType GCI_U_TYPE
	var i int
	//var ind int
	for i = 1; i <= col_count; i++ {
		columnType = gci_get_result_info_type(col_info, i)
		
		switch columnType {
		case U_TYPE_CHAR, U_TYPE_STRING, U_TYPE_NCHAR/*, U_TYPE_VARNCHAR*/:
			var data string
			res, data, _ = gci_get_data_string(rows.s.req, i)
			if res < 0 {
				return fmt.Errorf("get_data err : %d, %d\n", res, i)
			}
			dest[i - 1] = data
		case U_TYPE_INT, U_TYPE_NUMERIC, U_TYPE_SHORT:
			var data int
			res, data, _ = gci_get_data_int(rows.s.req, i)
			dest[int(i - 1)] = data
 		case U_TYPE_FLOAT:
			var data float64
			res, data, _ = gci_get_data_float(rows.s.req, i)
			dest[int(i - 1)] = data
		case U_TYPE_DOUBLE:
			var data float64
			res, data, _ = gci_get_data_double(rows.s.req, i)
			dest[int(i - 1)] = data
		case U_TYPE_BIT, U_TYPE_VARBIT:
			var data GCI_BIT
			res, data, _ = gci_get_data_bit(rows.s.req, i)
			dest[int(i - 1)] = data
		case U_TYPE_DATE, U_TYPE_TIME, U_TYPE_TIMESTAMP:
			var data GCI_DATE
			res, data, _ = gci_get_data_date(rows.s.req, i)
			dest[int(i - 1)] = data
		//case C.CCI_U_TYPE_OBJECT, C.CCI_U_TYPE_RESULTSET:
			//log.Println("cci_a_type_set")
		case U_TYPE_BIGINT:
			var data int64
			res, data, _ = gci_get_data_bigint(rows.s.req, i)
			dest[int(i - 1)] = data
		case U_TYPE_BLOB:
			var org_data, data GCI_BLOB
			var size int64 

			res, org_data, _ = gci_get_data_blob(rows.s.req, i)
			size = gci_blob_size(org_data)
			data, err = gci_blob_read(rows.s.c.con, org_data, 0, size)

			gci_blob_free(org_data)
			dest[int(i - 1)] = data
		case U_TYPE_CLOB:
			/*var clob C.T_CCI_CLOB
			var size C.longlong
			var buf string
			cBuf := C.CString(buf)
			C.cci_get_data(rows.s.req, i, C.CCI_A_TYPE_CLOB, unsafe.Pointer(&clob), &ind)
			size = C.cci_clob_size(clob)
			C.cci_clob_read(rows.s.c.con, clob, 0, C.int(size), cBuf, &cci_error)
			_clob := CCI_CLOB { _CLOB : C.GoString(cBuf) }
			C.free(unsafe.Pointer(cBuf))
			C.cci_clob_free(clob)
			dest[int(i - 1)] = _clob*/
		default:
			/*
			if int(C.ex_cci_is_collection_type(columnType)) == 1 {
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
			}*/
		}
	}
	return nil
}

