package services_test

import (
	"fmt"
	"testing"
	"women-in-media-article-entity-analysis/internal/services"
)

func TestStoreCorrections(t *testing.T) {
	m := make(map[string]string)
	m["jonnytest"] = "Male"
	err := services.StoreCorrections(m)
	if err != nil {
		t.Error(err)
	}
}

func TestGetNames(t *testing.T) {
	names, err := services.GetNames()
	if err != nil {
		t.Error(err)
	}

	// See what the map has now
	fmt.Printf("mp is now: %+v\n", names)

	for key, value := range names {
		fmt.Println("key:", key, "value:", value)
	}

}
