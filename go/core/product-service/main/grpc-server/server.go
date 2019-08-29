package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"product-service/protocol/rest"
	grpc "product-service/protocol/rpc"
	v1 "product-service/service"
)

func main() {
	ctx := context.Background()

	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s",
	// 	"root",
	// 	"123456",
	// 	"localhost:3306",
	// 	"product",
	// 	"parseTime=true")

	// mydb := &v1.MySQLService{}

	// mydb.Open(ctx, dsn)

	// artdb := v1.NewArticleDB(mydb)
	// fmt.Println(artdb.Readall(ctx))

	db, err := sql.Open("mysql", "root:123456@/product?charset=utf8")
	fmt.Print(db)
	if err == nil {
		fmt.Print("1\n")
		rows, _ := db.Query("select title FROM artcile;")
		fmt.Print(rows)
		var title string
		if err := rows.Scan(title); err != nil {

		}
		fmt.Print(title)

		rows.Close()
	}

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
