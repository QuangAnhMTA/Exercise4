package main

import (
	"exercise4/database"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/xid"
)

var db *database.Db = new(database.Db)

func bai1() {
	var err error
	err = db.ConnectBD()
	err = db.CreateTable()
	u := database.User{
		Id:      xid.New().String(),
		Name:    "Nguyen Van A",
		Birth:   time.Now().UnixNano(),
		Created: time.Now().UnixNano(),
		Updated: time.Now().UnixNano(),
		Job:     "Hs",
	}
	lst := make([]database.User, 0)
	err, lst = db.ListUser()

	if err != nil {
		panic(err)
	}
	for i, user := range lst {
		log.Println(i, user)
	}
	err = db.InsertUser(&u)
	if err != nil {
		panic(err)
	}
	user := new(database.User)
	err, user = db.FindUser("bs7t60pcti8quec45td0")
	if err != nil {
		panic(err)
	}
	log.Println(user)
	err = db.UpdateUser("bs7t60pcti8quec45td0")
	if err != nil {
		panic(err)
	}
}
func bai2() {
	var err error
	err = db.ConnectBD()
	err = db.UpdateUser_Birth("bs80ea1cti8ra6i72ucg")
	if err != nil {
		panic(err)
	}
}
func bai3() {
	// var err error
	// _, err = db.ConnectBD()
	// for i := 0; i < 100; i++ {
	// 	u := database.User{}
	// 	u.Id = xid.New().String()
	// 	u.Name = "Tran Van C"
	// 	db.Engine.Insert(&u)
	// }
	// if err != nil {
	// 	panic(err)
	// }

}
func b3_2() {
	var err error
	err = db.ConnectBD()
	UserData := make(chan *database.User)
	var wg sync.WaitGroup
	rows, err := db.Engine.Rows(database.User{})
	//defer rows.Close()
	bean := new(database.User)
	for rows.Next() {
		wg.Add(1)
		err = rows.Scan(bean)
		UserData <- bean
	}
	numberWorker := 2
	done := make([]chan bool, numberWorker)
	for i := 0; i < numberWorker; i++ {
		done[i] = make(chan bool)
		go PrintUser(UserData, &wg, done[i])
	}

	if err != nil {
		panic(err)
	}
	wg.Wait()
	for _, x := range done {
		x <- true
	}
	time.Sleep(10 * time.Second)
}
func PrintUser(chanUser chan *database.User, wg *sync.WaitGroup, done chan bool) {
	for {
		select {
		case user := <-chanUser:
			log.Printf("Name %v : %v xong!\n", user.Id, user.Name)
			wg.Done()
		case <-done:
			log.Println("Da doc het du lieu")
			break
		}
	}
}
func main() {
	// bai1()
	//bai2()
	//bai3()
	b3_2()
	defer db.Engine.Close()
}
