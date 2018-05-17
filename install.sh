#!/bin/bash
curl https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_darwin_amd64-1.9.64.zip
unzip go_appengine_sdk_darwin_amd64-1.9.64.zip

mkdir -p $PWD/path
GOPATH=$PWD/path

go get google.golang.org/appengine/datastore
go get golang.org/x/oauth2/google
