#!/usr/bin/env pwsh
# mise description="GUI版のビルド (Windows)"
# mise alias="buildg-win"

go build -ldflags="-H=windowsgui" -o dateftp-gui.exe ./cmd/dateftp-gui
