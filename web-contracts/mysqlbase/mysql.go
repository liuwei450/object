package mysqlbase

import (
	"errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
	"web-contracts/config"
)

// DB 全局数据库连接实例（假设已通过 Init() 初始化）
var DB *gorm.DB

func InitDb() {
	var err error
	dsn := config.MYSQL_USER + ":" + config.MYSQL_PASSWORD + "@(" + config.MYSQL_URL + ")/" + config.MYSQL_DATABASE + "?" + config.MYSQL_CONFIG
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数（例如：User -> user，而非 users）
		},
	})
	if err != nil {
		log.Printf("MySQL 连接失败: %v, DSN: %s", err, dsn) // 生产环境删除 DSN 打印或脱敏
		panic(err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("获取连接池失败: %v", err)
		panic(err)
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns) // 从配置读取，例如 10
	sqlDB.SetMaxOpenConns(config.MaxOpenConns) // 从配置读取，例如 100
	sqlDB.SetConnMaxLifetime(time.Hour)        // 连接最大存活时间

	log.Println("MySQL 初始化成功")
	DB = db
}

// CloseConn 关闭数据库连接池，释放资源
// 返回错误：连接池关闭失败时返回错误
func CloseConn() error {
	if DB == nil {
		return errors.New("database connection is not initialized (DB is nil)")
	}

	// 1. 从 GORM 中获取底层 *sql.DB 实例
	sqlDB, err := DB.DB()
	if err != nil {
		return errors.New("failed to get underlying sql.DB: " + err.Error())
	}

	// 2. 关闭连接池（两步关闭法：先停止接受新连接，再等待现有连接完成）
	// Step 1: 设置最大连接数为 0，阻止新连接创建
	sqlDB.SetMaxOpenConns(0)
	// Step 2: 关闭连接池，等待现有连接最多 5 秒（超时后强制关闭）
	err = sqlDB.Close()
	if err != nil {
		return errors.New("failed to close database connection pool: " + err.Error())
	}

	// 3. 验证连接是否已关闭（可选，用于确认）
	if err := sqlDB.Ping(); err != nil {
		log.Println("database connection closed successfully")
	} else {
		return errors.New("database connection still active after close")
	}

	// 4. 重置全局 DB 实例（避免后续误用已关闭的连接）
	DB = nil
	return nil
}
