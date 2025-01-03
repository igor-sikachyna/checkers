#!/usr/bin/bash

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -i 's/"stake"/"upawn"/g' /root/.checkers/config/genesis.json
echo -e desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/([0-9]+)stake/\1upawn/g' /root/.checkers/config/app.toml
echo -e desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
    | xargs -I {} \
    docker run --rm -i \
    -v $(pwd)/prod-sim/{}:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/^chain-id = .*$/chain-id = "checkers-1"/g' \
    /root/.checkers/config/client.toml

docker run --rm -i \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    --entrypoint sed \
    checkersd_i \
    -Ei 's/"chain_id": "checkers"/"chain_id": "checkers-1"/g' \
    /root/.checkers/config/genesis.json