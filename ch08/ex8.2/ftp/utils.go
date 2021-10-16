package ftp

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"text/tabwriter"
)

func writeFTPData(address string, data io.Reader) error {
	c, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	defer c.Close()

	io.Copy(c, data)

	return nil
}

func listDir(path string) (*bytes.Buffer, error) {
	b := new(bytes.Buffer)
	tw := new(tabwriter.Writer).Init(b, 0, 8, 2, ' ', 0)

	header := []string{"Permissions", "Size", "Date Modified", "Name"}
	border := []string{}
	for _, h := range header {
		border = append(border, strings.Repeat("-", len(h)))
	}
	fmt.Fprintf(tw, strings.Join(header, "\t")+"\r\n")
	fmt.Fprintf(tw, strings.Join(border, "\t")+"\r\n")

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			return nil, err
		}

		fmt.Fprintf(tw, "%v\t%v\t%v\t%v\r\n", info.Mode().String(), info.Size(), info.ModTime().String(), file.Name())
	}

	tw.Flush()
	return b, nil
}
