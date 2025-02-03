package auth

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/helpers"
	model "github.com/mahirjain_10/stock-alert-app/backend/internal/models"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
	"github.com/mahirjain_10/stock-alert-app/backend/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context, r *gin.Engine, app *types.App) {
	var user types.RegisterUser

	// Bind the incoming JSON request to 'user' struct
	if !helpers.BindAndValidateJSON(c, &user) {
		return
	}

	// DEBUG : Printing user and user.ID
	fmt.Println(user.ID)
	fmt.Println(user)

	vError := validator.ValidateRegisterUser(user)
	fmt.Println("error ", vError)

	fmt.Println(len(vError))
	if len(vError) != 0 {
		helpers.SendResponse(c, http.StatusBadRequest, "Validation error", nil, vError, false)
		return
	}
	// Check if a user with the given email already exists
	// retrievedUser, err := model.FindUserByEmail(app, user.Email)
	// if err != nil {
	// 	helpers.SendResponse(c, http.StatusInternalServerError, "Error while checking user existence", nil, nil, false)
	// 	return
	// }

	// if retrievedUser.ID != "" {
	// 	helpers.SendResponse(c, http.StatusConflict, "User already exists with given email", nil, nil, false)
	// 	return
	// }

	// Set a new UUID for the user
	user.ID = uuid.New().String()
	fmt.Printf("Generated UUID: %s\n", user.ID)

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Error while hashing the password: %v\n", err)
		helpers.SendResponse(c, http.StatusInternalServerError, "Internal server error during password hashing", nil, nil, false)
		return
	}
	user.Password = string(hashedPassword)

	// Save the new user to the database
	err = model.InsertUser(app, user)
	if err != nil {
		if strings.Contains(err.Error(), "Error 1062 (23000): Duplicate entry") {
			helpers.SendResponse(c, http.StatusConflict, "User already exists with given email", nil, nil, false)
			return
		}
		helpers.SendResponse(c, http.StatusInternalServerError, "Error while saving the user", nil, nil, false)
		return
	}

	// Respond with success if the user was created
	helpers.SendResponse(c, http.StatusCreated, "User account created successfully", nil, nil, true)

}

func LoginUser(c *gin.Context, r *gin.Engine, app *types.App) {
    var user types.LoginUser
    
    // Unmarshall JSON
    if !helpers.BindAndValidateJSON(c,&user){
        return
    }

    // Validation of user object
	vError := validator.ValidateLoginUser(user)
    if len(vError) != 0 {
		helpers.SendResponse(c, http.StatusBadRequest, "Validation error", nil, vError, false)
		return
	}

    // Find user with the help of Email
    retrievedUser, err := model.FindUserByEmail(app, user.Email)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			helpers.SendResponse(c, http.StatusNotFound, "User with given email not found , Please create a new account", nil, nil, false)
			return
		}
		helpers.SendResponse(c, http.StatusInternalServerError, "Error while checking user existence", nil, nil, false)
		return
	}

    // Send 404 response if user is not found
    if retrievedUser.ID == "" {
		helpers.SendResponse(c, http.StatusNotFound, "User with given email not found , Please create a new account", nil, nil, false)
		return
	}

    // Compare Password
    err = bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password),[]byte(user.Password)); 
    if err !=nil {
        if err == bcrypt.ErrMismatchedHashAndPassword {
            // Handle incorrect password
            helpers.SendResponse(c, http.StatusUnauthorized, "Incorrect password", nil, nil, false)
            return
        }
        // Handle other potential errors 
        helpers.SendResponse(c, http.StatusInternalServerError, "Error comparing passwords", nil, nil, false)
        return
    }
    data := map[string]interface{}{
       "id":retrievedUser.ID,
    }
    helpers.SendResponse(c,http.StatusOK,"User logged in successfully",data,nil,true)
    
}
