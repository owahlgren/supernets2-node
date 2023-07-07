#!/bin/sh

set -e

gen() {
    local package=$1

    abigen --bin bin/${package}.bin --abi abi/${package}.abi --pkg=${package} --out=${package}/${package}.go
}

gen supernets2
gen supernets2bridge
gen matic
gen mockverifier
gen supernets2globalexitroot
gen supernets2datacommittee