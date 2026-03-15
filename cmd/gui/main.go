package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("DateFTP")
	myWindow.Resize(fyne.NewSize(500, 400))

	// 入力フィールド（PlayHolderとして例を表示）
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("例: 192.168.3.2")

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("例: 4006")

	userEntry := widget.NewEntry()
	userEntry.SetPlaceHolder("例: pc")

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("例: 132934")

	downloadDirLabel := widget.NewLabel("未選択")
	var selectedDir string

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
	)

	// ダウンロード実行ボタン
	startBtn := widget.NewButton("ダウンロード開始", func() {
		if hostEntry.Text == "" || portEntry.Text == "" || userEntry.Text == "" || passEntry.Text == "" || selectedDir == "" {
			dialog.ShowInformation("エラー", "すべての項目を入力し、ダウンロード先を選択してください。", myWindow)
			return
		}

		log.Printf("Connecting to %s:%s as %s", hostEntry.Text, portEntry.Text, userEntry.Text)
		log.Printf("Download destination: %s", selectedDir)

		dialog.ShowInformation("開始", "ダウンロード処理を呼び出します（連携未実装）", myWindow)
	})

	content := container.NewVBox(
		widget.NewLabel("FTPサーバー設定"),
		form,
		widget.NewLabel("ダウンロード先設定"),
		container.NewHBox(dirBtn, downloadDirLabel),
		startBtn,
	)

	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
