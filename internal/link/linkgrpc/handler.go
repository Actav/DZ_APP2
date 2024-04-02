package linkgrpc

import (
	"context"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/internal/database"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
)

var _ pb.LinkServiceServer = (*Handler)(nil)

func New(linksRepository linksRepository, timeout time.Duration) *Handler {
	return &Handler{linksRepository: linksRepository, timeout: timeout}
}

type Handler struct {
	pb.UnimplementedLinkServiceServer
	linksRepository linksRepository
	timeout         time.Duration
}

func (h Handler) GetLinkByUserID(ctx context.Context, id *pb.GetLinksByUserId) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	resp, err := h.linksRepository.FindByUserID(ctx, id.UserId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	links := make([]*pb.Link, 0, len(resp))
	for _, link := range resp {
		l := pb.Link{
			Id:        link.ID.String(),
			Title:     link.Title,
			Url:       link.URL,
			Images:    link.Images,
			Tags:      link.Tags,
			UserId:    link.UserID,
			CreatedAt: link.CreatedAt.Format(time.RFC3339),
			UpdatedAt: link.UpdatedAt.Format(time.RFC3339),
		}
		links = append(links, &l)
	}

	return &pb.ListLinkResponse{Links: links}, nil
}

func (h Handler) mustEmbedUnimplementedLinkServiceServer() {
	// TODO implement me
	panic("implement me")
}

func (h Handler) CreateLink(ctx context.Context, request *pb.CreateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	creationData := database.CreateLinkReq{
		ID:     id,
		URL:    request.Url,
		Title:  request.Title,
		Tags:   request.Tags,
		Images: request.Images,
		UserID: request.UserId,
	}
	if _, err = h.linksRepository.Create(ctx, creationData); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) GetLink(ctx context.Context, request *pb.GetLinkRequest) (*pb.Link, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err := h.linksRepository.FindByID(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	gettingData := &pb.Link{
		Id:        resp.ID.String(),
		Title:     resp.Title,
		Url:       resp.URL,
		Images:    resp.Images,
		Tags:      resp.Tags,
		UserId:    resp.UserID,
		CreatedAt: resp.CreatedAt.Format(time.RFC3339),
		UpdatedAt: resp.UpdatedAt.Format(time.RFC3339),
	}

	return gettingData, nil
}

func (h Handler) UpdateLink(ctx context.Context, request *pb.UpdateLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	creationData := database.CreateLinkReq{
		ID:     id,
		URL:    request.Url,
		Title:  request.Title,
		Tags:   request.Tags,
		Images: request.Images,
		UserID: request.UserId,
	}
	if _, err = h.linksRepository.Create(ctx, creationData); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) DeleteLink(ctx context.Context, request *pb.DeleteLinkRequest) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(request.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := h.linksRepository.Delete(ctx, id); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.Empty{}, nil
}

func (h Handler) ListLinks(ctx context.Context, request *pb.Empty) (*pb.ListLinkResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeout)
	defer cancel()

	resp, err := h.linksRepository.FindAll(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	links := make([]*pb.Link, 0, len(resp))
	for _, link := range resp {
		l := pb.Link{
			Id:        link.ID.String(),
			Title:     link.Title,
			Url:       link.URL,
			Images:    link.Images,
			Tags:      link.Tags,
			UserId:    link.UserID,
			CreatedAt: link.CreatedAt.Format(time.RFC3339),
			UpdatedAt: link.UpdatedAt.Format(time.RFC3339),
		}
		links = append(links, &l)
	}

	return &pb.ListLinkResponse{Links: links}, nil
}
