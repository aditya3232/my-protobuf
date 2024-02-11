# Change this to your own Go module
GO_MODULE := my-protobuf


.PHONY: tidy
tidy:
	go mod tidy


.PHONY: clean
clean:
ifeq ($(OS), Windows_NT)
	if exist "protogen" rd /s /q protogen	
else
	rm -fR ./protogen
endif


.PHONY: protoc
# (un)comment or add folder that contains proto as required
protoc:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	./proto/basic/*.proto \
	./proto/dummy/*.proto \
	./proto/jobsearch/*.proto \

# digunakan untuk hapus folder protogen, compile file proto, dan download dependency go
.PHONY: build
build: clean protoc tidy

# menjalankan main.go
.PHONY: run
run:
	go run main.go

#  menjalankan build & run
.PHONY: execute
execute: build run


.PHONY: protoc-validate
protoc-validate:
	protoc --validate_out="lang=go:./generated" --go_opt=module=${GO_MODULE} --go_out=. ./proto/car/*.proto
