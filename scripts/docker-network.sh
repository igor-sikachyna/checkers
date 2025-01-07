#!/usr/bin/bash

ALICE_VAL_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/persistent_peers = ""/persistent_peers = "'$ALICE_VAL_KEY'@val-alice:26656"/g' \
    /root/.checkers/config/config.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/private_peer_ids = ""/private_peer_ids = "'$ALICE_VAL_KEY'"/g' \
    /root/.checkers/config/config.toml

BOB_SENTRY_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-bob:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)
CAROL_NODE_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/node-carol:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/seeds = ""/seeds = "'$BOB_SENTRY_KEY'@sentry-bob:26656,'$CAROL_NODE_KEY'@node-carol:26656"/g' \
    /root/.checkers/config/config.toml

ALICE_SENTRY_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-alice:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/val-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/persistent_peers = ""/persistent_peers = "'$ALICE_SENTRY_KEY'@sentry-alice:26656"/g' \
    /root/.checkers/config/config.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/val-bob:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/persistent_peers = ""/persistent_peers = "'$BOB_SENTRY_KEY'@sentry-bob:26656"/g' \
    /root/.checkers/config/config.toml

BOB_VAL_KEY=$(docker run --rm -i \
    -v $(pwd)/prod-sim/val-bob:/root/.checkers \
    checkersd_i \
    tendermint show-node-id)

docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-bob:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/seeds = ""/seeds = "'$ALICE_SENTRY_KEY'@sentry-alice:26656,'$CAROL_NODE_KEY'@node-carol:26656"/g' \
    /root/.checkers/config/config.toml
docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-bob:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/persistent_peers = ""/persistent_peers = "'$BOB_VAL_KEY'@val-bob:26656"/g' \
    /root/.checkers/config/config.toml
docker run --rm -i \
    -v $(pwd)/prod-sim/sentry-bob:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/private_peer_ids = ""/private_peer_ids = "'$BOB_VAL_KEY'"/g' \
    /root/.checkers/config/config.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/node-carol:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/seeds = ""/seeds = "'$ALICE_SENTRY_KEY'@sentry-alice:26656,'$BOB_SENTRY_KEY'@sentry-bob:26656"/g' \
    /root/.checkers/config/config.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/node-carol:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei '0,/^laddr = .*$/{s/^laddr = .*$/laddr = "tcp:\/\/0.0.0.0:26657"/}' \
    /root/.checkers/config/config.toml

echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' \
    /root/.checkers/config/config.toml

echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' \
    /root/.checkers/config/app.toml

# Does not do anything
echo -e node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^enable-unsafe-cors = false/enable-unsafe-cors = true/g' \
    /root/.checkers/config/app.toml
