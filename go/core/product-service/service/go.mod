module product-service/service

go 1.12

require (
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	google.golang.org/grpc v1.21.1
	product-service/api v0.0.0-00010101000000-000000000000
)

replace (
	cloud.google.com/go => github.com/googleapis/google-cloud-go v0.44.3
	cloud.google.com/go/datastore => github.com/googleapis/google-cloud-go/datastore v1.0.0
	golang.org/x/build => github.com/golang/build v0.0.0-20190822182523-83da5c54f18b
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190820162420-60c769a6c586
	golang.org/x/exp => github.com/golang/exp v0.0.0-20190731235908-ec7cb31e5a56
	golang.org/x/image => github.com/golang/image v0.0.0-20190823064033-3a9bac650e44
	golang.org/x/lint => github.com/golang/lint v0.0.0-20190409202823-959b441ac422
	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20190826170111-cafc553e1ac5
	golang.org/x/mod => github.com/golang/mod v0.1.0
	golang.org/x/net => github.com/golang/net v0.0.0-20190813141303-74dc4d7220e7
	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/perf => github.com/golang/perf v0.0.0-20190823172224-ecb187b06eb0
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190826190057-c7b8b68b1456
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/time => github.com/golang/time v0.0.0-20190308202827-9d24e82272b4
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190826234050-71894ab67ee3
	golang.org/x/xerrors => github.com/golang/xerrors v0.0.0-20190717185122-a985d3407aa7
	google.golang.org/api => github.com/googleapis/google-api-go-client v0.9.0
	google.golang.org/appengine => github.com/golang/appengine v1.6.1
	google.golang.org/genproto => github.com/google/go-genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc => github.com/grpc/grpc-go v1.23.0
	product-service/api => ../api
)
