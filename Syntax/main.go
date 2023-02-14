package main

import (
	"fmt"
	"strconv"
)

func main() {

	var NumItems, NumFields, IntValue int
	var Name, Type, StrValue, LookFor string
	var BoolValue bool
	// var ListItems []interface{}
	var ListItems []map[string][]interface{}
	var ListValue []interface{}
	//ListItems := map[string]interface{}{}

	// type Items []struct{}
	// type Item struct{}
	// var f interface{}
	// f = map[string]interface{}{}
	// f = map[string]interface{}{
	// 	"Name": "Wednesday",
	// 	"Age":  6,
	// 	"Parents": []interface{}{
	// 		"Gomez",
	// 		"Morticia",
	// 	},
	// }
	// m := f.(map[string]interface{})
	// m["middle"] = "her"

	fmt.Print("Enter a number of items: ")
	fmt.Scan(&NumItems)
	fmt.Print("Enter a number of item's fields: ")
	fmt.Scan(&NumFields)

	for i := 1; i <= NumItems; i++ {
		m := map[string][]interface{}{}

		// fmt.Println(m)

		for j := 1; j <= NumFields; j++ {
			fmt.Printf("Enter a %v item's %v field name: ", i, j)
			fmt.Scan(&Name)
			fmt.Printf("Enter a %v item's %v field type: ", i, j)
			fmt.Scan(&Type)
			fmt.Printf("Enter a %v item's %v field value: ", i, j)
			fmt.Scan(&StrValue)

			switch {
			case Type == "int" || Type == "integer":
				fmt.Sscan(StrValue, &IntValue)
				m[Name] = append(m[Name], IntValue)
				//m[Name] = IntValue
				//ListItems[Name] = IntValue
				// fmt.Println(m, ListItems)
			case Type == "bool" || Type == "boolean":
				BoolValue, _ = strconv.ParseBool(StrValue)
				m[Name] = append(m[Name], BoolValue)
				//m[Name] = BoolValue
				//ListItems[Name] = BoolValue
			default:
				m[Name] = append(m[Name], StrValue)
				//m[Name] = StrValue
				//ListItems[Name] = StrValue
			}
		}

		fmt.Println(m, ListItems)
		// ListItems = append(ListItems, m)
		ListItems = append(ListItems, m)
		//fmt.Println(ListItems...)
		fmt.Println(ListItems)
		//fmt.Println(m, l)
	}

	fmt.Print("Enter a field name you look for: ")
	fmt.Scan(&LookFor)
	for _, item := range ListItems {
		for key, value := range item {
			if key == LookFor || LookFor == "all" {
				for _, vv := range value {
					ListValue = append(ListValue, vv)
				}
				//ListValue = append(ListValue, value)
				// fmt.Println(value, ListValue)
			}
		}
		if LookFor == "all" {
			fmt.Println(ListValue...)
			ListValue = []interface{}{}
		}
	}
	fmt.Println(ListValue...)

	//fmt.Print("Enter a field name you look for: ")
	//fmt.Scan(&LookFor)
	//for key, value := range ListItems {
	//	if key == LookFor || LookFor == "all" {
	//		ListValue = append(ListValue, value)
	//		// fmt.Println(value, ListValue)
	//	}
	//}
	//if LookFor == "all" {
	//	fmt.Println(ListValue...)
	//	// ListValue = []interface{}{}
	//}
	//fmt.Println(ListValue...)
}
