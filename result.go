package cubrid

import (
	"fmt"

	"github.com/jaehoonkim/cubrid/gci"
)

type cubridResult struct {
	c *cubridConn
}

func (result *cubridResult) LastInsertId() (int64, error) {
	var last_id int64
	var err gci.GCI_ERROR

	last_id, err = gci.Get_last_insert_id(result.c.con)
	if last_id < 0 {
		return 0, fmt.Errorf("cci_get_last_insert_id err: %d", err.Code)
	}

	return last_id, nil
}

func (result *cubridResult) RowsAffected() (int64, error) {
	var row_count int64
	var err gci.GCI_ERROR
	row_count, err = gci.Row_count(result.c.con)
	if err.Code < 0 {
		return 0, fmt.Errorf("cci_row_count err: %d", err.Code)
	}
	return row_count, nil
}
