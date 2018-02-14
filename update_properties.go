package main

import (
  "strings"
  "bufio"
  "fmt"
  "regexp"
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

func padWithSpace(str, pad string, length int) string {
	for {
		str += pad
		if len(str) > length {
			return str[0:length]
		}
	}
}

func ingest(filename string) (err error) {
  file, err := os.OpenFile(filename, os.O_RDWR, 0666)
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

      if !isPrefix { // If we've reached the end of the line or we're at the EOF, break
        break
      }
    }

    if err == io.EOF || err != nil {
      break
    }

    line    := buffer.String()
    pattern := regexp.MustCompile(`^#([^=]+)=(.+)$`)
    matches := pattern.FindAllStringSubmatch(line, -1) // matches is [][]string

    for _, string := range matches {
      n := strings.Replace(strings.TrimSpace(string[1]), "#", "", -1)
      v := strings.TrimSpace(string[2])
      e := strings.ToLower(strings.Replace(n, ".", "/", -1))

      // newline creation for if block
      nl := "{{if exists \"/" + e + "\" -}}\n"
      nl += padWithSpace(n, " ", 80) + " = {{getv \"/" + e + "\" \"" + v + "\"}}\n"
      nl += "{{end -}}"

      fmt.Println(nl)
      buffer.WriteString(nl)
      file.Sync()
    }

    // dont proccess comments
    if (isCommentOrBlank(line)) {
      if (!strings.Contains(line, "=")) {
        fmt.Println(line)
        buffer.WriteString(line)
        file.Sync()
      }
    } else {
      a := strings.Split(line, "=")
      n := strings.TrimSpace(a[0])
      v := strings.TrimSpace(a[1])
      e := strings.ToLower(strings.Replace(n, ".", "/", -1))

      nl := padWithSpace(n, " ", 80) + " = {{getv \"/" + e + "\" \"" + v + "\"}}"
      fmt.Println(nl)
      buffer.WriteString(nl)
      file.Sync()
    }
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
}
