# ☁️ Yandex Cloud Client Go: Your Ultimate Go SDK for Yandex Cloud

![Yandex Cloud Client Go](https://github.com/user-attachments/assets/f75920a6-c0cd-4da3-9223-5d7661ad3e47)

### 🚀 The Modern, Developer-First Go SDK for the Yandex Cloud API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/) [![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE) [![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fyandex--cloud--client--go-181717?style=flat&logo=github)](https://github.com/tigusigalpa/yandex-cloud-client-go)

**🌐 Language:** English | [Русский](README.md)

Welcome to the **Yandex Cloud Client for Go**! We've built this SDK to be the most powerful, elegant, and developer-friendly tool for interacting with the Yandex Cloud API. If you're looking to manage organizations, clouds, folders, or handle IAM authentication, our library empowers you to do it all with clean, idiomatic Go. Our goal is to provide an intuitive and efficient experience, so you can focus on what truly matters: building amazing applications.

---

## ✨ What Makes Our SDK Special?

We've packed our SDK with features designed to make your development process as smooth as possible.

<table>
<tr>
<td width="50%">

### 🔐 Seamless Authentication & Top-Notch Security

- **Effortless OAuth 2.0**: Full support for OAuth 2.0 tokens.
- **Automatic IAM Tokens**: We handle the generation of IAM tokens for you.
- **Intelligent Caching**: Smart, auto-refreshing cache for optimal performance.
- **Worry-Free Token Expiry**: Automatic management of 12-hour token lifecycles.
- **Built for Concurrency**: Completely thread-safe operations.

### 🏢 Comprehensive Resource Management

- **Organizations**: Full CRUD operations and fine-grained access control.
- **Clouds**: Complete lifecycle management, from creation to deletion.
- **Folders**: All the operations and permission management you need.
- **Service Accounts**: Full control over the lifecycle and access of service accounts.
- **User Accounts**: Easily fetch user information by ID or login.
- **API Keys**: Create and manage API keys with ease.
- **Refresh Tokens**: Full control over the token lifecycle.

</td>
<td width="50%">

### 🎯 Designed for Go Developers

- **Truly Idiomatic Go**: Code that feels natural and intuitive.
- **Type-Safe API**: Catch errors at compile time, not runtime.
- **Context-Aware**: Ready for cancellation and deadlines with `context` support.
- **Lean & Mean**: Minimal dependencies to keep your projects light.
- **Goroutine-Safe**: Built for concurrent applications from the ground up.

### 💎 Uncompromising Code Quality

- **Modern Go**: Built with Go 1.21+ and leveraging generics.
- **Robust Error Handling**: Comprehensive and easy-to-use error types.
- **Clean by Design**: A clean architecture for maintainability and extensibility.
- **Superbly Documented**: Clear and thorough documentation to guide you.

</td>
</tr>
</table>

## 📋 Before You Begin: Requirements

| Requirement | Version |
|-------------|---------|
| Go          | 1.21+   |

## 🚀 Getting Started in Minutes

### Step 1: Installation

It's as simple as running:

```bash
go get github.com/tigusigalpa/yandex-cloud-client-go
```

### Step 2: Grab Your OAuth Token

<details>
<summary>📝 Here's how to get your OAuth token</summary>

1.  Head over to the [Yandex OAuth page](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb).
2.  Authorize the application to access your account.
3.  Copy the token that appears.
4.  You're ready to use it in your code!

💡 **Pro Tip**: For better security, store your tokens in environment variables instead of hardcoding them!

For more details, check out the [official Yandex Cloud documentation](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token).

</details>

### Step 3: A Taste of the Basics

Here's a simple example to get you up and running:

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

## 💻 Practical Usage Examples

### Working with the Organizations API

```go
// List your organizations with a page size of 100
pageSize := 100
organizations, err := client.Organizations().List(ctx, &pageSize, nil)

// Get a specific organization by its ID
org, err := client.Organizations().Get(ctx, "organization_id")

// Update an organization's name and description
updateData := map[string]interface{}{
    "name":        "A Brand New Name",
    "description": "A fresh new description",
}
org, err = client.Organizations().Update(ctx, "organization_id", updateData)

// Grant a user the 'editor' role in an organization
result, err := client.Organizations().AddRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// And, just as easily, remove that role
result, err = client.Organizations().RemoveRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// See who has access to what
bindings, err := client.Organizations().ListAccessBindings(ctx, "organization_id", nil, nil)
```

### Managing Your Clouds

```go
// List clouds within a specific organization
orgID := "org_id"
pageSize := 100
clouds, err := client.Clouds().List(ctx, &orgID, &pageSize, nil)

// Get the details of a single cloud
cloud, err := client.Clouds().Get(ctx, "cloud_id")

// Create a new production cloud
description := "This is our main production cloud"
labels := map[string]string{"env": "production"}
cloud, err = client.Clouds().Create(
    ctx,
    "org_id",
    "My Production Cloud",
    &description,
    labels,
)

// Update a cloud's details
updateData := map[string]interface{}{
    "name":        "Updated Cloud Name",
    "description": "An even better description",
}
cloud, err = client.Clouds().Update(ctx, "cloud_id", updateData)

// Time to say goodbye: delete a cloud
result, err := client.Clouds().Delete(ctx, "cloud_id")

// Grant a user the 'editor' role for a cloud
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

## 🔐 Advanced Access Control

### Granting Multiple Roles at Once

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
// Completely replace all existing access bindings with a new set
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
// Find a user's ID by their Yandex Passport login
user, err := client.YandexPassportUserAccounts().GetByLogin(ctx, "username@yandex.ru")
if err != nil {
    log.Fatal(err)
}
userID := user["id"].(string)

// Now, use that ID to grant them a role in a folder
result, err := client.Folders().AddRole(
    ctx,
    "folder_id",
    userID,
    "ai.languageModels.user",
    "userAccount",
)

// Or grant them a role in a cloud
result, err = client.Clouds().AddRole(
    ctx,
    "cloud_id",
    userID,
    "editor",
    "userAccount",
)
```

---

## ⚠️ Graceful Error Handling

Our SDK provides detailed error types to help you handle issues gracefully.

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
        // Oh no, something's wrong with your credentials!
        log.Printf("Authentication failed: %v", e)
    case *errors.ValidationError:
        // Looks like you sent some invalid data.
        log.Printf("Validation error: %v", e)
    case *errors.APIError:
        // The Yandex Cloud API returned an error.
        log.Printf("API error with status %d: %v", e.StatusCode, e)
    default:
        // For all other kinds of errors.
        log.Printf("An unexpected error occurred: %v", err)
    }
}
```

---

## 🧪 Testing Your Code

We're big believers in testing! Here's how you can run the tests for this SDK:

```bash
# Run all tests
go test ./...

