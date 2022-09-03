dev_server:
	cd ./apps/server && go run .

build_server:
	cd ./apps/server && go build -o ../../dist/server_windows_amd64.exe .

preview_server:
	./dist/server_windows_amd64.exe

test_server:
	cd ./apps/server && go test -v ./...