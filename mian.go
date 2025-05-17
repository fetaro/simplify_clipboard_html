package main

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Cocoa
#import <Cocoa/Cocoa.h>

char* getHTMLFromPasteboard() {
    NSPasteboard *pasteboard = [NSPasteboard generalPasteboard];
    NSString *html = [pasteboard stringForType:@"public.html"];
    if (html == nil) {
        return NULL;
    }
    return strdup([html UTF8String]);
}
*/
import "C"
import (
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
	clipboardHTML := C.getHTMLFromPasteboard()
	if clipboardHTML == nil {
		fmt.Println("クリップボードにHTMLデータがありません")
		return
	}
	defer C.free(unsafe.Pointer(clipboardHTML))

	// Cの文字列をGoの文字列に変換
	text := C.GoString(clipboardHTML)

	// HTMLをパース
	doc, err := html.Parse(strings.NewReader(text))
	if err != nil {
		fmt.Println("HTMLのパースに失敗しました:", err)
		return
	}

	// 属性を除去
	removeAttributes(doc)

	// 処理済みのHTMLを出力
	var buf strings.Builder
	if err := html.Render(&buf, doc); err != nil {
		fmt.Println("HTMLのレンダリングに失敗しました:", err)
		return
	}
	fmt.Print(buf.String())
}
