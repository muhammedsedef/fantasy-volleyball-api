package repository

import (
	"context"
	"fantasy-volleyball-api/internal/fantasy-volleyball-api-command-service/domain"
	logger "fantasy-volleyball-api/pkg/log"
	"github.com/couchbase/gocb/v2"
	"reflect"
)

type IUserRepository interface {
	Upsert(context context.Context, entity *domain.User) error
}

type userRepository struct {
	CollectDeliveryBucket  *gocb.Bucket
	CollectDeliveryCluster *gocb.Cluster
	logger                 logger.Logger
}

func NewUserRepository(cluster *gocb.Cluster) IUserRepository {
	return &userRepository{
		CollectDeliveryBucket:  cluster.Bucket(""),
		CollectDeliveryCluster: cluster,
		logger:                 logger.GetLogger(reflect.TypeOf((*userRepository)(nil))),
	}
}
func (repository userRepository) Upsert(ctx context.Context, entity *domain.User) error {
	repository.logger.InfoWithContext(ctx, "userRepository.Upsert - Started - entity: %#v", entity)

	return nil
}
