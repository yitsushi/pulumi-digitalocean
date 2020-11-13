// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.DigitalOcean
{
    /// <summary>
    /// Provides a resource which can be used to create a custom Image from a URL
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using DigitalOcean = Pulumi.DigitalOcean;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var flatcar = new DigitalOcean.CustomImage("flatcar", new DigitalOcean.CustomImageArgs
    ///         {
    ///             Regions = 
    ///             {
    ///                 "nyc3",
    ///             },
    ///             Url = "https://stable.release.flatcar-linux.net/amd64-usr/2605.7.0/flatcar_production_digitalocean_image.bin.bz2",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class CustomImage : Pulumi.CustomResource
    {
        [Output("createdAt")]
        public Output<string> CreatedAt { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        [Output("distribution")]
        public Output<string?> Distribution { get; private set; } = null!;

        [Output("imageId")]
        public Output<int> ImageId { get; private set; } = null!;

        [Output("minDiskSize")]
        public Output<int> MinDiskSize { get; private set; } = null!;

        /// <summary>
        /// A name for the Custom Image.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("public")]
        public Output<bool> Public { get; private set; } = null!;

        /// <summary>
        /// A list of regions. (Currently only one is supported)
        /// </summary>
        [Output("regions")]
        public Output<ImmutableArray<string>> Regions { get; private set; } = null!;

        [Output("sizeGigabytes")]
        public Output<double> SizeGigabytes { get; private set; } = null!;

        [Output("slug")]
        public Output<string> Slug { get; private set; } = null!;

        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        [Output("tags")]
        public Output<ImmutableArray<string>> Tags { get; private set; } = null!;

        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// A URL from which the custom Linux virtual machine image may be retrieved.
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;


        /// <summary>
        /// Create a CustomImage resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CustomImage(string name, CustomImageArgs args, CustomResourceOptions? options = null)
            : base("digitalocean:index/customImage:CustomImage", name, args ?? new CustomImageArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CustomImage(string name, Input<string> id, CustomImageState? state = null, CustomResourceOptions? options = null)
            : base("digitalocean:index/customImage:CustomImage", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing CustomImage resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CustomImage Get(string name, Input<string> id, CustomImageState? state = null, CustomResourceOptions? options = null)
        {
            return new CustomImage(name, id, state, options);
        }
    }

    public sealed class CustomImageArgs : Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("distribution")]
        public Input<string>? Distribution { get; set; }

        /// <summary>
        /// A name for the Custom Image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("regions", required: true)]
        private InputList<string>? _regions;

        /// <summary>
        /// A list of regions. (Currently only one is supported)
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        /// <summary>
        /// A URL from which the custom Linux virtual machine image may be retrieved.
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        public CustomImageArgs()
        {
        }
    }

    public sealed class CustomImageState : Pulumi.ResourceArgs
    {
        [Input("createdAt")]
        public Input<string>? CreatedAt { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("distribution")]
        public Input<string>? Distribution { get; set; }

        [Input("imageId")]
        public Input<int>? ImageId { get; set; }

        [Input("minDiskSize")]
        public Input<int>? MinDiskSize { get; set; }

        /// <summary>
        /// A name for the Custom Image.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("public")]
        public Input<bool>? Public { get; set; }

        [Input("regions")]
        private InputList<string>? _regions;

        /// <summary>
        /// A list of regions. (Currently only one is supported)
        /// </summary>
        public InputList<string> Regions
        {
            get => _regions ?? (_regions = new InputList<string>());
            set => _regions = value;
        }

        [Input("sizeGigabytes")]
        public Input<double>? SizeGigabytes { get; set; }

        [Input("slug")]
        public Input<string>? Slug { get; set; }

        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tags")]
        private InputList<string>? _tags;
        public InputList<string> Tags
        {
            get => _tags ?? (_tags = new InputList<string>());
            set => _tags = value;
        }

        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// A URL from which the custom Linux virtual machine image may be retrieved.
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        public CustomImageState()
        {
        }
    }
}
