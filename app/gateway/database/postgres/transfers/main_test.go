package transfer

import (
	"os"
	"stoneBanking/app/gateway/database/postgres/pgtest"
	"testing"

	"github.com/jackc/pgx/v4/pgxpool"
)

var testPool *pgxpool.Pool

func TestMain(m *testing.M) {
	os.Exit(pgtest.SetupTests(m))
}
