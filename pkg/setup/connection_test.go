package setup

import (
	"testing"

	_ "github.com/lib/pq"
)

func TestTableConnect(t *testing.T) {

	conns := []struct {
		db  string
		ctx string
	}{
		{"postgres", "host=localhost port=5432 user=nra_wsl password=password dbname=db_nra_wsl sslmode=disable"},
		//		{"pos", "host=localhost port=5432 user=nra_wsl password=password dbname=db_nra_wsl sslmode=disable"},
		//		{"postgres", "host=localhost port=5432 user=nra_wsl password=password dbname=db_nra_wsl"},
	}

	for _, conn := range conns {
		if status, err := Connect(conn.db, conn.ctx); err != nil {
			t.Errorf("Connection FAILED: %s", err)
		} else if status == 0 {
			t.Logf("Connection ESTABLISHED")
		} else {
			t.Errorf("Unknown state: status %d, err %s", status, err)
		}
	}

}
