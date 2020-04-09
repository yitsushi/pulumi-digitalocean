// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package digitalocean

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

func LookupKubernetesCluster(ctx *pulumi.Context, args *LookupKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupKubernetesClusterResult, error) {
	var rv LookupKubernetesClusterResult
	err := ctx.Invoke("digitalocean:index/getKubernetesCluster:getKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getKubernetesCluster.
type LookupKubernetesClusterArgs struct {
	// The name of Kubernetes cluster.
	Name string `pulumi:"name"`
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
}

// A collection of values returned by getKubernetesCluster.
type LookupKubernetesClusterResult struct {
	// The range of IP addresses in the overlay network of the Kubernetes cluster.
	ClusterSubnet string `pulumi:"clusterSubnet"`
	// The date and time when the Kubernetes cluster was created.
	CreatedAt string `pulumi:"createdAt"`
	// The base URL of the API server on the Kubernetes master node.
	Endpoint string `pulumi:"endpoint"`
	// id is the provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The public IPv4 address of the Kubernetes master node.
	Ipv4Address string                           `pulumi:"ipv4Address"`
	KubeConfigs []GetKubernetesClusterKubeConfig `pulumi:"kubeConfigs"`
	Name        string                           `pulumi:"name"`
	// A list of node pools associated with the cluster. Each node pool exports the following attributes:
	// - `id` -  The unique ID that can be used to identify and reference the node pool.
	// - `name` - The name of the node pool.
	// - `size` - The slug identifier for the type of Droplet used as workers in the node pool.
	// - `nodeCount` - The number of Droplet instances in the node pool.
	// - `actualNodeCount` - The actual number of nodes in the node pool, which is especially useful when auto-scaling is enabled.
	// - `autoScale` - A boolean indicating whether auto-scaling is enabled on the node pool.
	// - `minNodes` - If auto-scaling is enabled, this represents the minimum number of nodes that the node pool can be scaled down to.
	// - `maxNodes` - If auto-scaling is enabled, this represents the maximum number of nodes that the node pool can be scaled up to.
	// - `tags` - A list of tag names applied to the node pool.
	// - `labels` - A map of key/value pairs applied to nodes in the pool. The labels are exposed in the Kubernetes API as labels in the metadata of the corresponding [Node resources](https://kubernetes.io/docs/concepts/architecture/nodes/).
	// - `nodes` - A list of nodes in the pool. Each node exports the following attributes:
	// + `id` -  A unique ID that can be used to identify and reference the node.
	// + `name` - The auto-generated name for the node.
	// + `status` -  A string indicating the current status of the individual node.
	// + `createdAt` - The date and time when the node was created.
	// + `updatedAt` - The date and time when the node was last updated.
	NodePools []GetKubernetesClusterNodePool `pulumi:"nodePools"`
	// The slug identifier for the region where the Kubernetes cluster is located.
	Region string `pulumi:"region"`
	// The range of assignable IP addresses for services running in the Kubernetes cluster.
	ServiceSubnet string `pulumi:"serviceSubnet"`
	// A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.
	Status string `pulumi:"status"`
	// A list of tag names to be applied to the Kubernetes cluster.
	Tags []string `pulumi:"tags"`
	// The date and time when the Kubernetes cluster was last updated.
	// * `kube_config.0` - A representation of the Kubernetes cluster's kubeconfig with the following attributes:
	// - `rawConfig` - The full contents of the Kubernetes cluster's kubeconfig file.
	// - `host` - The URL of the API server on the Kubernetes master node.
	// - `clusterCaCertificate` - The base64 encoded public certificate for the cluster's certificate authority.
	// - `token` - The DigitalOcean API access token used by clients to access the cluster.
	// - `clientKey` - The base64 encoded private key used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
	// - `clientCertificate` - The base64 encoded public certificate used by clients to access the cluster. Only available if token authentication is not supported on your cluster.
	// - `expiresAt` - The date and time when the credentials will expire and need to be regenerated.
	UpdatedAt string `pulumi:"updatedAt"`
	// The slug identifier for the version of Kubernetes used for the cluster.
	Version string `pulumi:"version"`
}
