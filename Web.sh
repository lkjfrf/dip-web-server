#!/bin/sh

if [ "$1" = "start" ]; then
	sudo nohup ./LoginLauncher &
elif [ "$1" = "stop" ]; then
	pkill -9 -ef ./LoginLauncher
elif [ "$1" = "status" ]; then
	ps -ef | grep ./LoginLauncher
elif [ "$1" = "log" ]; then
	tail -f nohup.out
fi

exit 0
