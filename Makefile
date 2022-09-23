create:
	protoc --go_out=. --go_opt=paths=source_relative \
        --go-grpc_out=. --go-grpc_opt=paths=source_relative \
        proto/rusprofile.proto
	protoc -I . --grpc-gateway_out ./ \
        --grpc-gateway_opt logtostderr=true \
        --grpc-gateway_opt paths=source_relative \
        --grpc-gateway_opt generate_unbound_methods=true \
        proto/rusprofile.proto

clean:
	rm proto/*.go

run:
	GOOS=linux go build -a -tags netgo -o rusprofile cmd/*.go
	docker build -t syalsr .
	docker run -p 8080:8080 -p 8081:8081 -p 9000:9000 syalsr

swag:
	swag init -g cmd/*.go