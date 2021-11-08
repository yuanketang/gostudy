package utils

import (
	"net/http"
	"os"
)

func OutputHtml(w http.ResponseWriter, r *http.Request, fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	stat, _ := file.Stat()
	http.ServeContent(w, r, file.Name(), stat.ModTime(), file)
	return nil
}
