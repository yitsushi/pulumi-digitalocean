// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
	"github.com/pulumi/pulumi/sdk/go/pulumi/config"
)

// The URL to use for the DigitalOcean API.
func GetApiEndpoint(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "digitalocean:apiEndpoint")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("https://api.digitalocean.com", nil, "DIGITALOCEAN_API_URL").(string); ok {
		return dv
	}
	return v
}

// The access key ID for Spaces API operations.
func GetSpacesAccessId(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "digitalocean:spacesAccessId")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "SPACES_ACCESS_KEY_ID").(string); ok {
		return dv
	}
	return v
}

// The secret access key for Spaces API operations.
func GetSpacesSecretKey(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "digitalocean:spacesSecretKey")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "SPACES_SECRET_ACCESS_KEY").(string); ok {
		return dv
	}
	return v
}

// The token key for API operations.
func GetToken(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "digitalocean:token")
	if err == nil {
		return v
	}
	if dv, ok := getEnvOrDefault("", nil, "DIGITALOCEAN_TOKEN").(string); ok {
		return dv
	}
	return v
}
