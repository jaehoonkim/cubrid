package cubrid

import (
	"fmt"
)

type cubridResult struct {
	c *cubridConn
}


func (result *cubridResult) LastInsertId() (int64, error) {
	var last_id int64
	var cci_error CCI_ERROR

	last_id, cci_error = gci_get_last_insert_id(result.c.con)
	if last_id < 0 {
		return 0, fmt.Errorf("cci_get_last_insert_id err: %d", last_id)
	}

	return last_id, nil
}

func (result *cubridResult) RowsAffected() (int64, error) {
	var row_count int64
	var cci_error CCI_ERROR
	row_count, cci_error = gci_row_count(result.c.con)
	if cci_error.err_code < 0 {
		return 0, fmt.Errorf("cci_row_count err: %d", cci_error.err_code)
	}
	return row_count, nil
}

