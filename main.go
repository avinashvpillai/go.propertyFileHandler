package main

import (
	"fmt"
	"log"
	"os"

	"github.com/magiconair/properties"
)

var propertyFiles = []string{"./application.properties"}
var props = properties.MustLoadFiles(propertyFiles, properties.UTF8, true)

func main() {
	var val = getProp("state")
	fmt.Println("property value for state : ", val)

	saveDataToProperties()

	val = getProp("state")
	fmt.Println("new property value for state : ", val)
}

func reloadProperties() {
	fmt.Println("reloading properties..")
	props = properties.MustLoadFiles(propertyFiles, properties.UTF8, true)
}

func getProp(key string) string {
	value := ""
	var ok bool
	value, ok = props.Get(key)
	if !ok {
		return ""
	}
	return value
}

func saveDataToProperties() {
	file, err := os.OpenFile("application.properties", os.O_RDWR, 0644)
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	defer file.Close()
	_, err = file.WriteAt([]byte("state=Kerala"), 0) // Write at 0 beginning
	if err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}
	reloadProperties()

}
