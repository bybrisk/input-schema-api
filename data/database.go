package data

import (
	//"fmt"
	//"go.mongodb.org/mongo-driver/bson"
	"github.com/shashank404error/shashankMongo"
	log "github.com/sirupsen/logrus"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateSchemaDB(d *CreateSchemaRequest) error {
	collectionName := shashankMongo.DatabaseName.Collection("bot-schema")
	
	var actionHandlerArray []string
	for _,v:=range d.ActionHandlersObjectArray{
		actionHandlerArray = append(actionHandlerArray,v.ActionHandler)
	}

	schema := BotSchemaDataStructure{
		BusinessID: d.BusinessID,
		Initialise: InitialiseStructSchema{
			Message : d.GreetingStory,
		    ActionHandlers: actionHandlerArray, 
		    GreetingCard: CardObject{
				CardPic: d.GreetingCardPicURL,
				Title: d.GreetingTitle,
				SubTitle: d.GreetingSubTitle,
				Story: d.GreetingStory,
				TopRightData: d.GreetingTopRightData,
				BottomRightData: d.GreetingBottomRightData, 
			},
		},
	}


	_, insertErr := collectionName.InsertOne(shashankMongo.CtxForDB, schema)
	if insertErr != nil {
		log.Error("CreateBotSchemaMongo ERROR:")
		log.Error(insertErr)
	} 

	return insertErr
}
