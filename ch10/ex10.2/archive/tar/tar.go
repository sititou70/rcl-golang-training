package tar

import (
	"archive/tar"
	"io"
	"os"

	"ex10.2/archive"
)

func init() {
	// https://filesignatures.net/index.php?page=search&search=TAR&mode=EXT
	archive.RegisterFormat(archive.Format{
		Signature: []byte("ustar"),
		Offset:    257,
		FileNames: FileNames,
	})
}

func FileNames(f *os.File) ([]string, error) {
	reader := tar.NewReader(f)

	fileNames := []string{}
	for {
		header, err := reader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		fileNames = append(fileNames, header.Name)
	}

	return fileNames, nil
}
