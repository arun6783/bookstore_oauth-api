package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		panic(err)
	}

	fmt.Println("connection to cassandra successful")

}

func GetSession() *gocql.Session {
	return session
}
