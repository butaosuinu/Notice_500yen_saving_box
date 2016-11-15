package saving

import (
	"strconv"
)

// SaveOpenBox
// 貯金箱を開いた際の処理
func SaveOpenBox(time string) {
	nowBalanceData := GetNowBalance()
	nowBalance, _ := strconv.Atoi(nowBalanceData.Balance)
	InsertOpenBoxRecord(time, nowBalance)
	ResetBalance(time)
}
