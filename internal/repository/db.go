package repository

import (
	"context"
	"github.com/jackc/pgx/v5"
	env "tax-api/internal"
)

var db, err = pgx.Connect(context.Background(), env.GetDBUrlEnv())

// todo check out guide
