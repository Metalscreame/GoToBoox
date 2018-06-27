package repository

type TagsRepository interface {

GetListOfTags() (tags []Tags, err error)

}