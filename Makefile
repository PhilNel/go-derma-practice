SPECIALS_CMD_DIR=cmd/specials-handler
LAMBDA_NAME=go-derma-specials-handler
BINARY_NAME=bootstrap
BUCKET_NAME=nelskin-practice-dev-assets-af-south-1

.PHONY: run build vendor test fmt package upload deploy clean

run:
	go run $(SPECIALS_CMD_DIR)/main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BINARY_NAME) ./$(SPECIALS_CMD_DIR)

vendor:
	go mod tidy && go mod vendor

test:
	go test ./...

fmt:
	go fmt ./...

package: build
	zip -j $(LAMBDA_NAME).zip $(BINARY_NAME)

upload:
	aws s3 cp $(LAMBDA_NAME).zip s3://$(BUCKET_NAME)/$(LAMBDA_NAME).zip

deploy: package upload

clean:
	rm -f $(BINARY_NAME) $(LAMBDA_NAME).zip


