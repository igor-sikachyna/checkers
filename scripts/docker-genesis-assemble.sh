#!/usr/bin/bash

sudo bash -c "cp prod-sim/desk-bob/config/gentx/gentx-* \
    prod-sim/desk-alice/config/gentx"

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i genesis collect-gentxs

docker run --rm -it \
    -v $(pwd)/prod-sim/desk-alice:/root/.checkers \
    checkersd_i \
    genesis validate

sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/desk-bob/config
sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/node-carol/config
sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/sentry-alice/config
sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/sentry-bob/config
sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/val-alice/config
sudo cp prod-sim/desk-alice/config/genesis.json prod-sim/val-bob/config