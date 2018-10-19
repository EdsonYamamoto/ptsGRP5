package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/src/model"
	srv "github.com/heroku/go-getting-started/src/service"
)

//GETAllCollection ...
/*
	Realiza o acesso no banco de dados retornando todas os documentos das collections
*/
func GETAllCollection(c *gin.Context) {
	name := c.Param("name")
	serviceUtil := srv.UtilBaseService{}
	mapString, err := serviceUtil.GetAllDataFromCollection(name)
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, mapString)
}

//POSTSaveCollection
/*
	Metodo destinado a salvar informações que cheguem como JSON no banco
*/
func POSTSaveCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.DocJSON
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	var nomeCollection string
	nomeCollection = documento["collection"].(string)

	var variaveis model.ModelDocumento
	variaveis.DataCriacao = time.Now()
	variaveis.DataUpdate = time.Now()
	variaveis.Deletado = false
	variaveis.Documento = documento["documento"].(interface{})

	srvUtil := srv.UtilBaseService{}

	//tentativas de armazear no banco
	var booleanaEnvioDados bool
	booleanaEnvioDados = false

	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		err = srvUtil.SaveUniqueInterfaceDataIntoCollection(nomeCollection, variaveis)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao gravar no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao gravar no banco",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Gravado no banco",
	})
	return
}

func POSTGetAllDocumentsNotDeletedFromCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.DocJSON
	var response []map[string](interface{})
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	nomeCollection := documento["collection"].(string)

	//tentativas de armazear no banco
	var booleanaEnvioDados bool
	booleanaEnvioDados = false

	serviceUtil := srv.UtilBaseService{}
	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		response, err = serviceUtil.GetAllDataNotDeletedFromCollection(nomeCollection)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao gravar no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao gravar no banco",
		})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

//POSTGetAllSearchFromCollection ...
/*
	Busca por dados que vem em de forma especifica da web
*/
func POSTGetAllSearchFromCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.ModelSearchDatabase
	var response []map[string](interface{})
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	nomeCollection := documento.Collection

	//tentativas de armazear no banco
	var booleanaEnvioDados bool
	booleanaEnvioDados = false

	serviceUtil := srv.UtilBaseService{}
	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		response, err = serviceUtil.GetAllSearchFromCollection(nomeCollection, documento.Where)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao gravar no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao gravar no banco",
		})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

func POSTGetOneSearchFromCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.UniqueData
	var response map[string](interface{})
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//tentativas de armazear no banco
	var booleanaEnvioDados bool
	booleanaEnvioDados = false

	serviceUtil := srv.UtilBaseService{}
	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		response, err = serviceUtil.GetOneDataByDoc(documento.Collection, documento.ID)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao gravar no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao gravar no banco",
		})
		return
	}

	c.JSON(http.StatusOK, response)
	return
}

//POSTUpdateCollection ...
/*
	Update dos dados
*/
func POSTUpdateCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.DocJSON
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	nomeCollection := documento["collection"].(string)
	log.Println(nomeCollection)
}

//POSTdeleteDocumentCollection ...
/*
	Deletar logicamente um documento
*/
func POSTdeleteDocumentCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.DocJSON
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	nomeCollection := documento["collection"].(string)
	idDocumento := documento["id"].(string)
	log.Println(nomeCollection)
	srvUtil := srv.UtilBaseService{}
	booleanaEnvioDados := false
	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		err = srvUtil.SetDocumentDeleteFromCollection(nomeCollection, idDocumento)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao dar update no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao dar update no banco",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alterado com sucesso"})
}

//POSTupdateDocumentCollection ...
/*
	Deletar logicamente um documento
*/
func POSTupdateDocumentCollection(c *gin.Context) {
	//Recebimento de dados JSON
	var documento model.DocJSON
	err := c.Bind(&documento)
	if err != nil {
		log.Println("Verificar o JSON: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Verificar o JSON",
		})
		return
	}

	//Separação dos Dados na estrutura
	nomeCollection := documento["collection"].(string)
	idDocumento := documento["id"].(string)
	documentoJSON := documento["documento"].(interface{})

	srvUtil := srv.UtilBaseService{}
	booleanaEnvioDados := false
	apiFirebase, err := srv.GetAPI("firebase")
	if err != nil {
		log.Println("Erro ao encontrar API firebase")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro na API firebase",
		})
	}
	for index := 0; index < apiFirebase.APINumeroTentativa && !booleanaEnvioDados; index++ {
		err = srvUtil.SetDocumentFromCollection(nomeCollection, idDocumento, documentoJSON)
		if err == nil {
			booleanaEnvioDados = true
		} else {
			time.Sleep(time.Millisecond * 500)
		}
	}
	if err != nil {
		log.Println("Erro ao dar update no banco: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro ao dar update no banco",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alterado com sucesso"})
}
