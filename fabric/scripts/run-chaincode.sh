export CHANNEL_NAME=mychannel
export CC_NAME=trust
export CC_VERSION=1.0.0

#create channel
peer channel create -o orderer.tfs.at2chain.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/channel.tx --tls --cafile $ORDERER_CA

#peer0.org1.com join the channel
peer channel join -b mychannel.block

#update org1 anchor peer
peer channel update -o orderer.tfs.at2chain.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls true --cafile ${ORDERER_CA}

#peer0.org2.com join the channel
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/users/Admin@org2.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org2.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/peers/peer0.org2.tfs.at2chain.com/tls/ca.crt
peer channel join -b mychannel.block

#update org2 anchor peer
peer channel update -o orderer.tfs.at2chain.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls true --cafile ${ORDERER_CA}


#peer0.org3.com join the channel
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/users/Admin@org3.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org3.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/peers/peer0.org3.tfs.at2chain.com/tls/ca.crt
peer channel join -b mychannel.block

#update org3 anchor peer
peer channel update -o orderer.tfs.at2chain.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/${CORE_PEER_LOCALMSPID}anchors.tx --tls true --cafile ${ORDERER_CA}

#install chaincode on peer0.org1
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.tfs.at2chain.com/users/Admin@org1.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org1.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org1.tfs.at2chain.com/peers/peer0.org1.tfs.at2chain.com/tls/ca.crt

#peer chaincode install -p bitbucket.org/at2chain/chaincode-token -n $CC_NAME -v $CC_VERSION
peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION

#install chaincode on peer0.org2
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/users/Admin@org2.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org2.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org2.tfs.at2chain.com/peers/peer0.org2.tfs.at2chain.com/tls/ca.crt

#peer chaincode install -p bitbucket.org/at2chain/chaincode-token -n $CC_NAME -v $CC_VERSION
peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION

#install chaincode on peer0.org3
export CORE_PEER_MSPCONFIGPATH=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/users/Admin@org3.tfs.at2chain.com/msp
export CORE_PEER_ADDRESS=peer0.org3.tfs.at2chain.com:7051
export CORE_PEER_LOCALMSPID="Org3MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=/opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/org3.tfs.at2chain.com/peers/peer0.org3.tfs.at2chain.com/tls/ca.crt

#peer chaincode install -p bitbucket.org/at2chain/chaincode-token -n $CC_NAME -v $CC_VERSION
peer chaincode install -p bitbucket.org/at2chain/chaincode-court-file-cert/chaincode -n $CC_NAME -v $CC_VERSION

#instaince chaincode on peer0.org1
peer chaincode instantiate -n $CC_NAME --tls true --cafile ${ORDERER_CA} -c '{"Args":["init"]}' -v $CC_VERSION -C $CHANNEL_NAME -P "OR('Org1.member','Org2.member','Org3.member')"
