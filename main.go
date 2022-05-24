package main

import (
	"fmt"
	"os"
	"gopkg.in/yaml.v3"
)


type Generic interface{}

func main(){
     loadContent("./demo.yml")
}

func checkError(e error, contents...string) {
    if e != nil {
	fmt.Println(contents)
        panic(e)
    }
}

func loadContent(path string){
     data, err := os.ReadFile(path)
     checkError(err, "Reading file: " + path)
     containerData := map[string]Generic{}
     err2 := yaml.Unmarshal(data, &containerData)
     checkError(err2, "Unmarshal yaml from " + path)
     fmt.Println(containerData)
}
