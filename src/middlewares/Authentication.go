package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/kVenkat-brs/Go-Assignment-2/src/db"
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success":false,
			"message":"No Api key found",
		})
	}
	usersCollection :=db.GetCollection("users")
	filter  := bson.M{
		"api_key": reqHeaderKey,
	}
	var userApiKey models.User
	if err:= usersCollection.FindOne(c.Context(),filter).Decode(&userApiKey);err!=nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"success":false,
			"message":"Invalid Api Key",
			"error":err,
		})
	}

	return c.Next()
	
}

func AuthorizationMiddleware(c fiber.Ctx) error {
	jwtsecret := []byte(os.Getenv("JWT_SECRET"))
	tokenString := c.Get("Authorization")

	if tokenString =="" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":false,
			"message":"No token found, Login again",
		})
		
	}
	token,err:= jwt.Parse(tokenString,func(t *jwt.Token) (interface{}, error) {return jwtsecret,nil})

	if err!=nil || !token.Valid {
		fmt.Println(err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":false,
			"message":"Unauthorized: invalid token!!",
		})
	}
	claims := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix())>claims["exp"].(float64){
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success":false,
			"message":"Token Expired!",
			
		})
	}

	return c.Next()
}