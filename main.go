package main

import (
	"os"
	"os/user"
	"net/http"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

func main() {
	user, _ := user.Current()
	filePath := user.HomeDir + "/.config/aah/aahelp.yaml"
	userFilePath := user.HomeDir + "/.aahelp.yaml"

	if _, err := os.Stat(filePath); err != nil {
		err := DownloadFile(filePath, "https://raw.githubusercontent.com/TheAndroidMaster/AAH/master/aahelp.yaml")
		if (err == nil) {
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
		err = yaml.Unmarshal([]byte(file), &m)
		if err == nil {
			if (userErr == nil) { 
				m2 := make(map[interface{}]interface{})
				err = yaml.Unmarshal([]byte(userFile), &m2)
				if err == nil {
					m = MergeMap(m, m2)
				} else {
					fmt.Printf("your file ~/.aahelp.yaml is not formatted correctly: %s\n", err)
				}
			}
		
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

			PrintMap(nil, m)			
		} else {
			fmt.Printf("err %v parsing file\n", err)
		}
	} else {
		fmt.Printf("err reading file\n")
	}
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

func PrintMap(key, val interface{}) {
	if v, ok := val.(map[interface{}]interface{}); ok {
		if key != nil {
			fmt.Printf("\n--> %s\n", key)
		}
		
		for k, val := range v {
			PrintMap(k.(string), val)
		}
		fmt.Printf("\n")
	} else {
		fmt.Printf("%-15s| \t\t%s\n", key, val)
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
