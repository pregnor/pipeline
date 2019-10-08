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
// PolicyRule A rule that defines and decision value if the match is found true for a given image.
type PolicyRule struct {
	Id string `json:"id,omitempty"`
	Gate string `json:"gate"`
	Trigger string `json:"trigger"`
	Action string `json:"action"`
	Params []PolicyRuleParams `json:"params,omitempty"`
}
