package api

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo"
)

type SqlHandler struct {
	Conn *sql.DB
}
type sumData struct {
	DATE        string  `json:"date_"`
	TIME        string  `json:"time_"`
	SIGNAL      int     `json:"signal"`
	DIFF_TIME   float32 `json:"diff_time"`
	SIGNAL_TYPE string  `json:"signal_type"`
}
type recordTable struct {
	ID     int    `json:"id"`
	DATE   string `json:"date"`
	TIME   string `json:"time_"`
	MCNAME string `json:"mc_name"`
	SIGNAL int    `json:"signal_type"`
	RESULT string `json:"result"`
	NGCODE string `json:"ng_code"`
}
type summaryRecordTable struct {
	Avg_ct  float32 `json:"avg_ct"`
	Avg_wt  float32 `json:"avg_wt"`
	S_total float32 `json:"s_total"`
	S_ct    float32 `json:"s_ct"`
	S_wt    float32 `json:"s_wt"`
	S_ngct  float32 `json:"s_ngct"`
	S_loss  float32 `json:"s_loss"`
	S_na    float32 `json:"s_na"`
	P_ct    float32 `json:"p_ct"`
	P_wt    float32 `json:"p_wt"`
	P_ngct  float32 `json:"p_ngct"`
	P_loss  float32 `json:"p_loss"`
	P_na    float32 `json:"p_na"`
}

var sqlCon = new(SqlHandler)

func ConnectDB() {
	condb, errdb := sql.Open("mssql", "server=localhost;user id=sa;password=EPPE7348;database=OperationRecord")
	if errdb != nil {
		fmt.Println("Error open db:", errdb.Error())
	}
	//defer condb.Close()
	fmt.Println("Connect DB OK")
	sqlCon.Conn = condb
}

func GetDistinctMCname(c echo.Context) (err error) {
	var resultMCname []string
	sqlTxt := fmt.Sprintf("SELECT DISTINCT mc_name FROM [OperationRecord].[dbo].[RecordTable] ORDER BY mc_name")
	fmt.Println(sqlTxt)
	rows, err := sqlCon.Conn.Query(sqlTxt)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data string
		errdat := rows.Scan(&data)

		//errdat := rows.Scan(&data.ID, &data.DATE, &data.TIME, &data.MCNAME, &data.SIGNAL, &data.RESULT, &data.NGCODE)
		if errdat != nil {
			log.Fatal()
		}
		/*d, _ := time.Parse(time.RFC3339, data.DATE)
		data.DATE = d.Format("2006-01-02")
		t, _ := time.Parse(time.RFC3339, data.TIME)
		data.TIME = t.Format("15:04:05.000")*/
		resultMCname = append(resultMCname, data)
	}
	u := resultMCname

	fmt.Println("GetDistinctMCname complete")
	return c.JSON(http.StatusOK, u)
}

func GetRecordData(c echo.Context) (err error) {
	var resultSumData []sumData

	detail := c.Param("detail")
	fmt.Println(detail)
	detailArr := strings.Split(detail, "&")
	mcname := detailArr[0]
	st_date := detailArr[1]
	shift := detailArr[2]
	st_time := detailArr[3]
	en_time := detailArr[4]

	sqlTxt := fmt.Sprintf("set nocount on; EXEC [dbo].[SummaryOperationRecord] '%s','%s','%s','%s','%s'",
		mcname, st_date, shift, st_time, en_time)
	fmt.Println(sqlTxt)
	rows, err := sqlCon.Conn.Query(sqlTxt)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data sumData
		errdat := rows.Scan(&data.DATE, &data.TIME, &data.SIGNAL, &data.DIFF_TIME, &data.SIGNAL_TYPE)

		//errdat := rows.Scan(&data.ID, &data.DATE, &data.TIME, &data.MCNAME, &data.SIGNAL, &data.RESULT, &data.NGCODE)
		if errdat != nil {
			log.Fatal()
		}
		d, _ := time.Parse(time.RFC3339, data.DATE)
		data.DATE = d.Format("2006-01-02")
		t, _ := time.Parse(time.RFC3339, data.TIME)
		data.TIME = t.Format("15:04:05.000")
		resultSumData = append(resultSumData, data)
	}
	u := resultSumData

	fmt.Println("GetRecordData complete")
	return c.JSON(http.StatusOK, u)
}

func GetSummaryData(c echo.Context) (err error) {
	var resultSumData []summaryRecordTable

	detail := c.Param("detail")
	fmt.Println(detail)
	detailArr := strings.Split(detail, "&")
	mcname := detailArr[0]
	st_date := detailArr[1]
	shift := detailArr[2]
	st_time := detailArr[3]
	en_time := detailArr[4]

	sqlTxt := fmt.Sprintf("set nocount on; EXEC [dbo].[SummaryOperationRecord] '%s','%s','%s','%s','%s'",
		mcname, st_date, shift, st_time, en_time)
	fmt.Println(sqlTxt)
	rows, err := sqlCon.Conn.Query(sqlTxt)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var data summaryRecordTable
		errdat := rows.Scan(
			&data.Avg_ct,
			&data.Avg_wt,
			&data.S_total,
			&data.S_ct,
			&data.S_wt,
			&data.S_ngct,
			&data.S_loss,
			&data.S_na,
			&data.P_ct,
			&data.P_wt,
			&data.P_ngct,
			&data.P_loss,
			&data.P_na)

		//errdat := rows.Scan(&data.ID, &data.DATE, &data.TIME, &data.MCNAME, &data.SIGNAL, &data.RESULT, &data.NGCODE)
		if errdat != nil {
			log.Fatal()
		}
		/*d, _ := time.Parse(time.RFC3339, data.DATE)
		data.DATE = d.Format("2006-01-02")
		t, _ := time.Parse(time.RFC3339, data.TIME)
		data.TIME = t.Format("15:04:05.000")*/
		resultSumData = append(resultSumData, data)
	}
	u := resultSumData

	fmt.Println("GetSummaryData complete")
	return c.JSON(http.StatusOK, u)
}
