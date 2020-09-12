package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type caseInsensitive struct {
  values []string
}

func (ci caseInsensitive) Len() int {
  return len(ci.values)
}

func (ci caseInsensitive) Less(i, j int) bool {
  return strings.ToLower(ci.values[i]) < strings.ToLower(ci.values[j])
}

func (ci caseInsensitive) Swap(i,j int) {
  ci.values[i], ci.values[j] = ci.values[j], ci.values[i]
}

func walkDir(path, indent string) (dirs, files int, err error) {
  fi, err := os.Stat(path)
  if err != nil {
    return 0, 0, fmt.Errorf("stat %s: %v", path, err)
  }
  if !fi.IsDir() {
    return 0, 1, nil
  }
  dir, err := os.Open(path)
  if err != nil {
    return 1, 0, fmt.Errorf("open: %s %v", path, err)
  }
  names, err := dir.Readdirnames(-1)
  _ = dir.Close() // It's safe to ignore this error
  if err != nil {
    return 1, 0, fmt.Errorf("read dir names %s: %v", path, err)
  }
  names = removeHiddenFiles(names)
  sort.Sort(caseInsensitive{names})
  add := "│   "
  for i, name := range names {
    if i == len(names)-1 {
      fmt.Println(indent + "└── " + name)
      add = "  "
    } else {
      fmt.Println(indent + "├── " + name)
    }
    d, _, err := walkDir(filepath.Join(path, name), indent + add)
    if err != nil {
      log.Println(err)
    }
    dirs, files = dirs+d, files+d
  }
  return dirs+1, files, nil
}

func removeHiddenFiles(files []string) []string {
  var unHiddenFiles []string
  for _, f := range files {
    if f[0] != '.' {
      unHiddenFiles = append(unHiddenFiles, f)
    }
  }
  return unHiddenFiles
}

func main() {
  path := "../"
  path, err := filepath.Abs(path)
  if err != nil {
    log.Fatalf("absolute %s: %v", path, err)
  }
  _, _, err = walkDir(path, "")
  if err != nil {
    log.Fatal(err)
  }
}
