package tarextract

import (
	"github.com/mholt/archiver"
)

func Extract(srcFile string, destination string) error {
	err := archiver.Unarchive(srcFile, destination)
	return err
}
