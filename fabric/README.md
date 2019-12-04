## 智能合约接口调用示例

AddRecord
参数(bizId,fileHash,creationDate,metadataJson)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["AddRecord","bizId","fileHashOriginal","creationDate","{\"storage\":\"BSEvent\",\"dataUri\":\"path\",\"preFileKey\":\"\",\"fileType\":\"jpg\",\"org\":\"BS\",\"uploader\":\"mm\",\"lawyer\":\"lawyerWang\",\"status\":\"0\",\"description\":\"divorce\"}"]}'
```
OriginalFileKeyIdSearch
参数（originalFileKeyId）
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["OriginalFileKeyIdSearch","fileHashOriginal-ef86f997b92d0dc6a274cf3229a46b13abf4e014562116c7824b3bb380c0dff5"]}'
```
GetRecord
参数：(keyId,operator)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["GetRecord","fileHashOriginal-b18819724df5124bfb6743f38fca1fa9ffc1f09a1bb4eceacf4d61c7411796ac-fileHash1-4b78b1165319422ced8a4cd3200b1b36398d96e50698930682c51f5dc13a2242","operator"]}'
```
GetAttestation
参数：(fileHash,keyId,operator)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust02 -c '{"Args":["GetAttestation","fileHash","fileHash-org-operator-e3c6546c27a527652f8cc87a5987288697fb2cbb60fbfc8053b39db0dcec8329","operator"]}'
```
AddEvent
参数：(keyId, eventName, operator, timestamp)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust02 -c '{"Args":["AddEvent","fileHash-org-operator-be487469ee478b6528f625dd8c60f3edcd3fa445c10c0de26ccd6cc2da9e6746","eventName","operator","timestamp"]}'
```
SearchEvent
参数：(keyId,operator)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["SearchEvent","fileHashOriginal-ef86f997b92d0dc6a274cf3229a46b13abf4e014562116c7824b3bb380c0dff5","2",""]}'
```
Archive
参数：(keyId,operator)
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["Archive","fileHashOriginal-ef86f997b92d0dc6a274cf3229a46b13abf4e014562116c7824b3bb380c0dff5","operator"]}'
```
Search
参数（queryString,operator）
```
peer chaincode invoke -o orderer.tfs.at2chain.com:7050  --tls true --cafile ${ORDERER_CA} -C mychannel -n trust -c '{"Args":["Search","{\"selector\":{\"$and\":[{\"bizId\":\"string\"},{\"creationDate\":{\"$gte\":\"1547270552\"}},{\"creationDate\":{\"$lte\":\"1547270552\"}}]}}","2",""]}'```
```

## 升级链码版本
```
#1.0.1 is the new version of chaincode
sh scripts/upgrade-chaincode.sh 1.0.1

```

##清理链数据
```
 # stop the compose if not yet
 docker-compose down
 
 # remove all containers
 docker rm $(docker ps -a -q)
 system docker prune

 #remove all images
 docker rmi -f $(docker images -q)
 
 #remove chain data
 rm -rf /mnt/fabric
 
 # reboot
 sudo reboot
```