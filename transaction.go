package cubrid

import "fmt"

type cubridTx struct {
	c *cubridConn
}
func (tx *cubridTx) Commit() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err GCI_ERROR

	res, err= gci_end_tran(conn_handle, TRAN_COMMIT)
	if res < 0 {
		gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Err_code)
	}
	return nil
}

func (tx *cubridTx) Rollback() error {
	var conn_handle int
	conn_handle = tx.c.con
	var res int
	var err GCI_ERROR

	res, err = gci_end_tran(conn_handle, TRAN_ROLLBACK)
	if res < 0 {
		gci_disconnect(conn_handle)
		return fmt.Errorf("commit error : %d", err.Err_code)
	}
	return nil
}

