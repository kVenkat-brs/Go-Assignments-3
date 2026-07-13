package controllers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/kVenkat-brs/Go-Assignment-2/src/db"
	utils "github.com/kVenkat-brs/Go-Assignments-3/src/Utils"
	"github.com/kVenkat-brs/Go-Assignments-3/src/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"golang.org/x/crypto/bcrypt"
)

//Helper function to Genarating Api key
func GenereateAPIKey()(string,error){
	bytes := make([]byte,34)

	_,err := rand.Read(bytes)
	if err!=nil {
		return "",err
	}

	return hex.EncodeToString(bytes),nil
}

// For Creating User
func CreateUser(c fiber.Ctx) error {

	var request models.User

	if err :=c.Bind().Body(&request);err!=nil{

		return utils.Error(c,fiber.StatusBadRequest,"Invalid Body")
	}

	users :=db.GetCollection("users")
	var existingUser models.User 

	if err := users.FindOne(c.Context(),bson.M{"email":request.Email}).Decode(&existingUser);err == nil{
			 return utils.Error(c,fiber.StatusConflict,"User Already Exists!")
	}

	hashedPassword,err:= bcrypt.GenerateFromPassword([]byte(request.Password),14)
	if err!=nil {
		return utils.Error(c,fiber.StatusInternalServerError,"Invalid Password!!")
	}

	request.Password = string(hashedPassword)

	result, err := users.InsertOne(c.Context(),request)

	if err!=nil {
		return utils.Error(c,fiber.StatusInternalServerError,"Unable to create User")
		
	}

	request.Id = result.InsertedID.(bson.ObjectID)

	// var user models.User

	// if err:=users.FindOne(c.Context(),bson.M{"_id":result.InsertedID}).Decode(&user);err!=nil{
	// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	// 	"success": false,
	// 	"message":"No user exixts",
	// 	"error":err,
		
	// 	})
	// }




	return utils.Success(c,fiber.StatusCreated,"User created successfully",request)




}

// User Login 


func Login(c fiber.Ctx) error{
	jwtsecret:=[]byte(os.Getenv("JWT_SECRET"))
	var body models.User

	if err := c.Bind().Body(&body);err!=nil{
		return utils.Error(c,fiber.StatusBadRequest,"Cannot Parse body")
	}

	users :=db.GetCollection("users")
	var existingUser models.User 

	if err := users.FindOne(c.Context(),bson.M{"email":body.Email}).Decode(&existingUser);err != nil{
			return utils.Error(c,fiber.StatusNotFound,"User not found")
	}


	if err:= bcrypt.CompareHashAndPassword([]byte(existingUser.Password),[]byte(body.Password));err!=nil{
		return utils.Error(c,fiber.StatusUnauthorized,"Invalid Password, try again with correct one!!")
	}

	ApiKey, err :=GenereateAPIKey()
	if err!=nil {
		return utils.Error(c,fiber.StatusInternalServerError,"Unable to Generate API Key!")
	}

	filter := bson.M{
		"_id":existingUser.Id,
	}

	
	existingUser.X_API_key =ApiKey

	token:= jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userId":existingUser.Id.Hex(),
		"exp":time.Now().Add(1*time.Hour).Unix(),
	})

	t,err := token.SignedString(jwtsecret)

	if err!=nil {
		fmt.Println(err)
		
		return utils.Error(c,fiber.StatusInternalServerError,"Unable to Create a Token")
		
	}

	result,err := users.UpdateOne(c.Context(),filter,bson.M{"$set":existingUser})

	if err!=nil {
		return utils.Error(c,fiber.StatusInternalServerError,"Unable to update Api key!!")
	}
	if result.MatchedCount == 0 {
		return utils.Error(c,fiber.StatusNotFound,"User Not Found")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":true,
		"message":"Login successful and Token created successfully",
		"token":t,
		"user name":existingUser.Name,
		"user_Api_key":existingUser.X_API_key,
	})



}