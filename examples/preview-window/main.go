package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/chroma/v2/quick"
	"github.com/koki-develop/go-fzf"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	f, err := fzf.New()
	if err != nil {
		log.Fatal(err)
	}

	idxs, err := f.Find(
		files,
		func(i int) string { return files[i].Name() },
		fzf.WithPreviewWindow(func(i, width, height int) string {
			info, _ := files[i].Info()
			// return fmt.Sprintf(
			// 	"Name: %s\nModTime: %s\nSize: %d bytes",
			// 	info.Name(), info.ModTime(), info.Size(),
			// )
			// vp := viewport.New(width, height)
			if info.IsDir() {
				return files[i].Name() + " is a dir"
			} else {
				content, err := os.ReadFile(files[i].Name())
				if err != nil {
					fmt.Println("Error reading file " + files[i].Name())
				}
				var b bytes.Buffer
				err = quick.Highlight(&b, string(content), "go", "terminal16m", "onedark")
				return b.String()
			}
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, i := range idxs {
		fmt.Println(files[i].Name())
	}
}
