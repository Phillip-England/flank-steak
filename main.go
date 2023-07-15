package main

import (
	"flank-steak/src/types"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {

	//==========================================================================
	// CONFIG
	//==========================================================================

	err := godotenv.Load()
	if err != nil {
		log.Panic(err.Error())
	}
	
	//==========================================================================
	// DATABASE
	//==========================================================================

	database := types.NewDatabase()
	err = database.InitTables()
	if err != nil {
		log.Fatal(err.Error())
	}

	//==========================================================================
	// ROUTER
	//==========================================================================

	r := gin.Default()
	r.LoadHTMLGlob("./templates/**/*")
	r.Static("/static", "./static")

	//==========================================================================
	// PAGES
	//==========================================================================

	r.GET("/", func(c *gin.Context) {
		_, err := c.Cookie(os.Getenv("SESSION_TOKEN_KEY"))
		if err == nil {
			c.Redirect(303, "/locations")
			return
		}
		c.HTML(200, "index.html", gin.H{
			"Banner": "CFA Tools",
			"LoginFormErr": c.Query("LoginFormErr"),
		})
	})

	r.GET("/signup", func(c *gin.Context) {
		_, err := c.Cookie(os.Getenv("SESSION_TOKEN_KEY"))
		if err == nil {
			c.Redirect(303, "/locations")
			return
		}
		c.HTML(200, "signup.html", gin.H{
			"Banner": "CFA Tools",
			"SignupFormErr": c.Query("SignupFormErr"),
		})
	})

	r.GET("/locations", func(c *gin.Context) {
		userModel := types.NewUserModel()
		err := userModel.Auth(c, database)
		if err != nil {
			c.Redirect(303, "/")
			return
		}
		locations, err := database.GetLocationsByUserID(userModel.ID)
		if err != nil {
			log.Panic(err.Error())
		}
		hasNoLocations := len(locations) == 0
		fmt.Println(hasNoLocations)
		c.HTML(200, "locations.html", gin.H{
			"LocationFormErr": c.Query("LocationFormErr"),
			"Locations": locations,
			"HasNoLocations": hasNoLocations,
			"Banner": "Locations Dashboard",
		})
	})

	r.GET("/location/:id", func(c *gin.Context) {
		userModel := types.NewUserModel()
		err := userModel.Auth(c, database)
		if err != nil {
			c.Redirect(303, "/")
			return
		}
		locationID := c.Params.ByName("id")
		locationModel, err := database.GetLocationByID(locationID)
		if err != nil {
			log.Panic(err.Error())
		}
		if (userModel.ID != locationModel.UserID) {
			c.Redirect(303, "/locations")
			return
		}
		c.HTML(200, "SingleLocation.html", gin.H{
			"Location": locationModel,
			"Banner": locationModel.Name,
		})
	})


	r.GET("/logout", func(c *gin.Context) {
		c.SetCookie(os.Getenv("SESSION_TOKEN_KEY"), "", -1, "/", "localhost", true, true)
		c.Redirect(303, "/")
	})

	//==========================================================================
	// ACTIONS
	//==========================================================================

	r.POST("/actions/signup", func(c *gin.Context) {
		userModel := types.NewUserModel()
		userModel.SetEmail(c.PostForm("email"))
		userModel, err = userModel.SetPassword(c.PostForm("password"))
		if err != nil {
			log.Fatal(err.Error())
		}
		err := userModel.Validate(database)
		if err != nil {
			c.Redirect(303, fmt.Sprintf("/signup?SignupFormErr=%s", err.Error()))
			return
		}
		err = userModel.Insert(database)
		if err != nil {
			log.Fatal(err.Error())
		}
		c.Redirect(303, "/")
	})

	r.POST("/actions/login", func(c *gin.Context) {
		userModel, err := types.NewUserModel().FindByEmail(database, c.PostForm("email"))
		if err != nil {
			c.Redirect(303, fmt.Sprintf("/?LoginFormErr=%s", "invalid credentials"))
			return
		}
		err = userModel.ComparePassword(c.PostForm("password"))
		if err != nil {
			c.Redirect(303, fmt.Sprintf("/?LoginFormErr=%s", "invalid credentials"))
			return
		}
		err = userModel.DeleteSessionsByUser(database)
		if err != nil {
			log.Fatal(err.Error())
		}
		sessionModel := types.NewSessionModel()
		err = sessionModel.Insert(database, userModel.ID)
		if err != nil {
			log.Fatal(err.Error())
		}
		c.SetCookie(os.Getenv("SESSION_TOKEN_KEY"), sessionModel.Token, 86400, "/", "localhost", true, true)
		c.Redirect(303, "/locations")
	})

	r.POST("/actions/location", func(c *gin.Context) {
		userModel := types.NewUserModel()
		err := userModel.Auth(c, database)
		if err != nil {
			c.Redirect(303, "/")
			return
		}
		fmt.Println(userModel)
		locationModel := types.NewLocationModel()
		err = locationModel.SetName(c.PostForm("name"))
		if err != nil {
			c.Redirect(303, fmt.Sprintf("/locations?LocationFormErr=%s", err.Error()))
			return
		}
		err = locationModel.SetNumber(c.PostForm("number"))
		if err != nil {
			c.Redirect(303, fmt.Sprintf("/locations?LocationFormErr=%s", err.Error()))
			return
		}
		locationModel.SetUserID(userModel.ID)
		locations, err := database.GetLocationsByUserID(userModel.ID)
		if err != nil {
			log.Fatal(err.Error())
		}
		if len(locations) >= 3 {
			c.Redirect(303, fmt.Sprintf("/locations?LocationFormErr=%s", "only allowed 3 locations per user"))
			return
		}
		err = locationModel.Insert(database)
		if err != nil {
			log.Panic(err.Error())
		}
		c.Redirect(303, "/locations")
	})

	//==========================================================================
	// RUNNING
	//==========================================================================

	r.Run()

}

