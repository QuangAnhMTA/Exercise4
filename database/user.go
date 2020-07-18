package database

import (
	"log"
	"time"
)

// user(id string, name string, birth int64, created int64, updated_at int64)

type User struct {
	Id      string `json: "id"`
	Name    string `json: "name"`
	Birth   int64  `json: "birth"`
	Created int64  `json; "create"`
	Updated int64  `json:"updated"`
	Job     string `json:"job"`
}

func (db *Db) InsertUser(u *User) error {
	var err error
	nb, err := db.Engine.Insert(u)
	if err != nil {
		return err
	}
	log.Println(nb, "user dc them vao bang user")
	var p = Point{UserId: u.Id, Points: 10}
	err = db.InsertPoint(&p)
	if err != nil {
		return err
	}
	return nil
}
func (db *Db) ListUser() (error, []User) {
	users := make([]User, 0)
	err := db.Engine.Find(&users)
	if err != nil {
		return err, nil
	}

	return nil, users
}
func (db *Db) FindUser(id string) (error, *User) {
	var u = new(User)
	_, err := db.Engine.Table(User{}).Where("Id = ?", id).Get(u)
	if err != nil {
		return err, nil
	}
	return nil, u
}
func (db *Db) UpdateUser(u User, Condition *User) error {
	_, err := db.Engine.Update(u, Condition)
	if err != nil {
		return err
	}
	return nil

}

func (db *Db) UpdateUser_Birth(id string) error {

	session := db.Engine.NewSession()
	defer session.Close()
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}
	var u = new(User)
	var p = new(Point)
	_, err := session.Table(User{}).Where("Id = ?", id).Get(u)
	if err != nil {
		session.Rollback()
	}
	u.Birth = time.Now().UnixNano()
	if _, err := session.Table(u).Where("Id = ?", id).Update(u); err != nil {
		session.Rollback()
	}

	_, err = session.Table(p).Where("Id = ?", id).Get(p)
	p.Points = p.Points + 10
	if _, err := session.Cols("Points").Where("User_id = ?", id).Update(&Point{Points: p.Points + 10}); err != nil {
		session.Rollback()
	}
	u.Name = u.Name + "updated"
	if _, err := session.Table(u).Where("Id = ?", id).Update(u); err != nil {
		session.Rollback()
	}
	return session.Commit()
}
