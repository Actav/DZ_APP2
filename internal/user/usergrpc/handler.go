package usergrpc

import (
	"context"
	"github.com/google/uuid"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/internal/database"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.UserServiceServer = (*Handler)(nil)

func New(usersRepository usersRepository, timeout time.Duration) *Handler {
	return &Handler{usersRepository: usersRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedUserServiceServer
	usersRepository usersRepository
	timeout         time.Duration
}

func (h Handler) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	creationData := database.CreateUserReq{
		ID:       id,
		Username: in.Username,
		Password: in.Password,
	}
	if _, err = h.usersRepository.Create(ctx, creationData); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.User, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := h.usersRepository.FindByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	gettingData := &pb.User{
		Id:        resp.ID.String(),
		Username:  resp.Username,
		Password:  resp.Password,
		CreatedAt: resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt: resp.UpdatedAt.Format(time.RFC3339),
	}

	return gettingData, nil
}

func (h Handler) UpdateUser(ctx context.Context, in *pb.UpdateUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	creationData := database.CreateUserReq{
		ID:       id,
		Username: in.Username,
		Password: in.Password,
	}
	if _, err = h.usersRepository.Create(ctx, creationData); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := uuid.Parse(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := h.usersRepository.DeleteByUserID(ctx, id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) ListUsers(ctx context.Context, in *pb.Empty) (*pb.ListUsersResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	resp, err := h.usersRepository.FindAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	users := make([]*pb.User, 0, len(resp))
	for _, user := range resp {
		u := pb.User{
			Id:        user.ID.String(),
			Username:  user.Username,
			Password:  user.Password,
			CreatedAt: user.CreatedAt.Format(time.RFC3339),
			UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
		}
		users = append(users, &u)
	}

	return &pb.ListUsersResponse{Users: users}, nil
}
