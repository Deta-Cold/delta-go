#!/bin/sh

set -ex

/usr/local/osxcross/bin/lipo \
   -create release/macos/build/detahardd-arm64 release/macos/build/detahardd-amd64 \
   -output release/macos/build/detahardd
