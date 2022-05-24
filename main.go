package main

import (
        "gopkg.in/yaml.v3"
        "log"
        "os"
"golang.org/x/exp/maps"
"strings"
)

type Generic interface{}

func main() {
        containerData := loadContent("test/demo.yml")
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
	      err2 := yaml.Unmarshal(data, &containerData)
        checkError(err2, "Unmarshal yaml from "+path)
        log.Println(containerData)
        return containerData
}

func evaluateTask(task interface{}){
  selectTask := func (clouserTask map[string]interface{}){

    maps.Keys(clouserTask)
    keys := strings.Join(maps.Keys(clouserTask)[:]," ")
   log.Println(keys)
   log.Println(strings.Trim(keys, "name"))
  }
  selectTask(task.(map[string]interface{}))
}

type Task struct {
        Name string
        body Executable
}

type Executable interface {
        Execute()
}

type Debug struct {
        msg string
}

func (task *Debug) Execute() {
        log.Println(task.msg)
}
