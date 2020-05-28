package utils

import "github.com/google/uuid"

// ExtractLastNCharsOfUUID returns last n character of a random UUID
func ExtractLastNCharsOfUUID(numOfChars int) string {
	uuid := uuid.New().String()
	uuid = uuid[len(uuid)-numOfChars:]
	return uuid
}
