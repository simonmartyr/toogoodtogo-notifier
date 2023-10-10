package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/simonmartyr/toogoodtogogo"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const (
	subject = "Too Good To Go Checker"
)

type TgtgNotifier struct {
	Config *Config
	Client *toogoodtogo.Client
}

func New(config *Config) *TgtgNotifier {
	credentials := config.Credentials
	return &TgtgNotifier{
		Config: config,
		Client: toogoodtogo.New(
			&http.Client{},
			toogoodtogo.WithAuth(credentials.UserId, credentials.AccessToken, credentials.RefreshToken, credentials.Cookie),
		),
	}
}

func (t *TgtgNotifier) SendNotification(items []*toogoodtogo.Item) error {
	log.Println("Preparing email")
	emailConfig := t.Config.EmailConfig
	emailContent, emailErr := CreateEmail(ParseItems(items))
	if emailErr != nil {
		return emailErr
	}
	emailMessage := fmt.Sprintf("To: %s\nSubject: %s\nContent-Type: text/html\n\n%s", emailConfig.To, subject, emailContent)
	cmd := exec.Command("msmtp", "-a", emailConfig.Account, emailConfig.To)
	cmd.Stdin = strings.NewReader(emailMessage)
	cmdErr := cmd.Run()
	if cmdErr != nil {
		return cmdErr
	}
	log.Println("Email sent")
	return nil
}

func (t *TgtgNotifier) Refresh() {
	log.Println("Refreshing auth")
	refErr := t.Client.RefreshAuthentication()
	if refErr != nil {
		fmt.Printf("Failed to get item %s", refErr.Error())
		return
	}
	credentials := t.Client.GetCredentials()
	t.Config.Credentials.AccessToken = credentials.AccessToken
	t.Config.Credentials.RefreshToken = credentials.RefreshToken
	t.Config.Credentials.Cookie = credentials.Cookie
	log.Println("Refresh successful")
}

func (t *TgtgNotifier) Save() {
	log.Println("Saving Config")
	toWrite, writeErr := json.Marshal(t.Config)
	if writeErr != nil {
		log.Fatal(writeErr)
		return
	}
	osWErr := os.WriteFile("tgtgconfig.json", toWrite, 0644)
	if osWErr != nil {
		log.Fatal(osWErr)
		return
	}
	log.Println("Save Complete")
}

func CreateEmail(content *[]PickupInfo) (string, error) {
	tmpl, templateErr := template.New("emailTemplate").Parse(emailTemplate)
	if templateErr != nil {
		return "", templateErr
	}
	var output bytes.Buffer
	emailContent := struct{ Items *[]PickupInfo }{Items: content}
	exeErr := tmpl.Execute(&output, emailContent)
	if exeErr != nil {
		return "", exeErr
	}
	return output.String(), nil
}
