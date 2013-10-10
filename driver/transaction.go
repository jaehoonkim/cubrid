package cubrid

import (
	"fmt"
	"github.com/sabzil/cubrid/gci"
)

type cubridTx struct {
	c *cubridConn
}
func (tx *cubridTx) Commit() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err gci.GCI_ERROR

	res, err= gci.Gci_end_tran(conn_handle, gci.TRAN_COMMIT)
	if res < 0 {
		gci.Gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Code)
	}
	return nil
}

func (tx *cubridTx) Rollback() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err gci.GCI_ERROR

	res, err = gci.Gci_end_tran(conn_handle, gci.TRAN_ROLLBACK)
	if res < 0 {
		gci.Gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Code)
	}
	return nil
}

