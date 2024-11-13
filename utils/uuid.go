// uuid_generator.go
package utils

import "github.com/google/uuid"

// GenerateUUID generates and returns a new UUID.
func GenerateUUID() (string, error) {
    // Generate a new UUID
    id, err := uuid.NewUUID()
    if err != nil {
        return "", err
    }
    return id.String(), nil
}
