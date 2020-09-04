package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/remeh/sizedwaitgroup"
)

var db *sql.DB

// Mssql should conform to structure of database.yaml
type Mssql struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}

//SQLs should conform to structure of database.yaml
type SQLs struct {
	SQL []string
}

//YamlConfig should conform to structure of database.yaml
type YamlConfig struct {
	Server      string
	Port        int
	User        string
	Password    string
	Database    string
	Iterations  int
	Concurrency int
	SQLs        []string
}

func executeSQL(db *sql.DB, tsql string) (int, time.Time, time.Duration, error) {

	ctx := context.Background()
	var err error
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	if db == nil {
		err = errors.New("db is nil")
		return 0, time.Now(), time.Since(time.Now()), err
	}

	starttime := time.Now()

	rows, err := db.QueryContext(ctx, tsql)
	endtime := time.Since(starttime)
	if err != nil {
		fmt.Println(err)
		return 0, time.Now(), time.Since(time.Now()), err
	}
	defer rows.Close()

	var count int
	for rows.Next() {
		count = count + 1
	}
	return count, starttime, endtime, nil
}

func readYAML(filename string) YamlConfig {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Error reading YAML file", err)
	}
	var yamlConfig YamlConfig
	err = yaml.Unmarshal(yamlFile, &yamlConfig)
	if err != nil {
		log.Fatal("Error parsing YAML file: ", err)
	}
	return yamlConfig
}

func main() {
	yamlMap := readYAML("database.yaml")
	//log.Printf("%#+v", yamlMap)
	//log.Printf("%v", yamlMap.Database)

	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		yamlMap.Server, yamlMap.User, yamlMap.Password, yamlMap.Port, yamlMap.Database)
	//log.Println(connString)

	db, err := sql.Open("mssql", connString)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Connected to database\n")

	CsvDelimiter := "~"
	fmt.Println("row" + CsvDelimiter + "startTime" + CsvDelimiter + "durationMicroseconds" + CsvDelimiter + "recordsReturned" + CsvDelimiter + "sql")

	swg := sizedwaitgroup.New(yamlMap.Concurrency)

	for iteration := 0; iteration < yamlMap.Iterations; iteration++ {
		for i, sql := range yamlMap.SQLs {
			swg.Add()
			go func(i int, sql string) {
				defer swg.Done()
				records, starttime, duration, err := executeSQL(db, sql)
				if err != nil {
					log.Println(err)
					log.Fatal("Error executing SQL statement", err.Error())
				}
				//fmt.Printf("SQL statement '%s' executed correctly in %d microseconds, and returned %d records\n", sql, duration.Microseconds(), records)
				fmt.Printf("%d%s%v%s%d%s%d%s%v\n", i, CsvDelimiter, starttime, CsvDelimiter, duration.Microseconds(), CsvDelimiter, records, CsvDelimiter, sql)
			}(i, sql)
		}
	}
	swg.Wait()
}
