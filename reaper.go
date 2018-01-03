package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"

	"github.com/golang/glog"
	"github.com/smartystreets/go-aws-auth"
)

var (
	logLevel    = os.Getenv("LOG_LEVEL")
	indexPrefix = os.Getenv("INDEX_PREFIX")
	maxCountStr = os.Getenv("MAX_COUNT")
	esURL       = os.Getenv("ES_URL")
	maxCount    = 30
	indexReaper awsESIndexReaper
)

type awsESQueryer interface {
	esRequest(string, string) ([]byte, error)
}

type awsESIndexReaper struct {
	indexList indexes
}

type indexSettings struct {
	Settings struct {
		Index struct {
			CreationDate string `json:"creation_date"`
			ProvidedName string `json:"provided_name"`
		} `json:"index"`
	} `json:"settings"`
}

type indexes map[string]indexSettings

func (es awsESIndexReaper) esRequest(url, action string) ([]byte, error) {

	client := new(http.Client)

	req, err := http.NewRequest(action, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	awsauth.Sign(req)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func makeESRequest(es awsESQueryer, url, action string) ([]byte, error) {
	return es.esRequest(url, action)
}

func (es awsESIndexReaper) sortIndexesByAge() ([]indexSettings, error) {
	indexesSlice := []indexSettings{}

	for _, v := range es.indexList {
		indexesSlice = append(indexesSlice, v)
	}

	sort.Slice(indexesSlice, func(i, j int) bool {
		return indexesSlice[i].Settings.Index.CreationDate < indexesSlice[j].Settings.Index.CreationDate
	})
	return indexesSlice, nil
}

func main() {
	flag.Parse()

	flag.Lookup("logtostderr").Value.Set("true")

	if logLevel != "" {
		flag.Lookup("v").Value.Set(logLevel)
	} else {
		flag.Lookup("v").Value.Set("2")
	}

	if esURL == "" {
		glog.Fatal("Set ES_URL to the elasticsearch URL")
	}

	if indexPrefix == "" {
		glog.Fatal("Set INDEX_PREFIX to the name of the indexes to match")
	}

	if maxCountStr != "" {
		t, _ := strconv.Atoi(maxCountStr)
		maxCount = t
	}

	glog.Infof("Index reaper started.  Will keep %d indexes matching ^%s on ES domain %s", maxCount, indexPrefix, esURL)

	body, err := makeESRequest(indexReaper, fmt.Sprintf("%s/%s*/_settings?pretty", esURL, indexPrefix), "GET")
	if err != nil {
		glog.Fatal(err)
	}

	err = json.Unmarshal(body, &indexReaper.indexList)
	if err != nil {
		glog.Fatal(err)
	}

	sortedIndexes, err := indexReaper.sortIndexesByAge()
	if err != nil {
		glog.Fatal(err)
	}
	if len(sortedIndexes) > maxCount {

		for _, e := range sortedIndexes[:len(sortedIndexes)-maxCount] {

			iName := e.Settings.Index.ProvidedName
			glog.Infof("Deleting index %s", iName)

			body, err := indexReaper.esRequest(fmt.Sprintf(esURL+"/"+iName), "DELETE")
			if err != nil {
				glog.Errorf("An error occurred deleting index %s: %s", iName, err)
			}

			glog.V(4).Infof("Delete response is %s", body)
		}
	}
}
