package middleware

import (
	"context"
	"net/http"
	"os"
	"time"

	"miniWiki/internal/auth/model"
	"miniWiki/pkg/transport"

	"github.com/sirupsen/logrus"
)

type accKey struct{}
type sKey struct{}

var (
	sessionCookieName = "session_id"
	secure            = os.Getenv("ENV") == "production"
)

type authServiceInterface interface {
	GetSession(ctx context.Context, sessionId string) (*model.Session, error)
	Refresh(ctx context.Context, request model.RefreshRequest) (*model.SessionResponse, error)
	GetAccountStatus(ctx context.Context, accountId int) (*bool, error)
}

func SessionMiddleware(auth authServiceInterface) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sCookie, err := r.Cookie(sessionCookieName)

			if err != nil || sCookie.Value == "" {
				transport.Respond(w, http.StatusUnauthorized, nil)
				return
			}

			session, err := auth.GetSession(r.Context(), sCookie.Value)

			if err != nil {
				transport.Respond(w, http.StatusInternalServerError, nil)
				return
			}

			if session.UserAgent != r.UserAgent() {
				transport.Respond(w, http.StatusUnauthorized, nil)
				return
			}
			var res *model.SessionResponse
			if session.ExpireAt.Before(time.Now()) {
				res, err = auth.Refresh(r.Context(), model.RefreshRequest{
					AccountId: session.AccountID,
					SessionId: session.SessionID,
					IpAddress: r.RemoteAddr,
				})
				if err != nil {
					logrus.WithContext(r.Context()).Error(err)
					transport.HandleErrorResponse(w, err)
					return
				}

				SetSessionCookie(w, res.SessionId)
			}

			status, err := auth.GetAccountStatus(r.Context(), session.AccountID)
			if err != nil {
				logrus.WithContext(r.Context()).Error(err)
			}
			if !*status {
				transport.Respond(w, http.StatusUnauthorized, nil)
				return
			}

			ctx := SetAccountId(r.Context(), session.AccountID)
			if res != nil {
				ctx = SetSessionId(ctx, res.SessionId)
			} else {
				ctx = SetSessionId(ctx, session.SessionID)
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetSessionId(r *http.Request) string {
	return r.Context().Value(sKey{}).(string)
}

func SetSessionId(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, sKey{}, val)
}

func SetAccountId(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, accKey{}, val)
}

func GetAccountId(r *http.Request) int {
	return r.Context().Value(accKey{}).(int)
}

func LogoutSession(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:   sessionCookieName,
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})
}

func SetSessionCookie(w http.ResponseWriter, sessionId string) {
	http.SetCookie(w, &http.Cookie{
		Name:   sessionCookieName,
		Value:  sessionId,
		Path:   "/",
		Secure: secure,
	})
}
