package core

import (
	"gogrpcgin/conf"
	"math/rand"
	"gogrpcgin/utils"
	"encoding/json"
	"reflect"
)

type Elastic struct {
	Host string
	IndexName string
	TypeName  string
}

type EsResponse struct {

	Error struct{
		Type string `json:"type"` 
	}

	Status int64	`json:"status"`

	Shards struct{
		Failed int64 	 `json:"failed"`
		Successful int64 `json:"successful"`
		total int64 	 `json:"total"`
	}

	Hits struct{
		Hits []struct{
			Score float64 					`json:"_score"`
			Source map[string]interface{} 	`json:"_source"`
		}
		MaxScore float64 					`json:"max_score"`
		Total 	 int64 						`json:"total"`
	}
}


func MasterES(indexName string, typeName string) *Elastic {

	return &Elastic{
		Host: conf.Conf.ES[indexName]["master"][rand.Intn(len(conf.Conf.ES[indexName]["master"]))],
		IndexName:indexName,
		TypeName:typeName,
	}
}

func SlaveES(indexName string, typeName string) *Elastic{

	return &Elastic{
		Host: conf.Conf.ES[indexName]["slave"][rand.Intn(len(conf.Conf.ES[indexName]["slave"]))],
		IndexName:indexName,
		TypeName:typeName,
	}
}

func (e *Elastic)Query(q string, v *EsResponse) {

	indexUrl := e.Host + "/" + e.IndexName +"/" + e.TypeName + "/_search/"

	r,_ := utils.HttPost(indexUrl, q);

	json.Unmarshal(r, v)

	// if  Error   panic
	s:= reflect.ValueOf(v).Elem()
	typeOfT := s.Type()
	for i:=0;i<s.NumField();i++ {
		if typeOfT.Field(i).Name == "Error"{
			panic("elasticsearch is error :" + v.Error.Type)
		}
	}
}
