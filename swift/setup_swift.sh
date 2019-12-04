#创建swift的宿主机绑定目录
mkdir -p $PWD/swift_data/1/node/sdb1 && mkdir -p $PWD/swift_data/1/node/sdb5
#运行swift docker 镜像
docker run -d  -v "$PWD/swift_data":/srv -p 5000:5000 -p 35357:35357 -p 8080:8080 --name swift_docker jeantil/openstack-keystone-swift:pike
#恢复swift 用户对srv目录权限
docker exec -it swift_docker chown -R swift:swift /srv/1/
#绑定swift 外网地址
docker exec -it swift_docker /swift/bin/register-swift-endpoint.sh http://192.168.44.129:8080/