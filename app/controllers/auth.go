package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/auth-core/app/repositories"
	"github.com/auth-core/app/types"
	"github.com/auth-core/config"
)

func extractToken(c *gin.Context) (string, bool) {
	// token := c.Request.Header.Get("Authorization")

	var clientID int64
	var api_key = ""
	var expiredAt int64
	currentTimestamp := GetCurrentTimeTimeZoneUnix(config.MustGetString("server.time_zone"))

	token := c.Query("key")
	// api_key = config.MustGetString(GetRunMode() + ".api_key") // USING STATIC IN CONFIG FILE

	// VALIDATE AND EXRTRACT TOKEN
	hmacSecretString := "" + GetRunMode() + "" + ".secret_key" // Value
	hmacSecret := []byte(hmacSecretString)
	tokenString := token
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(tokens *jwt.Token) (interface{}, error) {
		return []byte(hmacSecret), nil
	})

	if err != nil {
		fmt.Println("####### TOKEN INVALID", err.Error())
		return "", false
	}

	// EXTRACT TOKEN
	// for key, val := range claims {
	// 	fmt.Printf("Key: %v, value: %v\n", key, val)
	// }

	if claims != nil {
		clientID = int64(claims["client_id"].(float64))
		expiredAt = int64(claims["expired_at"].(float64))
	}

	fmt.Println("TIME", expiredAt, currentTimestamp)

	if expiredAt < currentTimestamp {
		return "", false
	}

	api_key, _ = RedisGet(R_KEY_CLIENT + strconv.Itoa(int(clientID)))

	if api_key == "" {
		clientToken, err := repositories.GetClientToken(int32(clientID))
		if err != nil {
			return "", false
		}

		if clientToken == nil {
			return "", false
		}

		err = RedisSet(R_KEY_CLIENT+strconv.Itoa(int(clientID)), clientToken.Token.String)
		if err != nil {
			return "", false
		}

		return clientToken.Token.String, true
	}

	if token != api_key {
		return token, false
	}

	return api_key, true
}

func Authenticate(c *gin.Context) {
	token, exist := extractToken(c)

	if !exist {
		c.AbortWithStatus(http.StatusForbidden)
		// showResponseError(c, "403", "Forbidden")
		// Keep return after call abort to stop current handler
		return
	}

	fmt.Println(token, exist)
}

func GetToken(c *gin.Context) {

	var requestGetToken types.RequestGetToken

	if err := c.BindJSON(&requestGetToken); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse)
	}

	// CREATE TOKEN
	token, expiredAt, err := CreateToken(&requestGetToken)
	if err != nil {
		fmt.Println(err.Error())
		gr := GeneralResponseSuccessBuild(time.Now().UTC(), false, "8007", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"general_response": gr,
		})
		return
	}

	var data types.ResponseGetToken
	data.Token = token
	data.ExpiredAt = int64(expiredAt)

	gr := GeneralResponseSuccessBuild(time.Now().UTC(), true, "000", "Success")

	c.JSON(http.StatusOK, gin.H{
		"general_response": gr,
		"data":             &data,
	})

	return
}

func CreateToken(data *types.RequestGetToken) (string, int, error) {

	client, err := repositories.GetClientByUsername(data)
	if err != nil {
		return "", 0, err
	}

	if client == nil {
		return "", 0, errors.New("Client not Found")
	}

	// CHECKING PASSWORD BCRYPT
	match := CheckPasswordHash(data.Password, client.Password.String)
	if match == false {
		return "", 0, errors.New("Wrong password please try again")
	}

	newTime := GetCurrentTimestampTimeZone(config.MustGetString("server.time_zone"))

	// afterDateTime := newTime.AddDate(100, 0, 0)
	afterDateTime := newTime.Local().AddDate(1, 0, 0) // ADD ONE DAY EXPIRED
	// afterDateTime := newTime.Local().Add(30 * time.Minute) // ADD TIME 30 MINUTES

	clientID := client.ID
	issuer := client.Name
	signingKey := []byte(GetRunMode() + ".secret_key")

	clientToken, err := repositories.GetClientToken(int32(clientID.Int64))
	if err != nil {
		return "", 0, err
	}

	if clientToken != nil {
		var updateClientToken types.UpdateClientToken
		updateClientToken.ClientID = clientToken.ClientID.Int64
		updateClientToken.DeletedAt = newTime
		updateClientToken.UpdatedAt = newTime
		updateClientToken.UpdatedBy = clientToken.ClientID.Int64

		err = repositories.UpdateClientToken(&updateClientToken)
		if err != nil {
			return "", 0, err
		}
	}

	type MyCustomClaims struct {
		ClientID  int64 `json:"client_id"`
		ExpiredAt int64 `json:"expired_at"`
		jwt.StandardClaims
	}

	// Create the Claims
	claims := MyCustomClaims{
		clientID.Int64,
		afterDateTime.Unix(),
		jwt.StandardClaims{
			ExpiresAt: 0,
			Issuer:    issuer.String,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	err = RedisSet(R_KEY_CLIENT+strconv.Itoa(int(clientID.Int64)), ss)
	if err != nil {
		return "", 0, err
	}

	var insertClientToken types.InsertClientToken
	insertClientToken.ClientID = clientID.Int64
	insertClientToken.Token = ss
	insertClientToken.ExpiredAt = afterDateTime.Unix()
	insertClientToken.CreatedBy = clientID.Int64
	insertClientToken.CreatedAt = newTime

	err = repositories.InsertClientToken(&insertClientToken)
	if err != nil {
		return "", 0, err
	}

	// GET DATA REDIS
	// value, err := RedisGet(R_KEY_CLIENT + strconv.Itoa(int(clientID.Int64)))
	// if err != nil {
	// 	return "", 0, err
	// }
	// fmt.Println("####", value)

	return ss, int(afterDateTime.Unix()), nil
}
