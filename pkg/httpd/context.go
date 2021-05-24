package httpd

import "github.com/boairy/koko/pkg/model"

type WebContext struct {
	User       *model.User
	Connection *Client
	Client     *Client
}
