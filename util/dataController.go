package util

import (
	"io"
	"log"
	"os"
	"strconv"
)

func WriteToFile(data int64) error {

	file, err := os.Create(".data/player.txt")
	if err != nil {
		log.Print("FOO")
		return err
		log.Print("BAR")
	}
	defer file.Close()
	strID := strconv.FormatInt(data, 10)

	if _, err := io.WriteString(file, strID); err != nil {
		return err
	}
	return file.Sync()
}