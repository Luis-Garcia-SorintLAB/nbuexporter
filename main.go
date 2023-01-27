package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

const (
	accept     = "application/vnd.netbackup+json;version=6.0"
	configPath = "./config"
)

// Read Domains configuration
func GetPrimariesFromFS(filesystem fs.FS) ([]Primary, error) {

	dir, err := fs.ReadDir(filesystem, ".")

	if err != nil {
		log.Fatal(err)
	}

	var primaries []Primary
	for _, f := range dir {
		primary, err := getPrimary(filesystem, f)
		if err != nil {
			return nil, err
		}
		primaries = append(primaries, primary)
	}

	return primaries, nil
}

func GetEntitiesFromFS(f fs.FS) ([]Entity, error) {

	var entity Entity
	var entities []Entity
	e, _ := fs.Glob(f, "entity-*.yaml")
	for _, file := range e {
		entityFile, err := fs.ReadFile(f, file)
		if err != nil {
			log.Fatalf("Error reading entity file %s : %s", file, err)
		}
		yaml.Unmarshal(entityFile, &entity)
		entities = append(entities, entity)
	}
	return entities, nil

}

//TODO Refactor GetPrimariesFromFS
// func GetPrimariesFromFS2(f fs.FS) ([]Primary, error) {

// 	var primary Primary
// 	var primaries []Primary
// 	e, _ := fs.Glob(f, "domain*.yaml")
// 	fmt.Println(e)
// 	for _, file := range e {
// 		fmt.Println(file)
// 		entityFile, err := fs.ReadFile(f, file)
// 		if err != nil {
// 			log.Fatalf("Error reading primary file %s : %s", file, err)
// 		}
// 		yaml.Unmarshal(entityFile, &primary)
// 		primaries = append(primaries, primary)
// 	}
// 	return primaries, nil

// }

// read configuration files for NBU domains
func getPrimary(filesystem fs.FS, f fs.DirEntry) (Primary, error) {
	primaryFile, err := fs.ReadFile(filesystem, f.Name())
	if err != nil {
		return Primary{}, err
	}
	var primary Primary
	yaml.Unmarshal(primaryFile, &primary)
	return primary, err
}

// Creates a Http Client
func NewHttpClient() http.Client {
	// Insecure https call
	tr := http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := &http.Client{
		Timeout:   30 * time.Second,
		Transport: &tr,
	}
	return *client
}

// Call Netbackup admin/jobs API and return a slice of them
func GetJobs(primary Primary) []Data {

	leap := time.Now().Add(-time.Hour * time.Duration(primary.TimeFrame.Num)).Format(time.RFC3339)
	var results []Data
	var result Data

	params := make(map[string]string)
	params["filter"] = "endTime ge " + leap
	//TODO: page[limit] value as primary parameter
	params["page[limit]"] = "1000"
	body := RunQuery("admin", "jobs", primary, params)
	if err := json.Unmarshal(body, &result); err != nil {
		log.Fatalf("Error parsing json %s", err)
	}
	results = append(results, result)
	return results

}

func RunQuery(api string, object string, primary Primary, params map[string]string) []byte {
	c := NewHttpClient()
	URL := "https://" + primary.Fqdn + "/netbackup/" + api + "/" + object
	req, _ := http.NewRequest(http.MethodGet, URL, nil)
	req.Header.Set("Accept", "application/vnd.netbackup+json;version=6.0")
	req.Header.Set("Authorization", primary.APIKey)
	q := req.URL.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	println(req.URL.String())
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("status %v", res.StatusCode)
	defer res.Body.Close()
	return body
}

func main() {

	filesystem := os.DirFS(configPath)
	//TODO: output file as parameter
	f, err := os.Create("./output/out.json")
	if err != nil {
		log.Fatalf("Error creating file %v", err)
	}
	defer f.Close()

	primaries, err := GetPrimariesFromFS(filesystem)
	if err != nil {
		log.Fatalf("error reading configuration files %s", err)
	}
	for {
		for i := 0; i < len(primaries); i++ {
			jobs := GetJobs(primaries[i])
			for i := 0; i < len(jobs); i++ {
				for j := 0; j < len(jobs[i].Data); j++ {
					job, _ := json.Marshal(jobs[i].Data[j].Attributes)
					_, err := fmt.Fprintf(f, "%s\n", job)
					if err != nil {
						log.Fatalf("Error writing data to file %s", err)
					}
					fmt.Printf("%s\n", job)
				}
			}
		}
		time.Sleep(1 * time.Hour)
	}
}
