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

	auth.NewAuth()
	auth.NewStore(os.Getenv("SECRET"))

	routes.SetupAuthRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
```
