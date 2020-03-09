package util

import (
	"github.com/gohouse/golib/file"
	"github.com/gohouse/golib/random"
)

func GetRandFileFromDir(dir string) (str string) {
	var allfiles []string
	var err error
	allfiles, err = file.GetAllFiles(dir)

	if err != nil {
		return
	}
	if len(allfiles) == 0 {
		return
	}
	return allfiles[random.RandBetween(0, len(allfiles)-1)]
}
