// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package digitalocean

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a DigitalOcean App resource.
//
// ## Example Usage
//
// To create an app, provide a [DigitalOcean app spec](https://www.digitalocean.com/docs/app-platform/references/app-specification-reference/) specifying the app's components.
// ### Basic Example
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
// 		_, err := digitalocean.NewApp(ctx, "golang_sample", &digitalocean.AppArgs{
// 			Spec: &digitalocean.AppSpecArgs{
// 				Name:   pulumi.String("golang-sample"),
// 				Region: pulumi.String("ams"),
// 				Services: digitalocean.AppSpecServiceArray{
// 					&digitalocean.AppSpecServiceArgs{
// 						EnvironmentSlug: pulumi.String("go"),
// 						Git: &digitalocean.AppSpecServiceGitArgs{
// 							Branch:       pulumi.String("main"),
// 							RepoCloneUrl: pulumi.String("https://github.com/digitalocean/sample-golang.git"),
// 						},
// 						InstanceCount:    pulumi.Int(1),
// 						InstanceSizeSlug: pulumi.String("professional-xs"),
// 						Name:             pulumi.String("go-service"),
// 					},
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
// ### Static Site Example
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
// 		_, err := digitalocean.NewApp(ctx, "static_ste_example", &digitalocean.AppArgs{
// 			Spec: &digitalocean.AppSpecArgs{
// 				Name:   pulumi.String("static-ste-example"),
// 				Region: pulumi.String("ams"),
// 				StaticSites: digitalocean.AppSpecStaticSiteArray{
// 					&digitalocean.AppSpecStaticSiteArgs{
// 						BuildCommand: pulumi.String("bundle exec jekyll build -d ./public"),
// 						Git: &digitalocean.AppSpecStaticSiteGitArgs{
// 							Branch:       pulumi.String("main"),
// 							RepoCloneUrl: pulumi.String("https://github.com/digitalocean/sample-jekyll.git"),
// 						},
// 						Name:      pulumi.String("sample-jekyll"),
// 						OutputDir: pulumi.String("/public"),
// 					},
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
// ### Multiple Components Example
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
// 		_, err := digitalocean.NewApp(ctx, "mono_repo_example", &digitalocean.AppArgs{
// 			Spec: &digitalocean.AppSpecArgs{
// 				Databases: digitalocean.AppSpecDatabaseArray{
// 					&digitalocean.AppSpecDatabaseArgs{
// 						Engine:     pulumi.String("PG"),
// 						Name:       pulumi.String("starter-db"),
// 						Production: pulumi.Bool(false),
// 					},
// 				},
// 				Domains: pulumi.StringArray{
// 					pulumi.String("foo.example.com"),
// 				},
// 				Name:   pulumi.String("mono-repo-example"),
// 				Region: pulumi.String("ams"),
// 				Services: digitalocean.AppSpecServiceArray{
// 					&digitalocean.AppSpecServiceArgs{
// 						EnvironmentSlug: pulumi.String("go"),
// 						Github: &digitalocean.AppSpecServiceGithubArgs{
// 							Branch:       pulumi.String("main"),
// 							DeployOnPush: pulumi.Bool(true),
// 							Repo:         pulumi.String("username/repo"),
// 						},
// 						HttpPort:         pulumi.Int(3000),
// 						InstanceCount:    pulumi.Int(2),
// 						InstanceSizeSlug: pulumi.String("professional-xs"),
// 						Name:             pulumi.String("api"),
// 						Routes: &digitalocean.AppSpecServiceRoutesArgs{
// 							Path: pulumi.String("/api"),
// 						},
// 						RunCommand: pulumi.String("bin/api"),
// 						SourceDir:  pulumi.String("api/"),
// 					},
// 				},
// 				StaticSites: digitalocean.AppSpecStaticSiteArray{
// 					&digitalocean.AppSpecStaticSiteArgs{
// 						BuildCommand: pulumi.String("npm run build"),
// 						Github: &digitalocean.AppSpecStaticSiteGithubArgs{
// 							Branch:       pulumi.String("main"),
// 							DeployOnPush: pulumi.Bool(true),
// 							Repo:         pulumi.String("username/repo"),
// 						},
// 						Name: pulumi.String("web"),
// 						Routes: &digitalocean.AppSpecStaticSiteRoutesArgs{
// 							Path: pulumi.String("/"),
// 						},
// 					},
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
// An app can be imported using its `id`, e.g.
//
// ```sh
//  $ pulumi import digitalocean:index/app:App myapp fb06ad00-351f-45c8-b5eb-13523c438661
// ```
type App struct {
	pulumi.CustomResourceState

	// The ID the app's currently active deployment.
	ActiveDeploymentId pulumi.StringOutput `pulumi:"activeDeploymentId"`
	// The date and time of when the app was created.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The default URL to access the app.
	DefaultIngress pulumi.StringOutput `pulumi:"defaultIngress"`
	// The live URL of the app.
	LiveUrl pulumi.StringOutput `pulumi:"liveUrl"`
	// A DigitalOcean App spec describing the app.
	Spec AppSpecPtrOutput `pulumi:"spec"`
	// The date and time of when the app was last updated.
	UpdatedAt pulumi.StringOutput `pulumi:"updatedAt"`
}

