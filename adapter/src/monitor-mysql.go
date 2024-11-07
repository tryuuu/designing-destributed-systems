// HTTPサーバーを立ち上げ、リクエストが来ると指定されたSQLクエリをMySQLデータベースで実行し、その結果を返す
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// コマンドライン引数・デフォルト値・説明
var (
	user   = flag.String("user", "", "The database user name")
	passwd = flag.String("password", "", "The database password")
	dtbs   = flag.String("database", "", "The database to connect to")
	query  = flag.String("query", "", "The test query")
	addr   = flag.String("address", "localhost:8080", "The address to listen on")
)

// 使用例:
// db-check -query="SELECT * from test-table" \
//          -user="root" \
//			-password="password" \
//			-database="test-db" \

func main() {
	// 各フラグのポインタに値を格納
	flag.Parse()
	// データベースに接続
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", *user, *passwd, *dtbs))
	// 接続に失敗した場合
	if err != nil {
		fmt.Printf("Error opening database: %v", err)
	}
	// クエリを実行するハンドラ
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// クエリを実行
		_, err := db.Exec(*query)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte(err.Error()))
			return
		}
		res.WriteHeader(http.StatusOK)
		res.Write([]byte("OK"))
		return
	})
	http.ListenAndServe(*addr, nil)
}
