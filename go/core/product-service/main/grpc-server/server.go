package main

import (
	"context"
	"fmt"
	"os"

	"product-service/protocol/rest"
	grpc "product-service/protocol/rpc"
	v1 "product-service/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
		"root",
		"123456",
		"localhost:3306",
		"product",
		"parseTime=true")

	mydb := &v1.MySQLService{}

	mydb.Open(ctx, dsn)

	artdb := v1.NewArticleDB(mydb)
	fmt.Println(artdb.Readall(ctx))

	// db, _ := sql.Open("mysql", "root:123456@tcp(localhost:3306)/test?charset=utf8")
	// conn, _ := db.Conn(ctx)

	// rows, _ := conn.QueryContext(ctx, `SELECT test_name FROM test_tbl;`)
	// rows.Next()
	// var testName string
	// if err := rows.Scan(&testName); err != nil {

	// }
	// fmt.Print(testName)

	// run HTTP gateway
	go func() {
		_ = rest.RunServer(ctx, "2536", "8080")
	}()

	server := v1.NewArtManagerServer()
	if err := grpc.RunServer(ctx, server, "2536"); err != nil {
		fmt.Fprintln(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
