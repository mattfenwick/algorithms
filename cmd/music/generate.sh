#!/bin/bash

set -euo pipefail
set -xv

go run main.go keys > keys.md

go run main.go scalesandchords > scales-and-chords.md
