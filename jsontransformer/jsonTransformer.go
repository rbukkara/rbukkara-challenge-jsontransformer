package jsontransformer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

// The function transforms the json input string based on the provided rules
func Transform(input []byte) (io.Reader, error) {

	var inputStruct map[string]interface{}
	err := json.Unmarshal(input, &inputStruct)

	if err != nil {
		return bytes.NewReader([]byte("")), err
	}

	result := make(map[string]interface{})

	for key, val := range inputStruct {
		fmt.Printf("Key: %s, Val: %v", key, val)
		value, err := transformDataTypes(trimMapKeys(val.(map[string]interface{})))
		if err != nil {
			fmt.Println(key, err)
		}
		if key == "" || value == nil {
			continue
		}
		formattedKey := strings.Trim(key, " ")
		result[formattedKey] = value
	}

	stringifiedOutput, _ := json.Marshal(result)

	// return string(stringifiedOutput), nil
	return bytes.NewReader(stringifiedOutput), nil
}

// transformDataTypes applies specific rules for individual types
func transformDataTypes(value map[string]interface{}) (interface{}, error) {

	if _, ok := value["N"]; ok {
		return transformNativeNumber(value["N"].(string))
	}

	if _, ok := value["S"]; ok {
		return transformNativeString(value["S"].(string))
	}

	if _, ok := value["BOOL"]; ok {
		return transformNativeBoolean(strings.Trim(value["BOOL"].(string), " "))
	}

	if _, ok := value["NULL"]; ok {
		return transformNativeNull(value["NULL"].(string))
	}

	if outerValue, ok := value["L"]; ok {
		return transformList(outerValue)
	}

	if _, ok := value["M"]; ok {
		return transformMaps(value["M"].(map[string]interface{}))
	}

	return nil, errors.New("invalid value")
}

// trims the trailing and leading spaces from input keys
func trimMapKeys(inputMap map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range inputMap {
		formattedKey := strings.Trim(key, " ")
		if formattedKey == "" {
			continue
		}
		result[formattedKey] = value
	}
	return result
}

func transformMaps(input map[string]interface{}) (interface{}, error) {
	mapEntries := map[string]interface{}{}
	for key, val := range input {
		switch val.(type) {
		case map[string]interface{}:
			native, err := transformDataTypes(trimMapKeys(val.(map[string]interface{})))
			if err != nil {
				fmt.Println("error decoding for this value", err)
			}
			if err != nil {
				continue
			}
			mapEntries[key] = native
		}
	}

	if len(mapEntries) == 0 {
		return nil, errors.New("no entries added to the M type")
	}

	return mapEntries, nil
}

func transformList(outerValue interface{}) (interface{}, error) {
	switch outerValue.(type) {
	case string:
		return nil, errors.New("invalid entry for lists")
	case []string:
		return nil, errors.New("invalid entry for lists")
	}
	var listEntries []interface{}
	for _, val := range outerValue.([]interface{}) {
		switch val.(type) {
		case map[string]interface{}:
			native, err := transformDataTypes(trimMapKeys(val.(map[string]interface{})))
			if err != nil {
				fmt.Println("error decoding for this value", err)
			}
			if native == nil || err != nil {
				continue
			}
			listEntries = append(listEntries, native)
		}
	}
	if len(listEntries) == 0 {
		return nil, errors.New("empty list")
	}
	return listEntries, nil
}

func transformNativeNumber(input string) (interface{}, error) {
	numericString := strings.Trim(input, " ")
	float1, err := strconv.ParseFloat(numericString, 64)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return float1, nil
}

func transformNativeString(input string) (interface{}, error) {
	stringValue := strings.Trim(input, " ")
	if stringValue == "" {
		return nil, errors.New("invalid String entry")
	}
	fmt.Println("stringValue: ", stringValue)
	t, err := time.Parse(time.RFC3339, stringValue)
	if err != nil {
		fmt.Println(err)
		return stringValue, nil
	}
	fmt.Println(t.Unix())
	return t.Unix(), nil
}

func transformNativeBoolean(input string) (interface{}, error) {
	stringValue := strings.Trim(input, " ")
	validBooleans := map[string]bool{
		"1":     true,
		"t":     true,
		"T":     true,
		"true":  true,
		"True":  true,
		"TRUE":  true,
		"0":     false,
		"f":     false,
		"F":     false,
		"false": false,
		"False": false,
		"FALSE": false,
	}

	if _, ok := validBooleans[stringValue]; !ok {
		return nil, errors.New("invalid boolean entry")
	}
	return validBooleans[stringValue], nil
}

func transformNativeNull(input string) (interface{}, error) {
	stringValue := strings.Trim(input, " ")
	validNulls := map[string]bool{
		"1":     true,
		"t":     true,
		"T":     true,
		"true":  true,
		"True":  true,
		"TRUE":  true,
		"0":     false,
		"f":     false,
		"F":     false,
		"false": false,
		"False": false,
		"FALSE": false,
	}
	if _, ok := validNulls[stringValue]; !ok {
		return nil, errors.New("invalid null entry")
	}
	return nil, nil
}
