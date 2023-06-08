package cache

import "github.com/go-redis/redis/v8"

func (db *db) GetDB() any {
	return db.getClient()
}

func (db *db) getClient() *redis.Client {
	return db.redis.GetClient()
}
