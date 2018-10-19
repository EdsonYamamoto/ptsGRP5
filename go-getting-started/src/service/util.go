package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"

	"github.com/heroku/go-getting-started/src/model"
)

/*
	Funções extras que ajudam na programação
*/
func GetAPI(apiName string) (model.API, error) {
	var allAPI []model.API
	var API model.API
	raw, err := ioutil.ReadFile("./appDoc.json")
	if err != nil {
		log.Fatalln(err)
	}
	json.Unmarshal(raw, &allAPI)
	for i := 0; i < len(allAPI); i++ {
		if allAPI[i].APIName == apiName {
			API = allAPI[i]
			return API, nil
		}
	}
	return API, errors.New("API nao encontrada")
}
