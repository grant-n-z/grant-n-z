#!/bin/bash

set -e -u -x

# Set auth
git config --local github.user tomoyane
git config --local github.token "${GITHUB_TOKEN}"

# gnzcacher bump
cd gnzcacher
cat grant_n_z_cacher.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > current_version.txt
current_version=$(cat current_version.txt)
major=$(cat current_version.txt| tr '.' '\n' | tail -n 3 | head -1)
minor=$(cat current_version.txt| tr '.' '\n' | tail -n 2 | head -1)
patch=$(cat current_version.txt| tr '.' '\n' | tail -n 1)
patch_bump=$((patch + 1))
update_version="${major}.${minor}.${patch_bump}"
sed -i -e "s/${current_version}/${update_version}/g" grant_n_z_cacher.yaml
rm current_version.txt grant_n_z_cacher.yaml-e
git add grant_n_z_cacher.yaml
git commit grant_n_z_cacher.yaml -m "bump version for grant_n_z_cacher"

# gnzserver bump
cd ../gnzserver
cat grant_n_z_server.yaml| grep version | sed 's/^[ \t]*//' | sed 's/version://' | sed 's/^[ \t]*//' > current_version.txt
current_version=`cat current_version.txt`
major=$(cat current_version.txt| tr '.' '\n' | tail -n 3 | head -1)
minor=$(cat current_version.txt| tr '.' '\n' | tail -n 2 | head -1)
patch=$(cat current_version.txt| tr '.' '\n' | tail -n 1)
patch_bump=$((patch + 1))
update_version="${major}.${minor}.${patch_bump}"
sed -i -e "s/${current_version}/${update_version}/g" grant_n_z_server.yaml
rm current_version.txt grant_n_z_server.yaml-e
git add grant_n_z_server.yaml
git commit grant_n_z_cacher.yaml -m "bump version for grant_n_z_server"
git push origin master
