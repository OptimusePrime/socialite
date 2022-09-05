
build_server:
	cd ./apps/server && go build -o ../../dist/server_windows_amd64.exe

preview_server:
	./dist/server_windows_amd64.exe start

start_server:
	cd ./apps/server && go run . start

test_server:
	cd ./apps/server && go test -v ./...

format_server:
	cd ./apps/server && go fmt .

get_server:
	cd ./apps/server && go get $(pkg)

mod_tidy:
	cd ./apps/server && go mod tidy