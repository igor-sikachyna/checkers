#!/usr/bin/bash

echo -e desk-alice'\n'desk-bob'\n'node-carol'\n'sentry-alice'\n'sentry-bob'\n'val-alice'\n'val-bob \
		| xargs -I {} \
		docker run --rm -i \
		-v $(pwd)/prod-sim/{}:/root/.checkers \
		checkersd_i \
		init checkers