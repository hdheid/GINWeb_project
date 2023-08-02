package Repository

import (
	"bufio"
	"encoding/json"
	"os"
)

type Topic struct {
	Id          int64  `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
}

type Post struct {
	Id          int64  `json:"id"`
	Parent_id   int64  `json:"parent_id"`
	Content     string `json:"content"`
	Create_time int64  `json:"create_time"`
}

var (
	TopicIndexMap map[int64]*Topic
	PostIndexMap  map[int64][]*Post
)

func Init(filePath string) error {
	if err := initTopicIndexMap(filePath); err != nil {
		return err
	}

	if err := initPostIndexMap(filePath); err != nil {
		return err
	}

	//for _, value := range PostIndexMap {
	//	for _, val := range value {
	//		fmt.Printf("%v\n", *val)
	//	}
	//}
	return nil
}

func initTopicIndexMap(filePath string) error {
	open, err := os.Open(filePath + "topic")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)
	topicTmpMAp := make(map[int64]*Topic)

	for scanner.Scan() {
		text := scanner.Text()
		var topic Topic
		if err := json.Unmarshal([]byte(text), &topic); err != nil {
			return err
		}
		topicTmpMAp[topic.Id] = &topic

		//fmt.Println("解析结果:")
		//fmt.Println(topic)
	}

	TopicIndexMap = topicTmpMAp

	return nil
}

func initPostIndexMap(filePath string) error {
	open, err := os.Open(filePath + "post")
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(open)
	postIndexMap := make(map[int64][]*Post)

	for scanner.Scan() {
		text := scanner.Text()
		var post Post
		if err := json.Unmarshal([]byte(text), &post); err != nil {
			return err
		}

		posts, ok := postIndexMap[post.Parent_id]
		if !ok {
			postIndexMap[post.Parent_id] = []*Post{&post}
			continue
		}
		posts = append(posts, &post)
		postIndexMap[post.Parent_id] = posts
	}

	PostIndexMap = postIndexMap

	return nil
}