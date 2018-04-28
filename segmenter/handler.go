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

	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	segmenter.LoadDict("vendor/github.com/go-ego/gse/data/dict/jp/dict.txt")

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC

	// reading our temp stdout
	fmt.Println("previous output:")
	fmt.Print(out)

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