// NewApp registers a new resource with the given unique name, arguments, and options.
func NewApp(ctx *pulumi.Context,
	name string, args *AppArgs, opts ...pulumi.ResourceOption) (*App, error) {
	if args == nil {
		args = &AppArgs{}
	}
	var resource App
	err := ctx.RegisterResource("digitalocean:index/app:App", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetApp gets an existing App resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetApp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AppState, opts ...pulumi.ResourceOption) (*App, error) {
	var resource App
	err := ctx.ReadResource("digitalocean:index/app:App", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering App resources.
type appState struct {
	// The ID the app's currently active deployment.
	ActiveDeploymentId *string `pulumi:"activeDeploymentId"`
	// The date and time of when the app was created.
	CreatedAt *string `pulumi:"createdAt"`
	// The default URL to access the app.
	DefaultIngress *string `pulumi:"defaultIngress"`
	// The live URL of the app.
	LiveUrl *string `pulumi:"liveUrl"`
	// A DigitalOcean App spec describing the app.
	Spec *AppSpec `pulumi:"spec"`
	// The date and time of when the app was last updated.
	UpdatedAt *string `pulumi:"updatedAt"`
}

type AppState struct {
	// The ID the app's currently active deployment.
	ActiveDeploymentId pulumi.StringPtrInput
	// The date and time of when the app was created.
	CreatedAt pulumi.StringPtrInput
	// The default URL to access the app.
	DefaultIngress pulumi.StringPtrInput
	// The live URL of the app.
	LiveUrl pulumi.StringPtrInput
	// A DigitalOcean App spec describing the app.
	Spec AppSpecPtrInput
	// The date and time of when the app was last updated.
	UpdatedAt pulumi.StringPtrInput
}

func (AppState) ElementType() reflect.Type {
	return reflect.TypeOf((*appState)(nil)).Elem()
}

type appArgs struct {
	// A DigitalOcean App spec describing the app.
	Spec *AppSpec `pulumi:"spec"`
}

// The set of arguments for constructing a App resource.
type AppArgs struct {
	// A DigitalOcean App spec describing the app.
	Spec AppSpecPtrInput
}

func (AppArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*appArgs)(nil)).Elem()
}

type AppInput interface {
	pulumi.Input

	ToAppOutput() AppOutput
	ToAppOutputWithContext(ctx context.Context) AppOutput
}

func (App) ElementType() reflect.Type {
	return reflect.TypeOf((*App)(nil)).Elem()
}

func (i App) ToAppOutput() AppOutput {
	return i.ToAppOutputWithContext(context.Background())
}

func (i App) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppOutput)
}

type AppOutput struct {
	*pulumi.OutputState
}

func (AppOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppOutput)(nil)).Elem()
}

func (o AppOutput) ToAppOutput() AppOutput {
	return o
}

func (o AppOutput) ToAppOutputWithContext(ctx context.Context) AppOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AppOutput{})
}
