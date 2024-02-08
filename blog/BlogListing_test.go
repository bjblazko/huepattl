package blog

import (
	"cloud.google.com/go/civil"
	"log"
	"testing"
	"time"
)

func TestList(t *testing.T) {
	table := BigQueryTable{
		ProjectId: "huepattl",
		Dataset:   "site",
		Table:     "blogs_test",
	}

	found, err := List(table)
	if err != nil {
		t.Errorf("failed with error: %s", err)
	}
	if len(found) < 1 {
		t.Errorf("no entries found while there should be some")
	}
	if found[0].Bucket != "Foo" {
		t.Errorf("expecting 'Foo' for 'Entry.Bucket', but got '%s'", found[0].Bucket)
	}
	for i := 0; i < len(found); i++ {
		log.Printf("found entry: %s", found[i])
	}
}

func TestInsert(t *testing.T) {
	table := BigQueryTable{
		ProjectId: "huepattl",
		Dataset:   "site",
		Table:     "blogs_test",
	}
	now := time.Now()
	items := []*Entry{
		{
			DateTime:   civil.DateTime{Date: civil.Date{Year: 1982, Month: 12, Day: 30}, Time: civil.TimeOf(now)},
			AuthorName: "Christina",
			Title:      "Test 1",
			Preview:    "Foo bar baz",
			Bucket:     "test",
			Filename:   "test1.md",
		},
		{
			DateTime:   civil.DateTime{Date: civil.Date{Year: 1977, Month: 12, Day: 9}, Time: civil.TimeOf(now)},
			AuthorName: "Timo",
			Title:      "Test 2",
			Preview:    "Foo2 bar2 baz2",
			Bucket:     "test", Filename: "test2.md"},
	}
	num, err := Insert(items, table)
	if err != nil {
		t.Errorf("failed: %s", err)
	}
	t.Logf("inserted %d rows", num)
}

func TestClear(t *testing.T) {
	table := BigQueryTable{
		ProjectId: "huepattl",
		Dataset:   "site",
		Table:     "foo",
	}
	err := DeleteTable(table)
	if err != nil {
		t.Errorf("failed: %s", err)
	}
}

func TestCreateTable(t *testing.T) {
	table := BigQueryTable{
		ProjectId: "huepattl",
		Dataset:   "site",
		Table:     "foo",
	}
	CreateTable(table)
}
