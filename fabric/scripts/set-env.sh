#!/bin/bash
export FABRIC_PATH=`sed '/^FABRIC_PATH=/!d;s/.*=//' tfs-env-config.yaml`

COUCHDB_PATH=`sed '/^COUCHDB_PATH=/!d;s/.*=//' tfs-env-config.yaml`
COUCHDB_USERNAME=`sed '/^COUCHDB_USERNAME=/!d;s/.*=//' tfs-env-config.yaml`
COUCHDB_PASSWORD=`sed '/^COUCHDB_PASSWORD=/!d;s/.*=//' tfs-env-config.yaml`

MYSQL_PATH=`sed '/^MYSQL_PATH=/!d;s/.*=//' tfs-env-config.yaml`
MYSQL_USERNAME=`sed '/^MYSQL_USERNAME=/!d;s/.*=//' tfs-env-config.yaml`
MYSQL_PASSWORD=`sed '/^MYSQL_PASSWORD=/!d;s/.*=//' tfs-env-config.yaml`
//TODO mysql connection Hostaddress Port Sid

SWIFT_PATH=`sed '/^SWIFT_PATH=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_USERNAME=`sed '/^SWIFT_OS_USERNAME=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_PASSWORD=`sed '/^SWIFT_OS_PASSWORD=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_PROJECT_NAME=`sed '/^SWIFT_OS_PROJECT_NAME=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_USER_DOMAIN_NAME=`sed '/^SWIFT_OS_USER_DOMAIN_NAME=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_PROJECT_DOMAIN_NAME=`sed '/^SWIFT_OS_PROJECT_DOMAIN_NAME=/!d;s/.*=//' tfs-env-config.yaml`
SWIFT_OS_AUTH_URL=`sed '/^SWIFT_OS_AUTH_URL=/!d;s/.*=//' tfs-env-config.yaml`

TFS_API_PATH=`sed '/^TFS_API_PATH=/!d;s/.*=//' tfs-env-config.yaml`
TFS_ADMAIN_API=`sed '/^TFS_ADMAIN_API=/!d;s/.*=//' tfs-env-config.yaml`

TFS_API_TO_FABRIC_CONFIG=`sed '/^TFS_API_TO_FABRIC_CONFIG=/!d;s/.*=//' tfs-env-config.yaml`
TFS_API_TO_SWIFT_CONFIG=`sed '/^TFS_API_TO_SWIFT_CONFIG=/!d;s/.*=//' tfs-env-config.yaml`
TFS_EXPLORER_TO_API_CONFIG=`sed '/^TFS_EXPLORER_TO_API_CONFIG=/!d;s/.*=//' tfs-env-config.yaml`

#非空校验

if [ "$FABRIC_PATH" = "" ]; then
   echo "FABRIC_PATH is not set!"
else
   echo "FABRIC_PATH is:" ${FABRIC_PATH}
fi

if [ "$COUCHDB_USERNAME" = "" ]; then
   echo "COUCHDB_USERNAME is not set!"
else
   echo "COUCHDB_USERNAME is:" ${COUCHDB_USERNAME}
fi

if [ "$COUCHDB_PASSWORD" = "" ]; then
   echo "COUCHDB_PASSWORD is not set!"
else
   echo "COUCHDB_PASSWORD is:" ${COUCHDB_PASSWORD}
fi

if [ "$MYSQL_USERNAME" = "" ]; then
   echo "MYSQL_USERNAME is not set!"
else
   echo "MYSQL_USERNAME is:" ${MYSQL_USERNAME}
fi

if [ "$MYSQL_PASSWORD" = "" ]; then
   echo "MYSQL_PASSWORD is not set!"
else
   echo "MYSQL_PASSWORD is:" ${MYSQL_PASSWORD}
fi

if [ "$SWIFT_PATH" = "" ]; then
   echo "SWIFT_PATH is not set!"
else
   echo "SWIFT_PATH is:" ${SWIFT_PATH}
fi

if [ "$SWIFT_OS_USERNAME" = "" ]; then
   echo "SWIFT_OS_USERNAME is not set!"
else
   echo "SWIFT_OS_USERNAME is:" ${SWIFT_OS_USERNAME}
fi

if [ "$SWIFT_OS_PASSWORD" = "" ]; then
   echo "SWIFT_OS_PASSWORD is not set!"
else
   echo "SWIFT_OS_PASSWORD is:" ${SWIFT_OS_PASSWORD}
fi

if [ "$SWIFT_OS_PROJECT_NAME" = "" ]; then
   echo "SWIFT_OS_PROJECT_NAME is not set!"
else
   echo "SWIFT_OS_PROJECT_NAME is:" ${SWIFT_OS_PROJECT_NAME}
fi

if [ "$SWIFT_OS_USER_DOMAIN_NAME" = "" ]; then
   echo "SWIFT_OS_USER_DOMAIN_NAME is not set!"
else
   echo "SWIFT_OS_USER_DOMAIN_NAME is:" ${SWIFT_OS_USER_DOMAIN_NAME}
fi

if [ "$SWIFT_OS_PROJECT_DOMAIN_NAME" = "" ]; then
   echo "SWIFT_OS_PROJECT_DOMAIN_NAME is not set!"
else
   echo "SWIFT_OS_PROJECT_DOMAIN_NAME is:" ${SWIFT_OS_PROJECT_DOMAIN_NAME}
fi

if [ "$SWIFT_OS_AUTH_URL" = "" ]; then
   echo "SWIFT_OS_AUTH_URL is not set!"
else
   echo "SWIFT_OS_AUTH_URL is:" ${SWIFT_OS_AUTH_URL}
fi

if [ "$TFS_API_PATH" = "" ]; then
   echo "TFS_API_PATH is not set!"
else
   echo "TFS_API_PATH is:" ${TFS_API_PATH}
fi

if [ "$TFS_ADMAIN_API" = "" ]; then
   echo "TFS_ADMAIN_API is not set!"
else
   echo "TFS_ADMAIN_API is:" ${TFS_ADMAIN_API}
fi

if [ "$TFS_API_TO_FABRIC_CONFIG" = "" ]; then
   echo "TFS_API_TO_FABRIC_CONFIG is not set!"
else
   echo "TFS_API_TO_FABRIC_CONFIG is:" ${TFS_API_TO_FABRIC_CONFIG}
fi

if [ "$TFS_API_TO_SWIFT_CONFIG" = "" ]; then
   echo "TFS_API_TO_SWIFT_CONFIG is not set!"
else
   echo "TFS_API_TO_SWIFT_CONFIG is:" ${TFS_API_TO_SWIFT_CONFIG}
fi

if [ "$TFS_EXPLORER_TO_API_CONFIG" = "" ]; then
   echo "TFS_EXPLORER_TO_API_CONFIG is not set!"
else
   echo "TFS_EXPLORER_TO_API_CONFIG is:" ${TFS_EXPLORER_TO_API_CONFIG}
fi