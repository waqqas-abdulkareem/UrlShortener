package urlshortener

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	database "github.com/w-k-s/short-url/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"testing"
	"time"
)

type URLRepositoryTestSuite struct {
	suite.Suite
	db      *database.Db
	urlRepo *URLRepository
	record  *URLRecord
}

func TestURLRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(URLRepositoryTestSuite))
}

func (suite *URLRepositoryTestSuite) SetupTest() {
	logger := log.New(os.Stdout, "short-url: ", log.Ldate|log.Ltime)
	suite.db = database.New("mongodb://localhost:27017/shorturl_test")
	suite.urlRepo = NewURLRepository(suite.db, logger)

	suite.record = &URLRecord{
		"http://www.example.com",
		"shrt",
		time.Now(),
	}
}

func (suite *URLRepositoryTestSuite) TearDownTest() {

	suite.db.Instance().
		C("urls").
		RemoveAll(bson.M{})

	defer suite.db.Close()

}

func (suite *URLRepositoryTestSuite) TestSaveRecordSucccessful() {

	_, err := suite.urlRepo.SaveRecord(suite.record)

	assert.Nil(suite.T(), err, "Expected: save record. Got: %s", err)
}

func (suite *URLRepositoryTestSuite) TestDuplicateRecordFails() {

	suite.urlRepo.SaveRecord(suite.record)
	_, err := suite.urlRepo.SaveRecord(suite.record)

	assert.True(suite.T(), mgo.IsDup(err), "Expected: duplication error. Got: %s", err)
}

func (suite *URLRepositoryTestSuite) TestFindExistingShortURL() {
	_, err := suite.urlRepo.SaveRecord(suite.record)
	if err != nil {
		panic(err)
	}

	result, err := suite.urlRepo.ShortURL(suite.record.LongURL)
	expectation := result != nil && result.ShortId == suite.record.ShortId
	assert.True(suite.T(), expectation, "Expected Matching ShortId '%s'. Got: '%v' (error: '%s')", suite.record.ShortId, result, err)
}

func (suite *URLRepositoryTestSuite) TestFindAbsentShortURL() {

	result, err := suite.urlRepo.ShortURL("http://www.nil.com")
	assert.NotNil(suite.T(), err, "Expected err when shortId not found. Got: nil. (record: %v)", result)
}

func (suite *URLRepositoryTestSuite) TestFindExistingLongURL() {
	_, err := suite.urlRepo.SaveRecord(suite.record)
	if err != nil {
		panic(err)
	}

	result, err := suite.urlRepo.LongURL(suite.record.ShortId)
	expectation := result != nil && result.LongURL == suite.record.LongURL

	assert.True(suite.T(), expectation, "Expected Matching LongURL '%s'. Got: '%v' (error: '%s')", suite.record.LongURL, result, err)
}

func (suite *URLRepositoryTestSuite) TestFindAbsentLongURL() {

	result, err := suite.urlRepo.LongURL("nil")
	assert.NotNil(suite.T(), err, "Expected err when longUrl not found. Got: nil. (record: %v)", result)

}