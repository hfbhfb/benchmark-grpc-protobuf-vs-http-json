

# 定义函数
build-docker(){

# 编译二进制服务器
go build

# 制作镜像
cat >Dockerfile_grpc <<EOF
FROM centos:centos7.6.1810
MAINTAINER hfb
ADD ./app-server-grpc /app-server-grpc
CMD [ "sh", "-c", "/app-server-grpc" ]
EOF

IMageName=benchmark-grpc
Version=v0.1
docker build -f ./Dockerfile_grpc -t $IMageName:$Version .

docker tag $IMageName:$Version docker.io/hefabao/$IMageName:$Version

### 上传到docker hub官方镜像仓库
docker push docker.io/hefabao/$IMageName:$Version

}


build-docker



