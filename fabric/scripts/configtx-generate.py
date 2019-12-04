#!/usr/bin/python
import yaml
from collections import OrderedDict

def represent_none(self, _):
    return self.represent_scalar('tag:yaml.org,2002:null', '')

yaml.add_representer(type(None), represent_none)


def setup_yaml():
  represent_dict_order = lambda self, data:  self.represent_mapping('tag:yaml.org,2002:map', data.items())
  yaml.add_representer(OrderedDict, represent_dict_order)
setup_yaml()


orgCount = 3
profilePrefixName = 'At2chain'
domainName = 'tfs.at2chain.com'
configtx = OrderedDict()


OrdererOrg =OrderedDict({'OrdererOrg':{
								  	'Name': 'OrdererOrg',
								  	'ID': 'OrdererMSP',
								  	'MSPDir': 'crypto-config/ordererOrganizations/tfs.at2chain.com/msp'}})
organizations = [OrdererOrg]


for i in range(1, orgCount + 1):
	Org = OrderedDict({'Org'+str(i): {
		'Name': 'Org'+str(i)+'MSP',
	   			'ID': 'Org'+str(i)+'MSP',
	             'MSPDir': 'crypto-config/peerOrganizations/org'+str(i)+'.'+domainName +'/msp',
	             'AnchorPeers': [OrderedDict({'Host': 'peer1.org{}.{}'.format(i,domainName),'port': '7051'})]
	}})
	organizations.append(Org)

configtx['Organizations'] = organizations


#Orderer
orderer = OrderedDict()
orderer = {
	'OrdererType': 'kafka',
	'Addresses': 'orderer.domainName:7050',
    'BatchTimeout': '2s',
    'BatchSize': {
 	                    'MaxMessageCount': 10,
                        'AbsoluteMaxBytes': '99 MB',
                        'PreferredMaxBytes': '512 KB'
                 },
    'Kafka': {'Brokers': [{'kafka1': '9092'}, {'kafka2': '9092'}]},
    'Organizations': None
}

configtx['Orderer'] = orderer

#Application
application = OrderedDict()
application['Organizations']= None

configtx['Application'] = application

#profiles
profiles = OrderedDict()
OrdererGenesis = OrderedDict()
OrdererGenesis['Orderer'] = {
	'OrdererType': 'kafka',
	'Addresses': 'orderer.domainName:7050',
    'BatchTimeout': '2s',
    'BatchSize': {
 	                    'MaxMessageCount': 10,
                        'AbsoluteMaxBytes': '99 MB',
                        'PreferredMaxBytes': '512 KB'
                 },
    'Kafka': {'Brokers': [{'kafka1': '9092'}, {'kafka2': '9092'}]},
    'Organizations': {
								  	'Name': 'OrdererOrg',
								  	'ID': 'OrdererMSP',
								  	'MSPDir': 'crypto-config/ordererOrganizations/tfs.at2chain.com/msp'}
    }

organizationsGenesis = []
for i in range(1, orgCount + 1):
	Org = OrderedDict({'Org'+str(i): {
		'Name': 'Org'+str(i)+'MSP',
	   			'ID': 'Org'+str(i)+'MSP',
	             'MSPDir': 'crypto-config/peerOrganizations/org'+str(i)+'.'+domainName +'/msp',
	             'AnchorPeers': [OrderedDict({'Host': 'peer1.org{}.{}'.format(i,domainName),'port': '7051'})]
	}})
	organizationsGenesis.append(Org)


Consortiums = OrderedDict()
SaservicesmpleConsortium = OrderedDict()
SaservicesmpleConsortium['Organizations'] = organizationsGenesis
Consortiums['Consortiums'] = SaservicesmpleConsortium
OrdererGenesis['Consortiums'] =  Consortiums
profiles[profilePrefixName+'OrdererGenesis'] = OrdererGenesis


organizationsChannel = []
for i in range(1, orgCount + 1):
	Org = OrderedDict({'Org'+str(i): {
		'Name': 'Org'+str(i)+'MSP',
	   			'ID': 'Org'+str(i)+'MSP',
	             'MSPDir': 'crypto-config/peerOrganizations/org'+str(i)+'.'+domainName +'/msp',
	             'AnchorPeers': [OrderedDict({'Host': 'peer1.org{}.{}'.format(i,domainName),'port': '7051'})]
	}})
	organizationsChannel.append(Org)
Channel = OrderedDict()
Channel['Consortium'] = profilePrefixName+'Consortium'
Channel['Application'] = None
Channel['Organizations'] = organizationsChannel
profiles[profilePrefixName+'Channel'] = Channel
configtx['Profiles'] = profiles

y = yaml.dump(configtx, indent=4, default_flow_style=False)
print(y)
