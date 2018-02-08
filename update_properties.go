package main

import (
    "fmt"
    "flag"
    "os"
)


func Check(e error) {
  if e != nil {
    panic(e)
  }
}


func Exists(filename string) bool {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    return false
  }
  return true
}


func main() {
  var filename string

  flag.StringVar(&filename, "f", "filename", "The file location to be parsed")
  if _, err := Exists(filename) {
    Check(err)
  }
}
