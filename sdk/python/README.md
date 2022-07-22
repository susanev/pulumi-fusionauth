# FusionAuth Resource Provider

The FusionAuth Resource Provider lets you manage [FusionAuth](http://fusionauth.io/) resources.

This is bridged using the [Terraform FusionAuth](https://github.com/gpsinsight/terraform-provider-fusionauth) package

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @theogravity/pulumi-fusionauth
```

or `yarn`:

```bash
yarn add @theogravity/pulumi-fusionauth
```

### Python

To use from Python, install using `pip`:

```bash
pip install theogravity_pulumi-fusionauth
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/theogravity/pulumi-fusionauth/sdk/go/fusionauth
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package theogravity.Fusionauth
```

## Configuration

The following configuration points are available for the `fusionauth` provider:

- `fusionauth:api_key` (environment: `FUSION_AUTH_API_KEY`) - the API key for `fusionauth`
- `fusionauth:host` (environment: `FUSION_AUTH_HOST_URL`) - the URL to the FusionAuth instance with the trailing slash omitted (ex: `https://instance.fusionauth.io`)