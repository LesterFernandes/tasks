module github.com/LesterFernandes/tasks/api-gateway

go 1.24.5

require (
	github.com/LesterFernandes/tasks/shared-protos v0.0.0-20250822172112-000e73a5de4d
	github.com/go-chi/chi/v5 v5.2.2
	github.com/rs/zerolog v1.34.0
	google.golang.org/grpc v1.75.0
)

require (
	github.com/LesterFernandes/tasks/utils v0.0.0-00010101000000-000000000000 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	golang.org/x/net v0.41.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.26.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250707201910-8d1bb00bc6a7 // indirect
	google.golang.org/protobuf v1.36.8 // indirect
)

replace github.com/LesterFernandes/tasks/utils => ../utils
