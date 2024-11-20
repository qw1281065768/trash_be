// Package model contains all the models required
// for a functional database management system
package model

import (
	"time"

	"gorm.io/gorm"
)

// User model - demo 待删除
type User struct {
	UserID    uint64         `gorm:"primaryKey" json:"userID,omitempty"`
	CreatedAt time.Time      `json:"createdAt,omitempty"`
	UpdatedAt time.Time      `json:"updatedAt,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName string         `json:"firstName,omitempty"`
	LastName  string         `json:"lastName,omitempty"`
	IDAuth    uint64         `json:"-"`
	Posts     []Post         `gorm:"foreignkey:IDUser;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"posts,omitempty"`
	Hobbies   []Hobby        `gorm:"many2many:user_hobbies" json:"hobbies,omitempty"`
}

// 线上表结构
/*type User1 struct {
	ID         int64     `gorm:"column:id" db:"id" json:"id" form:"id"`
	Username   string    `gorm:"column:username" db:"username" json:"username" form:"username"`
	Email      string    `gorm:"column:email" db:"email" json:"email" form:"email"`
	Phone      string    `gorm:"column:phone" db:"phone" json:"phone" form:"phone"`
	Passport   string    `gorm:"column:passport" db:"passport" json:"passport" form:"passport"`
	Role       string    `gorm:"column:role" db:"role" json:"role" form:"role"`
	Count      string    `gorm:"column:count" db:"count" json:"count" form:"count"`
	Vipdate    time.Time `gorm:"column:vipdate" db:"vipdate" json:"vipdate" form:"vipdate"`
	Createdat  time.Time `gorm:"column:createdat" db:"createdat" json:"createdat" form:"createdat"`
	Updatedat  time.Time `gorm:"column:updatedat" db:"updatedat" json:"updatedat" form:"updatedat"`
	Prompt     int64     `gorm:"column:prompt" db:"prompt" json:"prompt" form:"prompt"`
	Completion int64     `gorm:"column:completion" db:"completion" json:"completion" form:"completion"`
}*/

// 线上表结构
type User1 struct {
	ID         int64     `json:"ID,omitempty"`
	Username   string    `json:"username,omitempty"`
	Email      string    `json:"email,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	Passport   string    `json:"passport,omitempty"`
	Role       string    `json:"role,omitempty"`
	Count      string    `json:"count,omitempty"`
	Vipdate    time.Time `json:"vipdate"`
	CreatedAt  time.Time `gorm:"column:createdAt" db:"createdAt" json:"createdAt" form:"createdAt"`
	UpdatedAt  time.Time `gorm:"column:updatedAt" db:"updatedAt" json:"updatedAt" form:"updatedAt"`
	Prompt     int64     `json:"prompt,omitempty"`
	Completion int64     `json:"completion,omitempty"`
}

func (User1) TableName() string {
	return "users"
}

