package main

import (
	"log"
	"os"
	"strings"

	"github.com/jose78/gansible/conf"
	"golang.org/x/exp/maps"
	"gopkg.in/yaml.v3"
)

type Generic interface{}

var err error

func main() {
	path := conf.LoadFlags()
	containerData := loadContent(path)
	log.Println(containerData["tasks"])

	for _, item := range containerData["tasks"].([]interface{}) {
		evaluateTask(item)
	}
}



func checkError(e error, description string) {
	if e != nil {
		log.Printf("%s -> %w \n", description, e)
		panic(e)
	}
}

func loadContent(path string) map[string]Generic {
	data, err := os.ReadFile(path)
	checkError(err, "Reading file: "+path)
	containerData := map[string]Generic{}
	err = yaml.Unmarshal(data, &containerData)
	checkError(err, "Unmarshal yaml from "+path)
	log.Println(containerData)
	return containerData
}

func evaluateTask(task interface{}) {
	selectTask := func(clouserTask map[string]interface{}) (module, key string) {
		key = strings.Trim(strings.Trim(strings.Join(maps.Keys(clouserTask)[:], " "), "name"), " ")
		module = strings.Title(strings.ToLower(key))
		return
	}
	module, key := selectTask(task.(map[string]interface{}))
	log.Printf("key: %s, module: %s", key, module)

	// cache := conf.GetInstance()
	// cache.Set("debug", reflect.TypeOf(new(Debug)))
	//
	// instance ,_ := cache.Get("debug")
	// debug2 := (&instance).(Debug)
	//
	// log.Printf(debug2.msg)
}
