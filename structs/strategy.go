package structs

import (
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	StrategyNameLowerResponse  StrategyName = "lowerResponse"
	StrategyNameLowestResponse StrategyName = "lowestResponse"
	StrategyNameConfidenceBase StrategyName = "confidenceBase"
	StrategyNameMostConfidence StrategyName = "mostConfidence"
	StrategyNameLinearResponse StrategyName = "linearResponse"
)

type StrategyName string

type StrategyData struct {
	Id         *primitive.ObjectID    `json:"id,omitempty" bson:"_id,omitempty"`
	Meta       *primitive.ObjectID    `json:"meta" bson:"meta"`
	GamblerId  *primitive.ObjectID    `json:"gambler_id" bson:"gambler_id"`
	Name       StrategyName           `json:"name" bson:"name"`
	Properties map[string]interface{} `json:"properties" bson:"properties"`
}

func (s *StrategyData) GetProperty(name string) interface{} {
	if p, exist := s.Properties[name]; exist {
		return p
	} else {
		log.Warn("no property: ", name)
		return nil
	}
}

type StrategyMeta struct {
	Id          *primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        StrategyName        `json:"name" bson:"name"`
	Description string              `json:"description" bson:"description"`
	Properties  []struct {
		Name string `json:"name" bson:"name"`
		Type string `json:"type" bson:"type"` // int, float, string
	} `json:"properties" bson:"properties"`
}
