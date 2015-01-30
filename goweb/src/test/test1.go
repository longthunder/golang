// test1
package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	t := template.Must(template.ParseGlob("templates/**/*")).Lookup("index3.tpl")

	// Execute the template for each recipient.
	r := map[string]string{"title": "xxx website"}
	err := t.Execute(os.Stdout, r)
	if err != nil {
		log.Println("executing template:", err)
	}

}
