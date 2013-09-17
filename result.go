package cubrid

import (
	"fmt"
)

type cubridResult struct {
	c *cubridConn
}


func (result *cubridResult) LastInsertId() (int64, error) {
	var last_id int64
	var err GCI_ERROR

	last_id, err = Gci_get_last_insert_id(result.c.con)
	if last_id < 0 {
		return 0, fmt.Errorf("cci_get_last_insert_id err: %d", err.Err_code)
	}

	return last_id, nil
}

func (result *cubridResult) RowsAffected() (int64, error) {
	var row_count int64
	var err GCI_ERROR
	row_count, err = Gci_row_count(result.c.con)
	if err.Err_code < 0 {
		return 0, fmt.Errorf("cci_row_count err: %d", err.Err_code)
	}
	return row_count, nil
}

