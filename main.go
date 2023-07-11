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
			"LoginFormErr": c.Query("LoginFormErr"),
		})
	})

	//==========================================================================

	r.GET("/signup", func(c *gin.Context) {
		_, err := c.Cookie(os.Getenv("SESSION_TOKEN_KEY"))
		if err == nil {
			c.Redirect(303, "/locations")
			return
		}
		c.HTML(200, "signup.html", gin.H{
			"SignupFormErr": c.Query("SignupFormErr"),
		})
	})

	//==========================================================================

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
		numberOfLocations := len(locations)
		hasNoLocations := numberOfLocations == 0
		hasMoreThanThreeLocations := numberOfLocations > 3
		c.HTML(200, "locations.html", gin.H{
			"LocationFormErr": c.Query("LocationFormErr"),
			"Locations": locations[:3],
			"HasMoreThanThreeLocations": hasMoreThanThreeLocations,
			"HasNoLocations": hasNoLocations,
		})
	})

	//==========================================================================

	r.GET("/locations/all", func(c *gin.Context) {
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
		c.HTML(200, "AllLocations.html", gin.H{
			"LocationFormErr": c.Query("LocationFormErr"),
			"Locations": locations,
		})
	})

	//==========================================================================


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

	//==========================================================================

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

	//==========================================================================

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

