package main

import (
	"fmt"

	"github.com/magiconair/properties"
)

var propertyFiles = []string{"./application.properties"}

func main() {
	p := properties.MustLoadFiles(propertyFiles, properties.UTF8, true)
	var val = p.GetString("state", "default")
	fmt.Println("property value for Organization : ", val)
}
