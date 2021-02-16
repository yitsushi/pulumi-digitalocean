// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean database replica resource.
//
// ## Example Usage
// ### Create a new PostgreSQL database replica
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
// 		_, err = digitalocean.NewDatabaseReplica(ctx, "read_replica", &digitalocean.DatabaseReplicaArgs{
// 			ClusterId: postgres_example.ID(),
// 			Size:      pulumi.String("db-s-1vcpu-1gb"),
// 			Region:    pulumi.String("nyc1"),
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
// Database replicas can be imported using the `id` of the source database cluster and the `name` of the replica joined with a comma. For example
//
// ```sh
//  $ pulumi import digitalocean:index/databaseReplica:DatabaseReplica read-replica 245bcfd0-7f31-4ce6-a2bc-475a116cca97,read-replica
// ```
type DatabaseReplica struct {
	pulumi.CustomResourceState

	// The ID of the original source database cluster.
	ClusterId pulumi.StringOutput `pulumi:"clusterId"`
	// Name of the replica's default database.
	Database pulumi.StringOutput `pulumi:"database"`
	// Database replica's hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// The name for the database replica.
	Name pulumi.StringOutput `pulumi:"name"`
	// Password for the replica's default user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Network port that the database replica is listening on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost        pulumi.StringOutput `pulumi:"privateHost"`
	PrivateNetworkUuid pulumi.StringOutput `pulumi:"privateNetworkUuid"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringOutput `pulumi:"privateUri"`
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrOutput `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrOutput   `pulumi:"size"`
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The full URI for connecting to the database replica.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// Username for the replica's default user.
	User pulumi.StringOutput `pulumi:"user"`
}

// NewDatabaseReplica registers a new resource with the given unique name, arguments, and options.
func NewDatabaseReplica(ctx *pulumi.Context,
	name string, args *DatabaseReplicaArgs, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ClusterId == nil {
		return nil, errors.New("invalid value for required argument 'ClusterId'")
	}
	var resource DatabaseReplica
	err := ctx.RegisterResource("digitalocean:index/databaseReplica:DatabaseReplica", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseReplica gets an existing DatabaseReplica resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseReplica(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseReplicaState, opts ...pulumi.ResourceOption) (*DatabaseReplica, error) {
	var resource DatabaseReplica
	err := ctx.ReadResource("digitalocean:index/databaseReplica:DatabaseReplica", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseReplica resources.
type databaseReplicaState struct {
	// The ID of the original source database cluster.
	ClusterId *string `pulumi:"clusterId"`
	// Name of the replica's default database.
	Database *string `pulumi:"database"`
	// Database replica's hostname.
	Host *string `pulumi:"host"`
	// The name for the database replica.
	Name *string `pulumi:"name"`
	// Password for the replica's default user.
	Password *string `pulumi:"password"`
	// Network port that the database replica is listening on.
	Port *int `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost        *string `pulumi:"privateHost"`
	PrivateNetworkUuid *string `pulumi:"privateNetworkUuid"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri *string `pulumi:"privateUri"`
	// DigitalOcean region where the replica will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size *string  `pulumi:"size"`
	Tags []string `pulumi:"tags"`
	// The full URI for connecting to the database replica.
	Uri *string `pulumi:"uri"`
	// Username for the replica's default user.
	User *string `pulumi:"user"`
}

type DatabaseReplicaState struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringPtrInput
	// Name of the replica's default database.
	Database pulumi.StringPtrInput
	// Database replica's hostname.
	Host pulumi.StringPtrInput
	// The name for the database replica.
	Name pulumi.StringPtrInput
	// Password for the replica's default user.
	Password pulumi.StringPtrInput
	// Network port that the database replica is listening on.
	Port pulumi.IntPtrInput
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost        pulumi.StringPtrInput
	PrivateNetworkUuid pulumi.StringPtrInput
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringPtrInput
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
	// The full URI for connecting to the database replica.
	Uri pulumi.StringPtrInput
	// Username for the replica's default user.
	User pulumi.StringPtrInput
}

func (DatabaseReplicaState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaState)(nil)).Elem()
}

type databaseReplicaArgs struct {
	// The ID of the original source database cluster.
	ClusterId string `pulumi:"clusterId"`
	// The name for the database replica.
	Name               *string `pulumi:"name"`
	PrivateNetworkUuid *string `pulumi:"privateNetworkUuid"`
	// DigitalOcean region where the replica will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size *string  `pulumi:"size"`
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a DatabaseReplica resource.
type DatabaseReplicaArgs struct {
	// The ID of the original source database cluster.
	ClusterId pulumi.StringInput
	// The name for the database replica.
	Name               pulumi.StringPtrInput
	PrivateNetworkUuid pulumi.StringPtrInput
	// DigitalOcean region where the replica will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the replica (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrInput
	Tags pulumi.StringArrayInput
}

func (DatabaseReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseReplicaArgs)(nil)).Elem()
}

type DatabaseReplicaInput interface {
	pulumi.Input

	ToDatabaseReplicaOutput() DatabaseReplicaOutput
	ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput
}

func (*DatabaseReplica) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseReplica)(nil))
}

func (i *DatabaseReplica) ToDatabaseReplicaOutput() DatabaseReplicaOutput {
	return i.ToDatabaseReplicaOutputWithContext(context.Background())
}

func (i *DatabaseReplica) ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaOutput)
}

func (i *DatabaseReplica) ToDatabaseReplicaPtrOutput() DatabaseReplicaPtrOutput {
	return i.ToDatabaseReplicaPtrOutputWithContext(context.Background())
}

func (i *DatabaseReplica) ToDatabaseReplicaPtrOutputWithContext(ctx context.Context) DatabaseReplicaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaPtrOutput)
}

