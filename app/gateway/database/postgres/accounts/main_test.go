package account

import (
	"os"
	"testing"

	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func TestMain(m *testing.M) {
	os.Exit(pgtest.SetupTests(m))
}
