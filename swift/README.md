# How to start openstack swift dev env

```bash
docker run -d --rm  -p 5000:5000 -p 35357:35357 -p 8080:8080 --name keystone jeantil/openstack-keystone-swift:pike

docker exec -it keystone /swift/bin/register-swift-endpoint.sh http://106.14.114.220:8080/

docker exec -it keystone /swift/bin/register-swift-endpoint.sh http://106.14.114.220:35357/

docker exec -it keystone /swift/bin/register-swift-endpoint.sh http://106.14.114.220:5000/
```