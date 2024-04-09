package auth

type contextKey string

const UserKey contextKey = "userID"

func CreateJWT(secret []byte, userID int ) (string,error){
	expiration:=time.Second*time.Duration(configs.Envs.JWTExpiration)
	token:=jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"userID":    strconv.Itoa(int(userID)),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})
	tokenString, err:=token.SignedString(secret)
	if err!=nil{
		return "",err
	}
	return tokenString,err
}