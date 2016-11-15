package saving

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type SavingData struct {
	Time string `bson:"time"`
}

type Balance struct {
	Time    string `bson:"time"`
	Balance string `bson:"balance"`
}

// Saving
// 貯金結果をDBに保存する
func SaveSavingCount(time string) {
	session, _ := mgo.Dial("mongodb://localhost/notice_saving")
	defer session.Close()
	col := session.DB("notice_saving").C("saving_count")
	savingCount := &SavingData{
		Time: time,
	}

	err := col.Insert(savingCount)
	if err != nil {
		fmt.Printf("%+v \n", err)
	}

	fmt.Println(GetNowBalance())
	UpdateBalance(time)
	fmt.Println("save success")
}

// GetNowBalance
// 現在の貯金残高を返す
func GetNowBalance() *Balance {
	session, _ := mgo.Dial("mongodb://localhost/notice_saving")
	defer session.Close()
	col := session.DB("notice_saving").C("balance")
	b := new(Balance)
	query := col.Find(bson.M{}).Sort("-$natural")
	query.One(&b)

	return b
}

// UpdateBalance
// 貯金残高を更新する
func UpdateBalance(time string) {
	session, _ := mgo.Dial("mongodb://localhost/notice_saving")
	defer session.Close()
	col := session.DB("notice_saving").C("balance")
	nowBalanceData := GetNowBalance()
	nowBalance, _ := strconv.Atoi(nowBalanceData.Balance)
	fmt.Println(nowBalanceData.Time)

	if nowBalance == 0 || len(nowBalanceData.Time) == 0 {
		initBalance := &Balance{
			Time:    time,
			Balance: "500",
		}

		err := col.Insert(initBalance)
		if err != nil {
			fmt.Printf("%+v \n", err)
		}
	}

	selector := bson.M{"time": nowBalanceData.Time}

	newBalanceStr := strconv.Itoa(nowBalance + 500)
	newBalance := &Balance{
		Time:    time,
		Balance: newBalanceStr,
	}

	err := col.Update(selector, newBalance)

	if err != nil {
		if mgo.IsDup(err) {
			fmt.Printf("Duplicate key error \n")
		}
		if v, ok := err.(*mgo.LastError); ok {
			fmt.Printf("Code:%d N:%d UpdatedExisting:%t WTimeout:%t Waited:%d \n", v.Code, v.N, v.UpdatedExisting, v.WTimeout, v.Waited)
		} else {
			fmt.Printf("%+v \n", err)
		}
	}
}

// ResetBalance
// 貯金残高をリセットする
func ResetBalance(time string) {
	session, _ := mgo.Dial("mongodb://localhost/notice_saving")
	defer session.Close()
	col := session.DB("notice_saving").C("balance")

	balance := &Balance{
		Time:    time,
		Balance: strconv.Itoa(0),
	}
	err := col.Insert(balance)

	if err != nil {
		if mgo.IsDup(err) {
			fmt.Printf("Duplicate key error \n")
		}
		if v, ok := err.(*mgo.LastError); ok {
			fmt.Printf("Code:%d N:%d UpdatedExisting:%t WTimeout:%t Waited:%d \n", v.Code, v.N, v.UpdatedExisting, v.WTimeout, v.Waited)
		} else {
			fmt.Printf("%+v \n", err)
		}
	}
}
