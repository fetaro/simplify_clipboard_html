# simplify_clipboard_html_go

## 概要
macOSのクリップボードからHTML形式のデータを取得し、metaタグ以外の全てのタグの属性を除去したHTMLを出力するGoプログラムです。

## 特徴
- macOSのネイティブAPI（NSPasteboard）を直接利用
- 取得したHTMLからmetaタグ以外の全てのタグの属性を除去
- コマンドラインで簡単に実行可能
- デバッグモードで変換前後のHTMLを確認可能

## インストール方法

[simplufy_clipbpard_html](simplufy_clipbpard_html) をダウンロードして好きなフォルダにおいてください


## 実行方法
クリップボードに自校して下しあ

```
./simplify_clipboard_html
```

#### デバッグモード
変換前後のHTMLを確認したい場合は、以下のいずれかのオプションを使用します：
```
go run main.go --debug
```
または
```
go run main.go -d
```

## 使い方
1. クリップボードにHTML形式のデータをコピーします。
2. 上記のコマンドでプログラムを実行します。
3. metaタグ以外の全てのタグの属性が除去されたHTMLがクリップボードに書き込まれます。
4. デバッグモードを使用した場合は、変換前後のHTMLが標準出力に表示されます。

## 注意事項
- 本プログラムはmacOS専用です。
- クリップボードにHTMLデータが存在しない場合はエラーメッセージが表示されます。

## 自分でビルドする場合

### 必要要件
- macOS
- Go 1.23 以上

### 依存パッケージ
- golang.org/x/net/html

### インストール
```
go get golang.org/x/net/html
```

### ビルド方法
```
go build -o simplify_clipboard_html main.go
```

### 実行方法
```
go run main.go
```
またはビルド後：
```
./simplify_clipboard_html
```
