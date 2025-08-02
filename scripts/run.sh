#!/usr/bin/env bash


if [[ "$1" == "test" ]]; then
	shift
	zstdcat dataset/lichess_db_standard_rated_2013* | go run main.go $@
else
	zstdcat dataset/* | go run main.go $@
fi
