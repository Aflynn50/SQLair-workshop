package main

import (
	"database/sql"

	"github.com/canonical/sqlair"

	_ "github.com/mattn/go-sqlite3"
)

type Location struct {
	ID   int    `db:"room_id"`
	Name string `db:"name"`
	Team string `db:"team"`
}

type Person struct {
	Name string `db:"name"`
	ID   int    `db:"id"`
	Team string `db:"team"`
}

func main() {
	sqldb, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	db := sqlair.NewDB(sqldb)
	create := sqlair.MustPrepare(`
	CREATE TABLE person (
		name text,
		id integer,
		team text
	);
	CREATE TABLE location (
		room_id integer,
		name text,
		team text
	)`)
	err = db.Query(nil, create).Run()
	if err != nil {
		panic(err)
	}

	insertPerson := sqlair.MustPrepare("INSERT INTO person (name, id, team) VALUES ($Person.name, $Person.id, $Person.team);", Person{})

	var al = Person{"Alasatir", 1, "pals"}
	var ed = Person{"Ed", 2, "pals"}
	var gus = Person{"Gustavo", 3, "leadership"}
	var joe = Person{"Joe", 4, "juju"}
	var cole = Person{"Cole", 5, "pals"}
	var ben = Person{"Ben", 6, "charms"}
	var fred = Person{"Fred", 6, "kernos"}
	var people = []Person{al, ed, gus, joe, cole, ben, fred}
	for _, p := range people {
		err := db.Query(nil, insertPerson, p).Run()
		if err != nil {
			panic(err)
		}
	}

	insertLocation := sqlair.MustPrepare("INSERT INTO location (name, room_id, team) VALUES ($Location.name, $Location.room_id, $Location.team)", Location{})

	l1 := Location{1, "Congress Hall 1", "charms"}
	l2 := Location{100, "Marks Jacuzzi", "leadership"}
	l3 := Location{19, "Berlin 2", "juju"}
	l4 := Location{34, "Converted room #1065", "pals"}
	l5 := Location{8, "Converted room #1070", "kernos"}
	var locations = []Location{l1, l2, l3, l4, l5}
	for _, l := range locations {
		err := db.Query(nil, insertLocation, l).Run()
		if err != nil {
			panic(err)
		}
	}

	// Find out who is in room 19

	drop := sqlair.MustPrepare("DROP TABLE person; DROP TABLE location;")
	err = db.Query(nil, drop).Run()
	if err != nil {
		panic(err)
	}
}
