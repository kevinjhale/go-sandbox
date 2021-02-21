package wiki

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}