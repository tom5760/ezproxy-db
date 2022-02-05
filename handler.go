package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

const (
	responseCompressionLevel = 5

	headerAcceptEncoding            = "Accept-Encoding"
	headerContentTypeOptions        = "X-Content-Type-Options"
	headerCrossOriginEmbedderPolicy = "Cross-Origin-Embedder-Policy"
	headerCrossOriginOpenerPolicy   = "Cross-Origin-Opener-Policy"
	headerCrossOriginResourcePolicy = "Cross-Origin-Resource-Policy"
	headerFrameOptions              = "X-Frame-Options"
	headerRequestID                 = "X-Request-ID"
	headerVary                      = "Vary"

	contentTypeOptionsNoSniff = "nosniff"

	crossOriginEmbedderPolicyRequireCORP = "require-corp"

	crossOriginOpenerPolicySameOrigin = "same-origin"

	crossOriginResourcePolicySameOrigin  = "same-origin"
	crossOriginResourcePolicyCrossOrigin = "cross-origin"

	frameOptionsDeny = "DENY"
)

func newHandler(lg zerolog.Logger) http.Handler {
	mux := chi.NewMux()

	mux.Use(
		hlog.NewHandler(lg),

		recoverMiddleware,

		hlog.RequestIDHandler("request_id", headerRequestID),

		middleware.RealIP,

		hlog.RemoteAddrHandler("remote_addr"),
		hlog.MethodHandler("method"),
		hlog.URLHandler("url"),
		hlog.UserAgentHandler("user_agent"),
		hlog.AccessHandler(logAccessFn),

		middleware.CleanPath,
		middleware.Compress(responseCompressionLevel),

		middleware.SetHeader(headerVary, headerAcceptEncoding),

		// Some security-related headers.  See: https://web.dev/security-headers/
		// Generally default to the strictest settings.
		middleware.SetHeader(headerContentTypeOptions, contentTypeOptionsNoSniff),
		middleware.SetHeader(headerFrameOptions, frameOptionsDeny),
		middleware.SetHeader(headerCrossOriginOpenerPolicy, crossOriginOpenerPolicySameOrigin),
		middleware.SetHeader(headerCrossOriginEmbedderPolicy, crossOriginEmbedderPolicyRequireCORP),

		// HSTS will be set in nginx

		// This is overridden in API handlers that _are_ cross origin.
		middleware.SetHeader(headerCrossOriginResourcePolicy, crossOriginResourcePolicySameOrigin),
	)

	mux.MethodFunc(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return mux
}
