package rssgen

type Generator struct {
	ID          uint `gorm:"primary_key"`
	UUID        string
	UserID      uint
	Title       string
	Description string
	Url         string
	Rule        string
}

func (Generator) TableName() string {
	return "generator"
}
