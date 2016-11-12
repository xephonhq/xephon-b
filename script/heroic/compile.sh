#!/usr/bin/env bash

if [ -f "heroic-all.jar" ]; then
  echo "already compiled, remove the jar and re run the script"
  exit 0
fi

if [ ! -d "heroic" ]; then
  echo "heroic source not found, clone from github.com/spotify/heroic"
  git clone --depth 1 https://github.com/spotify/heroic.git
fi

cd heroic
# FIXME: the install-repackaged should detect its own path
tools/install-repackaged

mvn clean package

cp heroic-dist/target/heroic-dist-0.0.1-SNAPSHOT-shaded.jar ../heroic-all.jar
