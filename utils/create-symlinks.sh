#!/usr/bin/env bash


for f in /mnt/12tb-hdd/torrents/lichess*; do
	ln -s "$f" dataset/
done
