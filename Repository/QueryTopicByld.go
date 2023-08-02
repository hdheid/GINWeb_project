package Repository

import "sync"

type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})

	return topicDao
}

func (*TopicDao) QueryTopicById(id int64) *Topic {
	return TopicIndexMap[id]
}