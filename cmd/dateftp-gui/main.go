package main

import (
	"fmt"

	"github.com/j-sato-t/dateftp/pkg/ftpclient"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.NewWithID("com.github.dateftp.j-sato-t")
	myWindow := myApp.NewWindow("DateFTP")
	myWindow.Resize(fyne.NewSize(1600, 900))

	prefs := myApp.Preferences()

	// 入力フィールド（PlayHolderとして例を表示）
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("例: 192.168.3.2")
	hostEntry.SetText(prefs.StringWithFallback("host", ""))

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("例: 4006")
	portEntry.SetText(prefs.StringWithFallback("port", ""))

	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("例: pc")
	userEntry.SetText(prefs.StringWithFallback("user", ""))

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("例: 132934")

	rootPathEntry := widget.NewEntry()
	rootPathEntry.SetPlaceHolder("例: /device/DCIM/PHOTOGRAPHY_PRO")
	rootPathEntry.SetText(prefs.StringWithFallback("rootPath", ""))

	selectedDir := prefs.StringWithFallback("downloadDir", "")
	downloadDirLabelText := "未選択"
	if selectedDir != "" {
		downloadDirLabelText = selectedDir
	}
	downloadDirLabel := widget.NewLabel(downloadDirLabelText)

	// ディレクトリ選択用ボタン
	dirBtn := widget.NewButton("ダウンロード先を選択", func() {
		dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				return
			}
			selectedDir = uri.Path()
			downloadDirLabel.SetText(selectedDir)
		}, myWindow)
	})

	form := widget.NewForm(
		widget.NewFormItem("ホスト名", hostEntry),
		widget.NewFormItem("ポート番号", portEntry),
		widget.NewFormItem("ユーザー名", userEntry),
		widget.NewFormItem("パスワード", passEntry),
		widget.NewFormItem("ルートパス", rootPathEntry),
	)

	// ログ表示用データとリスト
	logData := binding.NewStringList()
	logList := widget.NewListWithData(logData,
		func() fyne.CanvasObject {
			return widget.NewLabel("Template log line to reserve height...")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})

	// ダウンロード実行ボタン
	startBtn := widget.NewButton("ダウンロード開始", func() {
		if hostEntry.Text == "" || portEntry.Text == "" || userEntry.Text == "" || passEntry.Text == "" || rootPathEntry.Text == "" || selectedDir == "" {
			dialog.ShowInformation("エラー", "すべての項目を入力してください。", myWindow)
			return
		}

		// 設定の保存
		prefs.SetString("host", hostEntry.Text)
		prefs.SetString("port", portEntry.Text)
		prefs.SetString("user", userEntry.Text)
		prefs.SetString("rootPath", rootPathEntry.Text)
		prefs.SetString("downloadDir", selectedDir)

		logData.Set([]string{})
		logData.Append("===============================")
		logData.Append(fmt.Sprintf("Connecting to %s:%s as %s", hostEntry.Text, portEntry.Text, userEntry.Text))
		logData.Append(fmt.Sprintf("Root path: %s", rootPathEntry.Text))
		logData.Append(fmt.Sprintf("Download destination: %s", selectedDir))

		conf := ftpclient.Config{
			Host:        hostEntry.Text,
			Port:        portEntry.Text,
			User:        userEntry.Text,
			Password:    passEntry.Text,
			RootPath:    rootPathEntry.Text,
			DownloadDir: selectedDir,
			LogFunc: func(msg string) {
				logData.Append(msg)
				logList.ScrollToBottom()
			},
		}

		// UI操作の無効化（開始ボタンの連打防止などが必要な場合はここでDisableにする）
		// ここでは簡略化のため、すぐにgoroutineで処理を開始します
		go func() {
			logData.Append("ダウンロード処理を開始します...")
			err := ftpclient.Download(conf)
			if err != nil {
				logData.Append(fmt.Sprintf("エラー発生: %v", err))
				dialog.ShowError(err, myWindow)
			} else {
				logData.Append("すべてのダウンロードが完了しました")
				dialog.ShowInformation("完了", "すべてのダウンロードが完了しました", myWindow)
			}
			logList.ScrollToBottom()
		}()
	})

	topContent := container.NewVBox(
		widget.NewLabel("FTPサーバー設定"),
		form,
		widget.NewLabel("ダウンロード先設定"),
		container.NewHBox(dirBtn, downloadDirLabel),
		startBtn,
		widget.NewLabel("実行ログ:"),
	)

	// Borderコンテナを使って、リストが残りの領域（Center）を占有するようにする
	content := container.NewBorder(topContent, nil, nil, nil, logList)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
