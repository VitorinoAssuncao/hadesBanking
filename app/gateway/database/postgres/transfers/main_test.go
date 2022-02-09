package transfer

import (
	"log"
	"os"
	"stoneBanking/app/gateway/database/postgres/pgtest"
	"testing"
)

func TestMain(m *testing.M) {
	teardown, err := pgtest.SetupTests()
	if err != nil {
		log.Fatalf(err.Error())
	}

	defer teardown()
	os.Exit(m.Run())
}
