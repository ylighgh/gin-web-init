#!/bin/sh

# 获取命令行输入的提交消息
if [ $# -eq 0 ]; then
  echo "请输入提交消息"
  exit 1
fi
commit_message="$*"

# 提交和推送Git更改
git add -A
git commit -m "$commit_message"
git push

# 获取当前git提交的短commit ID
commit_id=$(git rev-parse --short=8 HEAD)

# 定义Docker镜像的名称
image_name="harbor.axzo.cn/kube-ops/gin-web-init:$commit_id"

# 使用Docker构建镜像
docker build -t "$image_name" .

# 检查Docker构建是否成功
if [ $? -eq 0 ]; then
  echo "Docker镜像构建成功"
else
  echo "Docker镜像构建失败"
  exit 1
fi

docker push "$image_name"

# 检查推送镜像是否成功
if [ $? -eq 0 ]; then
  echo "推送镜像成功"
else
  echo "推送镜像失败"
  exit 1
fi

#
# 输出commit ID
echo "Commit ID: $commit_id"
