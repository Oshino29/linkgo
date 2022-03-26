package main

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func newStorage(path string) *Storage {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatalf("can't create or open database:\n%s\n", err.Error())
	}

	db.SetMaxOpenConns(1)

	s := &Storage{ db: db }
	s.Init()

	return s
}

func (s *Storage) Init() {
	query := `
			CREATE TABLE IF NOT EXISTS Item (
				ItemID INTEGER PRIMARY KEY,
				Title TEXT,
				Url VARCHAR(2083),
				Desc TEXT,
				CONSTRAINT UC_Item UNIQUE (Title, Url)
			);

			CREATE TABLE IF NOT EXISTS Tag (
				TagID INTEGER PRIMARY KEY,
				Name TEXT,
				ItemID INTEGER
			);
	`

	_, err := s.db.Exec(query)
	if err != nil {
		log.Fatalf("can't create table in database")
	}
}

func (s *Storage) saveItem(item Item) int64 {
	// insert information into table Item
	result, err := s.db.Exec(
		"INSERT INTO Item (Title, Url, Desc) VALUES (?, ?, ?)",
		item.Title, item.Url, item.Description)
	if err != nil { log.Fatalf("error when save item: %s",err.Error()) }

	// get id of inserted item
	item.ID, err = result.LastInsertId()
	if err != nil { log.Fatalf("error when recive id of saved item: %s", err.Error())}

	// insert tag-item relation into table tag 
	for tag := range item.Tags {
		_, err = s.db.Exec(
			"INSERT INTO Tag (Name, ItemID) VALUES (?, ?)",
			tag, item.ID)
		if err != nil { log.Fatalf("error when save item tags: %s", err.Error())}
	}
	return item.ID
}

// load all items from database and return a map with itemID as key
func (s *Storage) loadAllItems() *Items{
	items := make(Items, 0)
	var item Item

	// rows will be like ItemID Title Url Desc Tag
	// duplicated items in rows may encounter if the item has multiple tags
	query := `
			select Item.ItemID, Item.Title, Item.Url, Item.Desc, Tag.Name from Item
			left join Tag on Tag.ItemID = Item.ItemID
			`
	rows, err := s.db.Query(query)
	if err != nil {
		log.Fatal("returned error when query all items")
	}
	defer rows.Close()
	
	// iterate through queried rows
	var tag string
	for rows.Next() {
		// scan rows to item and a temp tag variable
		rows.Scan(&item.ID, &item.Title, &item.Url, &item.Description, &tag)
		if itemInMap, ok := items[item.ID]; !ok {
			// assign item to map[item.ID] if item not found in map
			item.Tags = append(item.Tags, tag)
			items[item.ID] = item
		} else {
			// if item found already in map
			// clone item out of map, append new tag to item and store back to map
			itemInMap.Tags = append(itemInMap.Tags, tag)
			items[item.ID] = itemInMap
		}

	}


	return &items
}

// don't need to query id of item after knowing about sql.Result.LastInsertedId
// func (s *Storage) queryItemID(item *Item) int {
// 	query := `
// 			SELECT ItemID from Item
// 			WHERE Title = ? AND Url = ?
// 	`

// 	var id int = -1

// 	row := s.db.QueryRow(query, item.Title, item.Url)
// 	switch err := row.Scan(&id); err {
// 	case sql.ErrNoRows:
// 	  log.Printf("no rows were returned by item:\n    %s\n    %s\n", item.Title, item.Url)
// 	case nil:
// 	  log.Printf("query for itemid returned:  %d", id)
// 	default:
// 	  panic(err)
// 	}
	
// 	return id
// }