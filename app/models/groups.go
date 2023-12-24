package models

import (
	"errors"
	"strings"

	"github.com/jniltinho/ftpdadmin/app/database"
	"gorm.io/gorm"
)

// Groups ProFTPd group table
type Groups struct {
	Groupname       string `gorm:"unique;column:groupname;type:varchar(32);not null;default:''" json:"groupname"`
	Gid             uint16 `gorm:"autoIncrement:true;primaryKey;column:gid;type:smallint(6) unsigned;not null" json:"gid"`
	Members         string `gorm:"column:members;type:varchar(255);not null;default:''" json:"members"`
	MainUsers       int    `gorm:"-" json:"user_count"`
	AdditionalUsers int    `gorm:"-" json:"additional_users"`
}

// TableName get sql table name.获取数据库表名
func (m *Groups) TableName() string {
	return "groups"
}

func (m *Groups) GetGroups() ([]Groups, error) {
	var groups []Groups
	database.DB().Find(&groups)
	return groups, nil
}

func (m *Groups) GetGroupByGid(gid uint16) (Groups, error) {
	var groups Groups
	err := database.DB().Where("gid = ?", gid).First(&groups).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return groups, ErrDataNotFound
	}
	return groups, nil
}

func GetGroupByGid(gid uint16) (Groups, error) {
	var group Groups
	err := database.DB().Where("gid = ?", gid).First(&group).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return group, ErrDataNotFound
	}

	return group, nil
}

func getGroupNameByGid(gid uint16) (Groups, error) {
	// Define User variable.
	var group Groups

	// Define query string.
	query := `SELECT groupname FROM groups WHERE gid = ?`

	// Send query to database.
	err := database.DB().Raw(query, gid).Scan(&group).Error
	if err != nil {
		// Return empty object and error.
		return group, err
	}

	// Return query result.
	return group, nil
}

func ListGroups() ([]Groups, error) {
	var groups []Groups
	database.DB().Find(&groups)

	for key, value := range groups {

		var countMembers int
		m := groups[key].Members
		if m != "" {
			countMembers = len(strings.Split(m, ","))
		}

		countMainUsers := countMainUsers(value.Gid)
		groups[key].MainUsers = countMainUsers
		groups[key].AdditionalUsers = countMembers
	}

	return groups, nil
}

func countMainUsers(gid uint16) int {
	var count int64
	database.DB().Model(&Users{}).Where("gid = ?", gid).Count(&count)
	return int(count)
}

func getAdditionalUsers(gid uint16) int {
	var count int64
	database.DB().Model(&Users{}).Where("gid = ?", gid).Count(&count)
	return int(count)
}
