package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>
#include <stdlib.h>

char* getHTMLFromPasteboard() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSString *html = [pasteboard stringForType:@"public.html"];
    if (html == nil) {
        return NULL;
    }
    return strdup([html UTF8String]);
}

void setHTMLToPasteboard(const char* html) {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    [pasteboard clearContents];
    NSString *htmlString = [NSString stringWithUTF8String:html];
    [pasteboard setString:htmlString forType:@"public.html"];
}
*/
import "C"
import (
	"flag"
	"fmt"
	"strings"
	"unsafe"

	"golang.org/x/net/html"
)

func removeAttributes(node *html.Node) {
	if node.Type == html.ElementNode && node.Data != "meta" {
		// metaタグ以外の要素ノードの場合、属性を全て削除
		node.Attr = nil
	}
	// 子ノードに対して再帰的に処理
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		removeAttributes(c)
	}
}

func main() {
	// コマンドライン引数の処理
	debug := flag.Bool("debug", false, "デバッグモードを有効にする")
	debugShort := flag.Bool("d", false, "デバッグモードを有効にする（-debugのショートカット）")
	flag.Parse()

	// デバッグモードの判定
	isDebug := *debug || *debugShort

	clipboardHTML := C.getHTMLFromPasteboard()
	if clipboardHTML == nil {
		fmt.Println("クリップボードにHTMLデータがありません")
		return
	}
	defer C.free(unsafe.Pointer(clipboardHTML))

	// Cの文字列をGoの文字列に変換
	originalText := C.GoString(clipboardHTML)

	// デバッグモードの場合、元のHTMLを表示
	if isDebug {
		fmt.Println("=== 元のHTML ===")
		fmt.Println(originalText)
		fmt.Println("===============")
	}

	// HTMLをパース
	doc, err := html.Parse(strings.NewReader(originalText))
	if err != nil {
		fmt.Println("HTMLのパースに失敗しました:", err)
		return
	}

	// 属性を除去
	removeAttributes(doc)

	// 処理済みのHTMLを文字列化
	var buf strings.Builder
	if err := html.Render(&buf, doc); err != nil {
		fmt.Println("HTMLのレンダリングに失敗しました:", err)
		return
	}
	result := buf.String()

	// デバッグモードの場合、変換後のHTMLを表示
	if isDebug {
		fmt.Println("=== 変換後のHTML ===")
		fmt.Println(result)
		fmt.Println("===================")
	}

	// クリップボードにHTML形式で書き込む
	cstr := C.CString(result)
	defer C.free(unsafe.Pointer(cstr))
	C.setHTMLToPasteboard(cstr)
}
