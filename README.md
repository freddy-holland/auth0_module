## **USAGE**

### Install:
`go get github.com/freddy-holland/auth0_module`

### Implementation:
```
server.go

package main

import (
	"os"

	"github.com/freddy-holland/auth0_module/auth"
	"github.com/freddy-holland/auth0_module/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	godotenv.Load(".env")
	e := echo.New()

	authCfg := auth.AuthConfig{
		Auth0Domain: os.Getenv("AUTH0_DOMAIN"),
		Auth0ClientID: os.Getenv("AUTH0_CLIENT_ID"),
		Auth0ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Auth0CallbackURL: os.Getenv("AUTH0_CALLBACK_URL"),
	}

	auth.NewAuth()
	auth.NewStore(os.Getenv("STORE_SECRET"))

	routes.SetupAuthRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
```
