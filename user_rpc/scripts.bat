@REM 根据数据库生成proto文件, 并将重定向的文件编码改为utf8
sql2pb -go_package ./pb -host localhost -package pb -password 101325 -port 3306 -schema gozero_study -service_name usercenter -user root | out-file ./pb/usercenter.proto -encoding utf8 

@REM 这一步之前必须要把该项目加入到go.work中
goctl rpc protoc ./pb/user.proto --go_out=./ --go-grpc_out=./ --zrpc_out=./ --style=go_zero

protoc -I ./ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./pb/userModel.proto
