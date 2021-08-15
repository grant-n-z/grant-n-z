build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o GrantNZCacher ./cmd
	mkdir -p build/linux-amd64
	mv ./GrantNZCacher build/linux-amd64/
	cp ./entrypoint.sh build/linux-amd64/
	cp ./grant_n_z_cacher.yaml build/linux-amd64/
	cp ./grant_n_z_cacher.txt build/linux-amd64/

build-linux-arm:
	GOOS=linux GOARCH=arm go build -o GrantNZCacher ./cmd
	mkdir -p build/linux-arm
	mv ./GrantNZCacher build/linux-arm/
	cp ./entrypoint.sh build/linux-arm/
	cp ./grant_n_z_cacher.yaml build/linux-arm/
	cp ./grant_n_z_cacher.txt build/linux-arm/

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o GrantNZCacher ./cmd
	mkdir -p build/darwin-amd64
	mv ./GrantNZCacher build/darwin-amd64/
	cp ./entrypoint.sh build/darwin-amd64/
	cp ./grant_n_z_cacher.yaml build/darwin-amd64/
	cp ./grant_n_z_cacher.txt build/darwin-amd64/

unit-test:
	go test -v ./...
