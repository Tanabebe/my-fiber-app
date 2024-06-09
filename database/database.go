package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type MyDb struct {
	Db *bun.DB
}

var DB MyDb

func Connect() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	fmt.Println(dsn)

	sqldb, err := sql.Open(os.Getenv("DB_DRIVER"), dsn)

	if err != nil {
		log.Fatal("database no connect", err)
	}

	database := bun.NewDB(sqldb, mysqldialect.New())

	database.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	DB = MyDb{
		Db: database,
	}
}
