package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/kVenkat-brs/Go-Assignment-2/src/db"
	utils "github.com/kVenkat-brs/Go-Assignments-3/src/Utils"
	"github.com/kVenkat-brs/Go-Assignments-3/src/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func RequestTimingMiddleware(c fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	duration := time.Since(start)

	fmt.Printf("%s %s took %d ms\n",
		c.Method(),
		c.Path(),
		duration.Milliseconds(),
	)

	return err
}

func ApiKeyMiddleware(c fiber.Ctx)error  {

	reqHeaderKey := c.Get("x-api-key")
	if reqHeaderKey == "" {
		return utils.Error(c,fiber.StatusUnauthorized,"No Api Key found")
	}
	usersCollection :=db.GetCollection("users")
	filter  := bson.M{
		"api_key": reqHeaderKey,
	}
	var userApiKey models.User
	if err:= usersCollection.FindOne(c.Context(),filter).Decode(&userApiKey);err!=nil{
		return utils.Error(c,fiber.StatusUnauthorized,"Invalid Api Key")
	}

	return c.Next()
	
}

func AuthorizationMiddleware(c fiber.Ctx) error {
	jwtsecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString := c.Get("Authorization")

	if tokenString =="" {
		return utils.Error(c,fiber.StatusUnauthorized,"No Token Found, Login Again!!")
		
		
	}
	token,err:= jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {return jwtsecret,nil})

	if err!=nil || !token.Valid {
		fmt.Println(err)
		return utils.Error(c,fiber.StatusUnauthorized,"invalid Token!!")
	}
	claims := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix())>claims["exp"].(float64){
		return utils.Error(c,fiber.StatusUnauthorized,"Token Expired!!")
	}

	return c.Next()
}