package auth

import (
	"context"
	ssov1 "github.com/DeMarDeXis/AuthProto/gen/go/sso"
	"sso/internal/domain/models"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Auth interface {
	GetToken(ctx context.Context, user models.UserToken) (string, error)
}

type serverAPI struct {
	ssov1.UnimplementedAuthServer
	auth Auth
}

func Register(gRPC *grpc.Server, auth Auth) {
	ssov1.RegisterAuthServer(gRPC, &serverAPI{auth: auth})
}

func (s *serverAPI) GetToken(
	ctx context.Context,
	req *ssov1.GetTokenRequest,
) (*ssov1.GetTokenResponse, error) {
	if err := validateID(req); err != nil {
		return nil, err
	}

	token, err := s.auth.GetToken(ctx, models.UserToken{UserID: req.GetUserId()})
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &ssov1.GetTokenResponse{
		Token: token,
	}, nil
}

// validateID validates the request ID. If it's == 0, returns an error.
func validateID(req *ssov1.GetTokenRequest) error {
	if req.GetUserId() == 0 {
		return status.Error(codes.InvalidArgument, "invalid user id")
	}
	return nil
}
