package main

import (
	"log"
	"reflect"
	"testing"
	"testing/fstest"

	"gopkg.in/yaml.v2"
)

func TestGetConfig(t *testing.T) {
	//must be spaces no tabs to get a well yaml formatted file
	var data1 = `
      fqdn: primary1.example.com
      apiKey: A7wqi7BrRQg-yNpz-C2eKV2npG0Ocy2_Ssv5OhjIOLnHq6_CCQwf-111111111111
      timeFrame:
        num: 8
        freq: hours`
	var data2 = `
      fqdn: primary2.example.com
      apiKey: A7wqi7BrRQg-yNpz-C2eKV2npG0Ocy2_Ssv5OhjIOLnHq6_CCQwf-111111111111
      timeFrame:
        num: 4
        freq: hours`

	fs := fstest.MapFS{
		"domain1.yml": {Data: []byte(data1)},
		"domain2.yml": {Data: []byte(data2)},
	}

	stubPrimaries := make([]Primary, 2)
	stubPrimaries[0].Fqdn = "primary1.example.com"
	stubPrimaries[0].APIKey = "A7wqi7BrRQg-yNpz-C2eKV2npG0Ocy2_Ssv5OhjIOLnHq6_CCQwf-111111111111"
	stubPrimaries[0].TimeFrame.Freq = "hours"
	stubPrimaries[0].TimeFrame.Num = 8
	stubPrimaries[1].APIKey = "A7wqi7BrRQg-yNpz-C2eKV2npG0Ocy2_Ssv5OhjIOLnHq6_CCQwf-111111111111"
	stubPrimaries[1].Fqdn = "primary2.example.com"
	stubPrimaries[1].TimeFrame.Freq = "hours"
	stubPrimaries[1].TimeFrame.Num = 4

	got, err := GetPrimariesFromFS(fs)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	assertDomains(t, got, stubPrimaries)

}
func TestGetEntitiesFromFS(t *testing.T) {
	t.Helper()

	var data = `
      name: Dummy_Entity
      primaries:
        - name: rpro02bck.aytonda.lan
        - name: rpro02bck.aytonda.lan
      workloads:
        - type: HYPERVISOR
          workloadDisplayName:
            - dummy.example.com
            - dummy2.example.com
        - type: SQL_SERVER
          instanceDatabaseName:
            - MSSQLSERVER\\DummyDB
            - MSSQLSERVER\\DummyDB2`

	fs := fstest.MapFS{
		"entity-1.yaml": {Data: []byte(data)},
	}
	got, _ := GetEntitiesFromFS(fs)
	var entity Entity
	var entities []Entity
	yaml.Unmarshal([]byte(data), &entity)
	entities = append(entities, entity)
	assertEntities(t, got, entities)
}

func assertDomains(t *testing.T, got []Primary, want []Primary) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertEntities(t *testing.T, got []Entity, want []Entity) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
