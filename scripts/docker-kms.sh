#!/usr/bin/bash

docker run --rm -it \
    -v $(pwd)/prod-sim/kms-alice:/root/tmkms \
    tmkms_i \
    init /root/tmkms

sudo cp $(pwd)/scripts/tmkms.toml $(pwd)/prod-sim/kms-alice

sudo bash -c "docker run --rm -t \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    checkersd_i \
    tendermint show-validator \
    | tr -d '\n' | tr -d '\r' \
    > prod-sim/desk-alice/config/pub_validator_key-val-alice.json"

sudo cp prod-sim/val-alice/config/priv_validator_key.json \
    prod-sim/desk-alice/config/priv_validator_key-val-alice.json
sudo mv prod-sim/val-alice/config/priv_validator_key.json \
    prod-sim/kms-alice/secrets/priv_validator_key-val-alice.json

docker run --rm -i \
    -v $(pwd)/prod-sim/kms-alice:/root/tmkms \
    -w /root/tmkms \
    tmkms_i \
    softsign import secrets/priv_validator_key-val-alice.json \
    secrets/val-alice-consensus.key

sudo cp prod-sim/sentry-alice/config/priv_validator_key.json \
    prod-sim/val-alice/config/

docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/priv_validator_laddr = ""/priv_validator_laddr = "tcp:\/\/0.0.0.0:26659"/g' \
    /root/.checkers/config/config.toml
docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^priv_validator_key_file/# priv_validator_key_file/g' \
    /root/.checkers/config/config.toml
docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^priv_validator_state_file/# priv_validator_state_file/g' \
    /root/.checkers/config/config.toml

sudo cp prod-sim/sentry-alice/config/priv_validator_key.json \
    prod-sim/val-alice/config