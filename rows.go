package cubrid

import (
	"database/sql/driver"
	"fmt"
	"io"
	"log"

	"gci"
)

type cubridRows struct {
	s *cubridStmt
}

func (rows *cubridRows) Columns() []string {
	log.Println("cubridRow:Columns")
	var col_info []gci.GCI_COL_INFO
	var col_count int
	var idx int

	col_info, _, col_count = gci.Get_result_info(rows.s.req)
	if col_info == nil {
		return nil
	}

	col_name := make([]string, int(col_count))
	for idx = 1; idx <= col_count; idx++ {
		col_name[idx-1] = gci.Get_result_info_name(col_info, idx)
	}

	return col_name
}

func (rows *cubridRows) Close() error {
	var res int

	res = gci.Close_req_handle(rows.s.req)
	if res < 0 {
		return fmt.Errorf("close_req_handle err : %d", res)
	}

	return nil
}

func (rows *cubridRows) Next(dest []driver.Value) error {
	var res int
	var err gci.GCI_ERROR
	var col_info []gci.GCI_COL_INFO
	var col_count int

	//log.Println("##### rows::Next #####")
	res, err = gci.Cursor(rows.s.req, 1, gci.GCI_CURSOR_CURRENT)
	if res == int(gci.GCI_ER_NO_MORE_DATA) {
		return io.EOF
	}
	if res < 0 {
		return fmt.Errorf("cursor err: %d, %s", err.Code, err.Msg)
	}

	res, err = gci.Fetch(rows.s.req)
	if res < 0 {
		return fmt.Errorf("fetch err: %d, %s", err.Code, err.Msg)
	}

	col_info, _, col_count = gci.Get_result_info(rows.s.req)
	var columnType gci.GCI_U_TYPE
	var i int
	//var ind int
	for i = 1; i <= col_count; i++ {
		columnType = gci.Get_result_info_type(col_info, i)

		switch columnType {
		case gci.U_TYPE_CHAR, gci.U_TYPE_STRING, gci.U_TYPE_NCHAR /*, U_TYPE_VARNCHAR*/ :
			var data string
			res, data, _ = gci.Get_data_string(rows.s.req, i)
			if res < 0 {
				return fmt.Errorf("get_data err : %d, %d\n", res, i)
			}
			dest[i-1] = data
		case gci.U_TYPE_INT, gci.U_TYPE_NUMERIC, gci.U_TYPE_SHORT:
			var data int
			res, data, _ = gci.Get_data_int(rows.s.req, i)
			dest[int(i-1)] = data
		case gci.U_TYPE_FLOAT:
			var data float64
			res, data, _ = gci.Get_data_float(rows.s.req, i)
			dest[int(i-1)] = data
		case gci.U_TYPE_DOUBLE:
			var data float64
			res, data, _ = gci.Get_data_double(rows.s.req, i)
			dest[int(i-1)] = data
		case gci.U_TYPE_BIT, gci.U_TYPE_VARBIT:
			var data gci.GCI_BIT
			res, data, _ = gci.Get_data_bit(rows.s.req, i)
			dest[int(i-1)] = data
		case gci.U_TYPE_DATE, gci.U_TYPE_TIME, gci.U_TYPE_TIMESTAMP:
			var data gci.GCI_DATE
			res, data, _ = gci.Get_data_date(rows.s.req, i)
			dest[int(i-1)] = data
		//case C.CCI_U_TYPE_OBJECT, C.CCI_U_TYPE_RESULTSET:
		//log.Println("cci_a_type_set")
		case gci.U_TYPE_BIGINT:
			var data int64
			res, data, _ = gci.Get_data_bigint(rows.s.req, i)
			dest[int(i-1)] = data
		case gci.U_TYPE_BLOB:
			var org_data, data gci.GCI_BLOB
			var size int64
			res, org_data, _ = gci.Get_data_blob(rows.s.req, i)
			size = gci.Blob_size(org_data)
			data, err = gci.Blob_read(rows.s.c.con, org_data, 0, size)

			gci.Blob_free(org_data)
			dest[int(i-1)] = data
		case gci.U_TYPE_CLOB:
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
			if int(gci.Is_collection_type(columnType)) == 1 {
				var data gci.GCI_SET
				var indicator int
				_, data, indicator = gci.Get_data_set(rows.s.req, i)
				if indicator == -1 {
					log.Println("set data is nil")
					return nil
				}
				dest[int(i-1)] = data
			}
		}
	}
	return nil
}
