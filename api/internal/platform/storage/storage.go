package storage

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"

	apiErrors "github.com/larlandis/shorturl/internal/pkg/errors"
)

type cassandraRepo struct {
	session *gocql.Session
}

func New(clusterName string) *cassandraRepo {
	cluster := gocql.NewCluster(clusterName)
	cluster.Consistency = gocql.Quorum
	var (
		session *gocql.Session
		err     error
	)
	// ugly hack don't this
	for {
		session, err = cluster.CreateSession()
		if err != nil {
			fmt.Println(err.Error() + " retrying...")
			time.Sleep(5 * time.Second)
		} else {
			break
		}
	}
	err = session.Query("create keyspace shorturl with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };").Exec()
	if err != nil {
		log.Println(err.Error())
	}
	err = session.Query("create table if not exists shorturl.hashes(input text,hash text,PRIMARY KEY (hash));").Exec()
	if err != nil {
		log.Println(err.Error())
	}
	return &cassandraRepo{
		session: session,
	}
}

func (c cassandraRepo) SavePair(_ context.Context, input string, hash string) error {
	return c.session.Query(
		`INSERT INTO shorturl.hashes (input, hash) VALUES (?, ?)`,
		input, hash).Exec()
}

func (c cassandraRepo) Search(_ context.Context, hash string) (url string, err error) {
	err = c.session.Query(
		`SELECT input FROM shorturl.hashes WHERE hash = ? LIMIT 1`,
		hash).
		Consistency(gocql.One).Scan(&url)
	if errors.Is(err, gocql.ErrNotFound) {
		return "", apiErrors.NotFoundError(err)
	}
	return url, err
}
