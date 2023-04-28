package main

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name       string
	CreditCard CreditCard `gorm:"foreignkey:UserName;references:name"`
}
type CreditCard struct {
	gorm.Model
	Number   string
	UserName string
}
type Customer struct {
	gorm.Model
}
type Cat struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}
type Dog struct {
	ID   int
	Name string
	Toy  Toy `gorm:"polymorphic:Owner;"`
}
type Toy struct {
	ID        int
	Name      string
	OwnerID   int
	OwnerType string
}
type Reader struct {
	gorm.Model
	ReaderName string
	Books      []Book `gorm:"foreignKey:BookReader;references:ReaderName"`
}
type Book struct {
	gorm.Model
	BookReader string
	BookName   string
}
type Developer struct {
	gorm.Model
	DeveloperName string
	Technologies  []Technology `gorm:"many2many:developer_technology"`
	Fun           string
}
type Technology struct {
	gorm.Model
	TechnologyName string
	Developers     []Developer `gorm:"many2many:developer_technology"`
}
type Employee struct {
	EmployeeName string
	CompanyID    int
	Company      Company
}
type Company struct {
	CompanyID   int
	CompanyName string
}

var cc1 = &CreditCard{
	Number:   "123",
	UserName: "Mark",
}
var cc2 = &CreditCard{
	Number:   "456",
	UserName: "Smith",
}
var user1 = &User{
	Name: "Smith",
}
var dev2 = Developer{DeveloperName: "Virat", Technologies: []Technology{{TechnologyName: "Python"}}}

func main() {
	databaseConnection()
	//hasOne()
	//Polymorph()
	//hasMany()
	//manyTomany()
	//GetAllDevs()
	//AssociationMode()
	//deleteMultiuser()
	//ReplaceAssociation()
	//CountAsso()
	//BathData()
	//preload()
	//nestedPreload()
	//deleteAsso()
	//multiJoin()
	//contextPrac()
	//contextTimeout()
	//SessionPrac()
	//createDevs()
	//findDev()
	//migrationPrac()
	//dropRef()
	//CreateDevUsingTx(db)
	//createView()
	//constraints()
	//preparestmtPrac()
	genericDB()
}

func databaseConnection() {
	dsn := "host=localhost user=postgres password=root dbname=associations port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  true,        // Disable color
		},
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		//DryRun: true,
		//PrepareStmt: true,
	})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("database connected successfully!!")
	}
	//db.AutoMigrate(&CreditCard{})
	//db.AutoMigrate(&User{})
	//db.AutoMigrate(&Dog{})
	//db.AutoMigrate(&Toy{})
	//db.AutoMigrate(&Book{})
	//db.AutoMigrate(&Reader{})

	//db.AutoMigrate(&Developer{})
	//db.AutoMigrate(&Technology{})
	//db.AutoMigrate(&Employee{})
	//db.AutoMigrate(&Company{})

}
func hasOne() {
	rows, _ := db.Table("users").Select("users.name,credit_cards.number").Joins("left join credit_cards on credit_cards.user_name=users.name ").Rows()
	for rows.Next() {
		fmt.Println(rows)

	}

}
func Polymorph() {
	db.Create(&Dog{Name: "dog1", Toy: Toy{Name: "toy1"}})

}
func hasMany() {
	//db.Create(&Reader{ReaderName: "kanishk", Books: []Book{{BookName: "history", BookReader: "kanishk"}, {BookName: "Golang", BookReader: "kanishk"}}})
	//db.Create(&Reader{ReaderName: "Shashwat", Books: []Book{{BookName: "Geography", BookReader: "Shashwat"}, {BookName: "Java", BookReader: "Shashwat"}}})
	rows, _ := db.Table("readers").Select("readers.reader_name,books.book_name").Joins("left join books on books.book_reader=readers.reader_name").Rows()
	for rows.Next() {
		fmt.Println(rows)
	}
}

func manyTomany() {
	//db.Create(&Developer{
	//	DeveloperName: "Steve",
	//	Technologies:  []Technology{{TechnologyName: "ios"}, {TechnologyName: "Android"}},
	//})
	//db.Create(&Technology{TechnologyName: "Golang", Developers: []Developer{{DeveloperName: "Shashwat"}, {DeveloperName: "Virat"}}})
	//dev1 := Developer{DeveloperName: "Bill", Technologies: []Technology{{TechnologyName: "C++"}}}

	db.Omit("Technologies").Create(&dev2)

}

