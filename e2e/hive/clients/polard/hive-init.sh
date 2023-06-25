#!/bin/bash
# SPDX-License-Identifier: BUSL-1.1
#
# Copyright (C) 2023, Berachain Foundation. All rights reserved.
# Use of this software is govered by the Business Source License included
# in the LICENSE file of this repository and at www.mariadb.com/bsl11.
#
# ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
# TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
# VERSIONS OF THE LICENSED WORK.
#
# THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
# LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
# LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
#
# TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
# AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
# EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
# MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
# TITLE.

KEYS[0]="dev0"
KEYS[1]="dev1"
KEYS[2]="dev2"
CHAINID="highbury_710-1"
MONIKER="localtestnet"
# Remember to change to other types of keyring like 'file' in-case exposing to outside world,
# otherwise your balance will be wiped quickly
# The keyring test does not require private key to steal tokens from you
KEYRING="test"
KEYALGO="eth_secp256k1"
LOGLEVEL="info"
# Set dedicated home directory for the ./bin/polard instance
HOMEDIR="/"
# to trace evm
#TRACE="--trace"
TRACE=""

# Path variables
CONFIG_TOML=$HOMEDIR/config/config.toml
APP_TOML=$HOMEDIR/config/app.toml
POLARIS_TOML=$HOMEDIR/config/polaris.toml
GENESIS=$HOMEDIR/config/genesis.json
TMP_GENESIS=$HOMEDIR/config/tmp_genesis.json

# used to exit on first error (any non-zero exit code)
set -e

# Reinstall daemon
# mage build

# Remove the previous folder

 # Set moniker and chain-id (Moniker can be anything, chain-id must be an integer)
polard init $MONIKER -o --chain-id $CHAINID --home "$HOMEDIR"
# Set client config
polard config set client keyring-backend $KEYRING --home "$HOMEDIR"
polard config set client chain-id "$CHAINID" --home "$HOMEDIR"

# If keys exist they should be deleted
for KEY in "${KEYS[@]}"; do
    polard keys add $KEY --keyring-backend $KEYRING --algo $KEYALGO --home "$HOMEDIR"
done

ETH_GENESIS_SOURCE=$(cat genesis.json)
temp_genesis=$(cat $GENESIS)
echo "eth_gen dump: "
echo $ETH_GENESIS_SOURCE

# Change eth_genesis in config/genesis.json
updated_genesis=$(echo "$temp_genesis" | jq --argjson eth_gen "$ETH_GENESIS_SOURCE" '.app_state["evm"] = $eth_gen')
echo "$updated_genesis" > "$GENESIS"

# Change parameter token denominations to ablack
jq '.app_state["staking"]["params"]["bond_denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["crisis"]["constant_fee"]["denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["gov"]["deposit_params"]["min_deposit"][0]["denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.app_state["mint"]["params"]["mint_denom"]="ablack"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"
jq '.consensus["params"]["block"]["max_gas"]="30000000"' "$GENESIS" >"$TMP_GENESIS" && mv "$TMP_GENESIS" "$GENESIS"

# Dump genesis
echo "genesis state:"
cat $GENESIS

# Allocate genesis accounts (cosmos formatted addresses)
for KEY in "${KEYS[@]}"; do
    polard genesis add-genesis-account $KEY 100000000000000000000000000ablack --keyring-backend $KEYRING --home "$HOMEDIR"
done

# Sign genesis transaction
polard genesis gentx ${KEYS[0]} 1000000000000000000000ablack --keyring-backend $KEYRING --chain-id $CHAINID --home "$HOMEDIR"
## In case you want to create multiple validators at genesis
## 1. Back to `./bin/polard keys add` step, init more keys
## 2. Back to `./bin/polard add-genesis-account` step, add balance for those
## 3. Clone this ~/../bin/polard home directory into some others, let's say `~/.cloned./bin/polard`
## 4. Run `gentx` in each of those folders
## 5. Copy the `gentx-*` folders under `~/.cloned./bin/polard/config/gentx/` folders into the original `~/../bin/polard/config/gentx`

# Collect genesis tx
polard genesis collect-gentxs --home "$HOMEDIR"

# Run this to ensure everything worked and that the genesis file is setup correctly
polard genesis validate-genesis --home "$HOMEDIR"

if [[ $1 == "pending" ]]; then
    echo "pending mode is on, please wait for the first block committed."
fi

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)m
polard start --pruning=nothing "$TRACE" --log_level $LOGLEVEL --api.enabled-unsafe-cors --api.enable --api.swagger --minimum-gas-prices=0.0001ablack --home "$HOMEDIR"
