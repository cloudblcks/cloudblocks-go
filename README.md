# Cloudblocks Credentials Generator Go SDK

## Authentication

To authenticate with Cloudblocks Credentials Generator Go SDK provide username and password as parameters when initialising the client: 

```go
client, err := cloudblocks_go.NewClient(username, password)
```

If you need to specify a custom base API URL (for example for testing purposes), provide it as a third parameter:

```go
client, err := cloudblocks_go.NewClient(username, password, apiURL)
```

## Usage

To request credentials, provide originId (ResourceID of the resource that requests access, typically the entity that is running the code) and targetId (ResourceID of the resource that you need to access): 

```go
creds, err := client.RequestCredentials(originId, targetId)
```

If there is an Access Role assigned to the origin resource that allows access to the target resource, credentials will be returned. Alternatively an error will be returned.
