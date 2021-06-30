#! /bin/bash

# check deps
curl --version > /dev/null 2>&1 | if [ $? -gt 0 ]; then echo "install cURL"; exit 1; fi
tar --version > /dev/null 2>&1 | if [ $? -gt 0 ]; then echo "install tar"; exit 1; fi

VER="v1.0"

VERSION=$1
if [ "${VERSION}" == "" ]; then
    VERSION=${VER}
fi

APP=github-downloader
DIR=${HOME}/.${APP}

DEST=/usr/bin
if [[ $(uname) == "Darwin" ]]; then DEST=/usr/local/bin ; fi

os() {
    os=$(uname)
    os=$(echo "${os}" | tr '[:upper:]' '[:lower:]')
    arch=$(uname -m)

    if [[ ${os} == "linux" || ${os} == "windows" ]]; then
        if [[ ${arch} == "x86_64" ]]; then
            arch="amd64"
        elif [[ ${arch} == "i386" ]]; then
            arch="386"
        elif [[ ${arch} == "aarch64" || ${arch} == "aarch64_be" || ${arch} == "armv8b" || ${arch} == "armv8l" ]]; then
            arch="arm64"
        elif [[ ${arch} != "arm" ]]; then
            echo "${os}-${arch} arch is not supported"
            exit 1
        fi
    fi

    echo "${os}-${arch}"
}

if [ ! -d ${DIR} ]; then mkdir -p ${DIR}; fi

echo "Downloading the app, version: ${VERSION}"

url="https://github.com/kislerdm/${APP}/releases/download/${VERSION}/${APP}-${VERSION}-$(os).tar.gz"

curl -SLo ${DIR}/${APP}-${VERSION}.tar.gz ${url}

echo "Extracting the app"

cd ${DIR}
tar -vxf ${APP}-${VERSION}.tar.gz
if [ $? -gt 0 ]; then
    echo "The app archive is corrupt, it's likely that the specified version (${VERSION}) does not exist"
    exit 1
fi

echo "Linking the app to ${DEST}"

if [ $(ls ${DEST} | grep -e github-downloader | wc -l) -gt 0 ]; then
    read  -n 1 -p "The binary is linked, would you like to overwrite? [Y/n] " answer
    if [[ ${answer} == "Y" || ${answer} == "y" ]]; then
        sudo rm ${DEST}/${APP}
    else
        echo
        echo "Bye!"
        exit 0
    fi
fi

sudo ln -sf ${DIR}/${APP} ${DEST}/${APP}

exec -l ${SHELL}
echo "Done. Bye!"
