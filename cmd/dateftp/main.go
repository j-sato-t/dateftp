package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"

	"github.com/j-sato-t/dateftp/pkg/ftpclient"
)

func main() {
	var (
		host        string
		port        string
		user        string
		password    string
		rootPath    string
		downloadDir string
	)

	// フラグの定義（短縮形も用意）
	flag.StringVarP(&host, "host", "h", "", "FTPホスト")
	flag.StringVarP(&port, "port", "P", "", "FTPポート")
	flag.StringVarP(&user, "user", "u", "", "FTPユーザー")
	flag.StringVarP(&password, "password", "p", "", "FTPパスワード")
	flag.StringVarP(&rootPath, "root-path", "r", "", "FTPルートパス")
	flag.StringVarP(&downloadDir, "download-dir", "d", "", "ダウンロードディレクトリ")

	flag.Parse()

	// 環境変数でのフォールバック
	if host == "" {
		host = os.Getenv("FTP_HOST")
	}
	if port == "" {
		port = os.Getenv("FTP_PORT")
	}
	if user == "" {
		user = os.Getenv("FTP_USER")
	}
	if password == "" {
		password = os.Getenv("FTP_PASSWORD")
	}
	if rootPath == "" {
		rootPath = os.Getenv("FTP_ROOT_PATH")
	}
	if downloadDir == "" {
		downloadDir = os.Getenv("FTP_DOWNLOAD_DIR")
	}

	// 必須パラメータのチェック
	if host == "" || port == "" || user == "" || password == "" || rootPath == "" {
		panic("host, port, user, password, root-path are required via flag or environment variable")
	}

	// downloadDirが無指定の場合はカレントディレクトリ
	if downloadDir == "" {
		cwd, err := os.Getwd()
		if err != nil {
			panic(fmt.Sprintf("カレントディレクトリの取得に失敗しました: %v", err))
		}
		downloadDir = cwd
	}

	// 絶対パスに変換
	absDownloadDir, err := filepath.Abs(downloadDir)
	if err != nil {
		panic(fmt.Sprintf("ダウンロード先ディレクトリの絶対パス取得に失敗しました: %v", err))
	}

	conf := ftpclient.Config{
		Host:        host,
		Port:        port,
		User:        user,
		Password:    password,
		RootPath:    rootPath,
		DownloadDir: absDownloadDir,
	}

	err = ftpclient.Download(conf)
	if err != nil {
		panic(fmt.Sprintf("ダウンロードに失敗しました: %v", err))
	}
	fmt.Println("すべてのダウンロードが完了しました")
}
