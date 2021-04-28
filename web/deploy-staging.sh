#!/usr/bin/env sh

# abort on errors
set -e

# build
# yarn build

cd docs/.vuepress/dist
echo 'staging.advenjourney.com' > CNAME
git init
git add -A
git commit -m 'deploy'

# if you are deploying to https://<USERNAME>.github.io
# git push -f git@github.com:<USERNAME>/<USERNAME>.github.io.git master

# if you are deploying to https://<USERNAME>.github.io/<REPO>
git push -f git@github.com:advenjourney/website.git master:staging

cd -