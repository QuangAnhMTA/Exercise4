package database

import "log"

//  point(user_id string, points int64, max_points int64)
type Point struct {
	UserId    string `json:"user_id"`
	Points    int64  `json:"Points"`
	MaxPoints int64  `json:"max_points"`
}

func (db *Db) InsertPoint(p *Point) error {
	var err error
	_, err = db.Engine.Insert(p)
	if err != nil {
		return err
	}
	log.Println("Them point thanh cong")
	return nil
}
func (db *Db) UpdatePoint(p *Point, condition *Point) error {
	_, err := db.Engine.Update(p, condition)
	if err != nil {
		return err
	}
	return nil

}
