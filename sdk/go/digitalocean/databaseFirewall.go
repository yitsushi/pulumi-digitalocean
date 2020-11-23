// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean database firewall resource allowing you to restrict
// connections to your database to trusted sources. You may limit connections to
// specific Droplets, Kubernetes clusters, or IP addresses.
//
// ## Example Usage
// ### Create a new database firewall allowing multiple IP addresses
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := digitalocean.NewDatabaseCluster(ctx, "postgres_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("pg"),
// 			Version:   pulumi.String("11"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Region:    pulumi.String("nyc1"),
// 			NodeCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDatabaseFirewall(ctx, "example_fw", &digitalocean.DatabaseFirewallArgs{
// 			ClusterId: postgres_example.ID(),
// 			Rules: digitalocean.DatabaseFirewallRuleArray{
// 				&digitalocean.DatabaseFirewallRuleArgs{
// 					Type:  pulumi.String("ip_addr"),
// 					Value: pulumi.String("192.168.1.1"),
// 				},
// 				&digitalocean.DatabaseFirewallRuleArgs{
// 					Type:  pulumi.String("ip_addr"),
// 					Value: pulumi.String("192.0.2.0"),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Create a new database firewall allowing a Droplet
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-digitalocean/sdk/v3/go/digitalocean"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		web, err := digitalocean.NewDroplet(ctx, "web", &digitalocean.DropletArgs{
// 			Size:   pulumi.String("s-1vcpu-1gb"),
// 			Image:  pulumi.String("centos-7-x64"),
// 			Region: pulumi.String("nyc3"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDatabaseCluster(ctx, "postgres_example", &digitalocean.DatabaseClusterArgs{
// 			Engine:    pulumi.String("pg"),
// 			Version:   pulumi.String("11"),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Region:    pulumi.String("nyc1"),
// 			NodeCount: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = digitalocean.NewDatabaseFirewall(ctx, "example_fw", &digitalocean.DatabaseFirewallArgs{
// 			ClusterId: postgres_example.ID(),
// 			Rules: digitalocean.DatabaseFirewallRuleArray{
// 				&digitalocean.DatabaseFirewallRuleArgs{
// 					Type:  pulumi.String("droplet"),
// 					Value: web.ID(),
// 				},
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Database firewalls can be imported using the `id` of the target database cluster For example
//
// ```sh
//  $ pulumi import digitalocean:index:DatabaseFirewall example-fw 5f55c6cd-863b-4907-99b8-7e09b0275d54
// ```
type DatabaseFirewall struct {
	pulumi.CustomResourceState

	// The ID of the target database cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
	Rules DatabaseFirewallRuleArrayOutput `pulumi:"rules"`
}

// NewDatabaseFirewall registers a new resource with the given unique name, arguments, and options.
func NewDatabaseFirewall(ctx *pulumi.Context,
	name string, args *DatabaseFirewallArgs, opts ...pulumi.ResourceOption) (*DatabaseFirewall, error) {
	if args == nil || args.ClusterId == nil {
		return nil, errors.New("missing required argument 'ClusterId'")
	}
	if args == nil || args.Rules == nil {
		return nil, errors.New("missing required argument 'Rules'")
	}
	if args == nil {
		args = &DatabaseFirewallArgs{}
	}
	var resource DatabaseFirewall
	err := ctx.RegisterResource("digitalocean:index:DatabaseFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseFirewall gets an existing DatabaseFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseFirewallState, opts ...pulumi.ResourceOption) (*DatabaseFirewall, error) {
	var resource DatabaseFirewall
	err := ctx.ReadResource("digitalocean:index:DatabaseFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseFirewall resources.
type databaseFirewallState struct {
	// The ID of the target database cluster.
	ClusterId *string `pulumi:"clusterId"`
	// A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
	Rules []DatabaseFirewallRule `pulumi:"rules"`
}

type DatabaseFirewallState struct {
	// The ID of the target database cluster.
	ClusterId pulumi.StringPtrInput
	// A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
	Rules DatabaseFirewallRuleArrayInput
}

func (DatabaseFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseFirewallState)(nil)).Elem()
}

type databaseFirewallArgs struct {
	// The ID of the target database cluster.
	ClusterId string `pulumi:"clusterId"`
	// A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
	Rules []DatabaseFirewallRule `pulumi:"rules"`
}

// The set of arguments for constructing a DatabaseFirewall resource.
type DatabaseFirewallArgs struct {
	// The ID of the target database cluster.
	ClusterId pulumi.StringInput
	// A rule specifying a resource allowed to access the database cluster. The following arguments must be specified:
	Rules DatabaseFirewallRuleArrayInput
}

func (DatabaseFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseFirewallArgs)(nil)).Elem()
}

type DatabaseFirewallInput interface {
	pulumi.Input

	ToDatabaseFirewallOutput() DatabaseFirewallOutput
	ToDatabaseFirewallOutputWithContext(ctx context.Context) DatabaseFirewallOutput
}

func (DatabaseFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseFirewall)(nil)).Elem()
}

func (i DatabaseFirewall) ToDatabaseFirewallOutput() DatabaseFirewallOutput {
	return i.ToDatabaseFirewallOutputWithContext(context.Background())
}

func (i DatabaseFirewall) ToDatabaseFirewallOutputWithContext(ctx context.Context) DatabaseFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseFirewallOutput)
}

type DatabaseFirewallOutput struct {
	*pulumi.OutputState
}

func (DatabaseFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseFirewallOutput)(nil)).Elem()
}

func (o DatabaseFirewallOutput) ToDatabaseFirewallOutput() DatabaseFirewallOutput {
	return o
}

func (o DatabaseFirewallOutput) ToDatabaseFirewallOutputWithContext(ctx context.Context) DatabaseFirewallOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DatabaseFirewallOutput{})
}
