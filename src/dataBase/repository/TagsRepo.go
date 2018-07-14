package repository

//TagsRepository is a repository interface that contains methods for tags that works with database
type TagsRepository interface {
	GetListOfTags() (tags []Tags, err error)
}
