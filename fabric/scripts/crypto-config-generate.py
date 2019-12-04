#!/usr/bin/python
import yaml
from collections import OrderedDict

def setup_yaml():
  represent_dict_order = lambda self, data:  self.represent_mapping('tag:yaml.org,2002:map', data.items())
  yaml.add_representer(OrderedDict, represent_dict_order)
setup_yaml()

crypto = OrderedDict()
orgCount = 3
Domain= 'tfs.at2chain.com'

#orderer
OrdererOrgs=OrderedDict()
Specs=OrderedDict()
OrdererOrgs['Name'] = 'Orderer'
OrdererOrgs['Domain'] = Domain
Specs['Hostname'] = 'orderer' + Domain
Specs['CommonName'] = 'orderer' + Domain
OrdererOrgs['Specs'] = [Specs]
crypto['OrdererOrgs'] = [OrdererOrgs]


#peer
PeerOrgs = OrderedDict()
for i in range(1,orgCount + 1):
	Specs=OrderedDict()
	PeerOrgs = OrderedDict()
	PeerOrgs['Name'] = 'Org'+str(i)
	PeerOrgs['Domain'] = 'Org'+str(i)+'.'+Domain
	Specs['Hostname'] = 'Org'+str(i)+'.'+Domain
	Specs['CommonName'] = 'Org'+str(i)+'.'+Domain
	PeerOrgs['Specs'] = [Specs]
	crypto['Org'+str(i)] = [PeerOrgs]


y = yaml.dump(crypto, indent=4, default_flow_style=False)
print(y)