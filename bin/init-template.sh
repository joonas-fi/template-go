#!/bin/bash -eu

id="$1"

mv cmd/foo "cmd/$id"

sed -i "s/foo/$id/g" bin/build.sh

sed -i "s/foo/$id/g" turbobob.json

sed -i "s/template-go/$id/g" go.mod
