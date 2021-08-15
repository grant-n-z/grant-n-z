build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o GrantNZServer ./cmd
	mkdir -p build/linux-amd64
	mv ./GrantNZServer build/linux-amd64/
	cp ./entrypoint.sh build/linux-amd64/
	cp ./grant_n_z_server.yaml build/linux-amd64/
	cp ./grant_n_z_server.txt build/linux-amd64/

build-linux-arm:
	GOOS=linux GOARCH=arm go build -o GrantNZServer ./cmd
	mkdir -p build/linux-arm
	mv ./GrantNZServer build/linux-arm/
	cp ./entrypoint.sh build/linux-arm/
	cp ./grant_n_z_server.yaml build/linux-arm/
	cp ./grant_n_z_server.txt build/linux-arm/

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o GrantNZServer ./cmd
	mkdir -p build/darwin-amd64
	mv ./GrantNZServer build/darwin-amd64/
	cp ./entrypoint.sh build/darwin-amd64/
	cp ./grant_n_z_server.yaml build/darwin-amd64/
	cp ./grant_n_z_server.txt build/darwin-amd64/

unit-test:
	go test -v ./...

e2e-test:
	go run ../e2e/cmd/cmd.go
