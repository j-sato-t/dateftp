# DateFTP

DateFTP は、FTPサーバーからファイルの更新日時を基準にディレクトリごとに一括ダウンロードするためのツールです。
このツールは **バイブコーディング（Vibe Coding）** によって作成されました。

## 目的

- スマホ内の写真やゲームのスクリーンショットなど **同じディレクトリに大量の同種ファイルがある** ケースを想定したftpダウンローダ
- スマホをFTPサーバとしてLANに公開するファイラアプリなどを使用する想定

## 機能

- ファイルの更新日時ごとにサブディレクトリを作成してダウンロード
  - 例: 20260321_123456_001.jpg → 2026/03/21/20260321_123456_001.jpg
- CLIおよびGUIの両方のインターフェースを提供

## CLI版 (`dateftp`)

### インストール方法

`go install` コマンドを使用してCLIツールをインストールできます。以下のコマンドを実行することで、 `dateftp` コマンドが利用可能になります。

```bash
go install github.com/j-sato-t/dateftp/cmd/dateftp@cmd/dateftp/v1.0.0
```
※ 実行にはGoがインストールされており、リポジトリのルートディレクトリで実行している必要があります。環境変数 `GOPATH/bin`（デフォルトは `~/go/bin` または `%USERPROFILE%\go\bin`）にパスが通っていることを確認してください。

### 実行方法

`dateftp` は、コマンドライン引数（フラグ）または環境変数を使用して設定を渡すことができます。

#### フラグでの実行方法

```bash
dateftp --host "192.168.1.100" \
        --port "21" \
        --user "ftpuser" \
        --password "ftppass" \
        --root-path "/path/to/remote/dir" \
        --download-dir "./downloads"
```

指定可能なフラグは以下の通りです：
- `-h, --host`: FTPホスト（必須）
- `-P, --port`: FTPポート（必須）
- `-u, --user`: FTPユーザー（必須）
- `-p, --password`: FTPパスワード（必須）
- `-r, --root-path`: FTPルートパス（必須）
- `-d, --download-dir`: ダウンロード先ディレクトリ（省略時はカレントディレクトリ）

#### 環境変数での実行方法

環境変数をエクスポートしてから実行することも可能です。フラグと環境変数の両方が存在する場合は、フラグが優先されます。

```bash
export FTP_HOST="192.168.1.100"
export FTP_PORT="21"
export FTP_USER="ftpuser"
export FTP_PASSWORD="ftppass"
export FTP_ROOT_PATH="/path/to/remote/dir"
export FTP_DOWNLOAD_DIR="./downloads"

dateftp
```

## GUI版 (`dateftp-gui`)

GUI版を利用することで、直感的な画面から設定を入力し、ダウンロードを実行することができます。また、便利な実行ログ表示機能も備えています。

### ビルド方法

Fyneを利用しているため、必要な依存関係がシステムにインストールされている必要があります（WindowsではCコンパイラなど）。
リポジトリのルートから以下のコマンドでビルドまたは実行を行います。

```bash
# 実行
go run ./cmd/dateftp-gui

# ビルド
go build -o dateftp-gui ./cmd/dateftp-gui
```

### 入力内容の保存について

GUI版では、一度入力した「ホスト名」「ポート番号」「ユーザー名」「ルートパス」「ダウンロード先」が自動的に保存され、次回起動時に復元されます。（※セキュリティ上の観点から、パスワードは保存されません）

**保存される場所の例:**
OSに組み込まれた設定保存の仕組み（FyneのPreferences）を利用しています。
- Windows: `C:\Users\<ユーザー名>\AppData\Roaming\com.github.dateftp.j-sato-t\preferences.json`
- macOS: `~/Library/Preferences/com.github.dateftp.j-sato-t/preferences.json`
- Linux: `~/.config/com.github.dateftp.j-sato-t/preferences.json`
