#!/usr/bin/bash

ALICE_NODE_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/persistent_peers = ""/persistent_peers = "'$ALICE_NODE_KEY'@val-alice:26656"/g' \
    /root/.checkers/config/config.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/private_peer_ids = ""/private_peer_ids = "'$ALICE_NODE_KEY'"/g' \
    /root/.checkers/config/config.toml

BOB_NODE_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-bob:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)
CAROL_NODE_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/node-carol:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/seeds = ""/seeds = "'$BOB_NODE_KEY'@sentry-bob:26656,'$CAROL_NODE_KEY'@node-carol:26656"/g' \
    /root/.checkers/config/config.toml