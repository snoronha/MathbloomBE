package models

import (
	_ "github.com/jinzhu/gorm"
)

type ItemDetail struct {
	ID                  uint    `json:"id" gorm:"primary_key"`
	Sku                 string  `json:"sku" sql:"type:varchar(32)"`
    UsItemId            string  `json:"USItemId" gorm:"type:varchar(16);unique_index"`
	OfferId             string  `json:"offerId" gorm:"type:varchar(64)"`
	Upc                 string  `json:"upc" gorm:"type:varchar(32)"`
	Rank                uint    `json:"rank"`
	Name                string  `json:"name"`
	MaxAllowed          uint    `json:"maxAllowed"`
	TaxCode             string  `json:"taxCode"`
	IsOutOfStock        string  `json:"isOutOfStock"`
	Thumbnail           string  `json:"thumbnail"`      // thumbnail
	Large               string  `json:"large"`          // large image
	IsAlcoholic         bool    `json:"isAlcoholic"`
	IsSnapEligible      uint    `json:"isSnapEligible"`
	PrimaryShelf        string  `json:"primaryShelf"`
	PrimaryAisle        string  `json:"primaryAisle"`
	PrimaryDepartment   string  `json:"primaryDepartment"`
	SalesUnit           string  `json:"salesUnit"`
	ProductUrl          string  `json:"productUrl"`
	Type                string  `json:"type"`
	ProductCode         string  `json:"productCode"`
	Brand               string  `json:"brand"`
	ProductType         string  `json:"productType"`
	ShortDescription    string  `json:"shortDescription" gorm:"type:varchar(2047)"`
	Description         string  `json:"description" gorm:"type:varchar(2047)"`
	ModelNum            string  `json:"modelNum"`
	AssembledInCountryOfOrigin  string `json:"assembledInCountryOfOrigin"`
	OriginOfComponents  string  `json:"originOfComponents"`
	Ingredients         string  `json:"ingredients" gorm:"type:varchar(511)"`
	AgeRestricted       bool    `json:"ageRestricted"`
	StorageType         string  `json:"storageType"`
	Weight              string  `json:"weight"`
	Rating              float64 `json:"rating"`
	ReviewsCount        uint    `json:"reviewsCount"`
	NutritionFacts      string  `json:"nutritionFacts" gorm:"type:varchar(2047)"`
	List                float64 `json:"list"`
    PriceUnitOfMeasure  string  `json:"priceUnitOfMeasure"`
    SalesUnitOfMeasure  string  `json:"salesUnitOfMeasure"`
	SalesQuantity       uint    `json:"salesQuantity"`
	IsRollback          string  `json:"isRollback"`
	IsClearance         string  `json:"isClearance"`
	Unit                float64 `json:"unit"`
	DisplayPrice        float64 `json:"displayPrice"`
	DisplayUnitPrice    string  `json:"displayUnitPrice"`
	IsInStock           bool    `json:"isInStock"`
}