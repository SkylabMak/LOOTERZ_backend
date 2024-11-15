// utils/content_type.go
package security

import (
	"encoding/json"
	"log"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/microcosm-cc/bluemonday"
)

// CheckContentType checks if the Content-Type of the request matches the expected type.
// If it doesn't match, it returns an error response.
// CheckContentType returns true if the Content-Type matches the expected type, otherwise false.
func CheckContentType(c *fiber.Ctx, expectedType string) bool {
	return c.Get("Content-Type") == expectedType
}

// CheckContentJSONType returns true if the Content-Type is "application/json", otherwise false.
func CheckContentJSONType(c *fiber.Ctx) bool {
	return c.Get("Content-Type") == "application/json"
}

func SanitizeInput(input string) string {
    p := bluemonday.StrictPolicy() 
    return p.Sanitize(input)
}

func SanitizeJSONBody(c *fiber.Ctx) bool {
	// Ensure Content-Type is application/json
	if c.Get("Content-Type") != "application/json" {
		return false
	}

	// Parse the JSON body into a map
	var data map[string]interface{}
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return false
	}

	// Create a strict policy for sanitization
	p := bluemonday.StrictPolicy()
	modified := false // Track if any value was sanitized

	// Sanitize all string fields and check for modifications
	for key, value := range data {
		if strValue, ok := value.(string); ok {
			log.Printf("Original value for key '%s': %s", key, strValue)
			sanitizedValue := p.Sanitize(strValue)
			if sanitizedValue != strValue {
				modified = true
				log.Printf("Sanitized value for key '%s': %s", key, sanitizedValue)
				return false
			}
		}
	}

	// Return false if any field was sanitized
	return !modified
}

func DetectSanitizeJSONBody(c *fiber.Ctx) (bool) {
	// Ensure Content-Type is application/json
	if c.Get("Content-Type") != "application/json" {
		return false
	}

	// Parse the JSON body into a map
	var data map[string]interface{}
	if err := json.Unmarshal(c.Body(), &data); err != nil {
		return false
	}

	// Check if the input contains unsafe content
	if !DetectUnsafeJSONBody(data) {
		log.Printf("UnsafeJSONBody '%s'", data)
		return false
	}

	return true
}

func DetectUnsafeString(input string) bool {
	// Example: Check for script tags
	return !govalidator.Contains(input, "<script>")
}

// DetectUnsafeJSONBody validates JSON for unsafe content
func DetectUnsafeJSONBody(data map[string]interface{}) bool {
	for _, value := range data {
		if strValue, ok := value.(string); ok {
			if !DetectUnsafeString(strValue) {
				return false // Unsafe content detected
			}
		}
	}
	return true // All content is safe
}