func GetAllDevs() {
	r, _ := db.Table("developers").Select("developers.developer_name,technologies.technology_name").Joins("JOIN developer_technology on developers.id=developer_technology.developer_id").Joins("join technologies on technologies.id=developer_technology.technology_id").Rows()
	for r.Next() {
		fmt.Println(r)
	}
}
func AssociationMode() {
	//dev3 := Developer{DeveloperName: "Henry", Technologies: []Technology{{TechnologyName: "php"}}}
	//var dev Developer
	//db.Find(&dev)
	//fmt.Println(dev)
	//db.Find(&dev2)
	//fmt.Println(dev2.ID)
	//var techs []Technology
	//db.Model(&dev2).Where("developers.id =?", 2).Association("Technologies").Append(&Technology{TechnologyName: "C"})
	//db.Model(Developer{}).Where("developers.id=?", "3").Association("Technologies").Append(&Technology{TechnologyName: "php"})
	//c := db.Model(&dev).Association("Technologies").Count()
	//fmt.Println(c)

	var techs []Technology
	var dev Developer
	db.Find(&dev)
	db.Model(&dev).Association("Technologies").Find(&techs)
	fmt.Println(techs)
	//db.Create(&dev3)
	//
	//db.Find(&dev3, 11)
	//db.Model(&dev3).Association("Technologies").Append(&Technology{TechnologyName: "JavaScript"})

	//fmt.Println(dev2)
	//fmt.Println(technologies)
}
func deleteMultiuser() {
	var devs []Developer
	db.Unscoped().Delete(&devs, []int{6, 7, 8, 9, 10})
}

func ReplaceAssociation() {
	//var dev Developer
	//db.Find(&dev)
	//TechnologyC := Technology{
	//	TechnologyName: "C",
	//}
	//TechnologyDocker := Technology{
	//	TechnologyName: "Docker",
	//}
	var dev Developer
	var techs []Technology
	//db.Model(&dev).Association("Technologies").Replace([]Technology{TechnologyC, TechnologyDocker})
	//db.Model(&dev).Association("Technologies").Replace(Technology{TechnologyName: "Docker"}, Technology)
	db.Find(&dev)
	db.Model(&dev).Association("Technologies").Find(&techs)
	fmt.Println(techs)
}

func CountAsso() {
	var dev []Developer
	var books []Technology
	//db.First(&dev)
	//db.Find(&dev)
	//db.Where("developers.id", 3).Find(&dev)
	//fmt.Println(db.Model(&dev).Where("code IN (1)", "").Association("Languages").Count())
	//fmt.Println(db.Model(&dev).Where("developer_technology.developer_id = 3").Association("Technologies").Count())
	//fmt.Println(books, "\n\n\n\n", dev)
	//sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
	//	return tx.Model(&dev).Association("Technologies")
	//})
	//fmt.Println(sql)
	//db.Debug().Where("id=?", 1).Find(&dev)
	//db.Find(&dev)
	//fmt.Println(dev)
	//db.Model(&dev).Association("Technologies").Find(&books)
	//fmt.Println(books)
	db.Find(&dev)

	db.Model(&dev).Association("Technologies").Find(&books)
	fmt.Println(books)
}

func BathData() {
	var devs []Developer
	var techs []Technology
	db.Find(&devs)
	db.Model(&devs).Association("Technologies").Find(&techs)
	fmt.Println(techs)
}

func preload() {
	var devs []Developer
	db.Preload("Technologies", "technology_name=?", "Docker").Find(&devs)
	fmt.Println(devs)
}
func nestedPreload() {
	var devs []Developer
	db.Preload("Technologies.Developers").Find(&devs)
	fmt.Println(devs)

}

func deleteAsso() {
	var dev Developer
	db.Find(&dev, 11)
	db.Unscoped().Select("Technologies").Delete(&dev)

}

func multiJoin() {
	var dev Developer
	db.Joins("developers").Joins("developer_technology").Joins("technologies").Find(&dev, 1)
	fmt.Println(dev)
}
func contextPrac() {
	var dev Developer
	ctx := context.Background()
	db.WithContext(ctx).Find(&dev, 5)
	fmt.Println(dev)

}
func contextTimeout() {
	var devs []Developer
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Microsecond)
	defer cancel()
	db.WithContext(ctx).Find(&devs)
	fmt.Println(devs)
}

