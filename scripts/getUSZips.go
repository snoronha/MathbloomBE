package main

import (
	"MathbloomBE/models"
	"MathbloomBE/util"
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	_ "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db := util.GetDB()
	db.AutoMigrate(&models.USZip{})

	csvFile, _ := os.Open("uszips.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	count := 0
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		if count > 0 {
			lat, _ := strconv.ParseFloat(line[1], 64)
			lng, _ := strconv.ParseFloat(line[2], 64)
			pop, _ := strconv.ParseUint(line[8], 10, 64)
			var uszip = models.USZip{
				Zip:        line[0],
				Lat:        lat,
				Lng:        lng,
				City:       line[3],
				StateID:    line[4],
				StateName:  line[5],
				Population: uint(pop),
				CountyName: line[11],
				TimeZone:   line[17],
			}
			db.Create(&uszip)
		}
		count++
	}

}
