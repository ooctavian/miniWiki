package middleware

import (
	"context"
	"net/http"
	"time"

	sQuery "miniWiki/domain/auth/query"
	"miniWiki/utils"
)

var (
	sKey   = "session_id"
	accKey = "acc_id"
)

func SessionMiddleware(querier sQuery.Querier) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			sCookie, err := r.Cookie(sKey)

			if err != nil || sCookie.Value == "" {
				utils.Respond(w, http.StatusUnauthorized, nil)
				return
			}
			session, err := querier.GetSession(r.Context(), sCookie.Value)

			if err != nil {
				utils.Respond(w, http.StatusInternalServerError, nil)
				return
			}

			if session.ExpireAt.Before(time.Now()) ||
				*session.UserAgent != r.UserAgent() {
				utils.Respond(w, http.StatusUnauthorized, nil)
				return
			}

			ctx := SetSessionId(r.Context(), session.SessionID)
			ctx = SetAccountId(ctx, *session.AccountID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetSessionId(r *http.Request) string {
	return r.Context().Value(sKey).(string)
}

func SetSessionId(ctx context.Context, val string) context.Context {
	return context.WithValue(ctx, sKey, val)
}

func SetAccountId(ctx context.Context, val int) context.Context {
	return context.WithValue(ctx, accKey, val)
}

func GetAccountId(r *http.Request) int {
	return r.Context().Value(accKey).(int)
}

func LogoutSession(w http.ResponseWriter) {
	c := &http.Cookie{
		Name:     sKey,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
	}

	http.SetCookie(w, c)
}
