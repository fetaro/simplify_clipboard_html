# slack_to_gdocs: Slackの箇条書きをGoogle Docsにきれいにコピーできるツール

(逆方向のツール[gdocs_to_slack](https://github.com/fetaro/gdocs_to_slack) もあります)

## 解決したい課題

例えばSlackに以下のような箇条書きがあったとしたときに

![img.png](docs/img.png)

これをコピーしてGoogle Docsに貼り付けると、

![img_4.png](docs/img_4.png)

このように、箇条書きの●や◯の記号が消えるし、文字の背景も変な色になる。

### このツールを使うと

このツールを使うことで、以下のようにきれいに貼り付けられる。

![img_5.png](docs/img_5.png)


## インストール方法

### 実行プログラムのダウンロード
Apple siliconのmacOSであれば、 [こちらの実行ファイル](https://raw.githubusercontent.com/fetaro/simplify_clipboard_html/refs/heads/main/dist/arm64/slack_to_gdocs) をダウンロードしてください。

Intel Macの場合は、[こちらの実行ファイル](https://raw.githubusercontent.com/fetaro/simplify_clipboard_html/refs/heads/main/dist/x86_64/slack_to_gdocs) をダウンロードしてください。

### 実行権限の付与
Macのターミナルを開いて、ダウンロードしたファイルに対して以下のコマンドを実行し実行権限を与えてください。

```
chmod 755 ./slack_to_gdocs
```


## 使い方

1. Slackの箇条書きをコピーして、クリップボードにデータがある状態にします

2. ツールを実行します。ターミナルで以下のように実行してください。

```bash
./slack_to_gdocs
```

クリップボードの内容が書き換わります

3. Google Docsにペーストしてください

### うまく動かないと思ったら

コマンドに -d オプションを付けて実行してください。処理内容が見られます。

```
./slack_to_gdocs -d
```
