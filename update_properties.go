package main

import (
  "strings"
  "bufio"
  "fmt"
  "flag"
  "os"
  "io"
  "bytes"
)

var (
  filename  string
)

func isCommentOrBlank(line string) bool {
  return strings.HasPrefix(line, "#") || "" == strings.TrimSpace(line)
}

func fileExists(filename string) bool {
  if _, err := os.Stat(filename); err == nil {
  	return true
  }
  return false
}

func limitLength(s string, length int) string {
  if len(s) < length {
    return s
  }
  return s[:length]
}


func ingest(filename string) (err error) {
  file, err := os.Open(filename)
  defer file.Close()

  if err != nil {
    return err
  }

  reader := bufio.NewReader(file) // Start reading from the file with a reader.

  for {
    var buffer bytes.Buffer
    var l []byte
    var isPrefix bool

    for {
      l, isPrefix, err = reader.ReadLine()
      buffer.Write(l)

      if !isPrefix || err != nil { // If we've reached the end of the line or we're at the EOF, break
        break
      }
    }

    if err == io.EOF {
      break
    }

    line := buffer.String()
    fmt.Printf(" > Read %d characters\n", len(line))
    fmt.Println(" > > " + limitLength(line, 50))      // Process the line here.
  }

  if err != io.EOF {
    fmt.Printf(" > Failed!: %v\n", err)
  }
  return
}

func main() {

  flag.StringVar(&filename, "f", "filename", "The file location to be parsed")
  flag.Parse()

  if !fileExists(filename) {
    fmt.Println(" > File does not exist\n")
  }
  ingest(filename)

  // ## needs to go up
  //   lineString  := string(bytes)
  //   pattern     := regexp.MustCompile(`^#([^=]+)=(.+)$`)
  //
  //   for _, match := range pattern.FindAllStringSubmatch(lineString, -1) {
  //     fmt.Printf("key=%s, value=%s\n", match[1], match[2])
  //   }
  //
  //   // for _, element := range RegSplit(lineString, "/^#([^=]+)=(.+)$/") {
  //   //   name  := element[1]
  //   //   value := element[2]
  //   //
  //   //   fmt.Printf("Name  %v\n", name)
  //   //   fmt.Printf("Value %v\n", value)
  //   // }
  //
  //   if (isCommentOrBlank(lineString)) {
  //     file.Seek(int64(-(len(lineString))), 1)
  //     file.WriteString(line)
  //   }
  //
  //     fmt.Printf(lineString + "\n")
  // }
}
