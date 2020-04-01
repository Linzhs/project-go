# include environment variables
include .env

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --slient

# exec: Run given command, wrapped with custom GOPATH. e.g; make exec run="go test ./..."
exec:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) $(run)

start:
	bash -c "trap 'make stop' EXIT; $(MAKE) compile start-server watch run='make compile start-server'"

stop: stop-server

# @关闭回声 @-关闭回声并忽略错误
compile:
	@-touch $(STDERR)
	@-rm $(STDERR)
	@-$(MAKE) -s go-compile 2> $(STDERR)
	@cat $(STDERR) | sed -e '1s/.*/\nError:\n/' | sed 's/make\[.*/ /' | sed "/^/s/^/     /" 1>&2

start-server:
	@echo "  >  $(PROJECTNAME) is available at $(ADDR)"
	@-$(GOBIN)/$(PROJECTNAME) 2>&1 & echo $$! > $(PID)
	@cat $(PID) | sed "/^/s/^/  \>  PID: /"

stop-server:
	@-touch $(PID)
	@-kill `cat $(PID)` 2> /dev/null || true
	@-rm $(PID)

restart-server: stop-server start-server

## watch: Run given command when code changes. e.g; make watch run="echo 'hey'"
watch:
	@yolo -i . -e vendor -e bin -c $(run)

install: go-get

gogoprotoc: go-get $(GOGOPROTOC)

gogo-generate:
	# Generate gogo, gRPC-Gateway, swagger, go-validators output.
	#
	# -I declares import folders, in order of importance
	# This is how proto resolves the protofile imports.
	# It will check for the protofile relative to each of these
	# folders and use the first one it finds.
	#
	# --gogo_out generates GoGo Protobuf output with gRPC plugin enabled.
	# --grpc-gateway_out generates gRPC-Gateway output.
	# --swagger_out generates an OpenAPI 2.0 specification for our gRPC-Gateway endpoints.
	# --govalidators_out generates Go validation files for our messages types, if specified.
	#
	# The lines starting with Mgoogle/... are proto import replacements,
	# which cause the generated file to import the specified packages
	# instead of the go_package's declared by the imported protof files.
	#
	# $$GOPATH/src is the output directory. It is relative to the GOPATH/src directory
	# since we've specified a go_package option relative to that directory.
	#
	# proto/example.proto is the location of the protofile we use.
	protoc \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/ \
		-I vendor/github.com/gogo/googleapis/ \
		-I vendor/ \
		--gogo_out=plugins=grpc,\
			Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,\
			Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,\
			Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
			Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
			Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
			Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
			Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types:\
			Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:.\
#     			--grpc-gateway_out=allow_patch_feature=false,\
#		 Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
#		 Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
#		 $GOPATH/src/ \
#				 --swagger_out=third_party/OpenAPI/ \
#				 --govalidators_out=gogoimport=true,\
#		 Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,\
#		 Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,\
#		 Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:\
#		 $GOPATH/src
	pkg/rpc/pb/*.proto
		# Workaround for https://github.com/grpc-ecosystem/grpc-gateway/issues/229.
        #sed -i.bak "s/empty.Empty/types.Empty/g" proto/example.pb.gw.go && rm proto/example.pb.gw.go.bak
        # Generate static assets for OpenAPI UI
        #statik -m -f -src third_party/OpenAPI/

go-compile: go-clean go-get go-build

go-build:
	@echo "  >  Building binary..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go build -o $(GOBIN)/$(PROJECTNAME) $(GOFILES)

go-generate:
	@echo "  >  Generating dependency files..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go generate $(generate)

go-get:
	@echo "  >  Checking if there is any missing dependencies..."
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go get -v $(get)

go-install:
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go install $(GOFILES)

go-clean:
	@echo "  >  Cleaning build cache"
	@GOPATH=$(GOPATH) GOBIN=$(GOBIN) go clean

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo