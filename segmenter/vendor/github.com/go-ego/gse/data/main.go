package main

import (
	"flag"
	"fmt"

	"github.com/go-ego/gse"
)

var (
	text = flag.String("text", "中国互联网历史上最大的一笔并购案", "要分词的文本")
)

func main() {
	flag.Parse()

	var seg gse.Segmenter
	seg.LoadDict("./data/dict/dictionary.txt")

	segments := seg.Segment([]byte(*text))
	fmt.Println(gse.ToString(segments, true))
}
