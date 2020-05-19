package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"strconv"
	"database/sql"
	"bigo/client"
	"bigo/model"
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

func InsertIntoMySQL() {
	db, err := sql.Open("mysql", schema)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	for i := 0; i < 10000; i++ {
		_sql := fmt.Sprintf(`INSERT INTO userinfo(id, name, age, gender, address, hobbies) VALUES(%d, "Adam", 20, "M", "Beijing", "programming")`, i)
		if _, err = db.Exec(_sql); err != nil {
			fmt.Println(err, "id=", i)
			return
		}
		fmt.Println(i, "th inserted")
	}
}

func InsertIntoBigo() {
	cli := client.NewClient("localhost:8080")
	defer cli.Conn.Close()

	bigoRequest := model.BigoRequest{
		CommandName: "HSET",
		Args:        []string{"", "name", "Adam", "age", "20", "gender", "M", "address", "Beijing", "hobbies", "programming"},
	}

	for i := 0; i < 10000; i++ {
		bigoRequest.Args[0] = strconv.Itoa(i)
		if err := cli.SendCommand(bigoRequest); err != nil {
			fmt.Println(err)
			return
		}

		if _, err := cli.ReadResponse(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(i, "th inserted")
	}
}


func main() {
	InsertIntoBigo()
}
