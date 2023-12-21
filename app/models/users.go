package models

import (
	"errors"
	"time"

	"github.com/jniltinho/ftpdadmin/app/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Users ProFTPd user table
type Users struct {
	ID             uint16    `gorm:"autoIncrement:true;primaryKey;column:id;type:smallint(2) unsigned;not null" json:"id"`
	Userid         string    `gorm:"unique;column:userid;type:varchar(64);not null;default:''" json:"userid"`
	UID            uint16    `gorm:"column:uid;type:smallint(6) unsigned;default:null" json:"uid"`
	Gid            uint16    `gorm:"column:gid;type:smallint(6) unsigned;default:null" json:"gid"`
	Passwd         string    `gorm:"column:passwd;type:varchar(265);not null;default:''" json:"passwd"`
	Homedir        string    `gorm:"column:homedir;type:varchar(255);not null;default:''" json:"homedir"`
	Comment        string    `gorm:"column:comment;type:varchar(255);not null;default:''" json:"comment"`
	Disabled       uint16    `gorm:"column:disabled;type:smallint(2) unsigned;not null;default:0" json:"disabled"`
	Shell          string    `gorm:"column:shell;type:varchar(32);not null;default:/bin/false" json:"shell"`
	SSHpubkey      string    `gorm:"column:sshpubkey;type:varchar(8192);not null;default:''" json:"sshpubkey"`
	Email          string    `gorm:"column:email;type:varchar(255);not null;default:''" json:"email"`
	Name           string    `gorm:"column:name;type:varchar(255);not null;default:''" json:"name"`
	Title          string    `gorm:"column:title;type:varchar(5);not null;default:''" json:"title"`
	Company        string    `gorm:"column:company;type:varchar(255);not null;default:''" json:"company"`
	QuotaType      string    `gorm:"column:quota_type;type:enum('user','group','class','all');not null;default:user" json:"quota_type"`
	PerSession     string    `gorm:"column:per_session;type:enum('false','true');not null;default:false" json:"per_session"`
	LimitType      string    `gorm:"column:limit_type;type:enum('soft','hard');not null;default:soft" json:"limit_type"`
	BytesInUsed    uint64    `gorm:"column:bytes_in_used;type:bigint(20) unsigned;not null;default:0" json:"bytes_in_used"`
	BytesOutUsed   uint64    `gorm:"column:bytes_out_used;type:bigint(20) unsigned;not null;default:0" json:"bytes_out_used"`
	BytesXferAvail uint32    `gorm:"column:bytes_xfer_avail;type:int(10) unsigned;not null;default:0" json:"bytes_xfer_avail"`
	FilesInUsed    uint64    `gorm:"column:files_in_used;type:bigint(20) unsigned;not null;default:0" json:"files_in_used"`
	FilesOutUsed   uint64    `gorm:"column:files_out_used;type:bigint(20) unsigned;not null;default:0" json:"files_out_used"`
	FilesXferAvail uint32    `gorm:"column:files_xfer_avail;type:int(10) unsigned;not null;default:0" json:"files_xfer_avail"`
	LoginCount     uint32    `gorm:"column:login_count;type:int(11) unsigned;not null;default:0" json:"login_count"`
	LastLogin      time.Time `gorm:"column:last_login;type:datetime;not null;default:0000-00-00 00:00:00" json:"last_login"`
	LastModified   time.Time `gorm:"column:last_modified;type:datetime;not null;default:0000-00-00 00:00:00" json:"last_modified"`
	Expiration     time.Time `gorm:"column:expiration;type:datetime;not null;default:0000-00-00 00:00:00" json:"expiration"`
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}

func (u *Users) GetUsers() ([]Users, error) {
	var users []Users
	database.DB().Find(&users)
	return users, nil
}

// GetFirstByID gets the user by his ID
func (u *Users) GetFirstByID(id string) error {
	err := database.DB().Where("id=?", id).First(u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

// GetFirstByEmail gets the user by his email
func (u *Users) GetFirstByEmail(email string) error {
	err := database.DB().Where("email=?", email).First(u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrDataNotFound
	}

	return err
}

// Create a new user
func (u *Users) Create() error {
	db := database.DB().Create(u)

	if db.Error != nil {
		return db.Error
	} else if db.RowsAffected == 0 {
		return ErrKeyConflict
	}

	return nil
}

// Signup a new user
func (u *Users) Signup() error {
	var user Users
	err := user.GetFirstByEmail(u.Email)

	if err == nil {
		return ErrUserExists
	} else if err != ErrDataNotFound {
		return err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Passwd), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// replace the plaintext password with ciphertext password
	u.Passwd = string(hash)

	return u.Create()
}

// Login a user
func (u *Users) Login(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Passwd), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

// LoginByEmailAndPassword login a user by his email and password
func LoginByEmailAndPassword(email, password string) (*Users, error) {
	var user Users
	err := user.GetFirstByEmail(email)
	if err != nil {
		return &user, err
	}

	return &user, user.Login(password)
}
