package api

import (
	"dd-server/utils"
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var licenseValidatorUrl string

var cache = map[string](map[string]string){}
var mutex = new(sync.Mutex)

func SetLicenseValidator(url string) {
	// TODO: check url.
	licenseValidatorUrl = url
}

// Call remote API to validate license and get the tags returned from API server.
// And add the response to metric's tags.
// For example, if validate successed, and return a tag{"user-id": "108"}
func validateLicense(license string) (map[string]string, error) {

	if licenseValidatorUrl == "" {
		return nil, nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	if v, ok := cache[license]; ok {
		return v, nil
	}
	log.Println("Get license from Server", license)

	data := map[string]string{"license": license}

	if resp, statusCode, err := utils.POST(licenseValidatorUrl, "", data, nil); statusCode == http.StatusOK {
		v, e := parseLicenseResult(resp)
		if e == nil && v != nil {
			cache[license] = v
		}
		return v, e
	} else {
		if err != nil {
			log.Printf("validate license error: %s", err.Error())
		}
		log.Printf("validate license error, resp: %s", resp)
		return nil, err
	}
}

func parseLicenseResult(resp string) (map[string]string, error) {

	var tags map[string]string
	if err := json.Unmarshal(([]byte)(resp), &tags); err != nil {
		log.Printf("parseLicenseResult: JSON Unmarshal error:", err)
		return nil, err
	}
	return tags, nil

}
