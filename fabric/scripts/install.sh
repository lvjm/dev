#!/bin/bash
read -p "please input organization count,default:2  " orgCount
orgCount=${orgCount:-2}
read -p "please input peer count in a organization default:2  " peerCountInOrg
peerCountInOrg=${peerCountInOrg:-2}
read -p "please input zookeeper count,default:3  " zkCount
zkCount=${zkCount:-3}
read -p "please input kafka count,default:4  " kafkaCount
kafkaCount=${kafkaCount:-3}
read -p "please input fabric domain name,default:at2chain.com  " domainName
domainName=${domainName:-"at2chain.com"}
read -p "please input fabric network name, " fabricNetworkName
read -p "please input ca username in Organization, default:admin  " registrarUserName
registrarUserName=${registrarUserName:-"admin"}
read -p "please input ca password in Organization, default:adminpw  " registrarUserPassword
registrarUserPassword=${registrarUserPassword:-"adminpw"}
read -p "please input username of couchdb  " couchdbUserName
read -p "please input password of couchdb  " couchdbUserPassword
read -p "please input channel name  " channelName

echo 'generate docker-compose.yaml'
python docker-compose-generate.py --orgCount ${orgCount} --zkCount ${zkCount} --kafkaCount ${kafkaCount} --peerCountInOrg ${peerCountInOrg} --fabricNetworkName ${fabricNetworkName} --couchdbUserName ${couchdbUserName} --couchdbUserPassword ${couchdbUserPassword} --domainName ${domainName} --registrarUserName ${registrarUserName} --registrarUserPassword ${registrarUserPassword}

echo 'generate network-config.yaml'
python network-config-generate.py --orgCount ${orgCount} --peerCountInOrg ${peerCountInOrg} --domainName ${domainName} --registrarUserName ${registrarUserName} --registrarUserPassword ${registrarUserPassword} --channelName ${channelName}



