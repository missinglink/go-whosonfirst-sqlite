#!/bin/sh

# SOMETHING SOMETHING SOMETHING lfs IS NOT INSTALLED YET

git lfs clone git@github.com:whosonfirst-data/${REPO}.git /usr/local/data/${REPO}

/bin/wof-sqlite-index -live-hard-die-fast -all -dsn /usr/local/data/${FNAME}-latest.db -mode repo ${REPO} 

bzip2 /usr/local/data/${FNAME}-latest.db

# SOMETHING SOMETHING SOMETHING awcli IS NOT INSTALLED YET

aws s3 --profile whosonfirst cp /usr/local/data/${FNAME}-latest.db s3://dist.whosnfirst.org/sqlite/
