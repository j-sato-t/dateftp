package ftpclient

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/jlaffaye/ftp"
)

type Config struct {
	Host        string
	Port        string
	User        string
	Password    string
	RootPath    string
	DownloadDir string
	LogFunc     func(string)
}

func Download(conf Config) error {
	addr := fmt.Sprintf("%s:%s", conf.Host, conf.Port)
	c, err := ftp.Dial(addr, ftp.DialWithTimeout(10*time.Second))
	if err != nil {
		return fmt.Errorf("FTPサーバー(%s)への接続に失敗しました: %w", addr, err)
	}
	defer c.Quit()

	err = c.Login(conf.User, conf.Password)
	if err != nil {
		return fmt.Errorf("FTPサーバーへのログインに失敗しました: %w", err)
	}

	rootBase := path.Base(conf.RootPath)

	return walkAndDownload(c, conf.RootPath, conf.DownloadDir, rootBase, ".", conf.LogFunc)
}

func walkAndDownload(c *ftp.ServerConn, currentFtpDir, baseDownloadDir, rootBase, relPath string, logFunc func(string)) error {
	entries, err := c.List(currentFtpDir)
	if err != nil {
		return fmt.Errorf("ディレクトリ(%s)の一覧取得に失敗しました: %w", currentFtpDir, err)
	}

	for _, entry := range entries {
		// カレントや親ディレクトリはスキップ
		if entry.Name == "." || entry.Name == ".." {
			continue
		}

		entryPath := path.Join(currentFtpDir, entry.Name)
		entryRelPath := path.Join(relPath, entry.Name)

		if entry.Type == ftp.EntryTypeFolder {
			err := walkAndDownload(c, entryPath, baseDownloadDir, rootBase, entryRelPath, logFunc)
			if err != nil {
				return err
			}
		} else if entry.Type == ftp.EntryTypeFile {
			ftptime := entry.Time

			year := fmt.Sprintf("%04d", ftptime.Year())
			month := fmt.Sprintf("%02d", ftptime.Month())
			day := fmt.Sprintf("%02d", ftptime.Day())

			fileRelDir := path.Dir(entryRelPath)
			if fileRelDir == "." {
				fileRelDir = ""
			}

			targetDir := filepath.Join(baseDownloadDir, rootBase, fileRelDir, year, month, day)
			err = os.MkdirAll(targetDir, 0755)
			if err != nil {
				return fmt.Errorf("ディレクトリ作成に失敗しました (%s): %w", targetDir, err)
			}

			targetFilePath := filepath.Join(targetDir, entry.Name)

			stat, err := os.Stat(targetFilePath)
			if err == nil {
				// 既存ファイルが存在する場合、ローカルが新しければスキップ
				if !stat.ModTime().Before(ftptime) {
					continue
				}
			}

			msg := fmt.Sprintf("ダウンロード中: %s -> %s", entryPath, targetFilePath)
			if logFunc != nil {
				logFunc(msg)
			} else {
				fmt.Println(msg)
			}

			err = downloadFile(c, entryPath, targetFilePath, ftptime)
			if err != nil {
				return fmt.Errorf("ファイルのダウンロードに失敗しました (%s): %w", entryPath, err)
			}
		}
	}
	return nil
}

func downloadFile(c *ftp.ServerConn, src, dest string, mtime time.Time) error {
	r, err := c.Retr(src)
	if err != nil {
		return err
	}
	defer r.Close()

	tempPath := dest + ".tmp"
	f, err := os.Create(tempPath)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, r)
	f.Close()
	if err != nil {
		os.Remove(tempPath)
		return err
	}

	err = os.Rename(tempPath, dest)
	if err != nil {
		os.Remove(tempPath)
		return err
	}

	// タイムスタンプをFTPの更新日時に変更する
	return os.Chtimes(dest, mtime, mtime)
}
