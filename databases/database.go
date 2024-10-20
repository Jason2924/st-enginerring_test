package databases

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	con "github.com/Jason2924/st-enginerring_test/config"
	ntt "github.com/Jason2924/st-enginerring_test/entities"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
	once sync.Once
)

type MysqlDatabase interface {
	Connect() *gorm.DB
	Close() error
	Ping(context.Context) error
}

type connection struct {
	idle     int
	maximum  int
	lifetime time.Duration
}

type mysqlDatabase struct {
	host         string
	name         string
	username     string
	password     string
	port         string
	migrateTable bool
	// importFile   bool
	connection *connection
}

func NewMysqlDatabase(msql *con.ConfigMysql) MysqlDatabase {
	pool := &connection{
		idle:     10,
		maximum:  20,
		lifetime: 1 * time.Hour,
	}
	return &mysqlDatabase{
		host:         msql.Host,
		name:         msql.Name,
		username:     msql.Username,
		password:     msql.Password,
		port:         "3306",
		migrateTable: msql.MigrateTable,
		connection:   pool,
	}
}

// creating a connection to database
func (dtb *mysqlDatabase) Connect() *gorm.DB {
	once.Do(func() {
		// connnecting to database
		link := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", "root", "mh22T7uDi61XC14Udw7b", dtb.host, dtb.port, dtb.name)
		var erro error
		conn, erro = gorm.Open(mysql.Open(link), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if erro != nil {
			log.Fatalln("Error occured while connecting database", erro)
		}
		// set conection pool
		if dtb.connection != nil {
			dtbs, erro := conn.DB()
			if erro != nil {
				log.Fatalln("Error occured while setting connection pool", erro)
			}
			dtbs.SetMaxIdleConns(dtb.connection.idle)
			dtbs.SetMaxOpenConns(dtb.connection.maximum)
			dtbs.SetConnMaxLifetime(dtb.connection.lifetime)
		}
		// auto migrate schema
		if dtb.migrateTable {
			// drop all existed tables first
			migr := conn.Migrator()
			migr.DropTable(&ntt.ProductSchema{})
			// then migrate new tables
			conn.Set("gorm:table_options", "ENGINE=InnoDB")
			if erro := conn.AutoMigrate(&ntt.ProductSchema{}); erro != nil {
				log.Fatalln("Error occured while migrating table", erro)
			}
		}
	})
	return conn
}

// closing a connection to database
func (dtb *mysqlDatabase) Close() error {
	if conn == nil {
		return nil
	}
	dtbs, erro := conn.DB()
	if erro != nil {
		return erro
	}
	return dtbs.Close()
}

// pinging a connection to database
func (dtb *mysqlDatabase) Ping(ctxt context.Context) error {
	dtbs, erro := dtb.Connect().DB()
	if erro != nil {
		return erro
	}
	return dtbs.PingContext(ctxt)
}
