module github.com/pulumi/pulumi-digitalocean/provider/v3

go 1.14

require (
	github.com/digitalocean/terraform-provider-digitalocean v1.22.2
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.19.0
	github.com/pulumi/pulumi/sdk/v2 v2.20.1-0.20210212181059-f4b0fa86fedc
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/digitalocean/terraform-provider-digitalocean => github.com/pulumi/terraform-provider-digitalocean v1.19.1-0.20210208113526-c641773437c2
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20201218231525-9cca98608a5e
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/keybase/go-crypto => github.com/keybase/go-crypto v0.0.0-20190416182011-b785b22cc757
	golang.org/x/crypto => golang.org/x/crypto v0.0.0-20200323165209-0ec3e9974c59
)
