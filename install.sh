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

echo "https://github.com/kislerdm/${APP}/releases/download/${VERSION}/${APP}-${VERSION}-$(uname)-$(uname -m).tar.gz"

curl -SLo ${DIR}/${APP}-${VERSION}.tar.gz "https://github.com/kislerdm/${APP}/releases/download/${VERSION}/${APP}-${VERSION}-$(uname)-$(uname -m).tar.gz"

echo "Extracting the app"

cd ${DIR}
tar -vxf ${APP}-${VERSION}.tar.gz
if [ $? -eq 1 ]; then
    echo "The app archive is corrupt, it's likely that the specified version (${VERSION}) does not exist"
    exit 1
fi

echo "Linking the app to /usr/local/bin"

if [ $(ls /usr/local/bin/github-downloader | wc -l) -gt 0 ]; then
    read  -n 1 -p "The file is linked, would you like to overwrite? [Y/n] " answer
    if [[ ${answer} == "Y" || ${answer} == "y" ]]; then
        sudo rm /usr/local/bin/${APP}
    else
        echo
        echo "Bye!"
        exit 0
    fi
fi

sudo ln -sf ${DIR}/${APP} /usr/local/bin/${APP}

exec -l ${SHELL}
