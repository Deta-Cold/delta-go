#!/bin/sh

set -e

cd $(dirname $0)

TARGET=$1
VERSION=$(cat /release/build/VERSION)

cd /release/build

install -D -m 0755 detahardd-$TARGET          ./usr/bin/detahardd
install -D -m 0644 /release/detahard.rules    ./lib/udev/rules.d/50-detahard.rules
install -D -m 0644 /release/detahardd.service ./usr/lib/systemd/system/detahardd.service

# prepare GPG signing environment
GPG_PRIVKEY=/release/privkey.asc
if [ -r $GPG_PRIVKEY ]; then
    export GPG_TTY=$(tty)
    export LC_ALL=C.UTF-8
    gpg --import /release/privkey.asc
    GPG_SIGN=gpg
    GPGSIGNKEY=$(gpg --list-keys --with-colons | grep '^pub' | cut -d ":" -f 5)
fi

NAME=detahard-bridge

rm -f *.tar.bz2
tar -cjf $NAME-$VERSION.tar.bz2 ./usr ./lib

for TYPE in "deb" "rpm"; do
    case "$TARGET-$TYPE" in
        linux-386-*)
            ARCH=i386
            ;;
        linux-amd64-deb)
            ARCH=amd64
            ;;
        linux-amd64-rpm)
            ARCH=x86_64
            ;;
        linux-arm64-deb)
            ARCH=arm64
            ;;
        linux-arm64-rpm)
            ARCH=aarch64
            ;;
    esac
    fpm \
        -s tar \
        -t $TYPE \
        -a $ARCH \
        -n $NAME \
        -v $VERSION \
        -d systemd \
        --license "LGPL-3.0" \
        --vendor "SatoshiLabs" \
        --description "Communication daemon for detahard" \
        --maintainer "SatoshiLabs <stick@satoshilabs.com>" \
        --url "https://detahard.io/" \
        --category "Productivity/Security" \
        --before-install /release/fpm.before-install.sh \
        --after-install /release/fpm.after-install.sh \
        --before-remove /release/fpm.before-remove.sh \
        $NAME-$VERSION.tar.bz2
    case "$TYPE-$GPG_SIGN" in
        deb-gpg)
            /release/dpkg-sig -k $GPGSIGNKEY --sign builder detahard-bridge_${VERSION}_${ARCH}.deb
            ;;
        rpm-gpg)
            rpm --addsign -D "%_gpg_name $GPGSIGNKEY" detahard-bridge-${VERSION}-1.${ARCH}.rpm
            ;;
    esac
done

rm -rf ./usr ./lib
