package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log/slog"
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
		slog.Error("Не передано имя файла")
		os.Exit(1)
	}

	if len(jsonKeys) == 0 {
		slog.Error("Нужно передать хотя бы один ключ для отображения")
		os.Exit(1)
	}

	variousStruct, err := readJson(jsonFileName)
	if err != nil {
		slog.Error("Возникла ошибка при чтении JSON", slog.String("ошибка", err.Error()))
		os.Exit(1)
	}

	printKeys(jsonKeys, variousStruct)
}

func readJson(jsonFileName string) (map[string]any, error) {
	var variousStruct map[string]any

	fileBytes, err := os.ReadFile(jsonFileName)
	if err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %w", err)
	}

	err = json.Unmarshal(fileBytes, &variousStruct)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %w", err)
	}

	return variousStruct, nil
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
