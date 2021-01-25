package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/reiver/go-porterstemmer"
)

func GetDB() *gorm.DB {
	fullPath := os.Getenv("GOPATH") + "/src/MathbloomBE"
	absFile, err := filepath.Abs(fullPath + "/conf.json")
	if err != nil {
		panic("Failed to get absolute path")
	}
	file, fileErr := ioutil.ReadFile(absFile)
	if fileErr != nil {
		panic("Failed to read conf.json")
	}
	type Conf struct {
		DBUser     string
		DBPassword string
		DBHost     string
		DBName     string
	}

	conf := Conf{}
	err = json.Unmarshal([]byte(file), &conf)
	if err != nil {
		panic("Failed to read unmarshal json")
	}

	dbStr := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.DBUser, conf.DBPassword, conf.DBHost, conf.DBName,
	)
	db, err := gorm.Open("mysql", dbStr)
	if err != nil {
		panic("Failed to connect: " + err.Error())
	}
	return db
}

func GetUrltext(url string, client http.Client) []byte {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return body
}

func CreateUploadFileStructure() {
	MAX_DIRS := 256
	fullPath := os.Getenv("GOPATH") + "/src/MathbloomBE/uploads"
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// uploads does not exist, create uploads and subdirectory structure
		err = os.Mkdir(fullPath, 0755)
		if err != nil {
			log.Fatal("Unable to create uploads directory: " + fullPath)
		}
		for i := 0; i < MAX_DIRS; i++ {
			hex1 := fmt.Sprintf("%02X", i)
			level1Path := fullPath + "/" + hex1
			if _, err = os.Stat(level1Path); os.IsNotExist(err) {
				err = os.Mkdir(level1Path, 0755)
				if err != nil {
					log.Fatal("Unable to create uploads subdirectory: " + level1Path)
				}
			}
			for j := 0; j < MAX_DIRS; j++ {
				hex2 := fmt.Sprintf("%02X", j)
				level2Path := fullPath + "/" + hex1 + "/" + hex2
				if _, err = os.Stat(level2Path); os.IsNotExist(err) {
					err = os.Mkdir(level2Path, 0755)
					if err != nil {
						log.Fatal("Unable to create uploads subdirectory: " + level2Path)
					}
				}
			}
		}
	}
}

func StemSentence(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z]+")
	strs := strings.Fields(str)
	if err != nil {
		log.Fatal(err)
	}
	stemmedStrs := []string{}
	for _, word := range strs {
		stemmedWord := reg.ReplaceAllString(word, "")
		stemmedWord = porterstemmer.StemString(strings.ToLower(stemmedWord))
		stemmedStrs = append(stemmedStrs, stemmedWord)
	}
	return strings.Join(stemmedStrs[:], " ")
}

// Distance in meters between (lat1, lng1) and (lat2, lng2)
func Distance(lat1, lng1, lat2, lng2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, ln1, la2, ln2, r float64
	la1 = lat1 * math.Pi / 180
	ln1 = lng1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	ln2 = lng2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(ln2-ln1)
	return 2 * r * math.Asin(math.Sqrt(h))
}

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}