type DatabaseReplicaPtrInput interface {
	pulumi.Input

	ToDatabaseReplicaPtrOutput() DatabaseReplicaPtrOutput
	ToDatabaseReplicaPtrOutputWithContext(ctx context.Context) DatabaseReplicaPtrOutput
}

type databaseReplicaPtrType DatabaseReplicaArgs

func (*databaseReplicaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReplica)(nil))
}

func (i *databaseReplicaPtrType) ToDatabaseReplicaPtrOutput() DatabaseReplicaPtrOutput {
	return i.ToDatabaseReplicaPtrOutputWithContext(context.Background())
}

func (i *databaseReplicaPtrType) ToDatabaseReplicaPtrOutputWithContext(ctx context.Context) DatabaseReplicaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaPtrOutput)
}

// DatabaseReplicaArrayInput is an input type that accepts DatabaseReplicaArray and DatabaseReplicaArrayOutput values.
// You can construct a concrete instance of `DatabaseReplicaArrayInput` via:
//
//          DatabaseReplicaArray{ DatabaseReplicaArgs{...} }
type DatabaseReplicaArrayInput interface {
	pulumi.Input

	ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput
	ToDatabaseReplicaArrayOutputWithContext(context.Context) DatabaseReplicaArrayOutput
}

type DatabaseReplicaArray []DatabaseReplicaInput

func (DatabaseReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*DatabaseReplica)(nil))
}

func (i DatabaseReplicaArray) ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput {
	return i.ToDatabaseReplicaArrayOutputWithContext(context.Background())
}

func (i DatabaseReplicaArray) ToDatabaseReplicaArrayOutputWithContext(ctx context.Context) DatabaseReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaArrayOutput)
}

// DatabaseReplicaMapInput is an input type that accepts DatabaseReplicaMap and DatabaseReplicaMapOutput values.
// You can construct a concrete instance of `DatabaseReplicaMapInput` via:
//
//          DatabaseReplicaMap{ "key": DatabaseReplicaArgs{...} }
type DatabaseReplicaMapInput interface {
	pulumi.Input

	ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput
	ToDatabaseReplicaMapOutputWithContext(context.Context) DatabaseReplicaMapOutput
}

type DatabaseReplicaMap map[string]DatabaseReplicaInput

func (DatabaseReplicaMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*DatabaseReplica)(nil))
}

func (i DatabaseReplicaMap) ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput {
	return i.ToDatabaseReplicaMapOutputWithContext(context.Background())
}

func (i DatabaseReplicaMap) ToDatabaseReplicaMapOutputWithContext(ctx context.Context) DatabaseReplicaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseReplicaMapOutput)
}

type DatabaseReplicaOutput struct {
	*pulumi.OutputState
}

func (DatabaseReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseReplica)(nil))
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaOutput() DatabaseReplicaOutput {
	return o
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaOutputWithContext(ctx context.Context) DatabaseReplicaOutput {
	return o
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaPtrOutput() DatabaseReplicaPtrOutput {
	return o.ToDatabaseReplicaPtrOutputWithContext(context.Background())
}

func (o DatabaseReplicaOutput) ToDatabaseReplicaPtrOutputWithContext(ctx context.Context) DatabaseReplicaPtrOutput {
	return o.ApplyT(func(v DatabaseReplica) *DatabaseReplica {
		return &v
	}).(DatabaseReplicaPtrOutput)
}

type DatabaseReplicaPtrOutput struct {
	*pulumi.OutputState
}

func (DatabaseReplicaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DatabaseReplica)(nil))
}

func (o DatabaseReplicaPtrOutput) ToDatabaseReplicaPtrOutput() DatabaseReplicaPtrOutput {
	return o
}

func (o DatabaseReplicaPtrOutput) ToDatabaseReplicaPtrOutputWithContext(ctx context.Context) DatabaseReplicaPtrOutput {
	return o
}

type DatabaseReplicaArrayOutput struct{ *pulumi.OutputState }

func (DatabaseReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseReplica)(nil))
}

func (o DatabaseReplicaArrayOutput) ToDatabaseReplicaArrayOutput() DatabaseReplicaArrayOutput {
	return o
}

func (o DatabaseReplicaArrayOutput) ToDatabaseReplicaArrayOutputWithContext(ctx context.Context) DatabaseReplicaArrayOutput {
	return o
}

func (o DatabaseReplicaArrayOutput) Index(i pulumi.IntInput) DatabaseReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseReplica {
		return vs[0].([]DatabaseReplica)[vs[1].(int)]
	}).(DatabaseReplicaOutput)
}

type DatabaseReplicaMapOutput struct{ *pulumi.OutputState }

func (DatabaseReplicaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DatabaseReplica)(nil))
}

func (o DatabaseReplicaMapOutput) ToDatabaseReplicaMapOutput() DatabaseReplicaMapOutput {
	return o
}

func (o DatabaseReplicaMapOutput) ToDatabaseReplicaMapOutputWithContext(ctx context.Context) DatabaseReplicaMapOutput {
	return o
}

func (o DatabaseReplicaMapOutput) MapIndex(k pulumi.StringInput) DatabaseReplicaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DatabaseReplica {
		return vs[0].(map[string]DatabaseReplica)[vs[1].(string)]
	}).(DatabaseReplicaOutput)
}

func init() {
	pulumi.RegisterOutputType(DatabaseReplicaOutput{})
	pulumi.RegisterOutputType(DatabaseReplicaPtrOutput{})
	pulumi.RegisterOutputType(DatabaseReplicaArrayOutput{})
	pulumi.RegisterOutputType(DatabaseReplicaMapOutput{})
}
