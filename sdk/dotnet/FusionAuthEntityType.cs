// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace theogravity.Fusionauth
{
    /// <summary>
    /// ## # Entity Type Resource
    /// 
    /// Entity Types categorize Entities. For example, an Entity Type could be `Device`, `API` or `Company`.
    /// 
    /// [Entity Type API](https://fusionauth.io/docs/v1/tech/apis/entity-management/entity-types/#create-an-entity-type)
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Fusionauth = theogravity.Fusionauth;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var company = new Fusionauth.FusionAuthEntityType("company", new Fusionauth.FusionAuthEntityTypeArgs
    ///         {
    ///             Data = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "createdBy", "jared@fusionauth.io" },
    ///             }),
    ///             JwtConfiguration = new Fusionauth.Inputs.FusionAuthEntityTypeJwtConfigurationArgs
    ///             {
    ///                 AccessTokenKeyId = "a7516c7c-6234-4021-b0b4-8870c807aeb2",
    ///                 Enabled = true,
    ///                 TimeToLiveInSeconds = 3600,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [FusionauthResourceType("fusionauth:index/fusionAuthEntityType:FusionAuthEntityType")]
    public partial class FusionAuthEntityType : Pulumi.CustomResource
    {
        /// <summary>
        /// An object that can hold any information about the Entity Type that should be persisted. Must be a
        /// JSON string.
        /// </summary>
        [Output("data")]
        public Output<string?> Data { get; private set; } = null!;

        /// <summary>
        /// The ID to use for the new Entity Type. If not specified a secure random UUID will be
        /// generated.
        /// </summary>
        [Output("entityTypeId")]
        public Output<string> EntityTypeId { get; private set; } = null!;

        /// <summary>
        /// A block to configure JSON Web Token (JWT) options.
        /// </summary>
        [Output("jwtConfiguration")]
        public Output<Outputs.FusionAuthEntityTypeJwtConfiguration> JwtConfiguration { get; private set; } = null!;

        /// <summary>
        /// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a FusionAuthEntityType resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public FusionAuthEntityType(string name, FusionAuthEntityTypeArgs? args = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthEntityType:FusionAuthEntityType", name, args ?? new FusionAuthEntityTypeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private FusionAuthEntityType(string name, Input<string> id, FusionAuthEntityTypeState? state = null, CustomResourceOptions? options = null)
            : base("fusionauth:index/fusionAuthEntityType:FusionAuthEntityType", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/theogravity/pulumi-fusionauth/releases/download/v${VERSION}",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing FusionAuthEntityType resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static FusionAuthEntityType Get(string name, Input<string> id, FusionAuthEntityTypeState? state = null, CustomResourceOptions? options = null)
        {
            return new FusionAuthEntityType(name, id, state, options);
        }
    }

    public sealed class FusionAuthEntityTypeArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object that can hold any information about the Entity Type that should be persisted. Must be a
        /// JSON string.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The ID to use for the new Entity Type. If not specified a secure random UUID will be
        /// generated.
        /// </summary>
        [Input("entityTypeId")]
        public Input<string>? EntityTypeId { get; set; }

        /// <summary>
        /// A block to configure JSON Web Token (JWT) options.
        /// </summary>
        [Input("jwtConfiguration")]
        public Input<Inputs.FusionAuthEntityTypeJwtConfigurationArgs>? JwtConfiguration { get; set; }

        /// <summary>
        /// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FusionAuthEntityTypeArgs()
        {
        }
    }

    public sealed class FusionAuthEntityTypeState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// An object that can hold any information about the Entity Type that should be persisted. Must be a
        /// JSON string.
        /// </summary>
        [Input("data")]
        public Input<string>? Data { get; set; }

        /// <summary>
        /// The ID to use for the new Entity Type. If not specified a secure random UUID will be
        /// generated.
        /// </summary>
        [Input("entityTypeId")]
        public Input<string>? EntityTypeId { get; set; }

        /// <summary>
        /// A block to configure JSON Web Token (JWT) options.
        /// </summary>
        [Input("jwtConfiguration")]
        public Input<Inputs.FusionAuthEntityTypeJwtConfigurationGetArgs>? JwtConfiguration { get; set; }

        /// <summary>
        /// A descriptive name for the entity type (i.e. `Customer` or `Email_Service`).
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public FusionAuthEntityTypeState()
        {
        }
    }
}