#!/usr/bin/bash

read -p "Enter desired password: " password

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i \
    keys \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    add alice

sudo bash -c "echo -n ${password} > prod-sim/desk-alice/keys/passphrase.txt"

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-bob:/root/.checkers \
    checkersd_i \
    keys \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    add bob

sudo bash -c "echo -n ${password} > prod-sim/desk-bob/keys/passphrase.txt"