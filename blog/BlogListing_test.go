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
	/*	{
		DateTime: civil.DateTime{
			Date: civil.Date{Year: 2024, Month: 2, Day: 3},
			Time: civil.TimeOf(now),
		},
		AuthorName: "Timo Böwing",
		Title:      "Stardate 2024-02-03, Saturday: Captain’s Log",
		Text:       "So after some twenty years of not really having a web site and being a hobbyist and\nprofessional software developer... I think it is time again to actually run a site again.\nI mean - isn't it a question of \"pride\"? ;)\n\nSo, as of now I do not know what will come here in the future. Maybe just a blog,\nmaybe I will provide information about my current and future projects here.\n\nTo be honest, my first idea to have a website again was not to have a website again,\nbut I wanted to develop one. From scratch, just for the fun of it and for learning.\n\nWell, currently, I am just rendering markdown based text here - dynamically rendered\neach time you visit this page. There are much smarter ways of achieving the goal...\n\nI could have done this instead:\n\n- build a site with Github Pages\n- use some other hosted CMS such as Wordpress, Joomla...\n- using a platform such as Squarespace, Wix, IONOS etc.\n- use a professional static page generator such as Hugo\n- just create some static HTML files and throw them on a VM in some datacenter\n- spin up a Raspberry Pi with or without runnung Docker, doing dynamic DNS routing to that box and run the pages using e.g.\n  - JVM with Quarkus, Spring Boot, Micronaut, Apache Tomcat, or just plain Jetty (JVM was and is my go-to stack for the past 22 years)\n  - doing similar but natively, such as building my Quarkus app using GraalVM\n  - spin up some Ruby, PHP, Django, LAMP, Sinatra, Node.js etc.\n- doing the above in some datacenter or even cloud provider\n- ...\n\nInstead - honestly - on the one hand I was tired of the operations part. I did not want to\nget my VM up to date, maintain and patch OS and middleware, play around with firewalls.\nTons of other stuff I am forgetting right now.\n\nSo, I decided to go to the dark side and host my site with a cloud provider. The serverless\nway. So, no dedicated Kubernetes cluster or VM, just plain Google Cloud Run.\nThis is basically a Kubernetes Knative deployment, running on an anonymous Kubernetes Cluster\nthat Google operates and I am not aware of.\n\nYou develop something, make a Dockerfile, and deploy it with some terminal command (in a pipleine or not).\n\nThat's for the ops part, I think more on that later.\n\nOK, so I decided to program my own stuff - knowing it would be inferior to anything else\nbut it is just fun. I decided to leave my feeling-at-home stack (JVM, Quarkus, Kotlin)\nand use something different. Something I was aware of and being one of the few stack\nI intentionally ignored: Go (golang).\n\nWhy? I mean there are other great programming languages, runtimes and frameworks - but\ncome on... Go uses capitals for functions! In Go you have `func DoSomething(arg string)` instead\nof `fun doSomething(arg: String)` like in most other languages.\n\nVerbs that start with uppercase letters! I had to ignore Go. Go uses upper and lower case\nto define the visibility scope of something. While other languages use reserved words like\n`public` and `private` for that, Go dooes this using case.\n\nI could have accepted all kind of styles differing from Java or Kotlin such as writing in\nKebap case, snake case etc. I get along with doing OO or functional... but this?\n\nWell, here we are, this page is presented employing Go. But that is up for future entries here.",
	},*/
}

func TestListing(t *testing.T) {
	connection, err := CreateClient(testTable.ProjectId)
	if err != nil {
		t.Errorf("failed getting connection %s", err)
	}

	//err = prepareTable(connection, true) // FIXME this should be true, but has timing problem?
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
