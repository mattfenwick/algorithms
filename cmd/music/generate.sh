#!/bin/bash

set -euo pipefail
set -xv

go run main.go markdown > keys.md
