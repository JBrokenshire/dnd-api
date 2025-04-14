#!/bin/sh
if [ -d $GOPATH/dnd-api.com/mdm/pkg/go-migrations/template/ ]
then
  cp -a $GOPATH/dnd-api.com/mdm/pkg/go-migrations/template/. ./
  exit 1
fi
if [ -d vendor/dnd-api.com/mdm/pkg/go-migrations/template ]
then
  cp -a vendor/dnd-api.com/mdm/pkg/go-migrations/template/. ./
  exit 1
fi
echo "Dependency path not found"
exit 0