/*
 * Product Info.
 *
 * The product info application uses the cloud provider APIs to asynchronously fetch and parse instance type attributes and prices, while storing the results in an in memory cache and making it available as structured data through a REST API.
 *
 * API version: 0.7.0
 * Contact: info@banzaicloud.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cloudinfo
// GetAttributeValuesPathParams GetAttributeValuesPathParams is a placeholder for the get attribute values route's path parameters
type GetAttributeValuesPathParams struct {
	// in:path
	Attribute string `json:"attribute,omitempty"`
	// in:path
	Provider string `json:"provider,omitempty"`
	// in:path
	Region string `json:"region,omitempty"`
	// in:path
	Service string `json:"service,omitempty"`
}
