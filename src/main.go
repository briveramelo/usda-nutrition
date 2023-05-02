package main

import (
	"brm/usda-nutrition/src/models"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	//results := utils.ReadCsvFile("/Users/brandonrivera-melo/Documents/Repos/usda-nutrition/data/ABBREV.txt")
	//var nuts []*models.Nutrition
	//for i := 0; i < len(results); i++ {
	//	nuts = append(nuts, models.New(results[i]))
	//}

	byteValue, _ := os.ReadFile("/Users/brandonrivera-melo/Documents/Repos/usda-nutrition/data/foundationDownload.json")
	var data models.NutritionData
	json.Unmarshal(byteValue, &data)
	var data2 []models.FoundationFoods
	fmt.Println(data2[0].DataType)
}
