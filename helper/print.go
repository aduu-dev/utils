package helper

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
)

// PrintAsJSON prints the given object as indented json.
func PrintAsJSON(i interface{}) {
	d := PrettyJSONString(i)
	fmt.Printf("%s:\n%s\n", reflect.TypeOf(i).String(), d)
}

// PrettyJSONString returns the given data structure as json.
//   Crashes if json marshalling failed.
func PrettyJSONString(i interface{}) string {
	return string(PrettyJSON(i))
}

// PrettyJSON returns the given data structure as json.
//   Crashes if json marshalling failed.
func PrettyJSON(i interface{}) []byte {
	d, err := json.MarshalIndent(i, "", "    ")
	if err != nil {
		log.Fatalln(err)
	}
	return d
}
