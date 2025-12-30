# slack_to_gdocs: Slackの箇条書きをGoogle Docsにきれいにコピーできるツール

(逆方向のツール[gdocs_to_slack](https://github.com/fetaro/gdocs_to_slack) もあります

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
Apple siliconのMacOSであれば、 [こちらの実行ファイル](https://raw.githubusercontent.com/fetaro/simplify_clipboard_html/refs/heads/main/dist/arm64/slack_to_gdocs) をダウンロードしてください。

Intell Macの場合は、[こちらの実行ファイル](https://raw.githubusercontent.com/fetaro/simplify_clipboard_html/refs/heads/main/dist/x86_64/slack_to_gdocs) をダウンロードしてください。

### 実行権限の付与
Macのターミナルを開いて、ダウンロードしたファイルに対して以下のコマンドを実行し実行権限を与えてください。

```
chmod 755 ./slack_to_gdocs
```


## 使い方

1. Slackの箇条書きをコピーして、クリップボードにデータが有る状態にします

2. ツールを実行します。ターミナルで以下のように実行してください。

```bash
./slack_to_gdocs
```

クリップボードの内容が書き換わります

3. Google Docsにペーストしてください

### うまく動かないと思ったら

コマンドに -d オプションを付けて実行してください。処理内容が見れます。

```
./slack_to_gdocs -d
```

出力結果
```
=== 変換前のクリップボードのHTMLデータ ===
<meta charset='utf-8'><ul data-stringify-type="unordered-list" data-list-tree="true" class="p-rich_text_list p-rich_text_list__bullet p-rich_text_list--nested" data-indent="0" data-border="0" style="box-sizing: inherit; margin: 0px 0px 0px 20px; padding: 0px; list-style-type: none; margin-inline-start: 0px; padding-inline-start: 0px; color: rgb(29, 28, 29); font-family: Slack-Lato, Slack-Fractions, appleLogo, sans-serif; font-size: 15px; font-style: normal; font-variant-ligatures: common-ligatures; font-variant-caps: normal; font-weight: 400; letter-spacing: normal; orphans: 2; text-align: left; text-indent: 0px; text-transform: none; widows: 2; word-spacing: 0px; -webkit-text-stroke-width: 0px; white-space: normal; background-color: rgb(255, 255, 255); text-decoration-thickness: initial; text-decoration-style: initial; text-decoration-color: initial;"><li data-stringify-indent="0" data-stringify-border="0" style="box-sizing: inherit; margin-inline-start: 28px; margin-bottom: 0px;">箇条書きレベル１<ul data-stringify-type="unordered-list" data-list-tree="true" class="p-rich_text_list p-rich_text_list__bullet p-rich_text_list--nested" data-indent="1" data-border="0" style="box-sizing: inherit; margin: 0px 0px 0px 20px; padding: 0px; list-style-type: none; margin-inline-start: 0px; padding-inline-start: 0px;"><li data-stringify-indent="1" data-stringify-border="0" style="box-sizing: inherit; margin-inline-start: 28px; margin-bottom: 0px;">箇条書きレベル２−１</li><li data-stringify-indent="1" data-stringify-border="0" style="box-sizing: inherit; margin-inline-start: 28px; margin-bottom: 0px;">箇条書きレベル２−２</li></ul></li></ul>
=====================================
=== 変換後のクリップボードのHTMLデータ ===
<html><head><meta charset="utf-8"/></head><body><ul><li>箇条書きレベル１<ul><li>箇条書きレベル２−１</li><li>箇条書きレベル２−２</li></ul></li></ul></body></html>
=====================================
クリップボードのHTMLデータをシンプルにしました
```
