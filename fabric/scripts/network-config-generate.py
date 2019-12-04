#!/usr/bin/python
import yaml
import os
from collections import OrderedDict
import argparse

class CustomDumper(yaml.Dumper):

    def increase_indent(self, flow=False, indentless=False):
        return super(CustomDumper, self).increase_indent(flow, False)


def represent_none(self, _):
    return self.represent_scalar('tag:yaml.org,2002:null', '')

yaml.add_representer(type(None), represent_none)

class quoted(str):
    pass

def quoted_presenter(dumper, data):
    return dumper.represent_scalar('tag:yaml.org,2002:str', data, style='"')
yaml.add_representer(quoted, quoted_presenter)

def setup_yaml():
  represent_dict_order = lambda self, data:  self.represent_mapping('tag:yaml.org,2002:map', data.items())
  yaml.add_representer(OrderedDict, represent_dict_order)    
setup_yaml()

class blockStyleLiteral(str):
    pass

def repr_str(dumper, data):
    if '\n' in data:
        return dumper.represent_scalar(u'tag:yaml.org,2002:str', data, style='|')
    return dumper.org_represent_str(data)

yaml.add_representer(blockStyleLiteral, repr_str)

parser = argparse.ArgumentParser()
parser.add_argument('--orgCount', required=True, type=int)
parser.add_argument('--peerCountInOrg', required=True, type=int)
parser.add_argument('--domainName', required=True, type=str)
parser.add_argument('--registrarUserName', required=True, type=str)
parser.add_argument('--registrarUserPassword', required=True, type=str)
parser.add_argument('--channelName', required=True, type=str)

args=parser.parse_args()
orgCount = args.orgCount
peerCountInOrg = args.peerCountInOrg
domainName = args.domainName
registrarUserName = args.registrarUserName
registrarUserPassword = args.registrarUserPassword
channelName = args.channelName
cryptoconfig = 'dev/crypto-config'


config = OrderedDict()
config['name'] = 'at2chain'
config['x-type'] = 'hlfv1'
config['x-loggingLevel'] = 'info'
config['description'] = 'the environment for {}'.format(config['name'] )
config['version'] = '1.0.0'

client = OrderedDict()
client['organization'] = 'Org1'
client['logging'] = OrderedDict({'level' : 'info'})
client['peer'] = OrderedDict({'timeout' : {'connection': '3s', 'queryResponse':'45s', 'executeTxResponse':'30s'}})
client['eventService'] = OrderedDict({'timeout':{'connection':'3s', 'registrationResponse':'3s'}})
client['orderer'] = OrderedDict({'timeout':{'connection':'3s', 'response':'5s'}})
client['cryptoconfig'] = {'path':cryptoconfig}
client['credentialStore'] = OrderedDict({'path' : '/tmp/hfc-kvs', 'cryptoStore':{'path':'/tmp/msp'}, 'wallet':'wallet-name'})
client['BCCSP'] = OrderedDict({'security' : {'enabled':True, 'default':{'provider': 'SW'}, 'hashAlgorithm':'SHA2', 'softVerify':True, 'ephemeral':False, 'level':'256'}})
client['tlsCerts'] = {'systemCertPool': False}
config['client'] = client

peersInChannel = OrderedDict()
for i in range(1, orgCount + 1):
  for j in range(1, peerCountInOrg + 1):
  	peersInChannel['peer{}.org{}.{}'.format(j, i, domainName)] = OrderedDict({'endorsingPeer' : True, 'chaincodeQuery' : True, 'ledgerQuery': True, 'eventSource':True})

channels = OrderedDict({channelName: {'orderers': ['orderer.{}'.format(domainName)], 'peers': peersInChannel}})
config['channels'] = channels

