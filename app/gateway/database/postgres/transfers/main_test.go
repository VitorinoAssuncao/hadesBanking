package transfer

import (
	"os"
	"stoneBanking/app/gateway/database/postgres/pgtest"
	"testing"
)

func TestMain(m *testing.M) {
	teardown := pgtest.SetupTests()
	defer teardown()
	os.Exit(m.Run())
}
