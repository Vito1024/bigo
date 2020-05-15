package benchmark

import (
	"bigo/client"
	"bigo/model"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

var (
	host     = "localhost"
	port     = "3306"
	dbname   = "user"
	user     = "mpb"
	password = "127.0.0.1"
	table    = "userinfo"
	schema   = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname)
)

var (
	db, err = sql.Open("mysql", schema)
	cli  = client.NewClient("localhost:8080")
)

func BenchmarkMySQL(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// select
		_sql := fmt.Sprintf(`SELECT * FROM %s WHERE id=?`, table)
		_, err = db.Query(_sql, "10000001")
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// insert
		_sql = fmt.Sprintf("INSERT INTO %s (id, name, age, gender, address, hobbies) VALUES(?, ?, ?, ?, ?, ?)", table)
		_, err = db.Exec(_sql, "10000006", "Adam", 21, "M", "New York", "programming")
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// update
		_sql = fmt.Sprintf("UPDATE %s SET age = ? WHERE id = ?", table)
		_, err = db.Exec(_sql, 22, "10000006")
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// delete
		_sql = fmt.Sprintf("DELETE FROM %s WHERE id = ?", table)
		_, err = db.Exec(_sql, "10000006")
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}
	}
}


func BenchmarkBigo(b *testing.B) {
	bq := model.BigoRequest{}
	for i := 0; i < b.N; i++ {
		// GET
		bq.CommandName = "HGET"
		bq.Args = []string{"user:10000001"}
		err = cli.SendCommand(bq)
		_, err = cli.ReadResponse()
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// SET
		bq. CommandName = "HSET"
		bq.Args = []string{"user:10000006", "name", "Adam", "age", "21", "gender", "M", "address", "New York", "hobbies", "programming"}
		err = cli.SendCommand(bq)
		_, err = cli.ReadResponse()
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// HSETFIELD
		bq.CommandName = "HSETFIELD"
		bq.Args = []string{"user:10000006", "age", "22"}
		err = cli.SendCommand(bq)
		_, err = cli.ReadResponse()
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}

		// DEL
		bq.CommandName = "DEL"
		bq.Args = []string{"user:10000006"}
		err = cli.SendCommand(bq)
		_, err = cli.ReadResponse()
		if err != nil {
			b.Fatal(err)
			b.Fail()
		}
	}
}