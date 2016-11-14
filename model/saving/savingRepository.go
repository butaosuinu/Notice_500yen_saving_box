package saving

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type SavingData struct {
	ID   bson.ObjectId `bson:"_id"`
	Time string        `bson:"time"`
}

// Saving
// 貯金結果をDBに保存する
func SaveSavingCount(time string) {
	session, _ := mgo.Dial("mongodb://localhost/notice_saving")
	defer session.Close()
	db := session.DB("notice_saving")
	savingCount := &SavingData{
		Time: time,
	}
	col := db.C("saving_count")
	col.Insert(savingCount)
	fmt.Println("save success")
}
