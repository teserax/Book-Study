package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func GetByURL(useCase service.UseCase, logLevel log.Level, handlerMetrics *metrics.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger := log.NewDefault(logLevel).WithField(helper.HeaderTraceIDName, r.Header.Get(helper.HeaderTraceIDName)).WithField("url", r.URL.String())
		logger.Debugf("=================================================")
		logger.Debugf("Start process request %s", r.URL.String())
		// detect "method"
		handlerName, err := helper.DetectHandler(mux.CurrentRoute(r))
		if err != nil {
			handlerName = helper.UnknownHeaderName
		}
		// validate method.
		logger.Debugf("Validate request method: %s", r.Method)
		if r.Method != http.MethodGet {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusMethodNotAllowed)).Inc()
			logger.Errorf(helper.ErrHTTPMethod, r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Write([]byte(fmt.Sprintf(helper.ErrHTTPMethod, r.Method)))
			return
		}

		// validate URL.
		vars := mux.Vars(r)
		_, existHashtype := vars["hashtype"]
		_, existURL := vars["url"]
		logger.Debugf("Validate request parameters. Hashtype exist: %t, URL exist: %t", existHashtype, existURL)
		if !existHashtype || !existURL {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusBadRequest)).Inc()
			logger.Errorf(helper.ErrHTTPURLFormat, r.URL.String())
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(helper.ErrHTTPURLFormat, r.URL.String())))
			return
		}

		// check if exist hashtype.
		logger.Debugf("Check if exist hashtype")
		var hashtype image.HashType
		hashtype, err = image.HashTypeByString(vars["hashtype"])
		if err != nil {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusBadRequest)).Inc()
			logger.Errorf(helper.ErrHashTypeFormat, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(helper.ErrHashTypeFormat, err)))
			return
		}

		// check if URL is valid
		logger.Debugf("Check parameter url")
		if _, err := url.Parse(vars["url"]); err != nil {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusBadRequest)).Inc()
			logger.Errorf(helper.ErrHTTPURLFormat, err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(helper.ErrHTTPURLFormat, err)))
			return
		} else if vars["url"] == "" {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusBadRequest)).Inc()
			logger.Errorf(helper.ErrHTTPURLEmpty)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf(helper.ErrHTTPURLEmpty)))
			return
		}

		// process request logic.
		logger.Debugf("Process request logic")
		ctx := context.WithValue(r.Context(), loaderHelper.ContextKey(loaderHelper.Headers), r.Header)
		hashInfo, err := useCase.GetHashByURL(ctx, vars["url"], hashtype, logger)
		if err != nil {
			// process apart NotFound
			if errors.Is(err, loader.ErrNotFound) {
				handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusNotFound)).Inc()
				// logger.Errorf(helper.ErrNotFound, err)
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf(helper.ErrHashByURL, err)))
				return
			} else if errors.Is(err, loader.ErrNotGotContent) {
				handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusNoContent)).Inc()
				logger.Errorf("Cant get content, err: %w", err)
				w.WriteHeader(http.StatusNoContent)
				w.Write([]byte(fmt.Sprintf(helper.ErrHashByURL, err)))
				return
			}
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusInternalServerError)).Inc()
			logger.Errorf(helper.ErrHashByURL, err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf(helper.ErrHashByURL, err)))
			return
		}

		// Validate hashInfo
		logger.Debugf("Validate hashInfo")
		if err := hashInfo.Validate(); err != nil {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusInternalServerError)).Inc()
			logger.Errorf("Error: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// marshal response.
		logger.Debugf("Marshal response")
		hashInfoJSON, err := json.Marshal(hashInfo)
		if err != nil {
			handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusInternalServerError)).Inc()
			logger.Errorf("Error: %w", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		logger.Debugf("Return response. StatusCode: %d, hash: %s", http.StatusOK, string(hashInfoJSON))
		w.Header().Set("Content-Type", "application/json")
		handlerMetrics.Processed.WithLabelValues(handlerName, strconv.Itoa(http.StatusOK)).Inc()
		w.WriteHeader(http.StatusOK)
		w.Write(hashInfoJSON)
	}
}
