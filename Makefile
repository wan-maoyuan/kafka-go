# 程序名称
NAME = kafka-go

# 版本号
VERSION = v1.0.0

# 目标输出目录
DIST_FOLDER = dist

# 版本构建目录
RELEASE_FOLDER = release

# protoc 文件夹
API_PROTO_FILES = $(shell find api -name *.proto)

# 构建附加选项
BUILD_OPTS := -ldflags "-s -w -X 'main.Version=${VERSION}'"

# 编译环境
BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64


# proto文件代码生成
.PHONY: api
api:
	protoc --proto_path=. 														\
 	       --go_out=paths=source_relative:. 									\
 	       --go-http_out=paths=source_relative:. 								\
 	       --go-grpc_out=paths=source_relative:. 								\
		   --go-errors_out=paths=source_relative:. 								\
	       $(API_PROTO_FILES)


.PHONY: build
build:
	${BUILD_ENV} go build ${BUILD_OPTS} -o ${DIST_FOLDER}/${NAME}/${NAME}


# 清理
.PHONY: clean
clean:
	-rm -rf $(DIST_FOLDER)/*
	-rm -f ${TEST_REPORT}
	-go clean 
	-go clean -cache