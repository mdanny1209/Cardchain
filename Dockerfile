FROM ignitehq/cli:0.23.0

USER root
RUN apt-get -y -qq update && \
	apt-get install -y -qq curl nodejs npm && \
	apt-get clean
#
# install jq to parse json within bash scripts
RUN curl -o /usr/local/bin/jq http://stedolan.github.io/jq/download/linux64/jq && \
  chmod +x /usr/local/bin/jq

RUN npm i pm2 -g

USER tendermint

EXPOSE 1317
EXPOSE 26657
EXPOSE 26658
EXPOSE 9090
EXPOSE 4500

WORKDIR .
COPY --chown=tendermint:tendermint . .

RUN chmod +x ./docker-run.sh


RUN ignite chain build
RUN ignite chain init

#RUN python3 ./scripts/migrate_with_data.py ./blockchain-data/exported_genesis.json ~/.Cardchain/config/genesis.json

ENTRYPOINT ./docker-run.sh
