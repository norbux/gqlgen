package models

import "github.com/norbux/gqlgen/integration/remote_api"

type Viewer struct {
	User *remote_api.User
}
