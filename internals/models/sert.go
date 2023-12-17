package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Sert struct {
	gorm.Model
	ID                  uuid.UUID `gorm:"type:uuid"`
	NameOrganization    string    `json:"textbox_name_organization,omitempty" gorm:"column:name_organization"`
	TypeManagement      string    `json:"listbox_type_management_without_id,omitempty" gorm:"column:type_management"`
	DateValidity        string    `json:"date_validity,omitempty" gorm:"column:date_validity"`
	Status              string    `json:"listbox_status,omitempty" gorm:"column:status"`
	IIN                 string    `json:"textbox_IIN,omitempty" gorm:"column:iin"`
	RegNumber           string    `json:"textbox_reg_number,omitempty" gorm:"column:reg_number"`
	DataFrom            string    `json:"date_from,omitempty" gorm:"column:data_from"`
	NameArrc            string    `json:"textbox_name_arrc,omitempty" gorm:"column:name_arrc"`
	EntityExpert        string    `json:"entity_expert,omitempty" gorm:"column:entity_expert"`
	AddressOrganization string    `json:"textbox_adress_organization,omitempty" gorm:"column:address_organization"`
}
