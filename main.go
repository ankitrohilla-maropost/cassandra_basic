/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace test with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table test.test_tb(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on test.test_tb(timeline);
*/
package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the instance
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "test"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert
	if err := session.Query(`INSERT INTO test_tb (timeline, id, text) VALUES (?, ?, ?)`,
		"TL", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	if err := session.Query(`SELECT id, text FROM test_tb WHERE timeline = ? LIMIT 1`,
		"TL").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("test_tb content:", id, text)

	// list all content
	/*iter := session.Query(`SELECT id, text FROM test_tb WHERE timeline = ?`, "TL").Iter()
	for iter.Scan(&id, &text) {
		fmt.Println("test_tb content:", id, text)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}*/
}