package v1

import (
	"context"
	"database/sql"
	"fmt"
	v1 "product-service/api"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ArticleDB is the API for Article.
type ArticleDB struct {
	dbo *MySQLService
}

// DbOperation is a API database operation
type DbOperation interface {
	Open(connstr string) (*sql.DB, error)
	Connect(context.Context) (*sql.Conn, error)
	Close()
}

// MySQLService is MySql database service
type MySQLService struct {
	db *sql.DB
}

// Open the MySql database service
func (s *MySQLService) Open(ctx context.Context, connstr string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connstr)
	if err != nil {
		fmt.Printf("Open database fail:%v\n", err)
		return nil, status.Error(codes.Unknown, "Open database fail:\n"+err.Error())
	}
	s.db = db
	fmt.Print(db)
	return db, nil
}

// Connect the MySql database service
func (s *MySQLService) Connect(ctx context.Context) (*sql.Conn, error) {

	conn, err := s.db.Conn(ctx)
	if err != nil {
		fmt.Printf("Connect database fail:%v\n", err)
		return nil, status.Error(codes.Unknown, "Connect database fail:\n"+err.Error())
	}
	return conn, nil
}

// Close the MySql database service
func (s *MySQLService) Close() {
	s.db.Close()
}

// NewArticleDB creates ArticleDB service
func NewArticleDB(dbo *MySQLService) *ArticleDB {
	return &ArticleDB{dbo: dbo}
}

//Readall is a API for read the all Articles.
func (a *ArticleDB) Readall(ctx context.Context) []*v1.Article {
	dbconn, _ := a.dbo.Connect(ctx)
	defer dbconn.Close()
	rows, err := a.dbo.db.Query("SELECT * FROM artcile")
	fmt.Print(dbconn)
	defer rows.Close()
	//fmt.Print(rows)

	list := []*v1.Article{}
	var reminder time.Time
	for rows.Next() {
		fmt.Print(1)
		ar := new(v1.Article)

		if err := rows.Scan(&ar.Id, &ar.Title, &ar.Author, &ar.Category, reminder); err != nil {
			fmt.Print(2)
			return nil
		}
		ar.LastUpdated, err = ptypes.TimestampProto(reminder)
		if err != nil {
			return nil
		}
		list = append(list, ar)
		fmt.Print(ar.Id)
	}
	return list
}

//ReadForID is a API for read the Article.
func (a *ArticleDB) ReadForID(ctx context.Context, id int) *v1.Article {
	return nil
}

//DeletForID is a API for delete the Article.
func (a *ArticleDB) DeletForID(ctx context.Context, id int) {

}

//Insert is a API for add the Article.
func (a *ArticleDB) Insert(ctx context.Context, article *v1.Article) {

}

//Updata is a API for updata the Article.
func (a *ArticleDB) Updata(ctx context.Context, id int, article *v1.Article) {

}
