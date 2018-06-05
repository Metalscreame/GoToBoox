package repository

type BookRepository interface {
	GetAll() ([]interface{}, error)
	GetByCategory(categoryID int) ([]Book, error)
	GetByID(bookID int) (b Book, err error)
	GetMostPopularBooks(id int) ([]Book, error)
}

//For connection to HerokuDatabase
/*func openDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://zrlfyamblttpom:e2c0e8832ea228e6b15e553ce69f7cb2c0ff4d646ff0f284245ce77cc78b437b@ec2-54-247-111-19.eu-west-1.compute.amazonaws.com:5432/d7ckgvm53enhum")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Print("Can not connect to database: ", err)
	}

	return db
}
*/





