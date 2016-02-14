package utils

import (
	"bytes"
	// "fmt"
	"github.com/opesun/goquery"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func boolimgurl(url string) (b bool, filepath string) {
	path := strings.Split(url, "/")
	last := strings.Split(path[len(path)-1], ".")
	fn := ""
	if strings.ToUpper(last[1]) == "JS" || strings.ToUpper(last[1]) == "CSS" {
		for i, v := range path {
			if v != ".." {
				fn += v + "/"
			}
			if i == len(path)-1 {
				fn = fn[:len(fn)-1]
			}
		}
		return true, fn
	} else {
		return false, fn
	}
}
