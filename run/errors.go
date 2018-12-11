package server

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrCreateDomainFailed   = status.New(codes.Internal, "Creation failed")
	ErrDeleteDomainFailed   = status.New(codes.Internal, "Deletion failed")
	ErrGenerationUUIDFailed = status.New(codes.Internal, "UUID creation failed")
	ErrListDomainsFailed    = status.New(codes.Internal, "List domains failed")
)

var (
	ErrFailedGetCA = errors.New("failed to get CA")
)
