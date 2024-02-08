package blog

import (
	"cloud.google.com/go/civil"
	"fmt"
	"log"
	"testing"
	"time"
)

var testTable = BigQueryTable{
	ProjectId: "huepattl",
	Dataset:   "site",
	Table:     "blogs_test",
}

var now = time.Now()
var testEntries = []*Entry{
	{
		DateTime: civil.DateTime{
			Date: civil.Date{Year: 1982, Month: 12, Day: 30},
			Time: civil.TimeOf(now),
		},
		AuthorName: "Christina",
		Title:      "Test 1",
		Text:       "Foo bar baz",
	},
	{
		DateTime: civil.DateTime{
			Date: civil.Date{Year: 1977, Month: 12, Day: 9},
			Time: civil.TimeOf(now),
		},
		AuthorName: "Timo",
		Title:      "Test 2",
		Text:       "Foo2 bar2 baz2",
	},
}

func TestListing(t *testing.T) {
	connection, err := CreateClient(testTable.ProjectId)
	if err != nil {
		t.Errorf("failed getting connection %s", err)
	}

	err = prepareTable(connection, false) // FIXME this should be true, but has timing problem?
	if err != nil {
		t.Errorf("failed preparing table with error %s", err)
	}

	// read data
	found, err := List(connection, testTable)
	if err != nil {
		t.Errorf("failed reading with error: %s", err)
	}

	// assert
	if len(found) < 1 {
		t.Errorf("no entries found while there should be some")
	}
	if found[0].Title != "Test 1" {
		t.Errorf("expecting 'Test 1' for 'Entry.Title', but got '%s'", found[0].Title)
	}
	for i := 0; i < len(found); i++ {
		log.Printf("found entry: %s", found[i])
	}

}

func prepareTable(connection Connection, create bool) error {
	// create temp. table
	if create == true {
		err := DeleteTable(connection, testTable)
		if err != nil {
			log.Printf("failed deleting table '%s' with error: %s - maybe OK since it was not there from previous run", testTable, err)
		}

		err = CreateTable(connection, testTable)
		if err != nil {
			return fmt.Errorf("failed creating table '%s' with error: %s", testTable, err)
		}
	}

	// fill with data
	num, err := Insert(testEntries, connection, testTable)
	if err != nil {
		return fmt.Errorf("failed writing data with error: %s", err)
	}
	log.Printf("inserted %d rows", num)
	return nil
}
