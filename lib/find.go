package lib

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const rfcUrlTemplate = "https://www.rfc-editor.org/rfc/rfc%d.txt"

func Find(rfcNum int) (string, error) {
	var (
		cacheDir, err = getCacheDir()
		useCache      = true
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[WARNING] Can't resolve the cache directory. RFCs won't be cached.\n")
		useCache = false
	}

	if useCache {
		docPath := filepath.Join(cacheDir, fmt.Sprintf("%d.txt", rfcNum))
		document, err := getFromCache(docPath)
		if err == nil {
			// Found in cache, return
			return document, nil
		}

		document, err = fetchDoc(rfcNum)
		if err != nil {
			return "", err
		}

		err = writeToCache(document, docPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "[WARNING] Couldn't save doc to cache (attempted at '%s')", docPath)
		}

		return document, nil
	}

	document, err := fetchDoc(rfcNum)
	if err != nil {
		return "", err
	}

	return document, nil
}

func getCacheDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	rfcCache := filepath.Join(cacheDir, "rfc")

	err = os.MkdirAll(rfcCache, 0755)
	if err != nil {
		return "", err
	}

	return rfcCache, nil
}

// fetchDoc fetches the RFC document from the internet using an HTTP GET request.
func fetchDoc(rfcNum int) (string, error) {
	res, err := http.Get(fmt.Sprintf(rfcUrlTemplate, rfcNum))
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	documentBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(documentBytes), nil
}

// getFromCache returns the document's content or error, if the document wasn't found
// or couldn't be opened.
func getFromCache(docPath string) (string, error) {
	document, err := os.ReadFile(docPath)
	if err != nil {
		return "", err
	}

	return string(document), nil
}

// writeToCache writes the document into the passed-in path, which should be
// the path for the RFC document inside the cache directory.
func writeToCache(doc string, path string) error {
	return os.WriteFile(path, []byte(doc), 0644)
}
