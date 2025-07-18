// cmd/fractalvault/main.go
package main

import (
"flag"
"log"
"os"

"fractalvault/internal/fractalvault"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := fractalvault.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
