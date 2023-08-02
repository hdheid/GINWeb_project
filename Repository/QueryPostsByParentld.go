package Repository

import "sync"

type PostDao struct {
}

var (
	postDao  *PostDao
	PostOnce sync.Once
)

func NewPostDaoInstance() *PostDao {
	topicOnce.Do(
		func() {
			postDao = &PostDao{}
		})

	return postDao
}

func (*PostDao) QueryPostById(id int64) []*Post {
	return PostIndexMap[id]
}
