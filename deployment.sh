#!/bin/bash
PROJECT_NAME="go-healthycheck"
CONFIG_FOLDER="configs"
TEMP_FOLDER="temp_deploy"
TARGET_DEST=${GOPATH}/bin/${PROJECT_NAME}

echo "creating temp folder"
mkdir -p ./${TEMP_FOLDER}

echo "build go package to temp folder"
cp -r ${CONFIG_FOLDER} ./${TEMP_FOLDER}
go build -o ./${TEMP_FOLDER}/${PROJECT_NAME}
echo "build succeed"

echo "remove old package from target path: ${TARGET_DEST}"
rm -rf ${TARGET_DEST}/

echo "move new package to target path: ${TARGET_DEST}"
mv ./${TEMP_FOLDER} ${TARGET_DEST}
