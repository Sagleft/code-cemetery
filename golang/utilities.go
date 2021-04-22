package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
)

//GetFloatPrecision returns the number of decimal places in a float
func GetFloatPrecision(f float64) int {
	return int(math.Floor(math.Log10(math.Floor(1 / f))))
}

//ReadFileToString read file to string
func ReadFileToString(filepath string) (string, error) {
	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file), err
}

//SaveStructToJSONFile save arbitrary structure to JSON file
func SaveStructToJSONFile(i interface{}, filepath string) (bool, error) {
	jsonr, err := json.Marshal(i)
	if err != nil {
		return false, err
	}
	var json string = string(jsonr)
	return SaveStringToFile(filepath, json)
}

//SaveStringToFile save arbitrary string to file
func SaveStringToFile(filepath string, content string) (bool, error) {
	file, err := os.Create(filepath)
	if err != nil {
		return false, err
	}
	defer file.Close()
	_, err = file.WriteString(content)
	if err != nil {
		return false, err
	}
	return true, nil
}

//IntArrFind find element index in arr. -1 if not found
func IntArrFind(a []int64, x int64) int {
	for i, v := range a {
		if x == v {
			return i
		}
	}
	return -1
}

//IntArrContains indicates whether x is contained in a.
func IntArrContains(a []int64, x int64) bool {
	for _, v := range a {
		if x == v {
			return true
		}
	}
	return false
}

//GoPrint print sth
func GoPrint(i interface{}) {
	fmt.Println(i)
}
