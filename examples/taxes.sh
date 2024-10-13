#!/bin/bash

set -xv
set -euo pipefail

go run ../cmd/main.go estimate --config ./taxes.json
