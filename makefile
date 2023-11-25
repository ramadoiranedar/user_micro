# .PHONY:

# Documentation HELPER
help:
	@echo "----------------------------------------------makefile user_micro----------------------------------------------"
	@echo "make: compile packages and dependencies"

	@echo "make make-gen: make folder /gen"
	@echo "make remove-gen: remove folder /gen"
	@echo "make remove-pkg: remove folder /pkg"
	@echo "make remove-cmd: remove folder /cmd"
	@echo "make remove-executable-server: remove file executable server at root directory /"

	@echo "make validate: validation OpenAPI swagger yaml file"

	@echo "make new-server-default: [ONCE] init for the first time generate server default main.go file"
	@echo "make new-server-custom: [ONCE] init for the first time generate server custom main.go file from ./main.go.dist"

	@echo "make gen: generate server and client after run the new-server"
	
	@echo "make gen-server-configureapi: generate server excluding main.go file with regenerate configureapi"
	@echo "make gen-server: generate server excluding main.go file without regenerate configureapi"

	@echo "make gen-client: generate client after generate server for this app and you can add more than one client for this"

	@echo "make clean: go cleaning and remove existing executable file server"

	@echo "make build: generate build file executable server"

	@echo "make run: serve REST API by the executable file server"
	@echo "make run-dev: run the server for development without build the file executable server"
	@echo "make all-dev: for easy generate all and run server for development"

	@echo "make run-api-doc: serve UI of API documenations Swagger OpenAPI"
	@echo "----------------------------------------------------------------------------------------------------------------"

make-gen:
	if [ ! -d gen ]; then mkdir gen; fi

remove-gen:
	if [ -d gen ]; then rm -R gen; fi

remove-pkg:
	if [ -d pkg ]; then rm -R pkg; fi

remove-cmd:	
	if [ -d cmd ]; then rm -R cmd; fi

remove-executable-server:
	rm -rf user-micro-server

validate:
	swagger validate ./api/swagger.yaml

new-server-default: validate make-gen
	swagger generate server --main-package=../../cmd/user-micro-server -A user-micro-server -t ./gen  -f ./api/swagger.yaml --principal models.Principal

new-server-custom: validate make-gen
	swagger generate server --main-package=../../cmd/user-micro-server -A user-micro-server -t ./gen  -f ./api/swagger.yaml --principal models.Principal
	cp main.go.custom cmd/user-micro-server/main.go

gen-custom: gen-server gen-client

gen: validate make-gen gen-client
	swagger generate server --exclude-main -A user-micro-server -t ./gen  -f ./api/swagger.yaml --principal models.Principal
	cp configure_user_micro_server.go.custom  gen/restapi/configure_user_micro_server.go

# Example:
#		swagger generate client -A user-micro-server -f ./api/swagger.yaml -c pkg/client -m ./gen /models --principal models.Principal
# 		swagger generate client -A for-example1-server -f ./api/example1.yaml -c pkg/client -m ./gen /model
# 		swagger generate client -A for-example2-server -f ./api/example2.yaml -c pkg/client -m ./gen /model
gen-client: validate	
	swagger generate client -A user-micro-server -f ./api/swagger.yaml -c pkg/client/user_micro_api_client -m ./gen/models --principal models.Principal
#	swagger generate client -A bpjs-server -f ./api/bpjs.yaml -c pkg/client/bpjs_client -m ./gen/model --principal models.Principal

clean: remove-executable-server
	go clean -i .

build:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/user-micro-server

run:
	./user-micro-server --port=8080 --host=0.0.0.0 --config="./config/app.yaml"

run-dev:
	go run ./cmd/user-micro-server/main.go --config=./config/app.yaml --host=0.0.0.0 --port=8080

all-dev: gen build run-dev

run-api-doc: validate
	swagger serve api/swagger.yaml --no-open --host=0.0.0.0 --port=8081 --base-path=/

# TODO: make for development/stagging/production
# ...
