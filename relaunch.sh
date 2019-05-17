#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/application/devOpsWebServer/
# 拉取最新代码
git pull https://github.com/sunhongyue4500/devOpsWebServer.git
cd webserver
# & 跑在后台
./webserver &