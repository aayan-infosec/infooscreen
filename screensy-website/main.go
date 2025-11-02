package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/text/language"
)

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"golang.org/x/text/language"
)

var (
	initOnce  sync.Once
	fileCache [][]byte
	matcher   language.Matcher
	fs        http.Handler
)

func Handler(w http.ResponseWriter, r *http.Request) {
	initOnce.Do(func() {
		var err error
		fileCache, matcher = fetchTranslations()
		fs = http.FileServer(http.Dir("."))
	})

	if r.URL.Path == "/" || r.URL.Path == "/index.html" {
		acceptLanguageHeader := r.Header.Get("Accept-Language")
		tags, _, _ := language.ParseAcceptLanguage(acceptLanguageHeader)
		_, idx, _ := matcher.Match(tags...)
		fileContent := fileCache[idx]

		http.ServeContent(w, r, "index.html", time.Time{}, bytes.NewReader(fileContent))
	} else {
		fs.ServeHTTP(w, r)
	}
}


func fetchTranslations() ([][]byte, language.Matcher) {
	// Get the filepaths of all translations
	filePaths, err := filepath.Glob("./translations/*.html")

	if err == filepath.ErrBadPattern {
		panic("Invalid pattern during fetchTranslations")
	}

	// Print a list of all found translation filepaths
	log.Printf("Registering the following %d translation files:", len(filePaths))
	for idx, filePath := range filePaths {
		log.Printf("%3d. %s\n", idx+1, filePath)
	}

	// Prepend "translations/en.html" to the filepaths, because it serves as
	// the ultimate fallback
	filePaths = append([]string{"translations/en.html"}, filePaths...)

	numTranslations := len(filePaths)
	fileNames := make([]string, numTranslations, numTranslations)
	fileCache := make([][]byte, numTranslations, numTranslations)
	languageTags := make([]language.Tag, numTranslations, numTranslations)

	for idx, filePath := range filePaths {

		// Get the filename (with extension)
		fileNames[idx] = filepath.Base(filePath)

		// Read the content of the file into the cache
		fileCache[idx], err = ioutil.ReadFile(filePath)
		if err != nil {
			panic("Could not read localisation file " + filePath)
		}

		// Get the basename (file name without extension)
		baseName := strings.TrimSuffix(fileNames[idx], filepath.Ext(fileNames[idx]))

		// Parse the basename as a language tag
		languageTags[idx] = language.MustParse(baseName)
	}

	return fileCache, language.NewMatcher(languageTags)
}


