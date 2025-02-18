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

// date time format layout
const YYYYMMDD = "2006-01-02"
const DDMMYYYY = "02.01.2006"

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

func PrintProductJSON(product_name string, product_version string) (string, error) {
	cycle_data, err := GetSingleCycle(product_name, product_version)
	if err != nil {
		return "", err
	}

	latestCycle, err := GetLatestCycle(product_name)
	if err != nil {
		return "", err
	}

	cycle_data.LastestCycle = latestCycle

	return PrettyPrintJSON(cycle_data), nil
}

func GetSingleCycle(product_name string, product_version string) (Product, error) {
	api_url := "https://endoflife.date/api/"

	version, err := semver.NewVersion(product_version)
	if err != nil {
		return Product{}, err
	}

	versionMajorMinor := fmt.Sprintf("%d.%d", version.Major(), version.Minor())

	resp, err := http.Get(api_url + product_name + "/" + versionMajorMinor + ".json")
	if err != nil {
		return Product{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Product{}, err
	}

	var release_cycle Product
	err = json.Unmarshal(body, &release_cycle)
	if err != nil {
		return Product{}, err
	}

	release_cycle.Name = product_name
	release_cycle.MyVersion = product_version

	if utils.TypeofObject(release_cycle.EndOfLifeDate) == "boolean" {
		if release_cycle.EndOfLifeDate.(bool) {
			release_cycle.EndOfLifeDate = "Discontinued."
		} else {
			release_cycle.EndOfLifeDate = "Supported."
		}
	} else if utils.TypeofObject(release_cycle.EndOfLifeDate) == "string" {

		t, err := time.Parse(YYYYMMDD, release_cycle.EndOfLifeDate.(string))
		if err != nil {
			return Product{}, err
		}

		release_cycle.EndOfLifeDate = t.Format(DDMMYYYY)
	} else {
		release_cycle.EndOfLifeDate = "Unknown"
	}

	return release_cycle, nil
}

func GetLatestCycle(product_name string) (string, error) {
	api_url := "https://endoflife.date/api/"

	resp, err := http.Get(api_url + product_name + ".json")
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
			fmt.Errorf("Error parsing version: %s", err)
		}

		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))

	latestCycle := vs[len(vs)-1].Original()

	return latestCycle, nil
}

func PrettyPrintJSON(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")

	return string(s)
}
