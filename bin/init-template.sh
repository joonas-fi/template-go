#!/bin/bash -eu

id="$1"

mv cmd/foo "cmd/$id"

sed -i "s/foo/$id/g" turbobob.json

sed -i "s/template-go/$id/g" go.mod

# start with a fresh timeline.
# this line is dangerous if we're not in a template-beginning state,
# but is mostly mitigated by the "$ mv cmd/foo" stopping this process.
rm -rf .git

# replace template-repo specific readme (= it contains instructions related to the template itself)
# with a blank one
rm README.md && echo -e "⬆️ For table of contents, click the above icon\n" > README.md

# thank you for your service (pats self in the back)
rm bin/init-template.sh
