package usecase

import "context"

type Repository interface {
	GetFooBar(ctx context.Context) (res []int64, err error)
}
type Cache interface {
}
