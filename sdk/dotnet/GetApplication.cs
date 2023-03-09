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
    public static class GetApplication
    {
        /// <summary>
        /// ## # Application Resource
        /// 
        /// [Applications API](https://fusionauth.io/docs/v1/tech/apis/applications)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fusionAuth = Fusionauth.GetApplication.Invoke(new()
        ///     {
        ///         Name = "FusionAuth",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetApplicationResult> InvokeAsync(GetApplicationArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetApplicationResult>("fusionauth:index/getApplication:getApplication", args ?? new GetApplicationArgs(), options.WithDefaults());

        /// <summary>
        /// ## # Application Resource
        /// 
        /// [Applications API](https://fusionauth.io/docs/v1/tech/apis/applications)
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using Pulumi;
        /// using Fusionauth = Pulumi.Fusionauth;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var fusionAuth = Fusionauth.GetApplication.Invoke(new()
        ///     {
        ///         Name = "FusionAuth",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetApplicationResult> Invoke(GetApplicationInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetApplicationResult>("fusionauth:index/getApplication:getApplication", args ?? new GetApplicationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetApplicationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetApplicationArgs()
        {
        }
        public static new GetApplicationArgs Empty => new GetApplicationArgs();
    }

    public sealed class GetApplicationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Application.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetApplicationInvokeArgs()
        {
        }
        public static new GetApplicationInvokeArgs Empty => new GetApplicationInvokeArgs();
    }


    [OutputType]
    public sealed class GetApplicationResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;

        [OutputConstructor]
        private GetApplicationResult(
            string id,

            string name)
        {
            Id = id;
            Name = name;
        }
    }
}
