package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
	"text/tabwriter"
	"time"
)

type DirStat struct {
	bytes int64
	num   int
}

func main() {
	if len(os.Args) != 2 {
		panic(fmt.Sprintf("usage: %s DIR_NAME\n", os.Args[0]))
	}

	rootDirEntris := dirents(os.Args[1])
	entris := make(chan EntryInfo)
	wg := sync.WaitGroup{}
	for _, ent := range rootDirEntris {
		if ent.IsDir() {
			wg.Add(1)
			dir := filepath.Join(os.Args[1], ent.Name())
			go walkDir(dir, dir, &wg, entris)
		}
	}

	go func() {
		wg.Wait()
		close(entris)
	}()

	dirstats := map[string]*DirStat{}
	go func() {
		for {
			printDirStatMap(dirstats)
			time.Sleep(time.Millisecond * 100)
		}
	}()

	for info := range entris {
		if dirstats[info.rootDir] == nil {
			dirstats[info.rootDir] = &DirStat{}
		}

		dirstats[info.rootDir].bytes += info.entrySize
		dirstats[info.rootDir].num++
	}

	printDirStatMap(dirstats)
}

func printDirStatMap(dirstats map[string]*DirStat) {
	keys := []string{}
	for key := range dirstats {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)

	header := []string{"Name", "Files Num", "Size"}
	border := []string{}
	for _, h := range header {
		border = append(border, strings.Repeat("-", len(h)))
	}
	fmt.Fprintf(tw, strings.Join(header, "\t")+"\t\n")
	fmt.Fprintf(tw, strings.Join(border, "\t")+"\t\n")
	for _, key := range keys {
		fmt.Fprintf(tw, "%s\t%d\t%s\t\n", key, dirstats[key].num, formatFileSize(dirstats[key].bytes))
	}

	fmt.Println("")
	tw.Flush()
}

func formatFileSize(bytes int64) string {
	unit := 1000

	adjustedBytes := float64(bytes)
	unitIndex := 0
	for ; int(adjustedBytes) > unit; adjustedBytes /= float64(unit) {
		unitIndex++
	}

	return fmt.Sprintf("%.1f %sB", adjustedBytes, []string{"", "k", "M", "G", "T", "P", "E"}[unitIndex])
}

type EntryInfo struct {
	rootDir   string
	entrySize int64
}

func walkDir(dir string, rootDir string, wg *sync.WaitGroup, entries chan<- EntryInfo) {
	defer wg.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, rootDir, wg, entries)
		} else {
			info, err := entry.Info()
			if err != nil {
				fmt.Fprintf(os.Stderr, "du: %v\n", err)
				continue
			}

			entries <- EntryInfo{rootDir, info.Size()}
		}
	}
}

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.DirEntry {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