func SessionPrac() {
	var dev Developer
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&dev, 1).Statement
	fmt.Println(stmt.SQL.String())

}
func preparestmtPrac() {

	var dev Developer
	var devs []Developer
	tx := db.Session(&gorm.Session{PrepareStmt: true})
	tx.First(&dev, 1)
	tx.Find(&devs)

	stmtManager, _ := tx.ConnPool.(*gorm.PreparedStmtDB)
	//fmt.Println(stmtManager)
	//stmtManager.Close()
	//fmt.Println(stmtManager.PreparedSQL)
	fmt.Println(stmtManager.Stmts)
	//for sql, stmt := range stmtManager.Stmts {
	//	fmt.Println(sql)
	//	fmt.Println(stmt)
	//	fmt.Println(stmt.Close())
	//}

}
func createDevs() {
	dev := Developer{
		DeveloperName: "John",
		Technologies: []Technology{
			{TechnologyName: "React"},
		},
	}
	db.Create(&dev)
}
func findDev() {
	var dev Developer
	db.Find(&dev, 12)

}
func (dev *Developer) AfterCreate(tx *gorm.DB) (err error) {
	if dev.DeveloperName == "John" {
		fmt.Println("hooks are called")
	}
	return
}
func (dev *Developer) AfterFind(tx *gorm.DB) (err error) {
	if dev.DeveloperName == "John" {
		fmt.Println("heloo")
		dev.DeveloperName = "Johny"
	}
	return
}
func CreateDevUsingTx(db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}
	if err := tx.Create(&Developer{DeveloperName: "Smith", Technologies: []Technology{{TechnologyName: ".NET"}}}).Error; err != nil {
		tx.Rollback()
	}
	return tx.Commit().Error

}
func dropRef() {

	db.Migrator().DropColumn(&Developer{}, "fun")
}
func migrationPrac() {
	db.Create(&Developer{
		DeveloperName: "Steve",
		Technologies:  []Technology{{TechnologyName: "ios"}, {TechnologyName: "Android"}},
	})
	db.Create(&Technology{TechnologyName: "Golang", Developers: []Developer{{DeveloperName: "Shashwat"}, {DeveloperName: "Virat"}}})
	dev1 := Developer{DeveloperName: "Bill", Technologies: []Technology{{TechnologyName: "C++"}}}

	db.Create(&dev1)
	dev3 := Developer{DeveloperName: "Henry", Technologies: []Technology{{TechnologyName: "php"}}}
	db.Create(&dev3)

	//db.Migrator().AddColumn(&Developer{}, "Fun")
	//fmt.Println(db.Migrator().HasColumn(&Developer{}, "Fun"))
	//db.Migrator().DropTable(&Developer{}, "Fun")

}
func createView() {
	que := db.Model(Developer{}).Where("developer_name=?", "Bill")
	db.Migrator().CreateView("view_dev", gorm.ViewOption{Query: que})

}

func constraints() {
	//type dummyDev struct {
	//	DeveloperName string `gorm:"check:name_checker,	Developer_Name<>'jinzhu'"`
	//}
	//type dummyDev2 struct {
	//	DeveloperName string `gorm:"check:name_checker,	Developer_Name<>'jinzhu'"`
	//}
	//db.AutoMigrate(&dummyDev{})
	//db.Migrator().CreateConstraint(&dummyDev2{}, "name_checker")
	//db.AutoMigrate(&dummyDev2{})
	//
	////db.Migrator().DropConstraint(&Developer{}, "name_checker")
	//fmt.Println(db.Migrator().HasConstraint(&dummyDev2{}, "name_checker"))
	////dev := Developer{DeveloperName: "jinzhu"}
	////db.Create(&dev)
	//type Per struct {
	//	Name string `gorm:"check:name_checker,Name<>'jinzhu'"`
	//}
	//db.AutoMigrate(&Per{})
	//type Emp struct {
	//	Name string
	//}
	//
	//db.Migrator().CreateConstraint(&Emp{}, "name_checker")
	//db.AutoMigrate(&Emp{})
	//db.Create(Emp{Name: "jinzhu"})
	//fmt.Println(db.Migrator().HasConstraint(&Emp{}, "name_checker"))

	type Engineer struct {
		gorm.Model
		EngineerName string
		CarID        int
	}

	type Car struct {
		gorm.Model

		Engineers []Engineer
	}
	//db.AutoMigrate(&Engineer{})
	//db.AutoMigrate(&Car{})
	db.Migrator().CreateConstraint(&Car{}, "Engineers")
	////db.Migrator().CreateConstraint(&Car{}, "fk_cars_engineers")

}
func genericDB() {
	sqlDB, _ := db.DB()
	fmt.Println(sqlDB.Ping())
	fmt.Println(sqlDB.Stats())
}
