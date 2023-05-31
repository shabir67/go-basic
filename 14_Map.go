package main

import "fmt"

func main() {

	person := map[string]string{
		"name":    "Muhammad",
		"address": "Lembang",
	}
	fmt.Println("1)------------ Basic Method at Map ---------------------------------------")

	fmt.Println("\n ")
	fmt.Println("  -Data Map   :", person)
	fmt.Println("  -Key name   :", person["name"])
	fmt.Println("  -Key address:", person["address"])
	fmt.Println("  -Length map :", len(person))
	fmt.Println("\n ")

	fmt.Println("1)------------ Basic Method at Map ---------------------------------------")
	fmt.Println("\n ")
	cars := make(map[string]string)
	cars["Brand"] = "Porsche"
	cars["Type"] = "911 GT3RS"
	cars["Fuel"] = "Diesel"
	fmt.Println("  -Cars map :", cars)
	fmt.Println("  -Cars len bef deletion :", len(cars))

	delete(cars, "Fuel")
	fmt.Println("  -Cars after delete :", cars)
	fmt.Println("  -Cars len aft deletion :", len(cars))
	fmt.Println("\n ")

}
