getent group detahardd >/dev/null || groupadd -r detahardd
getent group plugdev >/dev/null || groupadd -r plugdev
getent passwd detahardd >/dev/null || useradd -r -g detahardd -d /var -s /bin/false -c "detahard Bridge" detahardd
usermod -a -G plugdev detahardd
