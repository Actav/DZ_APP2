package v1

import (
	"encoding/json"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/api/apiv1"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/httpModul"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
	"net/http"
)

func newLinksHandler(linksClient linksClient) *linksHandler {
	return &linksHandler{client: linksClient}
}

type linksHandler struct {
	client linksClient
}

func (h *linksHandler) GetLinks(w http.ResponseWriter, r *http.Request) {
	resp, err := h.client.ListLinks(r.Context(), &pb.Empty{})
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	links := make([]apiv1.Link, 0, len(resp.Links))
	for _, link := range resp.Links {
		l := apiv1.Link{
			CreatedAt: link.CreatedAt,
			Id:        link.Id,
			Images:    link.Images,
			Tags:      link.Tags,
			Title:     link.Title,
			UpdatedAt: link.UpdatedAt,
			Url:       link.Url,
			UserId:    link.UserId,
		}
		links = append(links, l)
	}

	jsonData, err := json.Marshal(links)
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *linksHandler) PostLinks(w http.ResponseWriter, r *http.Request) {
	var createLinkReq apiv1.LinkCreate

	if err := json.NewDecoder(r.Body).Decode(&createLinkReq); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.BadRequest, http.StatusBadRequest)

		return
	}

	l := pb.CreateLinkRequest{
		Id:     createLinkReq.Id,
		Title:  createLinkReq.Title,
		Url:    createLinkReq.Url,
		Images: createLinkReq.Images,
		Tags:   createLinkReq.Tags,
		UserId: createLinkReq.UserId,
	}

	if _, err := h.client.CreateLink(r.Context(), &l); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *linksHandler) DeleteLinksId(w http.ResponseWriter, r *http.Request, id string) {
	if _, err := h.client.DeleteLink(r.Context(), &pb.DeleteLinkRequest{Id: id}); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *linksHandler) GetLinksId(w http.ResponseWriter, r *http.Request, id string) {
	resp, err := h.client.GetLink(r.Context(), &pb.GetLinkRequest{Id: id})
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	l := apiv1.Link{
		CreatedAt: resp.CreatedAt,
		Id:        resp.Id,
		Images:    resp.Images,
		Tags:      resp.Tags,
		Title:     resp.Title,
		UpdatedAt: resp.UpdatedAt,
		Url:       resp.Url,
		UserId:    resp.UserId,
	}

	jsonData, err := json.Marshal(l)
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *linksHandler) PutLinksId(w http.ResponseWriter, r *http.Request, id string) {
	var updateLinkReq apiv1.LinkCreate

	if err := json.NewDecoder(r.Body).Decode(&updateLinkReq); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.BadRequest, http.StatusBadRequest)

		return
	}

	l := pb.UpdateLinkRequest{
		Id:     updateLinkReq.Id,
		Images: updateLinkReq.Images,
		Tags:   updateLinkReq.Tags,
		Title:  updateLinkReq.Title,
		Url:    updateLinkReq.Url,
		UserId: updateLinkReq.UserId,
	}
	if _, err := h.client.UpdateLink(r.Context(), &l); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *linksHandler) GetLinksUserUserID(w http.ResponseWriter, r *http.Request, userID string) {
	resp, err := h.client.GetLinkByUserID(r.Context(), &pb.GetLinksByUserId{UserId: userID})
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.BadRequest, http.StatusBadRequest)

		return
	}

	links := make([]apiv1.Link, 0, len(resp.Links))
	for _, link := range resp.Links {
		l := apiv1.Link{
			CreatedAt: link.CreatedAt,
			Id:        link.Id,
			Images:    link.Images,
			Tags:      link.Tags,
			Title:     link.Title,
			UpdatedAt: link.UpdatedAt,
			Url:       link.Url,
			UserId:    link.UserId,
		}
		links = append(links, l)
	}

	jsonData, err := json.Marshal(links)
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
