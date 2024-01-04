package middleware

import (
	"context"
	"net/http"
	"strings"

	"CodeWithAzri/pkg/response"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
)

// FirebaseMiddleware represents Firebase middleware.
type FirebaseMiddleware struct {
	FirebaseApp *firebase.App
}

// NewFirebaseMiddleware creates a new FirebaseMiddleware instance.
func NewFirebaseMiddleware(firebaseApp *firebase.App) *FirebaseMiddleware {
	return &FirebaseMiddleware{
		FirebaseApp: firebaseApp,
	}
}

// FirebaseAuthMiddleware is middleware to handle Firebase authentication.
func (fa *FirebaseMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		client, err := initializeFirebaseAuthClient(fa.FirebaseApp)
		if err != nil {
			handleFirebaseAuthInitializationError(ctx)
			return
		}

		idToken := ctx.GetHeader("Authorization")
		if idToken == "" {
			handleAuthorizationHeaderMissingError(ctx)
			return
		}

		tokenParts := strings.Split(idToken, "Bearer ")
		if len(tokenParts) != 2 {
			handleInvalidTokenError(ctx)
			return
		}
		idToken = tokenParts[1]

		decoded, err := verifyIDToken(client, idToken)
		if err != nil {
			handleInvalidTokenError(ctx)
			return
		}

		ctx.Set("UserID", decoded.UID)
	}
}

// InitializeFirebaseAuthClient initializes the Firebase Auth client.
func initializeFirebaseAuthClient(app *firebase.App) (*firebaseAuth.Client, error) {
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}

// HandleFirebaseAuthInitializationError handles Firebase Auth initialization errors.
func handleFirebaseAuthInitializationError(ctx *gin.Context) {
	response.RespondErrorMessage(
		http.StatusInternalServerError,
		"Firebase Auth client initialization failed",
		ctx,
	)
}

// HandleAuthorizationHeaderMissingError handles missing authorization header errors.
func handleAuthorizationHeaderMissingError(ctx *gin.Context) {
	response.RespondErrorMessage(
		http.StatusUnauthorized,
		"Authorization header is required",
		ctx,
	)
}

// VerifyIDToken verifies the ID token using Firebase Auth client.
func verifyIDToken(client *firebaseAuth.Client, idToken string) (*firebaseAuth.Token, error) {
	decoded, err := client.VerifyIDToken(context.Background(), idToken)
	if err != nil {
		return nil, err
	}
	return decoded, nil
}

// HandleInvalidTokenError handles invalid token errors.
func handleInvalidTokenError(ctx *gin.Context) {
	response.RespondErrorMessage(
		http.StatusUnauthorized,
		"Token Invalid",
		ctx,
	)
}
