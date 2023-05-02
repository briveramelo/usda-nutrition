package models

import (
	"strconv"
	"strings"
)

type Nutrition struct {
	NDB_num     string  //5-digit Nutrient Databank number that uniquely identifies a food item. If this field is defined as numeric, the leading zero will be lost.
	Desc        string  //60-character abbreviated description of food item.
	Water       float64 //Water (g/100 g)
	Energy_Kcal float64 // Food energy (kcal/100 g)
	Protein     float64 //Protein (g/100 g)
	Lipid       float64 //Total lipid (fat) (g/100 g)
	Ash         float64 //Ash (g/100 g) inorganic impurities in the fuel (typically sand, nickel, aluminium, silicon, sodium, and vanadium), which can cause different kinds of problems. Typically the ash value is in the range of 0.03%–0.07% by mass
	Carbs       float64 //Carbohydrate, by difference (g/100 g)
	Fiber       float64 //Total dietary fiber (g/100 g)
	Sugar       float64 //Total sugars (g/100 g)
	Calcium     float64 //Calcium (mg/100 g)
	Iron        float64 //Iron (mg/100 g)
	Magnesium   float64 //Magnesium (mg/100 g)
	Phosphorous float64 //Phosphorus (mg/100 g)
	Potassium   float64 //Potassium (mg/100 g)
	Sodium      float64 //Sodium (mg/100 g)
	Zinc        float64 //Zink (mg/100 g)
	Copper      float64 //Copper (mg/100 g)
	Manganese   float64 //Manganese (mg/100 g)
	Selenium    float64 //Selenium (ug/100 g)
	VitC        float64 //Vitamin C (mg/100 g)
	Thiamin     float64 //Thiamin (mg/100 g)
	Riboflavin  float64 //Riboflavin (mg/100 g)
	Niacin      float64 //Niacin (mg/100 g)
	Panto       float64 //Pantothenic Acid (mg/100 g)
	VitB6       float64 //Vitamin B6 (mg/100 g)
	Folate      float64 //Folate, total (ug/100g)
	Folic_acid  float64 //Folic acid (ug/100g)
	Food_folate float64 //Food folate (ug/100g)
	Folate_DFE  float64 //ug dietary folate equivalents/100g)
	Choline     float64 //Choline, total (mg/100g)
	VitB12      float64 //Vitamin B12 (ug/100g)
	VitA        float64 //Vitamin A (IU/100g)
	VitaRae     float64 //Vitamin A ug retinol activity equivalents/100g)
	Retinol     float64 //Retinol (ug/100g)
	AlphaCarot  float64 //Alpha Carotene (ug/100g)
	BetaCarot   float64 //Beta Carotene (ug/100g)
	BetaCrypt   float64 //beta-cryptoxanthin (ug/100g)
	Lycopene    float64 //lycopene (ug/100g)
	LutAndZea   float64 //Lutein + Zeazanthin (ug/100g)
	VitE        float64 //Vitamin E (alpha tocopherol) (mg/100g)
	VitD_ug     float64 //Vitamin D (ug/100g)
	VitD_IU     float64 //Vitamin D (IU/100g)
	VitK        float64 //Vitamin K (phylloquinone) (μg/100 g)
	FA_Sat      float64 //Saturated fatty acid (g/100 g)
	FA_Mono     float64 //Monounsaturated fatty acids (g/100 g)
	FA_Poly     float64 //Polyunsaturated fatty acids (g/100 g)
	Cholesterol float64 //Cholesterol (mg/100 g)
	Gmwt1       float64 //First household weight for this item from the Weight file
	GmtwtDesc1  string  //Description of household weight number1
	Gmwt2       float64 //Second household weight for this item from the Weight file
	GmwtDesc2   string  //Description of household weight number2
	RefusePct   float64 //Percent refuse
}

func New(csv []string) *Nutrition {
	n := new(Nutrition)

	n.NDB_num = strings.Trim(csv[0], "~")
	n.Desc = strings.Trim(csv[1], "~")
	n.Water = parseFloat(csv[2])
	n.Energy_Kcal = parseFloat(csv[3])
	n.Protein = parseFloat(csv[4])
	n.Lipid = parseFloat(csv[5])
	n.Ash = parseFloat(csv[6])
	n.Carbs = parseFloat(csv[7])
	n.Fiber = parseFloat(csv[8])
	n.Sugar = parseFloat(csv[9])
	n.Calcium = parseFloat(csv[10])
	n.Iron = parseFloat(csv[11])
	n.Magnesium = parseFloat(csv[12])
	n.Phosphorous = parseFloat(csv[13])
	n.Potassium = parseFloat(csv[14])
	n.Sodium = parseFloat(csv[15])
	n.Zinc = parseFloat(csv[16])
	n.Copper = parseFloat(csv[17])
	n.Manganese = parseFloat(csv[18])
	n.Selenium = parseFloat(csv[19])
	n.VitC = parseFloat(csv[20])
	n.Thiamin = parseFloat(csv[21])
	n.Riboflavin = parseFloat(csv[22])
	n.Niacin = parseFloat(csv[23])
	n.Panto = parseFloat(csv[24])
	n.VitB6 = parseFloat(csv[25])
	n.Folate = parseFloat(csv[26])
	n.Folic_acid = parseFloat(csv[27])
	n.Food_folate = parseFloat(csv[28])
	n.Folate_DFE = parseFloat(csv[29])
	n.Choline = parseFloat(csv[30])
	n.VitB12 = parseFloat(csv[31])
	n.VitA = parseFloat(csv[32])
	n.VitaRae = parseFloat(csv[33])
	n.Retinol = parseFloat(csv[34])
	n.AlphaCarot = parseFloat(csv[35])
	n.BetaCarot = parseFloat(csv[36])
	n.BetaCrypt = parseFloat(csv[37])
	n.Lycopene = parseFloat(csv[38])
	n.LutAndZea = parseFloat(csv[39])
	n.VitE = parseFloat(csv[40])
	n.VitD_ug = parseFloat(csv[41])
	n.VitD_IU = parseFloat(csv[42])
	n.VitK = parseFloat(csv[43])
	n.FA_Sat = parseFloat(csv[44])
	n.FA_Mono = parseFloat(csv[45])
	n.FA_Poly = parseFloat(csv[46])
	n.Cholesterol = parseFloat(csv[47])
	n.Gmwt1 = parseFloat(csv[48])
	n.GmtwtDesc1 = strings.Trim(csv[49], "~")
	n.Gmwt2 = parseFloat(csv[50])
	n.GmwtDesc2 = strings.Trim(csv[51], "~")
	n.RefusePct = parseFloat(csv[52])

	return n
}

func parseFloat(val string) float64 {
	converted, _ := strconv.ParseFloat(val, 64)
	return converted
}
