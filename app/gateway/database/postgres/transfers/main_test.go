package transfer

import (
	"os"
	"stoneBanking/app/gateway/database/postgres/pgtest"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(pgtest.SetupTests(m))
}
