package data

import (
	"github.com/go-playground/validator/v10"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

//post request for creating a schema
type CreateSchemaRequest struct{

	// BusinessID of the business this schema is for
	//
	// required: true
	// max length: 1000
	BusinessID string `json: "businessID" validate:"required"`

	// Greeting message for the user
	//
	// required: true
	GreetingStory string `json:"greetingStory" validate:"required"`

	// Action Handlers object along with the details
	//
	//required: true
	ActionHandlersObjectArray []ActionHandlersObject `json: "actionHandlersObjectArray" validate:"required"`

	// Title of the greeting card the business wants the user to see (Mostly the name of the business)
	//
	//required: true
	GreetingTitle string `json:"greetingTitle" validate:"required"`

	//Subtitle of the greeting card (Mostly business vertical)
	//
	GreetingSubTitle string `json:"greetingSubtitle"`

	//Top right data on the greeting card (if some additional data)
	//
	// max length: 10
	GreetingTopRightData string `json:"greetingTopRightData"`

	// Bottom Right Data of the greeting card (if some additional data)
	//
	// max length: 10
	GreetingBottomRightData string `json:"greetingBottomRightData"`

	// URL of Greeting card pic
	//
	GreetingCardPicURL string `json:"greetingCardPicURL"`
}

type ActionHandlersObject struct {
	// Action Handler string
	// order
	// feedback
	// cancel
	//
	//required: true
	ActionHandler string `json: "actionHandler" validate:"required"`

	// Action Handler question object array
	//
	//required:true
	QuestionArray []QuestionObjectArray `json:"questionArray" validate:"required"`

	// Transaction success response from business
	// 
	EndingMessage string `json:"EndingMessage"`

	//addedActionHandler after the transaction
	//
	AddedActionHandler []string `addedActionHandler`
}

type QuestionObjectArray struct {

	//Question to be displayed
	//
	//required:true
	Question string `json:"question" validate:"required"`

	//Type of response expected (needed for setting the UI)
	// Number-Input 
	// Date-Input 
	// Quick-Reply
	// Text-Input
	// Location-Input
	// Polar-Input
	// Date-Input-Range
	// Quick-Reply-Advance
	//
	//required:true
	ResponseType string `json:"responseType" validate:"required"`
	
	//Context of the question (helps to prepare the payload)
	// key of the API needed to fire for actual action
	// 
	//required:true
	QuestionContext string `json:"questionContext" validate:"required"`

	//Type of question bot asks
	// Text
	// Card
	// Voice
	// Picture
	//
	//required:true
	QuestionType string `json:"questionType" validate:"required"`

	//Choice of response provided by the business for quick replies
	//
	CustomResponseChoice []string `json:"customResponseChoice"` 

	//Card object for the question
	//
	QestionCard CardObject `json:"questionCard"`

	//Card object for the answer
	//
	AnswerCard []CardObject `json:"answerCard"`
}

type CardObject struct{
    //URL of the pic on the card
	//
	CardPic string `json:"cardPic"`

	//Card title
	//
	Title string `json:"title"`

	//Subtitle of the card
	//
	SubTitle string `json:"subtitle"`

	//Story on the card
	//
	Story string `json:"story"`

	//Top right data on the card
	//
	TopRightData string `json:"topRightData"`

	//Bottom Right Data of the card
	//
	BottomRightData string `json:"bottomRightData"`
}

//post response for creating a schema
type CreateSchemaResponse struct {

    // BusinessID of the business
	//
	// required: true
	BusinessID string `json: "businessID"`

	//Message response
	//
	Message string `json:"message"`

	// status code
	// 200 Success
	// 401 Request Error
	// 501 Server Error
	// 502 Database Error
	//
	Status int64 `json:"status"`
}

type BotSchemaDataStructure struct {
	BusinessID string `json:"businessid"`
	Initialise InitialiseStructSchema `json:"initialise"`
	Action []struct{
		Handler string `json:"handler"`
		Data []QuestionObjectArray `json:"data"` 
	} `json:"action"`
	Response struct{
		Order struct {
			Message string `json:"message"`
			AddedActionHandler []string `addedActionHandler`
		}
	} `json"response"`
}

type InitialiseStructSchema struct{
	Message string `json:"message"`
	ActionHandlers []string `json:"actionHandlers"`
	GreetingCard CardObject `json:"greetingCard"`	
}

func (d *CreateSchemaRequest) ValidateCreateSchemaRequest() error {
	validate := validator.New()
	return validate.Struct(d)
}