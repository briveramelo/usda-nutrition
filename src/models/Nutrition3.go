package models

type NutritionData struct {
	FoundationFoods []FoundationFoods `json:"FoundationFoods" gorm:"column:FoundationFoods"`
}

type FoundationFoods struct {
	FoodNutrients []struct {
		Amount   float64 `json:"amount" gorm:"column:amount"`
		Nutrient struct {
			Number   string `json:"number" gorm:"column:number"`
			UnitName string `json:"unitName" gorm:"column:unitName"`
			Name     string `json:"name" gorm:"column:name"`
			Rank     int    `json:"rank" gorm:"column:rank"`
			ID       int    `json:"id" gorm:"column:id"`
		} `json:"nutrient" gorm:"column:nutrient"`
		Median                 float64 `json:"median" gorm:"column:median"`
		DataPoints             int     `json:"dataPoints" gorm:"column:dataPoints"`
		FoodNutrientDerivation struct {
			FoodNutrientSource struct {
				Code        string `json:"code" gorm:"column:code"`
				Description string `json:"description" gorm:"column:description"`
				ID          int    `json:"id" gorm:"column:id"`
			} `json:"foodNutrientSource" gorm:"column:foodNutrientSource"`
			Code        string `json:"code" gorm:"column:code"`
			Description string `json:"description" gorm:"column:description"`
		} `json:"foodNutrientDerivation" gorm:"column:foodNutrientDerivation"`
		ID   int    `json:"id" gorm:"column:id"`
		Type string `json:"type" gorm:"column:type"`
	} `json:"foodNutrients" gorm:"column:foodNutrients"`
	DataType    string `json:"dataType" gorm:"column:dataType"`
	NdbNumber   int    `json:"ndbNumber" gorm:"column:ndbNumber"`
	Description string `json:"description" gorm:"column:description"`
	FdcID       int    `json:"fdcId" gorm:"column:fdcId"`
	InputFoods  []struct {
		FoodDescription string `json:"foodDescription" gorm:"column:foodDescription"`
		ID              int    `json:"id" gorm:"column:id"`
		InputFood       struct {
			DataType     string `json:"dataType" gorm:"column:dataType"`
			Description  string `json:"description" gorm:"column:description"`
			FoodClass    string `json:"foodClass" gorm:"column:foodClass"`
			FdcID        int    `json:"fdcId" gorm:"column:fdcId"`
			FoodCategory struct {
				Code        string `json:"code" gorm:"column:code"`
				Description string `json:"description" gorm:"column:description"`
				ID          int    `json:"id" gorm:"column:id"`
			} `json:"foodCategory" gorm:"column:foodCategory"`
			PublicationDate string `json:"publicationDate" gorm:"column:publicationDate"`
		} `json:"inputFood" gorm:"column:inputFood"`
	} `json:"inputFoods" gorm:"column:inputFoods"`
	NutrientConversionFactors []struct {
		ProteinValue      float64 `json:"proteinValue" gorm:"column:proteinValue"`
		FatValue          float64 `json:"fatValue" gorm:"column:fatValue"`
		CarbohydrateValue float64 `json:"carbohydrateValue" gorm:"column:carbohydrateValue"`
		Type              string  `json:"type" gorm:"column:type"`
	} `json:"nutrientConversionFactors" gorm:"column:nutrientConversionFactors"`
	IsHistoricalReference bool `json:"isHistoricalReference" gorm:"column:isHistoricalReference"`
	FoodPortions          []struct {
		SequenceNumber int     `json:"sequenceNumber" gorm:"column:sequenceNumber"`
		Amount         float64 `json:"amount" gorm:"column:amount"`
		Modifier       string  `json:"modifier" gorm:"column:modifier"`
		MeasureUnit    struct {
			Name         string `json:"name" gorm:"column:name"`
			ID           int    `json:"id" gorm:"column:id"`
			Abbreviation string `json:"abbreviation" gorm:"column:abbreviation"`
		} `json:"measureUnit" gorm:"column:measureUnit"`
		GramWeight      float64 `json:"gramWeight" gorm:"column:gramWeight"`
		ID              int     `json:"id" gorm:"column:id"`
		MinYearAcquired int     `json:"minYearAcquired" gorm:"column:minYearAcquired"`
		Value           float64 `json:"value" gorm:"column:value"`
	} `json:"foodPortions" gorm:"column:foodPortions"`
	FoodClass    string `json:"foodClass" gorm:"column:foodClass"`
	FoodCategory struct {
		Description string `json:"description" gorm:"column:description"`
	} `json:"foodCategory" gorm:"column:foodCategory"`
	PublicationDate string `json:"publicationDate" gorm:"column:publicationDate"`
}
