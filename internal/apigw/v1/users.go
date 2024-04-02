package v1

import (
	"encoding/json"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/api/apiv1"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/httpModul"
	"gitlab.com/robotomize/gb-golang/homework/03-02-umanager/pkg/pb"
	"net/http"
)

func newUsersHandler(usersClient usersClient) *usersHandler {
	return &usersHandler{client: usersClient}
}

type usersHandler struct {
	client usersClient
}

func (h *usersHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	resp, err := h.client.ListUsers(r.Context(), &pb.Empty{})
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	users := make([]apiv1.User, 0, len(resp.Users))
	for _, user := range resp.Users {
		u := apiv1.User{
			CreatedAt: user.CreatedAt,
			Id:        user.Id,
			Password:  user.Password,
			UpdatedAt: user.UpdatedAt,
			Username:  user.Username,
		}
		users = append(users, u)
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *usersHandler) PostUsers(w http.ResponseWriter, r *http.Request) {
	var createUserReq apiv1.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&createUserReq); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.BadRequest, http.StatusBadRequest)

		return
	}

	u := pb.CreateUserRequest{
		Id:       createUserReq.Id,
		Username: createUserReq.Username,
		Password: createUserReq.Password,
	}
	if _, err := h.client.CreateUser(r.Context(), &u); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *usersHandler) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	if _, err := h.client.DeleteUser(r.Context(), &pb.DeleteUserRequest{Id: id}); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *usersHandler) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	resp, err := h.client.GetUser(r.Context(), &pb.GetUserRequest{Id: id})
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	u := apiv1.User{
		CreatedAt: resp.CreatedAt,
		Id:        resp.Id,
		Password:  resp.Password,
		UpdatedAt: resp.UpdatedAt,
		Username:  resp.Username,
	}

	jsonData, err := json.Marshal(u)
	if err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func (h *usersHandler) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	var updateUserReq apiv1.UserCreate

	if err := json.NewDecoder(r.Body).Decode(&updateUserReq); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.BadRequest, http.StatusBadRequest)

		return
	}

	u := pb.UpdateUserRequest{
		Id:       updateUserReq.Id,
		Username: updateUserReq.Username,
		Password: updateUserReq.Password,
	}
	if _, err := h.client.UpdateUser(r.Context(), &u); err != nil {
		httpModul.RespondWithError(w, err.Error(), apiv1.InternalServerError, http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
}
