/*
 * Anchore Engine API Server
 *
 * This is the Anchore Engine API. Provides the primary external API for users of the service.
 *
 * API version: 0.1.12
 * Contact: nurmi@anchore.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package anchore
// Policy struct for Policy
type Policy struct {
	Id string `json:"id"`
	Name string `json:"name,omitempty"`
	Comment string `json:"comment,omitempty"`
	Version string `json:"version"`
	Rules []PolicyRule `json:"rules,omitempty"`
}
