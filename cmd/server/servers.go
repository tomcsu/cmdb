package main

import (
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/jinzhu/gorm"
	"github.com/seizadi/cmdb/pkg/pb"
	"github.com/seizadi/cmdb/pkg/svc"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func CreateServer(logger *logrus.Logger, db *gorm.DB, interceptors []grpc.UnaryServerInterceptor) (*grpc.Server, error) {
	// create new gRPC grpcServer with middleware chain
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(interceptors...)))
	
	// register all of our services into the grpcServer
	s, err := svc.NewBasicServer(db)
	if err != nil {
		return nil, err
	}
	pb.RegisterCmdbServer(grpcServer, s)

	artifact, err := svc.NewArtifactsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterArtifactsServer(grpcServer, artifact)

	vault, err := svc.NewVaultsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVaultsServer(grpcServer, vault)

	version_tag, err := svc.NewVersionTagsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterVersionTagsServer(grpcServer, version_tag)

	deployment, err := svc.NewDeploymentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterDeploymentsServer(grpcServer, deployment)

	environment, err := svc.NewEnvironmentsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterEnvironmentsServer(grpcServer, environment)

	kube_cluster, err := svc.NewKubeClustersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterKubeClustersServer(grpcServer, kube_cluster)

	manifest, err := svc.NewManifestsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterManifestsServer(grpcServer, manifest)

	application, err := svc.NewApplicationsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterApplicationsServer(grpcServer, application)

	aws_rds_instance, err := svc.NewAwsRdsInstancesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsRdsInstancesServer(grpcServer, aws_rds_instance)

	aws_service, err := svc.NewAwsServicesServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterAwsServicesServer(grpcServer, aws_service)

	container, err := svc.NewContainersServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterContainersServer(grpcServer, container)

	region, err := svc.NewRegionsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterRegionsServer(grpcServer, region)

	secret, err := svc.NewSecretsServer()
	if err != nil {
		return nil, err
	}
	pb.RegisterSecretsServer(grpcServer, secret)

	return grpcServer, nil
}