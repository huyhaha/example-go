package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/hieunmce/example-go/domain"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.FindRequest{UserID: userID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return categoryEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req categoryEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}

	var req categoryEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.User.ID = userID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	userID, err := domain.UUIDFromString(chi.URLParam(r, "user_id"))
	if err != nil {
		return nil, err
	}
	return categoryEndpoint.DeleteRequest{UserID: userID}, nil
}
