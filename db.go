package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
var showlist []string
var count int
var removelist = &count

func main() {
	t := time.Now()
	db, err := sql.Open("sqlite3", "./mds123.db")
	checkErr(err)

	//插入数据
	/* 	stmt, err := db.Prepare("INSERT INTO users(name) values(?)")
	   	checkErr(err)

	   	res, err := stmt.Exec("魔鬼蘑菇人")
	   	checkErr(err)

	   	id, err := res.LastInsertId()
	   	checkErr(err)

	   	fmt.Println(id) */
	//更新数据
	/* 	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	   	checkErr(err)

	   	res, err = stmt.Exec("astaxieupdate", id)
	   	checkErr(err)

	   	affect, err := res.RowsAffected()
	   	checkErr(err)

	   	fmt.Println(affect)
	*/
	//查询数据
	rows, err := db.Query("SELECT * FROM users")
	checkErr(err)

	err2 := db.QueryRow("select count(*) FROM users ").Scan(&count)
	if err != nil {
		fmt.Println(err2)
	}
	for rows.Next() {
		var num int
		var username string
		err = rows.Scan(&num, &username)
		checkErr(err)
		fmt.Printf("%d"+":", num)
		fmt.Println(username)
		showlist = append(showlist, username)
	}

	//删除数据
	/* 	stmt, err = db.Prepare("delete from userinfo where uid=?")
	   	checkErr(err)

	   	res, err = stmt.Exec(id)
	   	checkErr(err)

	   	affect, err = res.RowsAffected()
	   	checkErr(err)

	   	fmt.Println(affect)
	*/
	db.Close()
	fmt.Printf("此次共有%d員戰士參加戰爭\n", count)
	/* 	for index := 0; index < 10; index++ {
		ra := GenerateRangeNum(1, count+1)
		fmt.Println("您選擇的戰士是:", showlist[ra-1])
	} */

	for {
		if count > 0 {
			ra := GenerateRangeNum(1, 2+1)
			ralist := GenerateRangeNum(1, count+1)
			fmt.Printf("此次共有%d員戰士參加戰爭，要去嗎?%d\n", count, ra)
			if ra == 1 {
				fmt.Printf("決定就是你了:%v，出擊吧!!\n", showlist[ralist-1])
				break
			} else {
				fmt.Printf("再練練吧:%v，孩子!!\n", showlist[ralist-1])
				removeslice(ralist - 1)
			}
		} else {
			fmt.Println("你真的太挑了(指)")
		}
	}

	elapsed := time.Since(t).String() //Program execution time
	fmt.Println("execution time: " + elapsed)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func GenerateRangeNum(min, max int) int {
	randNum := seededRand.Intn(max - min)
	randNum = (randNum + min)
	// fmt.Printf("rand is %v\n", randNum)
	return randNum
}

func removeslice(i int) {
	showlist = append(showlist[:i], showlist[i+1:]...)
	*removelist = len(showlist)
}
