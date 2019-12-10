package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/pkg/errors"
	"strings"
)

const NamesBucket = "bechdel-names"

func GetNames() (map[string]string, error) {

	sess, err := GetAwsSession("bechdel", "eu-west-1")

	if err != nil {
		return nil, errors.Wrap(err, "failed to marshall data for S3 upload")
	}

	downloader := s3manager.NewDownloader(sess)

	buffer := aws.NewWriteAtBuffer([]byte{})

	_, err = downloader.Download(buffer, &s3.GetObjectInput{
		Bucket: aws.String(NamesBucket),
		Key:    aws.String("names.json"),
	})

	if err != nil {
		return nil, errors.Wrap(err, "failed to download file")
	}

	mp := make(map[string]string)

	// Decode JSON into our map
	unmarshalError := json.Unmarshal(buffer.Bytes(), &mp)
	if unmarshalError != nil {
		println(err)
		return nil, errors.Wrap(err, "failed to unmarshall data from names object ")
	}

	return mp, nil
}

func MapGenderToGenderName(genderFromInput string) *string {
	var gender string
	if genderFromInput == "Male" {
		gender = "MaleName"
	} else if genderFromInput == "Female" {
		gender = "FemaleName"
	} else if genderFromInput == "NotName" {
		gender = "Place"
	}
	return &gender
}

func StoreCorrections(corrections map[string]string) error {
	sess, err := GetAwsSession("bechdel", "eu-west-1")

	if err != nil {
		return errors.Wrap(err, "failed to marshall data for S3 upload")
	}

	names, err := GetNames()

	if err != nil {
		return errors.Wrap(err, "Could not read names from names.json")
	}

	for key, value := range corrections {
		genderVal := MapGenderToGenderName(value)
		if genderVal != nil {
			keyWithoutDash := strings.Replace(key, "-", "", -1)
			names[keyWithoutDash] = *MapGenderToGenderName(value)
		}
	}

	marshalled, err := json.Marshal(names)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(NamesBucket),
		Key:    aws.String("names.json"),
		Body:   bytes.NewReader(marshalled),
	})

	if err == nil {
		fmt.Println("names upload successful")
	}

	return err
}
