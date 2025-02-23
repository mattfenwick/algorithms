#!/bin/bash

set -euo pipefail
set -xv

go run main.go keys --order fifths > keys.md

go run main.go keys --order chromatic > keys-chromatic.md

go run main.go scalesandchords > scales-and-chords.md
