#!/usr/bin/bash

ALICE=$(echo $(sudo cat prod-sim/desk-alice/keys/passphrase.txt) | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i \
    keys \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    show alice --address)

echo "alice: $ALICE"

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i \
    genesis add-genesis-account $ALICE 1000000000upawn

sudo mv prod-sim/desk-alice/config/genesis.json \
    prod-sim/desk-bob/config/

BOB=$(echo $(sudo cat prod-sim/desk-bob/keys/passphrase.txt) | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-bob:/root/.checkers \
    checkersd_i \
    keys \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    show bob --address)

echo "bob: $BOB"

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-bob:/root/.checkers \
    checkersd_i \
    genesis add-genesis-account $BOB 500000000upawn