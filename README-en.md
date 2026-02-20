# Yandex Cloud Client Go

![Yandex Cloud Client Go](https://github.com/user-attachments/assets/f75920a6-c0cd-4da3-9223-5d7661ad3e47)

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/) [![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE) [![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fyandex--cloud--client--go-181717?style=flat&logo=github)](https://github.com/tigusigalpa/yandex-cloud-client-go)

English | [Русский](README.md)

Go client for the [Yandex Cloud API](https://yandex.cloud/docs). Manage organizations, clouds, folders, service accounts, and IAM authentication with idiomatic Go.

---

## Features

### Authentication & Security

- OAuth 2.0 token support
- Automatic IAM token generation
- Token caching with auto-refresh (12h lifecycle)
- Thread-safe operations

### Resource Management

- **Organizations** — CRUD and access control
- **Clouds** — lifecycle management
- **Folders** — operations and permissions
- **Service Accounts** — lifecycle and access
- **User Accounts** — lookup by ID or login
- **API Keys** — create and manage
- **Refresh Tokens** — token lifecycle

### Developer Experience

- Idiomatic Go with `context` support
- Type-safe API
- Dependencies: minimal
- Goroutine-safe
- Go 1.21+ with generics
- Custom error types

## Requirements

- Go 1.21+

## Quick Start

### Installation

```bash
go get github.com/tigusigalpa/yandex-cloud-client-go
```

### Getting an OAuth Token

1. Go to the [Yandex OAuth page](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
2. Authorize the application
3. Copy the token

Store tokens in environment variables, not in code. See [Yandex Cloud docs](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token) for details.

### Basic Usage

```go
package main

import (
    "context"
    "fmt"
    "log"

    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
)

func main() {
    // Create a new context
    ctx := context.Background()

    // Initialize the client with your OAuth token
    client, err := yandexcloud.NewClient("your_oauth_token", nil)
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    // Let's list all your organizations
    organizations, err := client.Organizations().List(ctx, nil, nil)
    if err != nil {
        log.Fatalf("Failed to list organizations: %v", err)
    }

    fmt.Printf("Found organizations: %+v\n", organizations)

    // Now, let's list the clouds in a specific organization
    // Replace with your actual organization ID
    orgID := "your_organization_id"
    clouds, err := client.Clouds().List(ctx, &orgID, nil, nil)
    if err != nil {
        log.Fatalf("Failed to list clouds: %v", err)
    }

    fmt.Printf("Clouds in organization %s: %+v\n", orgID, clouds)
}
```

## Usage

### Organizations

```go
// List organizations with a page size of 100
pageSize := 100
organizations, err := client.Organizations().List(ctx, &pageSize, nil)

// Get a specific organization
org, err := client.Organizations().Get(ctx, "organization_id")

// Update organization
updateData := map[string]interface{}{
    "name":        "A Brand New Name",
    "description": "A fresh new description",
}
org, err = client.Organizations().Update(ctx, "organization_id", updateData)

// Add role
result, err := client.Organizations().AddRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Remove role
result, err = client.Organizations().RemoveRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// List access bindings
bindings, err := client.Organizations().ListAccessBindings(ctx, "organization_id", nil, nil)
```

### Clouds

```go
// List clouds within a specific organization
orgID := "org_id"
pageSize := 100
clouds, err := client.Clouds().List(ctx, &orgID, &pageSize, nil)

// Get cloud details
cloud, err := client.Clouds().Get(ctx, "cloud_id")

// Create cloud
description := "Production cloud"
labels := map[string]string{"env": "production"}
cloud, err = client.Clouds().Create(
    ctx,
    "org_id",
    "My Production Cloud",
    &description,
    labels,
)

// Update cloud
updateData := map[string]interface{}{
    "name":        "Updated Cloud Name",
    "description": "Updated description",
}
cloud, err = client.Clouds().Update(ctx, "cloud_id", updateData)

// Delete cloud
result, err := client.Clouds().Delete(ctx, "cloud_id")

// Add role
result, err = client.Clouds().AddRole(
    ctx,
    "cloud_id",
    "user_id",
    "editor",
    "userAccount",
)

// List all access bindings for a cloud
bindings, err := client.Clouds().ListAccessBindings(ctx, "cloud_id", nil, nil)
```

---

## Advanced Access Control

### Adding Multiple Roles

```go
// Add multiple roles to a folder in a single request
deltas := []map[string]interface{}{
    {
        "action": "ADD",
        "accessBinding": map[string]interface{}{
            "roleId": "editor",
            "subject": map[string]interface{}{
                "id":   "user_id_1",
                "type": "userAccount",
            },
        },
    },
    {
        "action": "ADD",
        "accessBinding": map[string]interface{}{
            "roleId": "viewer",
            "subject": map[string]interface{}{
                "id":   "user_id_2",
                "type": "userAccount",
            },
        },
    },
}

result, err := client.Folders().UpdateAccessBindings(ctx, "folder_id", deltas)
```

### Replacing All Access Bindings

```go
// Replace all access bindings
bindings := []map[string]interface{}{
    {
        "roleId": "admin",
        "subject": map[string]interface{}{
            "id":   "user_id",
            "type": "userAccount",
        },
    },
}

result, err := client.Clouds().SetAccessBindings(ctx, "cloud_id", bindings)
```

### Assigning Roles by User Login

```go
// Get user ID by Yandex Passport login
user, err := client.YandexPassportUserAccounts().GetByLogin(ctx, "username@yandex.ru")
if err != nil {
    log.Fatal(err)
}
userID := user["id"].(string)

// Assign role to folder
result, err := client.Folders().AddRole(
    ctx,
    "folder_id",
    userID,
    "ai.languageModels.user",
    "userAccount",
)

// Or assign role to cloud
result, err = client.Clouds().AddRole(
    ctx,
    "cloud_id",
    userID,
    "editor",
    "userAccount",
)
```

---

## Error Handling

```go
import (
    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
    "github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

client, err := yandexcloud.NewClient("your_oauth_token", nil)
if err != nil {
    log.Fatal(err)
}

_, err = client.Clouds().List(ctx, nil, nil, nil)
if err != nil {
    switch e := err.(type) {
    case *errors.AuthenticationError:
        log.Printf("Authentication failed: %v", e)
    case *errors.ValidationError:
        log.Printf("Validation error: %v", e)
    case *errors.APIError:
        log.Printf("API error (status %d): %v", e.StatusCode, e)
    default:
        log.Printf("Unexpected error: %v", err)
    }
}
```

---

## Testing

```bash
go test ./...
go test -cover ./...
go test -v ./...
```

---

## Contributing

```bash
git clone https://github.com/tigusigalpa/yandex-cloud-client-go.git
cd yandex-cloud-client-go
go mod download
cp .env.example .env
```

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/my-feature`)
3. Commit changes (`git commit -m 'Add feature'`)
4. Push to branch (`git push origin feature/my-feature`)
5. Open a Pull Request

Please follow Go conventions, add tests for new features, and keep documentation up to date.

---

## Security

If you discover a security vulnerability, please email **sovletig@gmail.com** instead of creating a public issue.

---

## Author

**Igor Sazonov** ([@tigusigalpa](https://github.com/tigusigalpa))

- **Email**: sovletig@gmail.com

---

## License

MIT License — see [LICENSE](LICENSE) for details.

---

## Related Packages

| Package                  | Description                | Links                                                                                                                                              |
|--------------------------|----------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| **YandexGPT Go**         | Go SDK for YandexGPT API   | [GitHub](https://github.com/tigusigalpa/yandexgpt-go)                                                                                             |
| **YandexGPT PHP**        | PHP SDK for YandexGPT API  | [GitHub](https://github.com/tigusigalpa/yandexgpt-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandexgpt-php)                     |
| **Yandex Cloud Client PHP** | PHP SDK for Yandex Cloud | [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandex-cloud-client-php) |

---

## Links

- [Yandex Cloud Documentation](https://yandex.cloud/docs)
- [Organization API Reference](https://yandex.cloud/ru/docs/organization/api-ref/)
- [Resource Manager API Reference](https://yandex.cloud/ru/docs/resource-manager/api-ref/)
- [IAM API Reference](https://yandex.cloud/ru/docs/iam/api-ref/)
- [GitHub Repository](https://github.com/tigusigalpa/yandex-cloud-client-go)
- [Issues](https://github.com/tigusigalpa/yandex-cloud-client-go/issues)
- [Discussions](https://github.com/tigusigalpa/yandex-cloud-client-go/discussions)
