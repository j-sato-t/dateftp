#!/usr/bin/env pwsh
# mise description="GUI版のビルド (Windows)"
# mise alias="buildg-win"

go build -o dateftp-gui.exe ./cmd/dateftp-gui
