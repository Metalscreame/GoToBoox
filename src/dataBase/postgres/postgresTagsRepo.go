package postgres

import (
	"database/sql"
	"github.com/metalscreame/GoToBoox/src/dataBase/repository"
	"log"
)

type tagsRepositoryPG struct {
	Db *sql.DB
}

//NewTagsRepository is a function to get New TagsRepository which uses given connection
func NewTagsRepository(Db *sql.DB) repository.TagsRepository {
	return &tagsRepositoryPG{Db}
}

func (p tagsRepositoryPG) GetListOfTags() (tags []repository.Tags, err error) {
	rows, err := p.Db.Query("SELECT tag_id, title  FROM gotoboox.tags LIMIT 25")
	if err != nil {
		log.Printf("Get %v", err)
		return
	}
	defer rows.Close()
	var tag repository.Tags
	for rows.Next() {

		if err := rows.Scan(&tag.ID, &tag.Title);
			err != nil {
			log.Printf("Get %v", err)
		}
		tags = append(tags, tag)
	}
	if err := rows.Err(); err != nil {
		log.Printf("Get %v", err)

	}
	return tags, nil
}
