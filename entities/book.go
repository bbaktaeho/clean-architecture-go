package entities

type Book struct {
	Id          uint64 `gorm:"primary_key:auro_increment" json:"id"`
	Title       string `gorm:"type:varchar(255)", json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserId      uint64 `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
