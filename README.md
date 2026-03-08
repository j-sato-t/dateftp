# dateftp
ftp download and sort of date

CLIで操作するFTPクライアントです。FTPサーバー上のファイルをダウンロードし、ファイルのタイムスタンプに基づいて「YYYY/MM/DD」形式のディレクトリ構造でローカルに保存します。

## 主な機能
- FTPサーバーからのファイルの再帰的ダウンロード
- ダウンロードしたファイルのタイムスタンプに基づいた年/月/日のディレクトリ自動生成
- ローカルに同名ファイルが既に存在する場合、タイムスタンプを比較して新しい場合のみダウンロードをスキップ
- `github.com/spf13/pflag` を用いた短縮形フラグでの柔軟な実行設定
- 環境変数でのパラメータのフォールバック対応

## 環境構築
- Go 1.23.0以上（`mise`による環境構築を推奨）

## インストールと依存関係の解決
リポジトリをクローン後、以下のコマンドで依存関係を解決してください。

```bash
go mod tidy
```

## 使用方法

フラグか環境変数を用いて接続先とダウンロード先を指定し、CLIツールを実行します。（フラグの指定が優先されます）

### フラグで指定して実行する例

```bash
go run cmd/cli/main.go \
  -host 192.168.1.10 -port 21 -user admin -password secret \
  -root-path /device/DCIM/PHOTOGRAPHY_PRO -download-dir /home/user/Downloads/dateftp
```

短縮形のフラグを使用することも可能です（`-h`, `-P`, `-u`, `-p`, `-r`, `-d`）。

### 環境変数を指定して実行する例

あらかじめ環境変数をセットしておいた後、以下のように引数なしで実行できます。

```bash
export FTP_HOST=192.168.1.10
export FTP_PORT=21
export FTP_USER=admin
export FTP_PASSWORD=secret
export FTP_ROOT_PATH=/data/ftp/photo
export FTP_DOWNLOAD_DIR=/home/user/Downloads/dateftp

go run cmd/cli/main.go
```

※ `-download-dir` と `FTP_DOWNLOAD_DIR` が指定されない場合は、自動でカレントディレクトリにダウンロードディレクトリが作成されます。
