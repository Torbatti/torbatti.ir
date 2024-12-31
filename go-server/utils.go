package main

import "os"

func ReadHtml(root_path string, hub_name string, part_name string) (string, error) {

	data, err := os.ReadFile(root_path + hub_name + "/" + part_name + ".html")
	if err != nil {
		return "", err
	}
	// fmt.Print(string(data))

	return string(data), nil
}

func ReadMd(root_path string, hub_name string, part_name string) (string, error) {

	data, err := os.ReadFile(root_path + hub_name + "/" + part_name + ".md")
	if err != nil {
		return "", err
	}
	// fmt.Print(string(data))

	return string(data), nil
}
