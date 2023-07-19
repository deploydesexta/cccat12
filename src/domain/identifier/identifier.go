package identifier

import "github.com/google/uuid"

func Unique() string {
	return uuid.New().String()
}
