#!/bin/bash -eu

id="$1"

mv cmd/foo "cmd/$id"

sed -i "s/foo/$id/g" bin/build.sh

sed -i "s/foo/$id/g" turbobob.json

sed -i "s/template-go/$id/g" go.mod

# start with a fresh timeline.
# this line is dangerous if we're not in a template-beginning state,
# but is mostly mitigated by the "$ mv cmd/foo" stopping this process.
rm -rf .git

# replace template-repo specific readme (= it contains instructions related to the template itself)
# with a blank one
rm README.md && touch README.md

# thank you for your service (pats self in the back)
rm bin/init-template.sh
