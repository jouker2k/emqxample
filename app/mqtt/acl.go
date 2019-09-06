package mqtt

import "github.com/jinzhu/gorm"

type MqttUser struct {
	ID          uint `gorm:"primary_key"`
	IsSuperuser bool
	Username    string
	Password    string
	Salt        string
}

func (MqttUser) TableName() string {
	return "mqtt_user"
}

type MqttAcl struct {
	ID       uint `gorm:"primary_key"`
	Allow    int
	Ipaddr   *string
	Username *string
	Clientid *string
	Access   int
	Topic    string
}

func (MqttAcl) TableName() string {
	return "mqtt_acl"
}

func aclDenyAllTopic(db *gorm.DB, t string) error {
	rule := &MqttAcl{
		Allow:    0,
		Ipaddr:   nil,
		Username: nil,
		Clientid: nil,
		Access:   3,
		Topic:    t}

	return db.Save(rule).Error
}

func aclAllowUserTopic(db *gorm.DB, u string, t string) error {
	rule := &MqttAcl{
		Allow:    1,
		Ipaddr:   nil,
		Username: &u,
		Clientid: nil,
		Access:   3,
		Topic:    t,
	}
	return db.Save(rule).Error
}
