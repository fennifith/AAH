package main

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {
	pwd, _ := os.Getwd()
	filePath := pwd + "/aahelp.yaml"
	if _, err := os.Stat("~/.aahelp.yaml"); err == nil {
		filePath = "~/.aahelp.yaml"
	} else if _, err = os.Stat(filePath); err != nil {
		fmt.Printf("default file '%v' doesn't exist\n", filePath)
		return
	}

	file, err := ioutil.ReadFile(filePath)
	if err == nil {
		m := make(map[interface{}]interface{})
		err = yaml.Unmarshal([]byte(file), &m)
		if err == nil {
			for i := 1; i < len(os.Args); i++ {
				if val, ok := m[os.Args[i]]; ok {
					if v, ok := val.(map[interface{}]interface{}); ok {
						m = v
					} else {
						fmt.Printf("%s: \t\t%s\n", os.Args[i], val)
						return
					}
				} else {
					fmt.Printf("couldn't find key '%s'\n\n--------------------\n", os.Args[i])
				}
			}

			print("values", m)			
		} else {
			fmt.Printf("err %v parsing file\n", err)
		}
	} else {
		fmt.Printf("err reading file\n")
	}
}

func print(key string, val interface{}) {
	if v, ok := val.(map[interface{}]interface{}); ok {
		fmt.Printf("\n--> %s\n", key)
		for k, val := range v {
			print(k.(string), val)
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("%-15s| \t\t%s\n", key, val)
	}
}
