@echo off
protoc --proto_path=. --proto_path=%USERPROFILE%\go\include --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative test.proto
