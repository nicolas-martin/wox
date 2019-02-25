# Go parameters
BINARY_NAME=wox

all: test build
build: 
		go build -o bin/$(BINARY_NAME)
test: 
		go test -v ./...
clean: 
		go clean
		rm -f $(BINARY_NAME)
run:
		go build -o $(BINARY_NAME) -v ./...
		./$(BINARY_NAME)
news3:
		aws s3api create-bucket \
		--bucket wox-test \
		--acl public-read 
#--region ca-central-1 
# --create-bucket-configuration LocationConstraint=ca-central-1

pubpolicy:
		aws s3api put-bucket-policy \
		--bucket wox-test \
		--policy file://pubpolicy.json