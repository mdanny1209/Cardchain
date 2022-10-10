#!/bin/bash

echo -e "'\033[0;31m' fasten your seatbelts '\033[0m'"
FAUCET_SECRET_KEY=""
CHAIN_ID=Cardchain



#Cardchaind unsafe-reset-all
echo  "Getting Genesis file..."
wget -O $HOME/.Cardchain/config/genesis.json "https://raw.githubusercontent.com/DecentralCardGame/Testnet/main/genesis.json"
echo  "Getting Addrbook file..."
wget -qO $HOME/.Cardchain/config/addrbook.json https://github.com/AlexToTheMoon/AM-Solutions/raw/main/addrbooks/addrbook-crowd.json
SEEDS=""
PEERS="a89083b131893ca8a379c9b18028e26fa250473c@159.69.11.174:36656"; \
sed -i.bak -e "s/^seeds *=.*/seeds = \"$SEEDS\"/; s/^persistent_peers *=.*/persistent_peers = \"$PEERS\"/" $HOME/.Cardchain/config/config.toml

SNAP_RPC="http://161.97.167.120:26657"

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
sed -i -e "s/^SECRET_KEY *=.*/SECRET_KEY = \"$FAUCET_SECRET_KEY\"/" go-faucet-master/.env
cd go-faucet-master
#screen -AmdS faucet go-faucet
./go-faucet &

Cardchaind start


#ignite chain serve

# the following line is evaluated if csd is terminated via pkill (docker-stop-and-export.sh)
#Cardchaind export > ~/.ignite/local-chains/Cardchain/exported_genesis.json
