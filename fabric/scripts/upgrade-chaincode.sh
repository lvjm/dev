#!/bin/bash

export CHANNEL_NAME=mychannel
export CC_NAME=trust
export CC_VERSION=$1


#export CHANNEL_NAME=mychannel
#export CC_NAME=trust
#export CC_VERSION=1.0.1

#install chaincode on peer0.org1
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.tfs.at2chain.com/users/Admin@org1.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org1.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.tfs.at2chain.com/peers/peer0.org1.tfs.at2chain.com/tls/ca.crt
peer chaincode install -p bitbucket.org/at2chain/chaincode-trust/chaincode -n $CC_NAME -v $CC_VERSION
#peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION
#peer chaincode install -p bitbucket.org/at2chain/chaincode-token -n $CC_NAME -v $CC_VERSION

#install chaincode on peer0.org2
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/users/Admin@org2.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org2.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/peers/peer0.org2.tfs.at2chain.com/tls/ca.crt
peer chaincode install -p bitbucket.org/at2chain/chaincode-trust/chaincode -n $CC_NAME -v $CC_VERSION
#peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION
#peer chaincode install -p bitbucket.org/at2chain/chaincode-token -n $CC_NAME -v $CC_VERSION

#install chaincode on peer0.org3
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/users/Admin@org3.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org3.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/peers/peer0.org3.tfs.at2chain.com/tls/ca.crt
peer chaincode install -p bitbucket.org/at2chain/chaincode-trust/chaincode -n $CC_NAME -v $CC_VERSION
#peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION

#peer chaincode upgrade -o orderer.tfs.at2chain.com:7050 -n $CC_NAME --tls true --cafile ${ORDERER_CA} -c '{"Args":["init"]}' -v $CC_VERSION -C $CHANNEL_NAME -P "OR('Org1MSP.member','Org2MSP.member','Org3MSP.member')"
peer chaincode upgrade -o orderer.tfs.at2chain.com:7050 -n $CC_NAME --tls true --cafile ${ORDERER_CA} -c '{"Args":["init"]}' -v $CC_VERSION -C $CHANNEL_NAME -P "OutOf(1,'Org1MSP.member','Org2MSP.member','Org3MSP.member')"