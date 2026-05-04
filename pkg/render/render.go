package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders templates using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	// create a template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
		return
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func createTemplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all the files named *page.gohtml from ./templates
	pages, err := filepath.Glob("./templates/*.gohtml")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.gohtml
	for _, page := range pages {
		name := filepath.Base(page)
		actualPage, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		match, err := filepath.Glob("./templates/*layout.gohtml")
		if err != nil {
			return myCache, err
		}

		if len(match) > 0 {
			// Where They Get Linked
			actualPage, err = actualPage.ParseGlob("./templates/*layout.gohtml")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = actualPage
	}

	return myCache, nil
}
