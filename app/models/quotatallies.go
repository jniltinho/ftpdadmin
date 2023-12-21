package models

// Quotatallies ProFTPd Quota table
type Quotatallies struct {
	Name          string `gorm:"column:name;type:varchar(30);not null;default:''" json:"name"`
	QuotaType     string `gorm:"column:quota_type;type:enum('user','group','class','all');not null;default:user" json:"quota_type"`
	BytesInUsed   uint32 `gorm:"column:bytes_in_used;type:int(10) unsigned;not null;default:0" json:"bytes_in_used"`
	BytesOutUsed  uint32 `gorm:"column:bytes_out_used;type:int(10) unsigned;not null;default:0" json:"bytes_out_used"`
	BytesXferUsed uint32 `gorm:"column:bytes_xfer_used;type:int(10) unsigned;not null;default:0" json:"bytes_xfer_used"`
	FilesInUsed   uint32 `gorm:"column:files_in_used;type:int(10) unsigned;not null;default:0" json:"files_in_used"`
	FilesOutUsed  uint32 `gorm:"column:files_out_used;type:int(10) unsigned;not null;default:0" json:"files_out_used"`
	FilesXferUsed uint32 `gorm:"column:files_xfer_used;type:int(10) unsigned;not null;default:0" json:"files_xfer_used"`
}

// TableName get sql table name.获取数据库表名
func (m *Quotatallies) TableName() string {
	return "quotatallies"
}
