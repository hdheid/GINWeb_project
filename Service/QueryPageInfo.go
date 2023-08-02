package Service

import (
	"GINWeb_project/Repository"
	"errors"
	"sync"
)

type PageInfo struct {
	Topic    *Repository.Topic
	PostList []*Repository.Post
}

type QueryPageInfoFlow struct {
	pageInfo PageInfo
}

func QueryPageInfo(topic int64) (*PageInfo, error) {
	pageInfo := QueryPageInfoFlow{}
	return pageInfo.Do(topic)
}

func (f *QueryPageInfoFlow) Do(topicId int64) (*PageInfo, error) {
	if err := f.prepareInfo(topicId); err != nil {
		return nil, err
	}
	if err := f.checkParam(); err != nil {
		return nil, err
	}
	return &f.pageInfo, nil
}

func (f *QueryPageInfoFlow) checkParam() error {
	if f.pageInfo.Topic.Id <= 0 {
		return errors.New("topic id must can't be less than or equal to zero")
	}
	return nil
}

func (f *QueryPageInfoFlow) prepareInfo(topicId int64) error {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		f.pageInfo.Topic = Repository.NewTopicDaoInstance().QueryTopicById(topicId)
	}()

	go func() {
		defer wg.Done()
		f.pageInfo.PostList = Repository.NewPostDaoInstance().QueryPostById(topicId)
	}()

	wg.Wait()
	return nil
}