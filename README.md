<div align="center">

# ☁️ Yandex Cloud Client Go

![Yandex Cloud Client Go](https://github.com/user-attachments/assets/f75920a6-c0cd-4da3-9223-5d7661ad3e47)

### 🚀 Modern Go SDK for Yandex Cloud API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fyandex--cloud--client--go-181717?style=flat&logo=github)](https://github.com/tigusigalpa/yandex-cloud-client-go)

**🌐 Language:** English | [Русский](README-ru.md)

**Powerful, elegant, and developer-friendly Go SDK for Yandex Cloud API.**

Manage organizations, clouds, folders, and IAM authentication with clean, idiomatic Go code.

</div>

---

## ✨ Features

<table>
<tr>
<td width="50%">

### 🔐 Authentication & Security

- **OAuth 2.0** token support
- **Automatic IAM** token generation
- **Smart caching** with auto-refresh
- **Token expiry** management (12h)
- **Thread-safe** operations

### 🏢 Resource Management

- **Organizations** - Full CRUD & access control
- **Clouds** - Complete lifecycle management
- **Folders** - Operations & permissions
- **Service Accounts** - Full lifecycle & access
- **User Accounts** - Get user info by ID or login
- **API Keys** - Create & manage API keys
- **Refresh Tokens** - Token lifecycle

</td>
<td width="50%">

### 🎯 Go Features

- **Idiomatic Go** code
- **Type-safe** API
- **Context support** ready
- **Minimal dependencies**
- **Goroutine-safe**

### 💎 Code Quality

- **Go 1.21+** with generics support
- **Full error handling**
- **Clean architecture**
- **Well documented**

</td>
</tr>
</table>

## 📋 Requirements

| Requirement | Version |
|-------------|---------|
| Go          | 1.21+   |

## 🚀 Quick Start

### Installation

```bash
go get github.com/tigusigalpa/yandex-cloud-client-go
```

### Get Your OAuth Token

<details>
<summary>📝 Click to see how to get OAuth token</summary>

1. Visit [Yandex OAuth](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
2. Authorize the application
3. Copy the token
4. Use it in your code

💡 **Tip**: Store tokens securely in environment variables!

For more details, see [Yandex Cloud Documentation](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token).

</details>

### Basic Usage

```go
package main

import (
    "fmt"
    "log"
    
    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
)

func main() {
    // Initialize client
    client, err := yandexcloud.NewClient("your_oauth_token", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // List all organizations
    organizations, err := client.Organizations().List(nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Organizations: %+v\n", organizations)
    
    // List clouds in organization
    orgID := "your_organization_id"
    clouds, err := client.Clouds().List(&orgID, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Clouds: %+v\n", clouds)
}
```

## 💻 Usage Examples

### Organizations API

```go
// List organizations
pageSize := 100
organizations, err := client.Organizations().List(&pageSize, nil)

// Get organization
org, err := client.Organizations().Get("organization_id")

// Update organization
updateData := map[string]interface{}{
    "name": "New Name",
    "description": "New Description",
}
org, err = client.Organizations().Update("organization_id", updateData)

// Add role to organization
result, err := client.Organizations().AddRole(
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Remove role from organization
result, err = client.Organizations().RemoveRole(
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// List access bindings
bindings, err := client.Organizations().ListAccessBindings("organization_id", nil, nil)
```

### Clouds API

```go
// List clouds
orgID := "org_id"
pageSize := 100
clouds, err := client.Clouds().List(&orgID, &pageSize, nil)

// Get cloud
cloud, err := client.Clouds().Get("cloud_id")

// Create cloud
description := "Production cloud"
labels := map[string]string{"env": "production"}
cloud, err = client.Clouds().Create(
    "org_id",
    "My Cloud",
    &description,
    labels,
)

// Update cloud
updateData := map[string]interface{}{
    "name": "Updated Name",
    "description": "Updated Description",
}
cloud, err = client.Clouds().Update("cloud_id", updateData)

// Delete cloud
result, err := client.Clouds().Delete("cloud_id")

// Add role to cloud
result, err = client.Clouds().AddRole(
    "cloud_id",
    "user_id",
    "editor",
    "userAccount",
)

// List access bindings
bindings, err := client.Clouds().ListAccessBindings("cloud_id", nil, nil)
```

### Folders API

```go
// List folders
pageSize := 100
folders, err := client.Folders().List("cloud_id", &pageSize, nil)

// Get folder
folder, err := client.Folders().Get("folder_id")

// Create folder
description := "Development folder"
labels := map[string]string{"team": "backend"}
folder, err = client.Folders().Create(
    "cloud_id",
    "My Folder",
    &description,
    labels,
)

// Update folder
updateData := map[string]interface{}{
    "name": "Updated Name",
}
folder, err = client.Folders().Update("folder_id", updateData)

// Delete folder
result, err := client.Folders().Delete("folder_id")

// List operations
operations, err := client.Folders().ListOperations("folder_id", nil, nil)

// Add role to folder
result, err = client.Folders().AddRole(
    "folder_id",
    "user_id",
    "ai.languageModels.user",
    "userAccount",
)

// List access bindings
bindings, err := client.Folders().ListAccessBindings("folder_id", nil, nil)
```

### Service Accounts API

```go
// List service accounts in folder
pageSize := 100
serviceAccounts, err := client.ServiceAccounts().List("folder_id", &pageSize, nil)

// Get service account
sa, err := client.ServiceAccounts().Get("service_account_id")

// Create service account
description := "Service account for API access"
sa, err = client.ServiceAccounts().Create(
    "folder_id",
    "my-service-account",
    &description,
)

// Update service account
updateData := map[string]interface{}{
    "name": "Updated name",
    "description": "Updated description",
}
sa, err = client.ServiceAccounts().Update("service_account_id", updateData)

// Delete service account
result, err := client.ServiceAccounts().Delete("service_account_id")

// Add role to service account
result, err = client.ServiceAccounts().AddRole(
    "service_account_id",
    "user_id",
    "editor",
    "userAccount",
)

// List access bindings
bindings, err := client.ServiceAccounts().ListAccessBindings("service_account_id", nil, nil)
```

### User Accounts API

```go
// Get user account by ID
user, err := client.UserAccounts().Get("user_account_id")

// Get user by Yandex Passport login (to get user ID for access control)
user, err = client.YandexPassportUserAccounts().GetByLogin("username")
// Returns: map with 'id', 'login', etc.

// Use the ID to assign roles
userID := user["id"].(string)
result, err := client.Folders().AddRole(
    "folder_id",
    userID,
    "editor",
    "userAccount",
)
```

### API Keys

```go
// List API keys for service account
pageSize := 100
keys, err := client.APIKeys().List("service_account_id", &pageSize, nil)

// Get API key
key, err := client.APIKeys().Get("api_key_id")

// Create API key (secret is shown only once!)
description := "API key for production"
key, err = client.APIKeys().Create("service_account_id", &description)
// Save key["secret"] immediately - it won't be shown again!

// Update API key
updateData := map[string]interface{}{
    "description": "Updated description",
}
key, err = client.APIKeys().Update("api_key_id", updateData)

// Delete API key
result, err := client.APIKeys().Delete("api_key_id")
```

### Refresh Tokens API

```go
// List refresh tokens
pageSize := 100
tokens, err := client.RefreshTokens().List(&pageSize, nil)

// Revoke refresh token
result, err := client.RefreshTokens().Revoke("token_id")
```

---

## 🔐 Advanced Access Control

### Adding Multiple Roles at Once

```go
// Add multiple roles to a folder
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

result, err := client.Folders().UpdateAccessBindings("folder_id", deltas)
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

result, err := client.Clouds().SetAccessBindings("cloud_id", bindings)
```

### Assigning Roles by User Login

```go
// Get user ID by Yandex Passport login
user, err := client.YandexPassportUserAccounts().GetByLogin("username@yandex.ru")
if err != nil {
    log.Fatal(err)
}
userID := user["id"].(string)

// Assign role to folder using the user ID
result, err := client.Folders().AddRole(
    "folder_id",
    userID,
    "ai.languageModels.user",
    "userAccount",
)

// Or assign to cloud
result, err = client.Clouds().AddRole(
    "cloud_id",
    userID,
    "editor",
    "userAccount",
)
```

---

## ⚠️ Error Handling

```go
import (
    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
    "github.com/tigusigalpa/yandex-cloud-client-go/errors"
)

client, err := yandexcloud.NewClient("oauth_token", nil)
if err != nil {
    log.Fatal(err)
}

clouds, err := client.Clouds().List(nil, nil, nil)
if err != nil {
    switch e := err.(type) {
    case *errors.AuthenticationError:
        // Handle authentication errors
        log.Printf("Authentication failed: %v", e)
    case *errors.ValidationError:
        // Handle validation errors
        log.Printf("Validation error: %v", e)
    case *errors.APIError:
        // Handle API errors
        log.Printf("API error (status %d): %v", e.StatusCode, e)
    default:
        // Handle other errors
        log.Printf("Error: %v", err)
    }
}
```

---

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

---

## 🤝 Contributing

We welcome contributions! Here's how you can help:

### Development Setup

```bash
# Clone repository
git clone https://github.com/tigusigalpa/yandex-cloud-client-go.git
cd yandex-cloud-client-go

# Install dependencies
go mod download

# Copy environment file
cp .env.example .env
```

### Contribution Guidelines

- ✅ **Follow Go conventions** and best practices
- ✅ **Write idiomatic Go** code
- ✅ **Write tests** for new features
- ✅ **Update documentation** as needed
- ✅ **One feature per PR** - keep it focused

### Pull Request Process

1. Fork the repository
2. Create feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests and ensure they pass
5. Commit changes (`git commit -m 'Add amazing feature'`)
6. Push to branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

---

## 🔒 Security

If you discover any security vulnerabilities, please email **sovletig@gmail.com** instead of using the issue tracker.

We take security seriously and will respond promptly.

---

## 📦 Deployment & Publishing

<details>
<summary>📋 Click to see deployment checklist</summary>

### Pre-Deployment

```bash
# Run tests
go test ./...

# Format code
go fmt ./...

# Run linter
go vet ./...

# Verify module
go mod verify
```

### GitHub Deployment

```bash
# Initialize repository
git init
git add .
git commit -m "Initial commit: v1.0.0"

# Push to GitHub
git remote add origin https://github.com/tigusigalpa/yandex-cloud-client-go.git
git branch -M main
git push -u origin main

# Create release
git tag v1.0.0
git push origin v1.0.0
```

### Version Numbering (Semantic Versioning)

- **MAJOR** (v1.x.x) - Breaking changes
- **MINOR** (vx.1.x) - New features, backwards-compatible
- **PATCH** (vx.x.1) - Bug fixes, backwards-compatible

</details>

---

## 👨‍💻 Author & Contributors

**Created with ❤️ by [Igor Sazonov](https://github.com/tigusigalpa)**

- 📧 Email: sovletig@gmail.com
- 🐙 GitHub: [@tigusigalpa](https://github.com/tigusigalpa)

### Contributors

Thanks to [all contributors](../../contributors) who help improve this package!

---

## 📄 License

This project is licensed under the **MIT License** - see the [LICENSE](LICENSE) file for details.

Free to use in personal and commercial projects. ✨

---

## 🔗 Related Packages

Explore our other Yandex Cloud packages:

| Package                  | Description           | Links                                                                                                                                              |
|--------------------------|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| **YandexGPT Go**         | SDK for YandexGPT API | [GitHub](https://github.com/tigusigalpa/yandexgpt-go)                                                                                             |
| **YandexGPT PHP**        | SDK for YandexGPT API | [GitHub](https://github.com/tigusigalpa/yandexgpt-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandexgpt-php)                     |
| **Yandex Cloud Client PHP** | PHP SDK for Yandex Cloud | [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandex-cloud-client-php) |

---

## 🔗 Useful Links

### Official Documentation

- 📖 [Yandex Cloud Documentation](https://yandex.cloud/docs)
- 🏢 [Organization API Reference](https://yandex.cloud/ru/docs/organization/api-ref/)
- ☁️ [Resource Manager API Reference](https://yandex.cloud/ru/docs/resource-manager/api-ref/)
- 🔐 [IAM API Reference](https://yandex.cloud/ru/docs/iam/api-ref/)

### Package Resources

- 🐙 [GitHub Repository](https://github.com/tigusigalpa/yandex-cloud-client-go)
- 🐛 [Issue Tracker](https://github.com/tigusigalpa/yandex-cloud-client-go/issues)
- 💬 [Discussions](https://github.com/tigusigalpa/yandex-cloud-client-go/discussions)

---

<div align="center">

### ⭐ Star us on GitHub!

If this package helped you, please consider giving it a star ⭐

**Made with ❤️ for the Go community**

[Report Bug](https://github.com/tigusigalpa/yandex-cloud-client-go/issues) • [Request Feature](https://github.com/tigusigalpa/yandex-cloud-client-go/issues) • [Contribute](https://github.com/tigusigalpa/yandex-cloud-client-go/pulls)

</div>
