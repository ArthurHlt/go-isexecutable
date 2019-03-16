# go-isexecutable

Lib for go to check if a given reader is an executable file.

## Usage

```go
package main

import (
	"github.com/ArthurHlt/go-isexecutable"
	"os"
	"fmt"
)

func main() {
	f, _ := os.Open("my-file-executable")
	// if closeAfterCheck is true and reader is a closable reader
	// isexecutable will close itself the reader after reading the first 4bytes of your file
	fmt.Println(isexecutable.IsExecutable(f, true))
	// show true
}
```