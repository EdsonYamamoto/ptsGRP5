package service

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/heroku/go-getting-started/src/model"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

//UtilBaseService ...
type UtilBaseService struct {
}

//DataBaseAccess ...
func (srvUtil *UtilBaseService) DataBaseAccess() (*firestore.Client, error) {
	opt := option.WithCredentialsFile("firebase.json")

	//opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	client, err := app.Firestore(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return client, err
}

//GetOneDataByDoc ...
func (srvUtil *UtilBaseService) GetOneDataByDoc(collection string, document string) (map[string]interface{}, error) {
	//retorna as informações do banco
	client, err := srvUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}
	result, err := client.Collection(collection).Doc(document).Get(context.Background())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	documento := map[string]interface{}{"id": document, "documento": result.Data()}
	//m := result.Data()
	defer client.Close()

	return documento, err
}

//GetAllDataFromCollection ...
func (srcUtil *UtilBaseService) GetAllDataFromCollection(collection string) ([]map[string]interface{}, error) {
	//retorna as informações do banco
	var results []map[string]interface{}

	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}
	//comando where é o mesmo do SQL SERVER
	//result := client.Collection(collection).Where("Deletado", "==", false).Documents(context.Background())
	result := client.Collection(collection).Where("Deletado", "==", false).Documents(context.Background())
	for {
		doc, err := result.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{"id": doc.Ref.ID, "documento": doc.Data()})
	}
	return results, err
}

//GetAllDataNotDeletedFromCollection ...
/*
	Retorna todos os dados não deletados de uma collection
*/
func (srcUtil *UtilBaseService) GetAllDataNotDeletedFromCollection(collection string) ([]map[string]interface{}, error) {
	//retorna as informações do banco
	var results []map[string]interface{}

	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}

	//result := client.Collection(collection).Where("Deletado", "==", false).Documents(context.Background())
	result := client.Collection(collection).Where("Deletado", "==", false).Documents(context.Background())
	for {
		doc, err := result.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{"id": doc.Ref.ID, "documento": doc.Data()})
	}
	return results, err
}

//GetAllSearchFromCollection ...
/*
	Retorna todos os dados não deletados de uma collection
*/
func (srcUtil *UtilBaseService) GetAllSearchFromCollection(collection string, where []model.ModelSearchDatabaseWhere) ([]map[string]interface{}, error) {
	//retorna as informações do banco
	var results []map[string]interface{}

	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return nil, err
	}

	result2 := client.Collection(collection).Where("Deletado", "==", false)

	for _, condicao := range where {
		result2 = result2.Where(condicao.Var1, condicao.Operador, condicao.Var2)
	}

	result3 := result2.Documents(context.Background())
	for {
		doc, err := result3.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		results = append(results, map[string]interface{}{"id": doc.Ref.ID, "documento": doc.Data()})
	}
	return results, err
}

//SaveUniqueInterfaceDataIntoCollection ...
func (srcUtil *UtilBaseService) SaveUniqueInterfaceDataIntoCollection(collection string, model interface{}) error {
	var mapString map[string]interface{}
	mapstructure.Decode(model, &mapString)
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}

	documentReference, writeResult, err := client.Collection(collection).Add(context.Background(), model)
	if err != nil {
		return err
	}
	log.Println("id do objeto salvo: " + documentReference.ID)
	log.Println("caminho salvo: " + documentReference.Parent.Path)
	log.Println(writeResult)
	return err
}

//SaveUniqueMapStringDataIntoCollection ...
func (srcUtil *UtilBaseService) SaveUniqueMapStringDataIntoCollection(collection string, mapString map[string]interface{}) error {
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}
	_, _, err = client.Collection(collection).Add(context.Background(), mapString)
	if err != nil {
		return err
	}
	return err
}

//SetDocumentoDeletedFromCollection ...
func (srcUtil *UtilBaseService) SetDocumentDeleteFromCollection(collection string, id string) error {
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}
	_, err = client.Collection(collection).Doc(id).Set(context.Background(), map[string]interface{}{"Deletado": true, "DataUpdate": time.Now()}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return err
}

//SetDocumentoDeletedFromCollection ...
func (srcUtil *UtilBaseService) SetDocumentFromCollection(collection string, id string, documento interface{}) error {
	client, err := srcUtil.DataBaseAccess()
	if err != nil {
		return err
	}
	_, err = client.Collection(collection).Doc(id).Set(context.Background(), map[string]interface{}{"DataUpdate": time.Now(), "Documento": documento}, firestore.MergeAll)
	if err != nil {
		return err
	}
	return err
}
