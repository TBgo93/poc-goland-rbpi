SHELL := pwsh.exe


run:
	go run main.go

re-build:
	go build main.go

compile:
	pwsh -Command { $$Env:GOOS = 'linux' ; $$Env:GOARCH = 'arm64' ; go build -o build/main main.go }