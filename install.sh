#! /bin/bash

# check deps
curl --version > /dev/null 2>&1 | if [ "$?" != "0" ]; then echo "install cURL"; exit 1; fi
tar --version > /dev/null 2>&1 | if [ "$?" != "0" ]; then echo "install tar"; exit 1; fi

VERSION=$1
if [ "${VERSION}" == "" ]; then
    VERSION="latest"
fi

APP=github-downloader
DIR=${HOME}/.${APP}

if [ ! -d ${DIR} ]; then mkdir -p ${DIR}; fi

echo "Downloading the app, version: ${VERSION}"

curl -SLo ${DIR}/${APP}.tar.gz "https://github.com/kislerdm/${APP}/releases/download/${VERSION}/${APP}-${VERSION}-$(uname)-$(uname -m).tar.gz"

echo "Extracting the app"

cd ${DIR}
tar -vxf ${APP}.tar.gz
if [ $? -eq 1 ]; then
    echo "The app archive is corrupt, it's likely that the specified version (${VERSION}) does not exist"
    exit 1
fi
rm ${APP}.tar.gz

echo "Linking the app to /usr/local/bin"

sudo ln -s ${DIR}/${APP} /usr/local/bin/${APP}
