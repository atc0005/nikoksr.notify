/*
 * Textmagic API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package TextMagic

type GetSenderSettingsResponse struct {
	User []SenderSettingsItem `json:"user"`
	Special []SenderSettingsItem `json:"special"`
	Other []SenderSettingsItem `json:"other"`
}