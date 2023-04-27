package repository

import (
	"context"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/domain"
	"github.com/couchbase/gocb/v2"
)

type IUserRepository interface {
	Upsert(context context.Context, entity *domain.User) error
}

type userRepository struct {
	CollectDeliveryBucket  *gocb.Bucket
	CollectDeliveryCluster *gocb.Cluster
}

func NewUserRepository(cluster *gocb.Cluster) IUserRepository {
	return &userRepository{
		CollectDeliveryBucket:  cluster.Bucket(""),
		CollectDeliveryCluster: cluster,
	}
}
func (repository userRepository) Upsert(context context.Context, entity *domain.User) error {
	//TODO implement me
	panic("implement me")
}
