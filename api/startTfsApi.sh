#!/usr/bin/env bash

java -jar tfs-api-app.jar\
 --spring.profiles.active=dev\
 --spring.datasource.url="jdbc:mysql://114.118.22.200:3306/tfs_api?useUnicode=true&characterEncoding=utf-8&useSSL=false"\
 --spring.datasource.username=root\
 --spring.datasource.password=At2plus#@212!\
 #--fabric.network.config-path=/opt/fabric-network-config.yaml\
 --fabric.chain-code.name=court\
 --fabric.chain-code.version=1.0.0\
 --openstack.swift.host=http://106.14.114.220:35357/v3