// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean database cluster resource.
//
//
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-digitalocean/blob/master/website/docs/r/database_cluster.html.markdown.
type DatabaseCluster struct {
	pulumi.CustomResourceState

	// Name of the cluster's default database.
	Database pulumi.StringOutput `pulumi:"database"`
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
	Engine pulumi.StringOutput `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrOutput `pulumi:"evictionPolicy"`
	// Database cluster's hostname.
	Host pulumi.StringOutput `pulumi:"host"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayOutput `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntOutput `pulumi:"nodeCount"`
	// Password for the cluster's default user.
	Password pulumi.StringOutput `pulumi:"password"`
	// Network port that the database cluster is listening on.
	Port pulumi.IntOutput `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringOutput `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringOutput `pulumi:"privateUri"`
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringOutput `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringOutput `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrOutput `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
	// The full URI for connecting to the database cluster.
	Uri pulumi.StringOutput `pulumi:"uri"`
	// The uniform resource name of the database cluster.
	Urn pulumi.StringOutput `pulumi:"urn"`
	// Username for the cluster's default user.
	User pulumi.StringOutput `pulumi:"user"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrOutput `pulumi:"version"`
}

// NewDatabaseCluster registers a new resource with the given unique name, arguments, and options.
func NewDatabaseCluster(ctx *pulumi.Context,
	name string, args *DatabaseClusterArgs, opts ...pulumi.ResourceOption) (*DatabaseCluster, error) {
	if args == nil || args.Engine == nil {
		return nil, errors.New("missing required argument 'Engine'")
	}
	if args == nil || args.NodeCount == nil {
		return nil, errors.New("missing required argument 'NodeCount'")
	}
	if args == nil || args.Region == nil {
		return nil, errors.New("missing required argument 'Region'")
	}
	if args == nil || args.Size == nil {
		return nil, errors.New("missing required argument 'Size'")
	}
	if args == nil {
		args = &DatabaseClusterArgs{}
	}
	var resource DatabaseCluster
	err := ctx.RegisterResource("digitalocean:index/databaseCluster:DatabaseCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDatabaseCluster gets an existing DatabaseCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDatabaseCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DatabaseClusterState, opts ...pulumi.ResourceOption) (*DatabaseCluster, error) {
	var resource DatabaseCluster
	err := ctx.ReadResource("digitalocean:index/databaseCluster:DatabaseCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DatabaseCluster resources.
type databaseClusterState struct {
	// Name of the cluster's default database.
	Database *string `pulumi:"database"`
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
	Engine *string `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Database cluster's hostname.
	Host *string `pulumi:"host"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows []DatabaseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name *string `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount *int `pulumi:"nodeCount"`
	// Password for the cluster's default user.
	Password *string `pulumi:"password"`
	// Network port that the database cluster is listening on.
	Port *int `pulumi:"port"`
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost *string `pulumi:"privateHost"`
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri *string `pulumi:"privateUri"`
	// DigitalOcean region where the cluster will reside.
	Region *string `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
	Size *string `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode *string `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags []string `pulumi:"tags"`
	// The full URI for connecting to the database cluster.
	Uri *string `pulumi:"uri"`
	// The uniform resource name of the database cluster.
	Urn *string `pulumi:"urn"`
	// Username for the cluster's default user.
	User *string `pulumi:"user"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version *string `pulumi:"version"`
}

type DatabaseClusterState struct {
	// Name of the cluster's default database.
	Database pulumi.StringPtrInput
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
	Engine pulumi.StringPtrInput
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrInput
	// Database cluster's hostname.
	Host pulumi.StringPtrInput
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayInput
	// The name of the database cluster.
	Name pulumi.StringPtrInput
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntPtrInput
	// Password for the cluster's default user.
	Password pulumi.StringPtrInput
	// Network port that the database cluster is listening on.
	Port pulumi.IntPtrInput
	// Same as `host`, but only accessible from resources within the account and in the same region.
	PrivateHost pulumi.StringPtrInput
	// Same as `uri`, but only accessible from resources within the account and in the same region.
	PrivateUri pulumi.StringPtrInput
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringPtrInput
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringPtrInput
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrInput
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayInput
	// The full URI for connecting to the database cluster.
	Uri pulumi.StringPtrInput
	// The uniform resource name of the database cluster.
	Urn pulumi.StringPtrInput
	// Username for the cluster's default user.
	User pulumi.StringPtrInput
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrInput
}

func (DatabaseClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseClusterState)(nil)).Elem()
}

type databaseClusterArgs struct {
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
	Engine string `pulumi:"engine"`
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy *string `pulumi:"evictionPolicy"`
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows []DatabaseClusterMaintenanceWindow `pulumi:"maintenanceWindows"`
	// The name of the database cluster.
	Name *string `pulumi:"name"`
	// Number of nodes that will be included in the cluster.
	NodeCount int `pulumi:"nodeCount"`
	// DigitalOcean region where the cluster will reside.
	Region string `pulumi:"region"`
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
	Size string `pulumi:"size"`
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode *string `pulumi:"sqlMode"`
	// A list of tag names to be applied to the database cluster.
	Tags []string `pulumi:"tags"`
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version *string `pulumi:"version"`
}

// The set of arguments for constructing a DatabaseCluster resource.
type DatabaseClusterArgs struct {
	// Database engine used by the cluster (ex. `pg` for PostreSQL, `mysql` for MySQL, or `redis` for Redis).
	Engine pulumi.StringInput
	// A string specifying the eviction policy for a Redis cluster. Valid values are: `noeviction`, `allkeysLru`, `allkeysRandom`, `volatileLru`, `volatileRandom`, or `volatileTtl`.
	EvictionPolicy pulumi.StringPtrInput
	// Defines when the automatic maintenance should be performed for the database cluster.
	MaintenanceWindows DatabaseClusterMaintenanceWindowArrayInput
	// The name of the database cluster.
	Name pulumi.StringPtrInput
	// Number of nodes that will be included in the cluster.
	NodeCount pulumi.IntInput
	// DigitalOcean region where the cluster will reside.
	Region pulumi.StringInput
	// Database Droplet size associated with the cluster (ex. `db-s-1vcpu-1gb`).
	Size pulumi.StringInput
	// A comma separated string specifying the  SQL modes for a MySQL cluster.
	SqlMode pulumi.StringPtrInput
	// A list of tag names to be applied to the database cluster.
	Tags pulumi.StringArrayInput
	// Engine version used by the cluster (ex. `11` for PostgreSQL 11).
	Version pulumi.StringPtrInput
}

func (DatabaseClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*databaseClusterArgs)(nil)).Elem()
}
