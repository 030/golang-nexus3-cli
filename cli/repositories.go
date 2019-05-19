package cli

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thedevsaddam/gojsonq"
)

func (n Nexus3) repositories() string {
	// Generated by curl-to-Go: https://mholt.github.io/curl-to-go

	req, err := http.NewRequest("GET", n.URL+"/service/rest/v1/repositories", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	var bodyString string
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString = string(bodyBytes)
	}
	return bodyString
}

func repositoryNamesJSON(json string) interface{} {
	jq := gojsonq.New().JSONString(json)
	jq.SortBy("name", "asc")
	name := jq.Pluck("name")
	return name
}

func (n Nexus3) repositoriesSlice() []interface{} {
	return repositoryNamesJSON(n.repositories()).([]interface{})
}

func (n Nexus3) RepositoryNames() {
	for _, name := range n.repositoriesSlice() {
		fmt.Printf("%s\n", name)
	}
}

func (n Nexus3) CountRepositories() {
	fmt.Println(len(n.repositoriesSlice()))
}

// Downloads retrieves artifacts from all repositories
func (n Nexus3) Downloads() error {
	for _, name := range n.repositoriesSlice() {
		n := Nexus3{URL: n.URL, User: n.User, Pass: n.Pass, Repository: name.(string)}
		err := n.StoreArtifactsOnDisk()
		if err != nil {
			return err
		}
	}
	return nil
}
