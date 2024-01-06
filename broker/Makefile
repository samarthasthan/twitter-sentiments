# Makefile

gen:
	@echo "Generating Go code from Protocol Buffers..."
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/tweet.proto
	@echo "Go code generation completed."

clean:
	@echo "Cleaning generated Go code..."
	@rm -f proto/*.go
	@echo "Clean completed."