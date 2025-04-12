package models

type Country struct {
	ID      uint   `gorm:"primaryKey"`
	Name    string `gorm:"not null"`
	Code    string `gorm:"size:2;unique;not null"`
	FlagURL string
}

type Tariff struct {
	ID            uint    `gorm:"primaryKey"`
	CountryID     uint    `gorm:"not null"`
	TargetCountry uint    `gorm:"not null"`
	Product       string  `gorm:"not null"`
	Type          string  `gorm:"not null"`
	Tariff        float64 `gorm:"not null"`
	LastUpdated   string
}
