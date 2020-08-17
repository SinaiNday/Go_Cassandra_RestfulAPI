package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Restful API using Go and Cassandra!")
}

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var Newstudent Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "wrong data")
	}
	json.Unmarshal(reqBody, &Newstudent)
	if err := Session.Query("INSERT INTO students(id, firstname, lastname, age) VALUES(?, ?, ?, ?)",
		Newstudent.ID, Newstudent.Firstname, Newstudent.Lastname, Newstudent.Age).Exec(); err != nil {
		fmt.Println("Error while inserting")
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusCreated)
	Conv, _ := json.MarshalIndent(Newstudent, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var students []Student
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM students").Iter()
	for iter.MapScan(m) {
		students = append(students, Student{
			ID:        m["id"].(int),
			Firstname: m["firstname"].(string),
			Lastname:  m["lastname"].(string),
			Age:       m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(students, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}
func GetOneStudent(w http.ResponseWriter, r *http.Request) {
	StudentID := mux.Vars(r)["id"]
	var students []Student
	m := map[string]interface{}{}

	iter := Session.Query("SELECT * FROM students WHERE id=?", StudentID).Iter()
	for iter.MapScan(m) {
		students = append(students, Student{
			ID:        m["id"].(int),
			Firstname: m["firstname"].(string),
			Lastname:  m["lastname"].(string),
			Age:       m["age"].(int),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(students, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))

}

func CountAllStudents(w http.ResponseWriter, r *http.Request) {

	var Count string
	err := Session.Query("SELECT count(*) FROM students").Scan(&Count)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s ", Count)

}
func DeleteOneStudent(w http.ResponseWriter, r *http.Request) {
	StudentID := mux.Vars(r)["id"]
	if err := Session.Query("DELETE FROM students WHERE id = ?", StudentID).Exec(); err != nil {
		fmt.Println("Error while deleting")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted successfully the student num %s ", StudentID)
}

func DeleteAllStudents(w http.ResponseWriter, r *http.Request) {

	if err := Session.Query("TRUNCATE students").Exec(); err != nil {
		fmt.Println("Error while deleting all students")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "deleted all successfully")

}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	StudentID := mux.Vars(r)["id"]
	var UpdateStudent Student
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data properly")
	}
	json.Unmarshal(reqBody, &UpdateStudent)
	if err := Session.Query("UPDATE students SET firstname = ?, lastname = ?, age = ? WHERE id = ?",
		UpdateStudent.Firstname, UpdateStudent.Lastname, UpdateStudent.Age, StudentID).Exec(); err != nil {
		fmt.Println("Error while updating")
		fmt.Println(err)
	}
	fmt.Fprintf(w, "updated successfully")

}
