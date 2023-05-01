package main

import (
	"database/sql"

	"github.com/canonical/sqlair"

	_ "github.com/mattn/go-sqlite3"
)

type Location struct {
	ID   int    `db:"room_id"`
	Name string `db:"name"`
}

type Person struct {
	Name string `db:"name"`
	ID   int    `db:"room_id"`
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
		room_id integer,
		team text
	);
	CREATE TABLE location (
		room_id integer,
		name text
	)`)
	err = db.Query(nil, create).Run()
	if err != nil {
		panic(err)
	}

	insertPerson := sqlair.MustPrepare("INSERT INTO person (name, room_id, team) VALUES ($Person.name, $Person.room_id, $Person.team);", Person{})
	insertLocation := sqlair.MustPrepare("INSERT INTO location (name, room_id) VALUES ($Location.name, $Location.room_id)", Location{})

	var al = Person{"Alasatir", 2, "pals"}
	var ed = Person{"Ed", 2, "pals"}
	var gus = Person{"Gustavo", 1, "leadership"}
	var joe = Person{"Joe", 3, "juju"}
	var people = []Person{al, ed, gus, joe}
	for _, p := range people {
		err := db.Query(nil, insertPerson, p).Run()
		if err != nil {
			panic(err)
		}
	}
	l1 := Location{1, "Palmovka"}
	l2 := Location{2, "Berlin"}
	l3 := Location{3, "London"}
	var locations = []Location{l1, l2, l3}
	for _, l := range locations {
		err := db.Query(nil, insertLocation, l).Run()
		if err != nil {
			panic(err)
		}
	}

	drop := sqlair.MustPrepare("DROP TABLE person; DROP TABLE location;")
	err = db.Query(nil, drop).Run()
	if err != nil {
		panic(err)
	}

}
