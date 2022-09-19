#! /bin/bash

grep ^gitPath remix.fixed.conf | cut -d '=' -f 2
grep ^esm_logPath remix.fixed.conf | cut -d '=' -f 2
grep ^tomcatVersionPath remix.fixed.conf | cut -d '=' -f 2

function kami() {
	ppp=$(grep ^gitPath remix.fixed.conf | cut -d '=' -f 2)
}


kami
echo "this : $ppp"