# Run tests and see your code coverage
go test -cover ./...

# Get all the details with verbose output
go test -v ./...
```

---

## 🤝 We Welcome Your Contributions!

We love community contributions! Here’s how you can help us make this SDK even better.

### Setting Up Your Development Environment

```bash
# First, clone the repository
git clone https://github.com/tigusigalpa/yandex-cloud-client-go.git
cd yandex-cloud-client-go

# Then, install all the dependencies
go mod download

# Create your own environment file
cp .env.example .env
```

### Our Contribution Guidelines

- ✅ **Follow Go's best practices** and conventions.
- ✅ **Write clean, idiomatic Go** code that others can easily understand.
- ✅ **Add tests** for any new features you introduce.
- ✅ **Keep the documentation updated** with your changes.
- ✅ **Focus on one feature per pull request** to keep things clear and manageable.

### The Pull Request Process

1.  Fork this repository.
2.  Create a new branch for your feature (`git checkout -b feature/my-awesome-feature`).
3.  Make your magic happen and commit your changes (`git commit -m 'feat: Add my awesome feature'`).
4.  Push your branch to your fork (`git push origin feature/my-awesome-feature`).
5.  Open a pull request and tell us about your great work!

---

## 🔒 Our Commitment to Security

If you discover a security vulnerability, please email us at **sovletig@gmail.com** instead of creating a public issue. We take security very seriously and will respond promptly to your report.

---

## 👨‍💻 Meet the Author & Our Amazing Contributors

**Crafted with ❤️ by [Igor Sazonov](https://github.com/tigusigalpa)**

- 📧 **Email**: sovletig@gmail.com
- 🐙 **GitHub**: [@tigusigalpa](https://github.com/tigusigalpa)

### Our Valued Contributors

A huge thank you to [all our contributors](../../contributors) who have helped improve this package!

---

## 📄 Our Open-Source License

This project is licensed under the **MIT License**. You can find all the details in the [LICENSE](LICENSE) file.

Feel free to use it in your personal and commercial projects. ✨

---

## 🔗 Explore Our Other Packages

We also have other packages for the Yandex Cloud ecosystem:

| Package                  | Description                               | Links                                                                                                                                              |
|--------------------------|-------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| **YandexGPT Go**         | A powerful Go SDK for the YandexGPT API.  | [GitHub](https://github.com/tigusigalpa/yandexgpt-go)                                                                                             |
| **YandexGPT PHP**        | A feature-rich PHP SDK for the YandexGPT API. | [GitHub](https://github.com/tigusigalpa/yandexgpt-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandexgpt-php)                     |
| **Yandex Cloud Client PHP** | A comprehensive PHP SDK for Yandex Cloud. | [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandex-cloud-client-php) |

---

## 🔗 Links You Might Find Useful

### Official Yandex Cloud Documentation

- 📖 [Yandex Cloud Documentation](https://yandex.cloud/docs)
- 🏢 [Organization API Reference](https://yandex.cloud/ru/docs/organization/api-ref/)
- ☁️ [Resource Manager API Reference](https://yandex.cloud/ru/docs/resource-manager/api-ref/)
- 🔐 [IAM API Reference](https://yandex.cloud/ru/docs/iam/api-ref/)

### Package Resources

- 🐙 [Our GitHub Repository](https://github.com/tigusigalpa/yandex-cloud-client-go)
- 🐛 [Found a Bug? Report It Here](https://github.com/tigusigalpa/yandex-cloud-client-go/issues)
- 💬 [Have Questions? Join the Discussion](https://github.com/tigusigalpa/yandex-cloud-client-go/discussions)
