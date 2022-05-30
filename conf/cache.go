package conf

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

type Cache map[string]reflect.Type
var container Cache

type ICache interface{
	Set(key string, value reflect.Type)
	Get(key string) (interface{}, error)
}

func (cache Cache) Set(key string, value reflect.Type) {
	cache[key] = value
	log.Println("Loaded module: %s", key)
}

/* Cache.Get: method to retrive an instance of type that match with the key
* key: 
*/
func (cache Cache) Get(key string) (interface{}, error) {
	if typeSelected, ok := cache[key]; ok {
		return reflect.New(typeSelected).Elem().Interface(), nil
	}
	return nil, errors.New(fmt.Sprintf("Error, the module %s d hasn`t been loaded ", key))
}

/* GetInstance: Singlenton function to return the cache obj.
*/
func GetInstance() ICache {
	if container == nil {
		container = make(Cache)
		log.Println("Generated cache of modules.")
	}
	return container
}
