package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func connectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=learnsql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	return db, err
}

type User struct {
	gorm.Model
	Name string
}

func main() {

	log.Println("welcome to gorm!")
	db, err := connectDB()
	if err != nil {
		panic(err)
	}
	fmt.Println("success connecting to db")
	var users []User
	var user User
	var res *gorm.DB

	// SELECT * FROM users u WHERE u.id = 1 AND u.deleted_at IS NULL ORDER BY u.id LIMIT 1
	res = db.First(&user, 1) //CAREFUL HERE: USE ONLY IF the struct has a DeletedAt field
	if res.Error != nil {
		log.Println(err) // it will give an error here. thats because First, Find, Last all include "WHERE deleted_at IS NOT NULL" if the go struct has a DeletedAt field (or embeds a gorm.Model), and 2 params are passed. Solution: Either include a deleted_at in your types and in your database, or use the option below
	}

	//SELECT * FROM users u WHERE u.id = 1 ORDER BY u.id LIMIT 1
	user.ID = 0 //if the user.ID is not 0 when FIND/FIND/LAST is ran, then automatically a WHERE id = user.ID is added to the query
	db.First(&user, "id = ?", 1)

	db.Where("name = ? OR 1 = 1", "abc").Find(&users)

	db.Where("email = ? AND name = ?", "abc@gm.com", "abc")

	db.Where("email = ?", "abc@gm.com").Or("name = ?", "abc").Find(&users)

	//limit, offset, order by

	db.Where("updatedAt >= ?", time.Now()).Limit(5).Offset(10).Order("name DESC").Find(&users)

	//grouping & having
	//option 1: store the results directly in an []groupingResultA

	/* var resultsA []groupingResultA
	r := db.Model(User{}).Select("age, count(*) as reps").Group("age").Find(&resultsA) //select age, count(*) as reps from users group by age
	fmt.Println(r.Error) */

	//option 2: process them one by one (use only when there may be TOO many groupingResults to store in a single array)

	type groupingResultA struct {
		Age  int
		Reps int
	}
	rows, err := db.Table("users").Select("age, count(*) as reps").Group("age").Having("reps > ?", 10).Rows()
	if err != nil {
		//handle err
		log.Println("err in option 2: ", err)
	}

	defer rows.Close()
	for rows.Next() {
		var currGroupingResult groupingResultA
		if err := rows.Scan(&currGroupingResult.Age, &currGroupingResult.Reps); err != nil {
			//handle scanning err
			log.Println("err scanning into grouping result: ", err)
		}
	}

	//joins
	type resultB struct {
		Name         string
		Age          int
		Company_name string
	}
	var resultsB []resultB
	db.Model(User{}).Select("name, age, company_name").Joins("JOIN companies ON users.company_id = companies.company_id").Find(&resultsB)

	//transactions

	err = db.Transaction(func(tx *gorm.DB) error {
		//wrapping user creation in a transaction.
		u := User{Name: "newUser"}
		if res := tx.Create(&u); res.Error != nil {
			return res.Error
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	//common stuff with res
	_ = res.RowsAffected
	_ = res.Error

}
