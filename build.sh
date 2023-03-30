# 如果vendor文件夹存在，就移动为/go
#if [ -d "vendor" ]; then
#  rm -rf /go
#  mv vendor /go
#fi
echo "/go 下的内容为:"
ls /go
go mod tidy
echo "开始编译"
go build -o ./temp/linux_amd64/main .
echo "编译完成"
#rm -rf vendor
#mv /go vendor

