package deps

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type Deps struct {
	DB  *pgxpool.Pool
	RDB *redis.Client
	S3  *s3.Client
}

func New(db *pgxpool.Pool, rdb *redis.Client, s3Client *s3.Client) *Deps {
	return &Deps{
		DB:  db,
		RDB: rdb,
		S3:  s3Client,
	}
}

func (d *Deps) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.RDB != nil {
		d.RDB.Close()
	}
}
