package services

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"women-in-media-article-entity-analysis/internal/models"
)

func GetGenderAnalysis(name string) (*models.GenderAnalysis, error) {

	names, namesErr := GetNames()

	if namesErr != nil {
		return nil, errors.Wrap(namesErr, "Could not get names")
	}

	values := map[string]interface{}{"name": name, "names": names}

	jsonValue, _ := json.Marshal(values)

	resp, err := http.Post("https://cat02lbi4d.execute-api.eu-west-1.amazonaws.com/PROD/getGenderAnalysis", "application/json", bytes.NewBuffer(jsonValue))
	defer resp.Body.Close()

	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshall s3 data")
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, errors.Wrap(err, "could not read body test")
	}
	var genderAnalysis = new(models.GenderAnalysis)
	//TODO: validate response
	genderAnalysisError := json.Unmarshal(body, &genderAnalysis)
	if genderAnalysisError != nil {
		return nil, errors.Wrap(genderAnalysisError, "could not parse response from CAPI")
	}
	return genderAnalysis, nil
}
