# 💎 go-errorify

A simple zero-dependencies multilingual error wrapper for your Go (Golang) apps or CLI.

## Languages

🇺🇸 English, 🇷🇺 Russian, 🇪🇸 Español, ...

## Requirements

- Go `1.11+`
- Go Modules

## Install

```console
go get -u github.com/koddr/go-errorify
```

## How to use `go-errorify`?

```go
package main

import (
	"github.com/koddr/go-errorify"
	"os"
)

func main() {
  // Start creating file
  file, err := os.Create("./folder/file_1.txt")
  // Checking error
  errorify.Simple(err)
  // Close file after finish
  defer file.Close()

  // ...
}
```

If something went wrong (no such file or directory), output is:

```console
✘ Error: open ./folder/file_1.txt: no such file or directory
```

![console_output](https://user-images.githubusercontent.com/11155743/70362590-8db79080-1896-11ea-80f8-bbbcb3de179d.png)

## Developers

- Idea and active development by [Vic Shóstak](https://github.com/koddr) (aka Koddr).

## Project assistance

If you want to say «thank you» or/and support active development `go-errorify`:

1. Add a GitHub Star to project.
2. Twit about project [on your Twitter](https://twitter.com/intent/tweet?text=A%20dead%20simple%20zero-dependencies%20%23multilingual%20error%20wrapper%20for%20your%20Go%20%28%40Golang%29%20apps%20or%20%23CLI%20%F0%9F%91%8D%20https%3A%2F%2Fgithub.com%2Fkoddr%2Fgo-errorify).
3. If you want, send to project's author some money via PayPal: [@paypal.me/koddr](https://paypal.me/koddr?locale.x=en_EN).

## License

MIT
