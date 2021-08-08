package dtrack

import (
	"context"
	"net/http"
	"net/url"
)

type UserService struct {
	client *Client
}

func (u UserService) ForceChangePassword(ctx context.Context, username, password, newPassword string) error {
	body := url.Values{}
	body.Set("username", username)
	body.Set("password", password)
	body.Set("newPassword", newPassword)
	body.Set("confirmPassword", newPassword)

	req, err := u.client.newRequest(ctx, http.MethodPost, "/api/v1/user/forceChangePassword", withBody(body))
	if err != nil {
		return err
	}

	req.Header.Set("Accept", "*/*")

	_, err = u.client.doRequest(req, nil)
	if err != nil {
		return err
	}

	return nil
}