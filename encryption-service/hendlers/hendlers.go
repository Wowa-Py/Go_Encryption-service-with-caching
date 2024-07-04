package handlers

import (
	"context"
	"encoding/json"
	"encryption-service/config"
	"encryption-service/utils"
	"net/http"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Request struct {
	Input string `json:"input"`
}

type Response struct {
	MD5Hash    string `json:"md5"`
	SHA256Hash string `json:"sha256"`
}

func EncryptHandler(cfg config.Config) http.HandlerFunc {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
	})

	return func(w http.ResponseWriter, r *http.Request) {
		var req Request
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		cacheKey := "hash:" + req.Input
		val, err := rdb.Get(ctx, cacheKey).Result()
		if err == redis.Nil {
			md5Hash := utils.HashMD5(req.Input)
			sha256Hash := utils.HashSHA256(req.Input)

			res := Response{MD5Hash: md5Hash, SHA256Hash: sha256Hash}
			resJSON, _ := json.Marshal(res)

			rdb.Set(ctx, cacheKey, resJSON, 10*time.Minute)

			w.Header().Set("Content-Type", "application/json")
			w.Write(resJSON)
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(val))
		}
	}
}
