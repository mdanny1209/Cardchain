#!/bin/bash

csd export --for-zero-height > /var/blockchain/genesis.json
/go/bin/csd start --rpc.laddr tcp://0.0.0.0:26657
