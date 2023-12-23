package helpers

import (
	"bytes"
	
	"encoding/json"
	"net/http"

	"sm-system/configs"
)

func SendMessage(userId, text string) error {
	apiURL := "https://api.telegram.org/bot" + configs.Config("TELEGRAM_BOT_KEY") + "/sendMessage"

	request, err := json.Marshal(map[string]string{
		"chat_id": userId,
		"text": text,
	}); if err != nil {
		return err
	}

	_, err = http.Post(apiURL, "application/json", bytes.NewBuffer(request)); if err != nil {
		return err
	}

	return nil
}

