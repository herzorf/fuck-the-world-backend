#!/bin/bash

# 服务器信息
SERVER_IP="42.192.105.150"
DEPLOY_PATH="/root/backend/fuck-the-world"

docker-compose build
docker save -o fuck-the-world.tar fuck-the-world
scp fuck-the-world.tar docker-compose.yml config.production.yaml root@$SERVER_IP:$DEPLOY_PATH
# 在服务器上执行部署
ssh root@$SERVER_IP << EOF
  cd $DEPLOY_PATH
  # 关闭并移除旧容器
  docker-compose down

  # 加载新的 Docker 镜像
  docker load -i fuck-the-world.tar

  # 重新启动容器
  docker-compose up -d
EOF

rm -f fuck-the-world.tar

echo "部署完成！"