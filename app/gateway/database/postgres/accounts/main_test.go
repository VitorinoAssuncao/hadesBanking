package account

import (
	"os"
	"testing"

	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func TestMain(m *testing.M) {
	teardown := pgtest.SetupTests()
	defer teardown()
	os.Exit(m.Run())
}
