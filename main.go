// Ortelius v11 User Microservice that handles creating and retrieving User Information
package main

import (
	"context"
	"encoding/json"

	_ "github.com/ortelius/scec-user/docs"

	driver "github.com/arangodb/go-driver/v2/arangodb"
	"github.com/arangodb/go-driver/v2/arangodb/shared"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/ortelius/scec-commons/database"
	"github.com/ortelius/scec-commons/model"
)

var logger = database.InitLogger()
var dbconn = database.InitializeDB("evidence")

// GetUsers godoc
// @Summary Get a List of Users
// @Description Get a list ofthe user.
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/user [get]
func GetUsers(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	// query all the users in the collection
	aql := `FOR user in evidence
			FILTER (user.objtype == 'User')
			RETURN user`

	// execute the query with no parameters
	if cursor, err = dbconn.Database.Query(ctx, aql, nil); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err) // log error
	}

	defer cursor.Close() // close the cursor when returning from this function

	var users []*model.User // define a list of users to be returned

	for cursor.HasMore() { // loop thru all of the documents

		user := model.NewUser()      // fetched user
		var meta driver.DocumentMeta // data about the fetch

		// fetch a document from the cursor
		if meta, err = cursor.ReadDocument(ctx, user); err != nil {
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
		users = append(users, user)                                          // add the user to the list
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key) // log the key
	}

	return c.JSON(users) // return the list of users in JSON format
}

// GetUser godoc
// @Summary Get a User
// @Description Get a user based on the _key or name.
// @Tags user
// @Accept */*
// @Produce json
// @Success 200
// @Router /msapi/user/:key [get]
func GetUser(c *fiber.Ctx) error {

	var cursor driver.Cursor       // db cursor for rows
	var err error                  // for error handling
	var ctx = context.Background() // use default database context

	key := c.Params("key")                // key from URL
	parameters := map[string]interface{}{ // parameters
		"key": key,
	}

	// query the users that match the key or name
	aql := `FOR user in evidence
			FILTER (user.name == @key or user._key == @key)
			RETURN user`

	// run the query with patameters
	if cursor, err = dbconn.Database.Query(ctx, aql, &driver.QueryOptions{BindVars: parameters}); err != nil {
		logger.Sugar().Errorf("Failed to run query: %v", err)
	}

	defer cursor.Close() // close the cursor when returning from this function

	user := model.NewUser() // define a user to be returned

	if cursor.HasMore() { // user found
		var meta driver.DocumentMeta // data about the fetch

		if meta, err = cursor.ReadDocument(ctx, user); err != nil { // fetch the document into the object
			logger.Sugar().Errorf("Failed to read document: %v", err)
		}
		logger.Sugar().Infof("Got doc with key '%s' from query\n", meta.Key)

	} else { // not found so get from NFT Storage
		if jsonStr, exists := database.MakeJSON(key); exists {
			if err := json.Unmarshal([]byte(jsonStr), user); err != nil { // convert the JSON string from LTF into the object
				logger.Sugar().Errorf("Failed to unmarshal from LTS: %v", err)
			}
		}
	}

	return c.JSON(user) // return the user in JSON format
}

// NewUser godoc
// @Summary Create a User
// @Description Create a new User and persist it
// @Tags user
// @Accept application/json
// @Produce json
// @Success 200
// @Router /msapi/user [post]
func NewUser(c *fiber.Ctx) error {

	var err error                  // for error handling
	var meta driver.DocumentMeta   // data about the document
	var ctx = context.Background() // use default database context
	user := model.NewUser()        // define a user to be returned

	if err = c.BodyParser(user); err != nil { // parse the JSON into the user object
		return c.Status(503).Send([]byte(err.Error()))
	}

	cid, dbStr := database.MakeNFT(user) // normalize the object into NFTs and JSON string for db persistence

	logger.Sugar().Infof("%s=%s\n", cid, dbStr) // log the new nft

	var resp driver.CollectionDocumentCreateResponse
	// add the user to the database.  Ignore if it already exists since it will be identical
	if resp, err = dbconn.Collection.CreateDocument(ctx, user); err != nil && !shared.IsConflict(err) {
		logger.Sugar().Errorf("Failed to create document: %v", err)
	}
	meta = resp.DocumentMeta
	logger.Sugar().Infof("Created document in collection '%s' in db '%s' key='%s'\n", dbconn.Collection.Name(), dbconn.Database.Name(), meta.Key)

	return c.JSON(user) // return the user object in JSON format.  This includes the new _key
}

// setupRoutes defines maps the routes to the functions
func setupRoutes(app *fiber.App) {

	app.Get("/swagger/*", swagger.HandlerDefault) // handle displaying the swagger
	app.Get("/msapi/user", GetUsers)              // list of users
	app.Get("/msapi/user/:key", GetUser)          // single user based on name or key
	app.Post("/msapi/user", NewUser)              // save a single user
}

// @title Ortelius v11 User Microservice
// @version 11.0.0
// @description RestAPI for the User Object
// @description ![Release](https://img.shields.io/github/v/release/ortelius/scec-user?sort=semver)
// @description ![license](https://img.shields.io/github/license/ortelius/.github)
// @description
// @description ![Build](https://img.shields.io/github/actions/workflow/status/ortelius/scec-user/build-push-chart.yml)
// @description [![MegaLinter](https://github.com/ortelius/scec-user/workflows/MegaLinter/badge.svg?branch=main)](https://github.com/ortelius/scec-user/actions?query=workflow%3AMegaLinter+branch%3Amain)
// @description ![CodeQL](https://github.com/ortelius/scec-user/workflows/CodeQL/badge.svg)
// @description [![OpenSSF-Scorecard](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user/badge)](https://api.securityscorecards.dev/projects/github.com/ortelius/scec-user)
// @description
// @description ![Discord](https://img.shields.io/discord/722468819091849316)

// @termsOfService http://swagger.io/terms/
// @contact.name Ortelius Google Group
// @contact.email ortelius-dev@googlegroups.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /msapi/user
func main() {
	port := ":" + database.GetEnvDefault("MS_PORT", "8080") // database port
	app := fiber.New()                                      // create a new fiber application
	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowOrigins: "*",
	}))

	setupRoutes(app) // define the routes for this microservice

	if err := app.Listen(port); err != nil { // start listening for incoming connections
		logger.Sugar().Fatalf("Failed get the microservice running: %v", err)
	}
}
