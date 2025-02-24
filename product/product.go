package product

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"time"

	"github.com/Masterminds/semver"
	"github.com/bvestli/check_endoflife/utils"
)

// YYYYMMDD date time format layout
const YYYYMMDD = "2006-01-02"

// DDMMYYYY date time format layout
const DDMMYYYY = "02.01.2006"

// Product defines the structure of the product data received from the API
type Product struct {
	Name          string `json:"name"`                   // Product name. string
	Cycle         string `json:"cycle,omitempty"`        // Release cycle version. string
	ReleaseDate   string `json:"releaseDate"`            // Release Date for the first release in this cycle. string<date>
	EndOfLifeDate any    `json:"eol"`                    // End of Life Date for this release cycle. string or boolean
	Latest        string `json:"latest"`                 // Latest release in this cycle
	Link          string `json:"link,omitempty"`         // Link to changelog for the latest release, if available. string or null
	Support       any    `json:"support"`                // Whether this release cycle has active support. string<date> or boolean
	Discontinued  any    `json:"discontinued,omitempty"` // Whether this cycle is now discontinued. string<date> or boolean
	MyVersion     string `json:"myversion"`              // Version of the product defined in the products-file. string
	LastestCycle  string `json:"latestcycle"`            // Latest release cycle. string
}

// FullProductData returns the full product data
func FullProductData(productName string, productVersion string) (Product, error) {
	cycleData, err := GetSingleCycle(productName, productVersion)
	if err != nil {
		return Product{}, err
	}

	latestCycle, err := GetLatestCycle(productName)
	if err != nil {
		return Product{}, err
	}

	cycleData.LastestCycle = latestCycle

	return cycleData, nil
}

// GetSingleCycle returns the single cycle for a product
func GetSingleCycle(productName string, productVersion string) (Product, error) {
	apiURL := "https://endoflife.date/api/"

	version, err := semver.NewVersion(productVersion)
	if err != nil {
		return Product{}, err
	}

	versionMajorMinor := fmt.Sprintf("%d.%d", version.Major(), version.Minor())

	resp, err := http.Get(apiURL + productName + "/" + versionMajorMinor + ".json")
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Product{}, err
	}

	var releaseCycle Product
	err = json.Unmarshal(body, &releaseCycle)
	if err != nil {
		return Product{}, err
	}

	releaseCycle.Name = productName
	releaseCycle.MyVersion = productVersion

	if utils.TypeofObject(releaseCycle.EndOfLifeDate) == "boolean" {
		if releaseCycle.EndOfLifeDate.(bool) {
			releaseCycle.EndOfLifeDate = "Discontinued."
		} else {
			releaseCycle.EndOfLifeDate = "Supported."
		}
	} else if utils.TypeofObject(releaseCycle.EndOfLifeDate) == "string" {

		t, err := time.Parse(YYYYMMDD, releaseCycle.EndOfLifeDate.(string))
		if err != nil {
			return Product{}, err
		}

		releaseCycle.EndOfLifeDate = t.Format(DDMMYYYY)
	} else {
		releaseCycle.EndOfLifeDate = "Unknown"
	}

	return releaseCycle, nil
}

// GetLatestCycle returns the latest cycle for a product
func GetLatestCycle(productName string) (string, error) {
	apiURL := "https://endoflife.date/api/"

	resp, err := http.Get(apiURL + productName + ".json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var products []Product
	err = json.Unmarshal(body, &products)
	if err != nil {
		return "", err
	}

	cycles := []string{}
	for _, product := range products {
		cycles = append(cycles, product.Cycle)
	}

	vs := make([]*semver.Version, len(cycles))
	for i, r := range cycles {
		v, err := semver.NewVersion(r)
		if err != nil {
			return "", fmt.Errorf("Error parsing version: %s", err)
		}

		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))

	latestCycle := vs[len(vs)-1].Original()

	return latestCycle, nil
}
