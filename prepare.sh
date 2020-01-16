#!/bin/bash
curl -L -o qemu-aarch64-static-v4.2.0-2.tar.gz https://github.com/multiarch/qemu-user-static/releases/download/v4.2.0-2/qemu-aarch64-static.tar.gz
tar xzf qemu-aarch64-static-v4.2.0-2.tar.gz
cp qemu-aarch64-static /usr/bin/
docker run --rm --privileged multiarch/qemu-user-static:register
