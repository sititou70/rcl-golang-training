package zip

import (
	"archive/zip"
	"os"

	"ex10.2/archive"
)

func init() {
	// https://filesignatures.net/index.php?page=search&search=ZIP&mode=EXT
	archive.RegisterFormat(archive.Format{
		Signature: []byte("PK"),
		Offset:    0,
		FileNames: FileNames,
	})
}

func FileNames(f *os.File) ([]string, error) {
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	reader, err := zip.NewReader(f, stat.Size())
	if err != nil {
		return nil, err
	}

	fileNames := []string{}
	for _, file := range reader.File {
		fileNames = append(fileNames, file.Name)
	}

	return fileNames, nil
}
