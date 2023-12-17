package workers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"sm-system/configs"
	"sm-system/internals/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func FetchDataAndSave(db *gorm.DB) error {
	url := "http://techreg.gov.kz/Synergy/rest/api/registry/data_ext?registryCode=certificates_reg_sm&pageNumber=0&countInPart=300000&fields=textbox_reg_number&fields=textbox_name_organization&fields=textbox_IIN&fields=date_from&fields=date_validity&fields=listbox_status&fields=listbox_type_management_without_id&fields=textbox_name_arrc&fields=entity_expert&fields=textbox_adress_organization"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Authorization", "Basic " + configs.Config("BASIC_AUTH_TOKEN"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ошибка при чтении тела ответа: %v", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		return fmt.Errorf("ошибка при разборе JSON: %v", err)
	}

	results := data["result"].([]interface{})
	for _, res := range results {
		result := res.(map[string]interface{})
		fieldValueData := result["fieldValue"].(map[string]interface{})

		iin := getString(fieldValueData["textbox_IIN"])
		regNumber := getString(fieldValueData["textbox_reg_number"])

		var existingSert models.Sert
		if err := db.Where("iin = ? AND reg_number = ?", iin, regNumber).First(&existingSert).Error; err != nil {
			if err != gorm.ErrRecordNotFound {
				log.Printf("Ошибка при поиске в базе данных: %v", err)
				continue
			}

			sert := models.Sert{
				ID:                  uuid.New(),
				NameOrganization:    getString(fieldValueData["textbox_name_organization"]),
				NameArrc:            getString(fieldValueData["textbox_name_arrc"]),
				TypeManagement:      getString(fieldValueData["listbox_type_management_without_id"]),
				DateValidity:        getString(fieldValueData["date_validity"]),
				AddressOrganization: getString(fieldValueData["textbox_adress_organization"]),
				Status:              getString(fieldValueData["listbox_status"]),
				EntityExpert:        getString(fieldValueData["entity_expert"]),
				IIN:                 iin,
				RegNumber:           regNumber,
				DataFrom:            getString(fieldValueData["date_from"]),
			}

			if err := db.Create(&sert).Error; err != nil {
				log.Printf("Ошибка при сохранении в базу данных: %v", err)
			}
		} else {
			log.Println("Сертификат уже существует в базе данных, пропускаем добавление")
		}
	}

	log.Printf("Данные успешно получены и сохранены")
	return nil
}

func getString(val interface{}) string {
	if val == nil {
		return ""
	}
	strVal, ok := val.(string)
	if !ok {
		return ""
	}
	return strVal
}
