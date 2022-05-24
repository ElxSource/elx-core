package main

import ("fmt"
"gopkg.in/yaml.v3")



var data = `---
tasks:
   - name: Trace
     debug:
        msg: "Hola mundo"
   - name: "Do something"
     echo: 
        msg: "Soy un test"
`
func main(){
     fmt.Print("Hola mahjo")
     containerData := map[string]interface{}
     err := yaml.Unmarshal([]byte(data), &containerData)
     if err != nil {
	     fmt.Println("KO:"+ err)
     }
     fmt.Println(containerData)
}

