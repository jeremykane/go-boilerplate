package entity

type FacilityType struct {
	BaseEntity
	Name       string `gorm:"column:name" json:"name"`
	BahasaName string `gorm:"column:bahasa_name" json:"bahasaName"`
}
