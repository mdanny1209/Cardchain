#!/bin/bash

echo "please give write permission to backup folder:"
sudo chmod o+w backup

#docker-compose exec blockchain pkill Cardchaind

echo "writing backup"

# this code finds the process id, it is no longer necessary
#PID=$(docker-compose exec blockchain pidof Cardchaind)
#echo PID:$PID
docker-compose exec blockchain pkill Cardchaind

while $(docker-compose exec blockchain pkill -0 Cardchaind); do
	sleep 1
done
