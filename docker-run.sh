#!/bin/bash

# TODO
#if false && test -f "/var/blockchain/exported_genesis.json"; then

echo -e "'\033[0;31m' fasten your seatbelts '\033[0m'"



NODE_MONIKER=ccRocks
CHAIN_ID=Cardchain

#PEERS="61f05a01167b1aec59275f74c3d7c3dc7e9388d4@45.136.28.158:26658"

#echo  "Downloading Binary..."
#curl https://get.ignite.com/DecentralCardGame/Cardchain@latest! | sudo bash


#echo  "Initializing Cardchain..."
#Cardchaind config chain-id $CHAIN_ID
#Cardchaind init $NODE_MONIKER --chain-id $CHAIN_ID


#echo  "Getting Genesis file..."
#cp $HOME/Testnet1/genesis.json $HOME/.Cardchain/config/genesis.json



#Cardchaind unsafe-reset-all
wget -O $HOME/.Cardchain/config/addrbook.json "https://raw.githubusercontent.com/StakeTake/guidecosmos/main/CrowdControl/Cardchain/addrbook.json"
SEEDS=""
PEERS="a89083b131893ca8a379c9b18028e26fa250473c@159.69.11.174:36656"; \
sed -i.bak -e "s/^seeds *=.*/seeds = \"$SEEDS\"/; s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" $HOME/.Cardchain/config/config.toml
SNAP_RPC="http://cc.stake-take.com:36657"
LATEST_HEIGHT=$(curl -s $SNAP_RPC/block | jq -r .result.block.header.height); \
BLOCK_HEIGHT=$((LATEST_HEIGHT - 2000)); \
TRUST_HASH=$(curl -s "$SNAP_RPC/block?height=$BLOCK_HEIGHT" | jq -r .result.block_id.hash)
echo $LATEST_HEIGHT $BLOCK_HEIGHT $TRUST_HASH

sed -i.bak -E "s|^(enable[[:space:]]+=[[:space:]]+).*$|\1true| ; \
s|^(rpc_servers[[:space:]]+=[[:space:]]+).*$|\1\"$SNAP_RPC,$SNAP_RPC\"| ; \
s|^(trust_height[[:space:]]+=[[:space:]]+).*$|\1$BLOCK_HEIGHT| ; \
s|^(trust_hash[[:space:]]+=[[:space:]]+).*$|\1\"$TRUST_HASH\"| ; \
s|^(seeds[[:space:]]+=[[:space:]]+).*$|\1\"\"|" $HOME/.Cardchain/config/config.toml

# config pruning
indexer="kv"
pruning="custom"
pruning_keep_recent="100"
pruning_keep_every="0"
pruning_interval="10"

sed -i -e "s/^indexer *=.*/indexer = \"$indexer\"/" $HOME/.Cardchain/config/config.toml
sed -i -e "s/^pruning *=.*/pruning = \"$pruning\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-keep-recent *=.*/pruning-keep-recent = \"$pruning_keep_recent\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-keep-every *=.*/pruning-keep-every = \"$pruning_keep_every\"/" $HOME/.Cardchain/config/app.toml
sed -i -e "s/^pruning-interval *=.*/pruning-interval = \"$pruning_interval\"/" $HOME/.Cardchain/config/app.toml

echo -e "'\033[0;31m' starting faucet '\033[0m'"
pm2 start ./faucetserver/server.js

Cardchaind start


# empty blocks would be nice, but this doesn't start the faucet
#Cardchaind start --consensus.create_empty_blocks false

#ignite chain serve

# the following line is evaluated if csd is terminated via pkill (docker-stop-and-export.sh)
#Cardchaind export > ~/.ignite/local-chains/Cardchain/exported_genesis.json
