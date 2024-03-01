package blog

import (
	"testing"
	"time"
)

var testRepo = RepositoryProperties{
	ProjectId:  "huepattl",
	Collection: "blogs_test",
}

var entries = []*BlogEntry{
	{
		Id:         "hello-world",
		Date:       time.Date(1977, 12, 9, 14, 35, 0, 0, time.UTC),
		Author:     "Timo",
		Tags:       "test, hello, world",
		Lead:       "Hello test blog entry #1",
		Title:      "Test 1",
		DocumentId: "1.md",
	},
	{
		Id:         "hello-world-again",
		Date:       time.Date(1982, 12, 30, 10, 58, 1, 2, time.UTC),
		Author:     "Christina",
		Tags:       "test, hello, world",
		Lead:       "Hello test blog entry #2",
		Title:      "Test 2",
		DocumentId: "2.md",
	},
}

func TestSave(t *testing.T) {
	Save(testRepo, entries[0])
	Save(testRepo, entries[1])
}

func TestList(t *testing.T) {
	found, err := List(testRepo)
	if err != nil {
		t.Errorf("Failed to list blog entries: %s", err)
	}

	t.Logf("Found: %s", found[0].Title)
}

func TestFindById(t *testing.T) {
	found, err := FindById(testRepo, "hello-world")
	if err != nil {
		t.Errorf("Failed to list blog entries: %s", err)
	}

	t.Logf("Found: %s", found.Title)
}
