package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" //import for side effects
	"log"
	"math/rand"
)

func main() {
	myDatabase := OpenDataBase("./Demo.db")
	defer myDatabase.Close()
	create_tables(myDatabase)
	add_sample_data(myDatabase)
}
func OpenDataBase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal(err)
	}
	return database
}

func create_tables(database *sql.DB) {
	createStatement1 := "CREATE TABLE IF NOT EXISTS students(    " +
		"banner_id INTEGER PRIMARY KEY," +
		"first_name TEXT NOT NULL," +
		"last_name TEXT NOT NULL," +
		"gpa REAL DEFAULT 0," +
		"credits INTEGER DEFAULT 0);"
	database.Exec(createStatement1)
	courseCreateStatement := "CREATE TABLE IF NOT EXISTS course(   " +
		" course_prefix TEXT NOT NULL,  " +
		"  course_number INTEGER NOT NULL,  " +
		"  cap INTEGER DEFAULT 20,    description TEXT,   " +
		" PRIMARY KEY(course_prefix, course_number)"
	database.Exec(courseCreateStatement)
	regcourseCreateStatement := "CREATE TABLE IF NOT EXISTS class_list(" +
		"registration_id INTEGER PRIMARY KEY," +
		"course_prefix TEXT NOT NULL," +
		"course_number INTEGER NOT NULL," +
		"banner_id INTEGER NOT NULL," +
		"registration_date TEXT," +
		"FOREIGN KEY (banner_id) REFERENCES student (banner_id)" +
		"ON DELETE CASCADE ON UPDATE NO ACTION," +
		"FOREIGN KEY (course_prefix, course_number) REFERENCES courses (course_prefix, course_number)" +
		"ON DELETE CASCADE ON UPDATE NO ACTION"
	database.Exec(regcourseCreateStatement)
}

func add_sample_data(database *sql.DB) {
	sampleNames := map[string]string{"John": "Santore", "Enping": "Li", "Margaret": "Black",
		"Seikyung": "Jung", "Haleh": "Khojasteh", "Abdul": "Sattar", "Paul": "Kim", "Laura": "Gross"}
	statement := "INSERT INTO STUDENTS (banner_id, first_name, last_name, gpa, credits)" +
		"  VALUES (?, ?, ?, ?, ?)"
	count := 1001
	for firstName, lastName := range sampleNames {
		randGPA := rand.Float32() + float32(rand.Intn(4))
		randCredits := rand.Intn(120)
		prepped_statement, err := database.Prepare(statement)
		if err != nil {
			//cowardly bail out since this is academia
			log.Fatal(err)
		}
		prepped_statement.Exec(count, firstName, lastName, randGPA, randCredits)
		count += 1
	}
}
