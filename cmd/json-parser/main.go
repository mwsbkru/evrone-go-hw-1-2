package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type keys []string

func (k *keys) String() string {
	return fmt.Sprintf("%v", *k)
}

func (k *keys) Set(value string) error {
	*k = append(*k, value)
	return nil
}

func main() {
	var jsonKeys keys
	var jsonFileName string

	flag.Var(&jsonKeys, "keys", "Поля json, которые нужно вывести")
	flag.StringVar(&jsonFileName, "json", "", "JSON-файл, из которого нужно извлечь ключи")
	flag.Parse()

	if jsonFileName == "" {
		fmt.Println("Не передано имя файла")
		os.Exit(1)
	}

	if len(jsonKeys) == 0 {
		fmt.Println("Нужно передать хотя бы один ключ для отображения")
		os.Exit(1)
	}

	variousStruct := readJson(jsonFileName)
	printKeys(jsonKeys, variousStruct)
}

func readJson(jsonFileName string) map[string]any {
	var variousStruct map[string]any

	fileBytes, err := os.ReadFile(jsonFileName)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %v. Ошибка: %e\n", jsonFileName, err)
		os.Exit(1)
	}

	err = json.Unmarshal(fileBytes, &variousStruct)
	if err != nil {
		fmt.Printf("Ошибка при парсинге JSON: %v\n", err)
		os.Exit(1)
	}

	return variousStruct
}

func printKeys(jsonKeys []string, variousStruct map[string]any) {
	for _, key := range jsonKeys {
		if val, ok := variousStruct[key]; ok {
			fmt.Printf("%s: %v\n", key, val)
		} else {
			fmt.Printf("%s: %v\n", key, "<Отсутствует в JSON>")
		}
	}
}
