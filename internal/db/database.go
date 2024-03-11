package db

import (
	"log"

	"govtech-onecv/internal/util"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Init() {
	pg_user, err := util.GetEnv("POSTGRES_USER")
	if err != nil {
		log.Fatalf(err.Error())
	}

	pg_password, err := util.GetEnv("POSTGRES_PASSWORD")
	if err != nil {
		log.Fatalf(err.Error())
	}

	pg_db, err := util.GetEnv("POSTGRES_DB")
	if err != nil {
		log.Fatalf(err.Error())
	}

	pg_port, err := util.GetEnv("POSTGRES_PORT")
	if err != nil {
		log.Fatalf(err.Error())
	}

	dsn := "host=db" + " user=" + pg_user + " password=" + pg_password + " dbname=" + pg_db + " port=" + pg_port + " sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	d.DB = db
}

func (d *Database) AutoMigrate() {
	d.DB.AutoMigrate(&TeacherSchema{})
	d.DB.AutoMigrate(&StudentSchema{})

}

var seedStudents = []StudentSchema{
	{Student: "studentmary@gmail.com", Suspend: false, Teachers: pq.StringArray{"teacherken@gmail.com"}},
	{Student: "studentjon@gmail.com", Suspend: false, Teachers: pq.StringArray{"teacherken@gmail.com", "teacherjoe@gmail.com"}},
	{Student: "studentbob@gmail.com", Suspend: false, Teachers: pq.StringArray{"teacherjoe@gmail.com"}},
	{Student: "studentmiche@gmail.com", Suspend: false, Teachers: pq.StringArray{"teacherjoe@gmail.com"}},
}

var seedTeachers = []TeacherSchema{
	{Teacher: "teacherken@gmail.com", Students: pq.StringArray{"studentmary@gmail.com", "studentjon@gmail.com"}},
	{Teacher: "teacherjoe@gmail.com", Students: pq.StringArray{"studentjon@gmail.com", "studentbob@gmail.com", "studentmiche@gmail.com"}},
}

func (d *Database) Seed() {
	for _, student := range seedStudents {
		d.DB.Create(&student)
	}

	for _, teacher := range seedTeachers {
		d.DB.Create(&teacher)
	}

	student1 := StudentSchema{}
	d.DB.Find(&student1)
	log.Println(student1)
}
