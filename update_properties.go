package main

import (
  "strings"
  "log"
  "flag"
  "os"
)

var (
  filename string
)

func isCommentOrBlank(line string) bool {
  return strings.HasPrefix(line, "#") || "" == strings.TrimSpace(line)
}


func Exists(filename string) bool {
  if _, err := os.Stat(filename); err == nil {
  	return true
  }
  return false
}


func main() {
  var filename string
  flag.StringVar(&filename, "f", "filename", "The file location to be parsed")
  flag.Parse()

  if !Exists(filename) {
    log.Fatal("File does not exist")
  }
}
