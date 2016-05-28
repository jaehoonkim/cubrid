package cubrid

import (
	"fmt"

	"gci"
)

type cubridTx struct {
	c *cubridConn
}

func (tx *cubridTx) Commit() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err gci.GCI_ERROR

	res, err = gci.End_tran(conn_handle, gci.TRAN_COMMIT)
	if res < 0 {
		gci.Disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Code)
	}
	return nil
}

func (tx *cubridTx) Rollback() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err gci.GCI_ERROR

	res, err = gci.End_tran(conn_handle, gci.TRAN_ROLLBACK)
	if res < 0 {
		gci.Disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Code)
	}
	return nil
}
