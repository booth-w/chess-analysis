#!/usr/bin/env bash


if [[ "$1" == "test" ]]; then
	shift
	zstdcat dataset/lichess_db_standard_rated_2013* | bin/chess-anal $@
else
	zstdcat dataset/* | bin/chess-anal $@
fi