organizations = OrderedDict()
for i in range(1, orgCount + 1):
	org = OrderedDict()
	org['mspid'] = 'Org{}MSP'.format(i)
	orgDomain = 'org{}.{}'.format(i, domainName)
	org['cryptoPath'] = '{}/peerOrganizations/{}/users/Admin@{}/msp'.format(cryptoconfig, orgDomain, orgDomain)
	org['peers'] = ['peer{}.org{}.{}'.format(j, i, domainName) for j in range(1, peerCountInOrg+1)]
	org['certificateAuthorities'] = ['ca-org{}'.format(i)]
	privateKeyDir = '../crypto-config/peerOrganizations/{}/users/Admin@{}/msp/keystore/'.format(orgDomain, orgDomain)
	signedCertFile = '../crypto-config/peerOrganizations/{}/users/Admin@{}/msp/signcerts/Admin@{}-cert.pem'.format(orgDomain, orgDomain, orgDomain)
	files = os.listdir(privateKeyDir)
	privateKey = open(privateKeyDir + files[0], 'r').read()
	signedCert = open(signedCertFile, 'r').read()
	org['adminPrivateKey'] = {'pem' : blockStyleLiteral(privateKey)}
	org['signedCert'] = {'pem' : blockStyleLiteral(signedCert)}
	organizations['Org{}'.format(i)] = org
organizations['ordererorg'] = OrderedDict({'mspID':'OrdererMSP', 'cryptoPath':'{}/ordererOrganizations/{}/users/Admin@{}/msp'.format(cryptoconfig, domainName, domainName)})
config['organizations'] = organizations

orderers = OrderedDict();
ordererDomain = 'orderer.{}'.format(domainName)
caCertFile = '../crypto-config/ordererOrganizations/{}/tlsca/tlsca.{}-cert.pem'.format(domainName, domainName)
caCert = open(caCertFile, 'r').read()
orderers[ordererDomain] = OrderedDict({'url': 'grpcs://{}:7050'.format(ordererDomain), 'grpcOptions':{'sslProvider':'openSSL', 'negotiationType':'TLS', 'hostnameOverride':ordererDomain, 'grpc-max-send-message-length':15}, 'tlsCACerts':{'pem': blockStyleLiteral(caCert)}});
config['orderers'] = orderers

peers = OrderedDict();
for i in range(1, orgCount + 1):
  for j in range(1, peerCountInOrg + 1):
  	peerDomain = 'peer{}.org{}.{}'.format(j, i, domainName)
  	tlsCACertFile = '../crypto-config/peerOrganizations/org{}.tfs.at2chain.com/tlsca/tlsca.org{}.{}-cert.pem'.format(i, i, domainName)
  	tlsCACert = open(tlsCACertFile, 'r').read()
  	peers[peerDomain] = OrderedDict({'url' : 'grpcs://{}:7051'.format(peerDomain), 'grpcOptions':{'grpc.http2.keepalive_time':15, 'negotiationType':'TLS', 'sslProvider':'openSSL'}, 'tlsCACerts':{'pem': blockStyleLiteral(tlsCACert)}})
config['peers'] = peers

certificateAuthorities = OrderedDict();
for i in range(1, orgCount + 1):
	caName = 'ca.org{}'.format(i)
	ca = OrderedDict();
	ca['url'] = 'https://{}.{}:7054'.format(caName, domainName)
	ca['httpOptions'] = {'verify': True}
	tlsCACertFile = '../crypto-config/peerOrganizations/org{}.{}/ca/{}.{}-cert.pem'.format(i, domainName, caName, domainName)
	tlsCACert = open(tlsCACertFile, 'r').read()
	ca['tlsCACerts'] = OrderedDict({'pem': blockStyleLiteral(tlsCACert)})
	ca['registrar'] = OrderedDict({'enrollId': registrarUserName, 'enrollSecret': registrarUserPassword})
	ca['caName'] = 'ca-core'
	certificateAuthorities[caName] = ca
config['certificateAuthorities'] = certificateAuthorities	

with open('fabric-network-config.yaml', 'w') as outfile:
	yaml.dump(config, outfile, Dumper=CustomDumper, width=1000, default_flow_style=False)