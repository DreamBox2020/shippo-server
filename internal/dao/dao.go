package dao

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"shippo-server/utils/config"
	"time"
)

type DaoGroup struct {
	User             *UserDao
	Temp             *TempDao
	Passport         *PassportDao
	Captcha          *CaptchaDao
	Album            *AlbumDao
	Role             *RoleDao
	PermissionAccess *PermissionAccessDao
	PermissionPolicy *PermissionPolicyDao
	Picture          *PictureDao
	WxArticle        *WxArticleDao
	WxCommentLike    *WxCommentLikeDao
	WxOffiaccount    *WxOffiaccountDao
	WxComment        *WxCommentDao
	WxPassport       *WxPassportDao
}

type Dao struct {
	db    *gorm.DB
	Group *DaoGroup
}

func New() *Dao {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DB.DSN, // DSN data source name
		DefaultStringSize:         256,           // string 类型字段的默认长度
		DisableDatetimePrecision:  true,          // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,          // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,          // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,         // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "shippo_",
			SingularTable: true, // 使用单数表名
		},
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
				LogLevel:                  logger.Info,            // Log level
				IgnoreRecordNotFoundError: false,
				Colorful:                  config.IsLocal(), // 非ide环境禁用彩色打印
			},
		),
	})

	if err != nil {
		panic(err)
	}

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	d := &Dao{
		db:    db,
		Group: nil,
	}
	d.initGroup()

	return d
}

func (t *Dao) initGroup() {
	t.Group = &DaoGroup{
		User:             NewUserDao(t),
		Temp:             NewTempDao(t),
		Passport:         NewPassportDao(t),
		Captcha:          NewCaptchaDao(t),
		Album:            NewAlbumDao(t),
		Role:             NewRoleDao(t),
		PermissionAccess: NewPermissionAccessDao(t),
		PermissionPolicy: NewPermissionPolicyDao(t),
		Picture:          NewPictureDao(t),
		WxArticle:        NewWxArticleDao(t),
		WxCommentLike:    NewWxCommentLikeDao(t),
		WxOffiaccount:    NewWxOffiaccountDao(t),
		WxComment:        NewWxCommentDao(t),
		WxPassport:       NewWxPassportDao(t),
	}
}

func (t *Dao) IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}

func newTest() *Dao {
	config.Init()
	return New()
}
