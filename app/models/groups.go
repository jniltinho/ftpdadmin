package models

// Groups ProFTPd group table
type Groups struct {
	Groupname string `gorm:"unique;column:groupname;type:varchar(32);not null;default:''" json:"groupname"`
	Gid       uint16 `gorm:"autoIncrement:true;primaryKey;column:gid;type:smallint(6) unsigned;not null" json:"gid"`
	Members   string `gorm:"column:members;type:varchar(255);not null;default:''" json:"members"`
}

// TableName get sql table name.获取数据库表名
func (m *Groups) TableName() string {
	return "groups"
}

func (m *Groups) GetGroups() ([]Groups, error) {
	var groups []Groups
	DB().Find(&groups)
	return groups, nil
}
