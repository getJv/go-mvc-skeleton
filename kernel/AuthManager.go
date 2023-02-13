package kernel

import (
	"fmt"
	"getjv/backend/app/http/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func GeneratePassword(rawPass string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(rawPass),10)
}

func ComparePassword(storedPass string,incomingPass string) error{
	return bcrypt.CompareHashAndPassword([]byte(storedPass),[]byte(incomingPass) )
}

func GenerateToken(user models.User) (string,error){
	
	hours,errGetMin := time.ParseDuration(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT")) 
	if errGetMin != nil{
		return "",errGetMin
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * hours).Unix(),
	})

	tokenString,err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil{
		return "",err
	}
	return tokenString,nil

}

func SaveAuthCookie(c *fiber.Ctx, tokenString string) error {

	timeLife,err := time.ParseDuration(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))
	if err != nil {
		return err
	}
	
	c.Cookie(&fiber.Cookie{
        Name:     "Authorization",
        Value:   tokenString,
        Expires:  time.Now().Add(timeLife * time.Hour),
        HTTPOnly: true,
        SameSite: "lax",
    })

	return nil

}



func RetrieveUserFromCookie(c *fiber.Ctx) (int,error){
	// Get auth cookie
	authCookie := c.Cookies("Authorization") 
	if authCookie == ""{
		return 0,nil
	}
	// verify alg
	token, err := jwt.Parse(authCookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})
	// verify is Valid Token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0,err
	} 

	// retrieve data from token
	userId := claims["sub"]
	expiration := claims["exp"]

	// Verify if it is expired
	if float64(time.Now().Unix()) > expiration.(float64){
		return 0,fmt.Errorf("Expired token")
	}

	return int(userId.(float64)),nil
}

func GetAuthUser(c *fiber.Ctx) (models.User,error){
	user := c.Locals("user")
	if user == nil{
		return models.User{},fmt.Errorf("No User Found")
	}

	return user.(models.User),nil
		
}