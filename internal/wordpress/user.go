package wordpress

import (
	"encoding/json"
	"fmt"

	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

type User struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Slug        string `json:"slug"`
}

func (api *API) GetUsers() []User {
	limit, err := request.CalculateMiB(8)

	if err != nil {
		return []User{}
	}

	response, err := request.Do(fmt.Sprintf("%swp/v2/users", api.api), api.userAgent, limit)

	if err != nil {
		return nil
	}

	var users []User

	err = json.Unmarshal([]byte(response), &users)

	if err != nil {
		return nil
	}

	return users
}
