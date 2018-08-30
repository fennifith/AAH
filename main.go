package main

import (
	"fmt"
	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"strings"
)

func main() {
	user, _ := user.Current()
	filePath := user.HomeDir + "/.config/aah/aahelp.yaml"
	userFilePath := user.HomeDir + "/.aahelp.yaml"

	if _, err := os.Stat(filePath); err != nil {
		err := DownloadFile(filePath, "https://raw.githubusercontent.com/TheAndroidMaster/AAH/master/aahelp.yaml")
		if err == nil {
			main()
		} else {
			fmt.Printf("tried to download aahelp.yaml from TheAndroidMaster/AAH, didn't work\n%s\n", err)
			fmt.Printf("please download the file to ~/.config/aah/aahelp.yaml yourself and the program will work\n")
		}

		return
	}

	file, err := ioutil.ReadFile(filePath)
	userFile, userErr := ioutil.ReadFile(userFilePath)
	if err == nil {
		m := make(map[interface{}]interface{})
		keys := ""
		err = yaml.Unmarshal([]byte(file), &m)
		if err == nil {
			if userErr == nil {
				m2 := make(map[interface{}]interface{})
				err = yaml.Unmarshal([]byte(userFile), &m2)
				if err == nil {
					m = MergeMap(m, m2)
				} else {
					fmt.Printf("your file ~/.aahelp.yaml is not formatted correctly: %s\n", err)
				}
			}

			for i := 1; i < len(os.Args); i++ {
				if key, val, ok := FindVal(os.Args[i], m); ok {
					keys += key + " "
					if v, ok := val.(map[interface{}]interface{}); ok {
						m = v
					} else {
						fmt.Printf("%s: \t\t", keys)
						color.New(color.FgWhite, color.Bold).Printf("%s\n", val)
						return
					}
				} else {
					fmt.Printf("couldn't find key '%s'\n\n--------------------\n", os.Args[i])
				}
			}

			indent := -1
			if len(keys) > 0 {
				color.New(color.FgBlue, color.Bold).Printf("%s:\n", keys)
				indent = 0
			}
			
			PrintMap(nil, m, indent)
		} else {
			fmt.Printf("err %v parsing file\n", err)
		}
	} else {
		fmt.Printf("err reading file\n")
	}
}

func FindVal(key string, m map[interface{}]interface{}) (string, interface{}, bool) {
	if val, ok := m[key]; ok {
		return key, val, true
	}

	for key2, v := range m {
		if k2, ok := key2.(string); ok && strings.HasPrefix(k2, key) {
			return k2, v, true
		}
	}

	return "", nil, false
}

func MergeMap(m1 map[interface{}]interface{}, m2 map[interface{}]interface{}) map[interface{}]interface{} {
	for k, v := range m2 {
		if val, ok := v.(map[interface{}]interface{}); ok && m1[k] != nil {
			if val2, ok := m1[k].(map[interface{}]interface{}); ok {
				m1[k] = MergeMap(val2, val)
			} else {
				m1[k] = v
			}
		} else {
			m1[k] = v
		}
	}

	return m1
}

func PrintMap(key, val interface{}, iter int) {
	indent := ""
	for i := 0; i < iter; i++ {
		indent += "  "
	}

	if v, ok := val.(map[interface{}]interface{}); ok {
		if key != nil {
			color.New(color.FgBlue, color.Bold).Printf("%s:\n", indent+key.(string))
		}

		for k, val := range v {
			PrintMap(k.(string), val, iter+1)
		}
	} else {
		fmt.Printf("%-30s", indent+key.(string)+":")
		color.New(color.FgWhite, color.Bold).Printf("%s\n", val)
	}
}

func DownloadFile(path string, url string) error {
	os.MkdirAll(path[:len(path)-11], os.ModePerm)

	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
