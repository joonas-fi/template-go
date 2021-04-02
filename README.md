A template for kicking off a Go project.

I have this `new-go-project.sh` script on my workstation:

```bash
#!/bin/bash -eu

git clone https://github.com/joonas-fi/template-go.git "$1"

# replace placeholders with concrete ones etc.
cd "$1" && bin/init-template.sh "$1"
```
