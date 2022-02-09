package account

import (
	"log"
	"os"
	"testing"

	"stoneBanking/app/gateway/database/postgres/pgtest"
)

func TestMain(m *testing.M) {
	teardown, err := pgtest.SetupTests()
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer teardown()
	os.Exit(m.Run())
}
