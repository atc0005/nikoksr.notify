/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type CreateTemplateInputObject struct {
	// Template name.
	Name string `json:"name,omitempty"`
	// Template text. May contain tags inside braces. See [Get timezones](https://docs.textmagic.com/#section/Custom-fields-list-(Merge-tags)).
	Content string `json:"content,omitempty"`
}
