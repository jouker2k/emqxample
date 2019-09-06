package mqtt

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

type Store struct {
	db *gorm.DB
}

func NewStore(h string, p string, u string, n string, pass string, migrate bool) *Store {
	db, err := gorm.Open(
		"postgres",
		"host="+h+" port="+p+" user="+u+" dbname="+n+" sslmode=disable password="+pass)
	if err != nil {
		log.Fatalf("Error cannot open database: %v \n", err)
	}
	s := &Store{db}

	if migrate {
		s.migrate()
	}

	if err := s.init(); err != nil {
		log.Fatalf("Cannot init: %s", err)
	}

	return s
}

func (s *Store) init() error {
	//be nicer
	a := &[]MqttUser{}
	s.db.Find(a)
	if len(*a) != 0 {
		return nil
	}

	aUser := &MqttUser{
		IsSuperuser: false,
		Username:    "admin",
		Password:    encryptPass("pass"),
		Salt:        "salt",
	}

	cUser := &MqttUser{
		IsSuperuser: false,
		Username:    "client",
		Password:    encryptPass("pass"),
		Salt:        "salt",
	}

	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Save(&aUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Save(&cUser).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := aclDenyAllTopic(tx, "#"); err != nil {
		tx.Rollback()
		return err
	}

	//todo manage versions
	if err := aclAllowUserTopic(tx, "admin", DeviceTopic("/v1.0", "1")); err != nil {
		tx.Rollback()
		return err
	}

	if err := aclAllowUserTopic(tx, "client", DeviceTopic("/v1.0", "1")); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *Store) migrate() {
	s.db.AutoMigrate(&MqttUser{})
	s.db.AutoMigrate(&MqttAcl{})
}
