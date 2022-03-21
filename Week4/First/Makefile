gen:
	# protoc --go_out=pb/ --go_opt=paths=source_relative proto/processor_message.proto
	# protoc --go_out=pb/ --go_opt=paths=source_relative proto/*.proto
	protoc --go_out=pb/ --go_opt=paths=source_relative --proto_path=proto proto/*.proto
clean:
	rm pb/*.go
run:
	go run main.go



# export PATH="$PATH:$(go env GOPATH)/bin"