#!/usr/bin/bash

sudo cp prod-sim/val-bob/config/priv_validator_key.json \
    prod-sim/desk-bob/config/priv_validator_key.json

echo $(sudo cat prod-sim/desk-bob/keys/passphrase.txt) | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-bob:/root/.checkers \
    checkersd_i \
    genesis gentx bob 40000000upawn \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    --account-number 0 --sequence 0 \
    --chain-id checkers-1 \
    --gas 1000000 \
    --gas-prices 0.1upawn

sudo mv prod-sim/desk-bob/config/genesis.json \
    prod-sim/desk-alice/config/

echo $(sudo cat prod-sim/desk-alice/keys/passphrase.txt) | docker run --rm -i \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i \
    genesis gentx alice 60000000upawn \
    --keyring-backend file --keyring-dir /root/.checkers/keys \
    --account-number 0 --sequence 0 \
    --pubkey $(cat prod-sim/desk-alice/config/pub_validator_key-val-alice.json) \
    --chain-id checkers-1 \
    --gas 1000000 \
    --gas-prices 0.1upawn