package function

import (
	"bytes"
	"fmt"

	"github.com/go-ego/gse"
	"io"
	"os"
)

// Handle a serverless request
func Handle(req []byte) string {
	var segmenter gse.Segmenter
	// Loading the default dictionary

	segmenter.LoadDict("vendor/github.com/go-ego/gse/data/dict/jp/dict.txt")

	// segmenter.LoadDict("your gopath"+"/src/github.com/go-ego/gse/data/dict/dictionary.txt")

	// Text Segmentation
	//text := []byte("你好世界, Hello world.")
	fmt.Println("hello world!")
	segments := segmenter.Segment(req)

	// Handle word segmentation results
	// Support for normal mode and search mode two participle,
	// see the comments in the code ToString function.
	//fmt.Println(gse.ToString(segments, false))
	return gse.ToString(segments, false)
}
