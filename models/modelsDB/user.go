package modelsDB

type User struct {
	UserID   string `gorm:"primaryKey;column:UserID"`
	UserName string `gorm:"column:UserName"`
	Head     bool   `gorm:"column:head"`
	RoomID   string `gorm:"column:roomID"`
	ImgID    string `gorm:"column:imgID"`
	Ready    bool   `gorm:"column:ready"`

	Room *Room `gorm:"foreignKey:RoomID"`
}

// TableName overrides the default table name used by GORM
func (User) TableName() string {
	return "User" // Use "User" as the table name
}
