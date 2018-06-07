package repository

type CategoryRepository interface{
	GetAllCategories () ([]Categories, error)
}
