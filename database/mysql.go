package database

import (
	"fmt"

	"promptpay-payment-gateway/configs"

	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mySqlDB struct {
	Client *gorm.DB
}

func NewMySqlDB(conf configs.MySQL, secrets configs.Secrets) (*mySqlDB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?checkConnLiveness=false&loc=Local&parseTime=true&readTimeout=%s&timeout=%s&writeTimeout=%s&maxAllowedPacket=0",
		secrets.MySqlUsername,
		secrets.MySqlPassword,
		conf.Host,
		conf.Database,
		conf.Timeout,
		conf.Timeout,
		conf.Timeout,
	)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, errors.Wrap(err, "gorm open")
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.Wrap(err, "get db")
	}

	sqlDB.SetMaxIdleConns(int(conf.MaxIdleConns))
	sqlDB.SetMaxOpenConns(*conf.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(conf.MaxLifetime)

	return &mySqlDB{db}, nil
}

func (db *mySqlDB) Close() error {
	mysql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "client db")
	}
	if err := mysql.Close(); err != nil {
		return errors.Wrap(err, "mysql close")
	}
	return nil
}

func (db *mySqlDB) Ping() error {
	sql, err := db.Client.DB()
	if err != nil {
		return errors.Wrap(err, "mysql")
	}
	err = sql.Ping()
	return errors.Wrap(err, "mysql")
}
