package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()
	e.GET("/", Test)
	e.Logger.Fatal(e.Start(":1323"))
}

type Question struct {
	gorm.Model
	Answer []Answer
	Text   string
}

type Answer struct {
	gorm.Model
	Text       string
	Eligible   bool
	QuestionID uint
}

func Test(c echo.Context) error {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&Question{}, &Answer{})

	// Create
	// createQuestionAnswer()

	// read question with andwer
	// question := questionAnswers()

	// read all answers
	// answers := Answers()

	// read all user with credit cards
	// questions := QuestionsAnswers()

	// read all users only
	// questions := questions()

	// update question
	// updateQuestion()

	// read question by id
	// question := questionById()

	// read answer by id
	// answer := answerById()

	// read answer by id with question
	// answer := answerByIdQuestion()

	// update answer by id and by id question
	// answer := updateAnswerByIdQuestion()

	// delete answer by id and by id question
	// del := deleteAnswerByIdQuestionAndByIdAswer()

	// delete question by id and delete all answers by id question
	del := deleteQuestionByIdAndDeleteAllAnswersByIdQuestion()

	return c.JSON(http.StatusOK, del)
}

func questions() *[]Question {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var questions []Question
	db.Find(&questions)

	return &questions
}

func QuestionsAnswers() *[]Question {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var questionsAnswers []Question
	db.Preload("Answer").Find(&questionsAnswers)

	return &questionsAnswers
}

func Answers() *[]Answer {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var answers []Answer
	db.Find(&answers)

	return &answers
}

func questionAnswers() *Question {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var question Question
	db.Preload("Answer").First(&question, 1)

	return &question
}

func createQuestionAnswer() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Create
	db.Create(&Question{Text: "Jinzhu3", Answer: []Answer{
		{Text: "1111-2222-3333-4444", Eligible: false},
		{Text: "1111-2222-3333-4444", Eligible: false},
		{Text: "1111-2222-3333-4444", Eligible: false},
		{Text: "1111-2222-3333-4444", Eligible: false},
		{Text: "5555-6666-7777-8800", Eligible: true},
	}})
	return db
}

func questionById() *Question {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var question Question
	db.First(&question, 1)

	return &question
}

func updateQuestion() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.Model(&Question{}).Where("id = ?", 1).Update("text", "Jinzhu2up")
	return db
}

func answerById() *Answer {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var answer Answer
	db.First(&answer, 1)

	return &answer
}

func answerByIdQuestion() *Answer {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var answer Answer
	db.Where("id = @id AND question_id = @question_id", sql.Named("id", 1), sql.Named("question_id", 1)).Find(&answer)
	return &answer
}

func updateAnswerByIdQuestion() *Answer {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var answer Answer
	db.Where("id = @id AND question_id = @question_id", sql.Named("id", 1), sql.Named("question_id", 1)).Find(&answer)
	db.Model(&Answer{}).Where("id = @id AND question_id = @question_id", sql.Named("id", 1), sql.Named("question_id", 1)).Update("text", "Jinzhu2up")
	return &answer
}

func deleteAnswerByIdQuestionAndByIdAswer() *Answer {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var answer Answer
	db.Where("id = @id AND question_id = @question_id", sql.Named("id", 2), sql.Named("question_id", 1)).Find(&answer)
	db.Delete(&answer)
	return &answer
}

func deleteQuestionByIdAndDeleteAllAnswersByIdQuestion() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.sqlite"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	var question Question
	db.Where("id = @id", sql.Named("id", 2)).Find(&question)
	db.Delete(&question)

	var answers []Answer
	db.Where("question_id = @question_id", sql.Named("question_id", 2)).Find(&answers)
	db.Delete(&answers)
	return db
}
