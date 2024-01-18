package middleware

import (
	"context"
	"net/http"
	"strings"

	"CodeWithAzri/pkg/response"

	firebase "firebase.google.com/go"
	firebaseAuth "firebase.google.com/go/auth"
)

type UserIDKey string

const (
	UserIDContextKey UserIDKey = "UserID"
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

// AuthMiddleware is middleware to handle Firebase authentication.
func (fa *FirebaseMiddleware) AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		client, err := initializeFirebaseAuthClient(fa.FirebaseApp)
		if err != nil {
			handleFirebaseAuthInitializationError(w)
			return
		}

		idToken := r.Header.Get("Authorization")
		if idToken == "" {
			handleAuthorizationHeaderMissingError(w)
			return
		}

		tokenParts := strings.Split(idToken, "Bearer ")
		if len(tokenParts) != 2 {
			handleInvalidTokenError(w)
			return
		}
		idToken = tokenParts[1]

		decoded, err := verifyIDToken(client, idToken)
		if err != nil {
			handleInvalidTokenError(w)
			return
		}

		if decoded == nil || decoded.UID == "" {
			handleInvalidTokenError(w)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDContextKey, decoded.UID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
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
func handleFirebaseAuthInitializationError(w http.ResponseWriter) {
	response.RespondErrorMessage(
		http.StatusInternalServerError,
		"Firebase Auth client initialization failed",
		w,
	)
}

// HandleAuthorizationHeaderMissingError handles missing authorization header errors.
func handleAuthorizationHeaderMissingError(w http.ResponseWriter) {
	response.RespondErrorMessage(
		http.StatusUnauthorized,
		"Authorization header is required",
		w,
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
func handleInvalidTokenError(w http.ResponseWriter) {
	response.RespondErrorMessage(
		http.StatusUnauthorized,
		"Token Invalid",
		w,
	)
}
