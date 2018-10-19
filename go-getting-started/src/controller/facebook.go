package controller

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/heroku/go-getting-started/src/model"
	srv "github.com/heroku/go-getting-started/src/service"
)

//EAAiOFq6ZCSIoBAKOJs7oYkYPnmRcrjVrtPsv078RnQxVHdXUhA7lEmwDwGaZATPDkyBjm6QguVSjGsL79dqOf71nMFPg7ZCLZArZBefGoZAGk20RRM7cZBJPC7oftr1Mv4mGX2lJjRiVA7prZCrnq9G64VnbsWn0wZBn3ZB39uZCnFALvM9n1YeuGhm
const (
	FACEBOOK_API = "https://graph.facebook.com/v2.6/me/messages?access_token="
	IMAGE        = "http://37.media.tumblr.com/e705e901302b5925ffb2bcf3cacb5bcd/tumblr_n6vxziSQD11slv6upo3_500.gif"
)

/*
	Realiza o acesso no banco de dados retornando todas os documentos das collections
*/
//GETFacebookWebhook ...
func GETFacebookWebhook(c *gin.Context) {
	challenge := c.Query("hub.challenge")
	mode := c.Query("hub.mode")
	token := c.Query("hub.verify_token")
	println(challenge)
	println(mode)
	println(token)

	if mode != "" && token == "you know nothing, john snot" {
		c.String(http.StatusOK, challenge)
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Message not supported",
		})
	}
}

//POSTFacebookWebhook ...
func POSTFacebookWebhook(c *gin.Context) {

	log.Println(time.Now())
	var callback model.FacebookCallback
	err := c.Bind(&callback)
	log.Println(callback)
	if err != nil {
		log.Println("Erro no webhook do facebook: " + err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Erro no webhook do facebook",
		})
		return
	}
	if callback.Object == "page" {
		for _, entry := range callback.Entry {
			for _, event := range entry.Messaging {
				ProcessMessage(event)
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Got your message",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Message not supported",
		})
	}
}

//ProcessMessage ...
func ProcessMessage(event model.FacebookMessaging) {
	client := &http.Client{}
	var response model.FacebookResponse
	if event.Message.Text == "GrupoAgendaMedicaSorocabaGrupoPTS5" {
		response = model.FacebookResponse{
			Recipient: model.FacebookUser{
				ID: event.Sender.ID,
			},
			Message: model.FacebookMessage{
				Text: "Ola você recebeu sua mensagem, de volta",
			},
		}

	} else if event.Message.Text == "sim" {
		log.Println("SIM")
		log.Println(time.Now())
	} else if event.Message.Text == "nao" {
		log.Println("NAO")
		log.Println(time.Now())
	} else {
		response = model.FacebookResponse{
			Recipient: model.FacebookUser{
				ID: event.Sender.ID,
			},
			Message: model.FacebookMessage{
				Attachment: &model.FacebookAttachment{
					Type: "image",
					Payload: model.FacebookPayload{
						URL: IMAGE,
					},
				},
			},
		}
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(&response)

	apiFace, err := srv.GetAPI("facebook")
	if err != nil {
		log.Println(err)
	}
	url := apiFace.APIServer + apiFace.APIPassword
	req, err := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
}
