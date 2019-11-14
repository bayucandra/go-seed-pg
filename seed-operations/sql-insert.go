package seed_operations

import (
	"bufio"
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/bayucandra/go-seed/db"
	"log"
	"os"
	"regexp"
	"strings"
)

func SeedAll( filePaths []string ) {
	for _, val := range filePaths {
		_ = csvParse(val)
	}
}

func csvParse( filePath string ) (err error) {

	r := regexp.MustCompile(`.*/(.*)`)
	fileName := r.FindStringSubmatch(filePath)[1]

	r = regexp.MustCompile(`^\d{0,3}\.\d{0,3}_(.*)\.[a-z A-Z]{0,}`)
	tableName := r.FindStringSubmatch(fileName)[1]

	f, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.New( fmt.Sprintf("open file error: %v", err) )
	}
	defer f.Close()

	sc := bufio.NewScanner(f)
	idx := 0
	var fieldNames []string

	for sc.Scan() {
		txt := sc.Text()  // GET the line string
		if idx == 0 {
			csvReader := csv.NewReader(strings.NewReader(txt))
			fieldNames, err = csvReader.Read()
			if err != nil {
				return
			}
		} else {
			csvReader := csv.NewReader(strings.NewReader(txt))
			var record []string
			record, err = csvReader.Read()
			if err != nil {
				return
			}
			err = sqlInsert(fieldNames, record, tableName)
			if err != nil {
				log.Fatal(err)
			}
		}
		idx++
	}
	if err = sc.Err(); err != nil {
		return
	}
	return
}

func sqlInsert( cols []string, record []string, table string ) (err error){

	colsStr := "( "

	for idx, val := range cols {
		if idx < len(cols) -1 {
			colsStr += fmt.Sprintf("%s, ", val)
		} else {
			colsStr += fmt.Sprintf("%s )", val)
		}
	}

	valStr := "( "
	for idx, val := range record {

		if val != "null" {
			val = fmt.Sprintf("'%v'", val)
		}

		if idx < len(cols) -1 {
			valStr += fmt.Sprintf("%v, ", val)
		} else {
			valStr += fmt.Sprintf("%v )", val)
		}
	}

	qry := fmt.Sprintf(
		`INSERT INTO "%s"."%s" %s VALUES %s`, os.Getenv("PG_SCHEMA"),
		table, colsStr, valStr)

	res, err := db.DBConn.Exec( qry )

	if res == nil {
		return errors.New(fmt.Sprintf("not inserted correctly, result is nil: %v", err))
	}

	rowInserted, _ := res.RowsAffected()
	if rowInserted != 1 {
		err = errors.New("not inserted correctly, affected rows is not 1")
	}

	return
}

func RowParse(rows *sql.Rows) (records map[int]map[string]interface{}, count int) {
	columns, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	columnCount := len(columns)

	records = make(map[int]map[string]interface{})

	for rows.Next() {

		values := make([]interface{}, columnCount)
		valuePtrs := make([]interface{}, columnCount)
		for i := range columns {
			valuePtrs[i] = &values[i]
		}

		err := rows.Scan(valuePtrs...)
		if err != nil {
			panic(err)
		}
		row := make(map[string]interface{})

		for i:= range columns {
			row[columns[i]] = values[i]
		}
		records[count] = row
		count++
	}

	return

}
