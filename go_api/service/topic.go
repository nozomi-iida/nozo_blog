package service

import (
	"github.com/nozomi-iida/nozo_blog/domain/topic"
	"github.com/nozomi-iida/nozo_blog/domain/topic/sqlite"
	"github.com/nozomi-iida/nozo_blog/entity"
)

type topicConfigurations func(tp *TopicService) error

type TopicService struct {
	tp topic.TopicRepository
}

func NewTopicService (cfgs ...topicConfigurations) (*TopicService, error) {
	ts := &TopicService{}

	for _, cfg := range cfgs {
		err := cfg(ts)
		if err != nil {
			return nil, err
		}
	}

	return ts, nil
}

func WithSqliteTopicRepository(fileString string) topicConfigurations {
	return func(ts *TopicService) error {
		s, err := sqlite.New(fileString)
		if err != nil {
			return err
		}
		ts.tp = s

		return nil
	}
}

func (ts *TopicService) Create(name string, description string) (entity.Topic, error)  {
	tp, err := entity.NewTopic(entity.Topic{Name: name, Description: description})	
	tp, err = ts.tp.Create(tp)
	if err != nil {
		return entity.Topic{}, err
	}

	return tp, nil
}
