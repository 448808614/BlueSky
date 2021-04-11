
# 安装Protobuf-Go
go get github.com/golang/protobuf/proto
go get -u github.com/golang/protobuf/protoc-gen-go
# gogo-proto使用了golang-protobuf的库

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
# Windows下没有mv命令，请手动重命名文件夹
