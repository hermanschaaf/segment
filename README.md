Segment
=======

A segmentation package for Go. Right now this only supports sentence-level segmentation, and works best with Germanic languages (such as English):

### Usage

```go
package main

import "github.com/hermanschaaf/segment"

func main() {
	paragraph := "Here is some text. \"Please split it. Then print it,\" she said. \"Then Mr. Johnson wants to speak with you\""
	sentences := Sentence(paragraph)
	for _, s := range sentences {
		fmt.Println(strings.Trim(s, " "))
	}
}

```

This automatically handles quoted text as you would expect, and will print all the sentences, one per line:

```
Here is some text.
"Please split it. Then print it," she said.
"Then Mr. Johnson wants to speak with you"
```