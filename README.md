# Yandex Cloud Client Go

![Yandex Cloud Client Go](https://github.com/user-attachments/assets/f75920a6-c0cd-4da3-9223-5d7661ad3e47)

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/) [![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE) [![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fyandex--cloud--client--go-181717?style=flat&logo=github)](https://github.com/tigusigalpa/yandex-cloud-client-go)

Русский | [English](README-en.md)

Go-клиент для [API Яндекс.Облака](https://yandex.cloud/docs). Управление организациями, облаками, каталогами, сервисными аккаунтами и IAM-аутентификацией на идиоматичном Go.

---

## Возможности

### Аутентификация и безопасность

- Поддержка токенов OAuth 2.0
- Автоматическая генерация IAM-токенов
- Кеширование токенов с автообновлением (12-часовой цикл)
- Потокобезопасные операции

### Управление ресурсами

- **Организации** — CRUD и контроль доступа
- **Облака** — управление жизненным циклом
- **Каталоги** — операции и разрешения
- **Сервисные аккаунты** — жизненный цикл и доступ
- **Аккаунты пользователей** — поиск по ID или логину
- **API-ключи** — создание и управление
- **Refresh-токены** — жизненный цикл токенов

### Удобство разработки

- Идиоматичный Go с поддержкой `context`
- Типобезопасный API
- Минимальные зависимости
- Goroutine-safe
- Go 1.21+ с дженериками
- Пользовательские типы ошибок

## Требования

- Go 1.21+

## Быстрый старт

### Установка

```bash
go get github.com/tigusigalpa/yandex-cloud-client-go
```

### Получение OAuth-токена

1. Перейдите на [страницу Яндекс.OAuth](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
2. Разрешите приложению доступ
3. Скопируйте токен

Храните токены в переменных окружения, а не в коде. Подробнее в [документации Яндекс.Облака](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token).

### Основное использование

```go
package main

import (
    "context"
    "fmt"
    "log"

    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
)

func main() {
    // Создаем новый контекст
    ctx := context.Background()

    // Инициализируем клиент с вашим OAuth-токеном
    client, err := yandexcloud.NewClient("your_oauth_token", nil)
    if err != nil {
        log.Fatalf("Не удалось создать клиент: %v", err)
    }

    // Давайте получим список всех ваших организаций
    organizations, err := client.Organizations().List(ctx, nil, nil)
    if err != nil {
        log.Fatalf("Не удалось получить список организаций: %v", err)
    }

    fmt.Printf("Найденные организации: %+v\n", organizations)

    // Теперь давайте получим список облаков в определенной организации
    // Замените на ваш реальный ID организации
    orgID := "your_organization_id"
    clouds, err := client.Clouds().List(ctx, &orgID, nil, nil)
    if err != nil {
        log.Fatalf("Не удалось получить список облаков: %v", err)
    }

    fmt.Printf("Облака в организации %s: %+v\n", orgID, clouds)
}
```

## Использование

### Организации

```go
// Список организаций с размером страницы 100
pageSize := 100
organizations, err := client.Organizations().List(ctx, &pageSize, nil)

// Получить организацию
org, err := client.Organizations().Get(ctx, "organization_id")

// Обновить организацию
updateData := map[string]interface{}{
    "name":        "Новое имя",
    "description": "Новое описание",
}
org, err = client.Organizations().Update(ctx, "organization_id", updateData)

// Добавить роль
result, err := client.Organizations().AddRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Удалить роль
result, err = client.Organizations().RemoveRole(
    ctx,
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Список привязок доступа
bindings, err := client.Organizations().ListAccessBindings(ctx, "organization_id", nil, nil)
```

### Облака

```go
// Получаем список облаков в определенной организации
orgID := "org_id"
pageSize := 100
clouds, err := client.Clouds().List(ctx, &orgID, &pageSize, nil)

// Получить детали облака
cloud, err := client.Clouds().Get(ctx, "cloud_id")

// Создать облако
description := "Production-облако"
labels := map[string]string{"env": "production"}
cloud, err = client.Clouds().Create(
    ctx,
    "org_id",
    "Мое продакшн-облако",
    &description,
    labels,
)

// Обновить облако
updateData := map[string]interface{}{
    "name":        "Обновленное имя",
    "description": "Обновленное описание",
}
cloud, err = client.Clouds().Update(ctx, "cloud_id", updateData)

// Удалить облако
result, err := client.Clouds().Delete(ctx, "cloud_id")

// Добавить роль
result, err = client.Clouds().AddRole(
    ctx,
    "cloud_id",
    "user_id",
    "editor",
    "userAccount",
)

// Получаем список всех привязок доступа для облака
bindings, err := client.Clouds().ListAccessBindings(ctx, "cloud_id", nil, nil)
```

---

## Расширенное управление доступом

### Добавление нескольких ролей

```go
// Добавляем несколько ролей в каталог одним запросом
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

### Замена всех привязок доступа

```go
// Заменить все привязки доступа
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

### Назначение ролей по логину пользователя

```go
// Получить ID пользователя по логину Яндекс.Паспорта
user, err := client.YandexPassportUserAccounts().GetByLogin(ctx, "username@yandex.ru")
if err != nil {
    log.Fatal(err)
}
userID := user["id"].(string)

// Назначить роль в каталоге
result, err := client.Folders().AddRole(
    ctx,
    "folder_id",
    userID,
    "ai.languageModels.user",
    "userAccount",
)

// Или назначить роль в облаке
result, err = client.Clouds().AddRole(
    ctx,
    "cloud_id",
    userID,
    "editor",
    "userAccount",
)
```

---

## Обработка ошибок

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
        log.Printf("Ошибка аутентификации: %v", e)
    case *errors.ValidationError:
        log.Printf("Ошибка валидации: %v", e)
    case *errors.APIError:
        log.Printf("Ошибка API (статус %d): %v", e.StatusCode, e)
    default:
        log.Printf("Непредвиденная ошибка: %v", err)
    }
}
```

---

## Тестирование

```bash
go test ./...
go test -cover ./...
go test -v ./...
```

---

## Участие в разработке

```bash
git clone https://github.com/tigusigalpa/yandex-cloud-client-go.git
cd yandex-cloud-client-go
go mod download
cp .env.example .env
```

1. Сделайте форк репозитория
2. Создайте ветку (`git checkout -b feature/my-feature`)
3. Закоммитьте изменения (`git commit -m 'Add feature'`)
4. Отправьте ветку (`git push origin feature/my-feature`)
5. Откройте Pull Request

Пожалуйста, следуйте соглашениям Go, добавляйте тесты для новых функций и обновляйте документацию.

---

## Безопасность

Если вы обнаружите уязвимость, напишите на **sovletig@gmail.com** вместо создания публичного issue.

---

## Автор

**Игорь Сазонов** ([@tigusigalpa](https://github.com/tigusigalpa))

- **Email**: sovletig@gmail.com

---

## Лицензия

MIT License — см. [LICENSE](LICENSE).

---

## Связанные пакеты

| Пакет                    | Описание                    | Ссылки                                                                                                                                             |
|--------------------------|-------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| **YandexGPT Go**         | Go SDK для API YandexGPT    | [GitHub](https://github.com/tigusigalpa/yandexgpt-go)                                                                                             |
| **YandexGPT PHP**        | PHP SDK для API YandexGPT   | [GitHub](https://github.com/tigusigalpa/yandexgpt-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandexgpt-php)                     |
| **Yandex Cloud Client PHP** | PHP SDK для Яндекс.Облака | [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandex-cloud-client-php) |

---

## Ссылки

- [Документация Яндекс.Облака](https://yandex.cloud/docs)
- [Справочник API организаций](https://yandex.cloud/ru/docs/organization/api-ref/)
- [Справочник API Resource Manager](https://yandex.cloud/ru/docs/resource-manager/api-ref/)
- [Справочник API IAM](https://yandex.cloud/ru/docs/iam/api-ref/)
- [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-go)
- [Issues](https://github.com/tigusigalpa/yandex-cloud-client-go/issues)
- [Обсуждения](https://github.com/tigusigalpa/yandex-cloud-client-go/discussions)