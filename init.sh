
# 安装Protobuf-Go
go get github.com/gogo/protobuf/proto
go get github.com/gogo/protobuf/protoc-gen-gogo
go get github.com/gogo/protobuf/protoc-gen-gofast
go get github.com/buger/jsonparser

mkdir src
cd src

# 初始化Protobuf
mkdir google.golang.org
cd google.golang.org
git clone https://github.com/protocolbuffers/protobuf-go.git
mv protobuf-go protobuf
