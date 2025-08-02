#!/usr/bin/env bash


go build -o bin/chess-anal main.go
hyperfine "zstdcat dataset/lichess_db_standard_rated_2013* | bin/chess-anal $@"
