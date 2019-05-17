#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/application/devOpsWebServer/
git pull https://github.com/sunhongyue4500/devOpsWebServer.git
cd webserver
# & 跑在后台
./webserver &