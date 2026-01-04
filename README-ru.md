<div align="center">

# ☁️ Yandex Cloud Client Go

![Yandex Cloud Client Go](https://github.com/user-attachments/assets/f75920a6-c0cd-4da3-9223-5d7661ad3e47)

### 🚀 Современный Go SDK для Yandex Cloud API

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat)](LICENSE)
[![GitHub](https://img.shields.io/badge/GitHub-tigusigalpa%2Fyandex--cloud--client--go-181717?style=flat&logo=github)](https://github.com/tigusigalpa/yandex-cloud-client-go)

**🌐 Язык:** Русский | [English](README.md)

**Мощный, элегантный и удобный для разработчиков Go SDK для Yandex Cloud API.**

Управляйте организациями, облаками, каталогами и IAM-аутентификацией с помощью чистого идиоматичного Go кода.

</div>

---

## ✨ Возможности

<table>
<tr>
<td width="50%">

### 🔐 Аутентификация и безопасность

- Поддержка **OAuth 2.0** токенов
- **Автоматическая генерация** IAM токенов
- **Умное кеширование** с автообновлением
- Управление **сроком действия** токенов (12ч)
- **Потокобезопасные** операции

### 🏢 Управление ресурсами

- **Организации** - полный CRUD и контроль доступа
- **Облака** - полное управление жизненным циклом
- **Каталоги** - операции и права доступа
- **Сервисные аккаунты** - полный жизненный цикл
- **Пользовательские аккаунты** - получение информации
- **API ключи** - создание и управление
- **Refresh токены** - управление жизненным циклом

</td>
<td width="50%">

### 🎯 Особенности Go

- **Идиоматичный Go** код
- **Типобезопасный** API
- Готовность к **Context**
- **Минимум зависимостей**
- **Goroutine-safe**

### 💎 Качество кода

- **Go 1.21+** с поддержкой дженериков
- **Полная обработка ошибок**
- **Чистая архитектура**
- **Хорошо документирован**

</td>
</tr>
</table>

## 📋 Требования

| Требование | Версия |
|------------|--------|
| Go         | 1.21+  |

## 🚀 Быстрый старт

### Установка

```bash
go get github.com/tigusigalpa/yandex-cloud-client-go
```

### Получение OAuth токена

<details>
<summary>📝 Нажмите, чтобы узнать, как получить OAuth токен</summary>

1. Перейдите на [Yandex OAuth](https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990aa636648e9b2ef855fa7bec2fb)
2. Авторизуйте приложение
3. Скопируйте токен
4. Используйте его в коде

💡 **Совет**: Храните токены безопасно в переменных окружения!

Подробнее см. в [документации Yandex Cloud](https://yandex.cloud/ru/docs/iam/concepts/authorization/oauth-token).

</details>

### Базовое использование

```go
package main

import (
    "fmt"
    "log"
    
    yandexcloud "github.com/tigusigalpa/yandex-cloud-client-go"
)

func main() {
    // Инициализация клиента
    client, err := yandexcloud.NewClient("your_oauth_token", nil)
    if err != nil {
        log.Fatal(err)
    }
    
    // Список всех организаций
    organizations, err := client.Organizations().List(nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Организации: %+v\n", organizations)
    
    // Список облаков в организации
    orgID := "your_organization_id"
    clouds, err := client.Clouds().List(&orgID, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Облака: %+v\n", clouds)
}
```

## 💻 Примеры использования

### API организаций

```go
// Список организаций
pageSize := 100
organizations, err := client.Organizations().List(&pageSize, nil)

// Получить организацию
org, err := client.Organizations().Get("organization_id")

// Обновить организацию
updateData := map[string]interface{}{
    "name": "Новое имя",
    "description": "Новое описание",
}
org, err = client.Organizations().Update("organization_id", updateData)

// Добавить роль к организации
result, err := client.Organizations().AddRole(
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Удалить роль из организации
result, err = client.Organizations().RemoveRole(
    "organization_id",
    "user_id",
    "editor",
    "userAccount",
)

// Список привязок доступа
bindings, err := client.Organizations().ListAccessBindings("organization_id", nil, nil)
```

### API облаков

```go
// Список облаков
orgID := "org_id"
pageSize := 100
clouds, err := client.Clouds().List(&orgID, &pageSize, nil)

// Получить облако
cloud, err := client.Clouds().Get("cloud_id")

// Создать облако
description := "Продакшн облако"
labels := map[string]string{"env": "production"}
cloud, err = client.Clouds().Create(
    "org_id",
    "Мое облако",
    &description,
    labels,
)

// Обновить облако
updateData := map[string]interface{}{
    "name": "Обновленное имя",
    "description": "Обновленное описание",
}
cloud, err = client.Clouds().Update("cloud_id", updateData)

// Удалить облако
result, err := client.Clouds().Delete("cloud_id")

// Добавить роль к облаку
result, err = client.Clouds().AddRole(
    "cloud_id",
    "user_id",
    "editor",
    "userAccount",
)

// Список привязок доступа
bindings, err := client.Clouds().ListAccessBindings("cloud_id", nil, nil)
```

### API каталогов

```go
// Список каталогов
pageSize := 100
folders, err := client.Folders().List("cloud_id", &pageSize, nil)

// Получить каталог
folder, err := client.Folders().Get("folder_id")

// Создать каталог
description := "Каталог разработки"
labels := map[string]string{"team": "backend"}
folder, err = client.Folders().Create(
    "cloud_id",
    "Мой каталог",
    &description,
    labels,
)

// Обновить каталог
updateData := map[string]interface{}{
    "name": "Обновленное имя",
}
folder, err = client.Folders().Update("folder_id", updateData)

// Удалить каталог
result, err := client.Folders().Delete("folder_id")

// Список операций
operations, err := client.Folders().ListOperations("folder_id", nil, nil)

// Добавить роль к каталогу
result, err = client.Folders().AddRole(
    "folder_id",
    "user_id",
    "ai.languageModels.user",
    "userAccount",
)

// Список привязок доступа
bindings, err := client.Folders().ListAccessBindings("folder_id", nil, nil)
```

### API сервисных аккаунтов

```go
// Список сервисных аккаунтов в каталоге
pageSize := 100
serviceAccounts, err := client.ServiceAccounts().List("folder_id", &pageSize, nil)

// Получить сервисный аккаунт
sa, err := client.ServiceAccounts().Get("service_account_id")

// Создать сервисный аккаунт
description := "Сервисный аккаунт для API"
sa, err = client.ServiceAccounts().Create(
    "folder_id",
    "my-service-account",
    &description,
)

// Обновить сервисный аккаунт
updateData := map[string]interface{}{
    "name": "Обновленное имя",
    "description": "Обновленное описание",
}
sa, err = client.ServiceAccounts().Update("service_account_id", updateData)

// Удалить сервисный аккаунт
result, err := client.ServiceAccounts().Delete("service_account_id")

// Добавить роль к сервисному аккаунту
result, err = client.ServiceAccounts().AddRole(
    "service_account_id",
    "user_id",
    "editor",
    "userAccount",
)

// Список привязок доступа
bindings, err := client.ServiceAccounts().ListAccessBindings("service_account_id", nil, nil)
```

### API пользовательских аккаунтов

```go
// Получить пользовательский аккаунт по ID
user, err := client.UserAccounts().Get("user_account_id")

// Получить пользователя по логину Яндекс.Паспорта
user, err = client.YandexPassportUserAccounts().GetByLogin("username")
// Возвращает: map с 'id', 'login' и т.д.

// Использовать ID для назначения ролей
userID := user["id"].(string)
result, err := client.Folders().AddRole(
    "folder_id",
    userID,
    "editor",
    "userAccount",
)
```

### API ключи

```go
// Список API ключей для сервисного аккаунта
pageSize := 100
keys, err := client.APIKeys().List("service_account_id", &pageSize, nil)

// Получить API ключ
key, err := client.APIKeys().Get("api_key_id")

// Создать API ключ (секрет показывается только один раз!)
description := "API ключ для продакшна"
key, err = client.APIKeys().Create("service_account_id", &description)
// Сохраните key["secret"] немедленно - он больше не будет показан!

// Обновить API ключ
updateData := map[string]interface{}{
    "description": "Обновленное описание",
}
key, err = client.APIKeys().Update("api_key_id", updateData)

// Удалить API ключ
result, err := client.APIKeys().Delete("api_key_id")
```

### API Refresh токенов

```go
// Список refresh токенов
pageSize := 100
tokens, err := client.RefreshTokens().List(&pageSize, nil)

// Отозвать refresh токен
result, err := client.RefreshTokens().Revoke("token_id")
```

---

## 🔐 Расширенное управление доступом

### Добавление нескольких ролей одновременно

```go
// Добавить несколько ролей к каталогу
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

result, err := client.Clouds().SetAccessBindings("cloud_id", bindings)
```

### Назначение ролей по логину пользователя

```go
// Получить ID пользователя по логину Яндекс.Паспорта
user, err := client.YandexPassportUserAccounts().GetByLogin("username@yandex.ru")
if err != nil {
    log.Fatal(err)
}
userID := user["id"].(string)

// Назначить роль к каталогу используя ID пользователя
result, err := client.Folders().AddRole(
    "folder_id",
    userID,
    "ai.languageModels.user",
    "userAccount",
)

// Или назначить к облаку
result, err = client.Clouds().AddRole(
    "cloud_id",
    userID,
    "editor",
    "userAccount",
)
```

---

## ⚠️ Обработка ошибок

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
        // Обработка ошибок аутентификации
        log.Printf("Ошибка аутентификации: %v", e)
    case *errors.ValidationError:
        // Обработка ошибок валидации
        log.Printf("Ошибка валидации: %v", e)
    case *errors.APIError:
        // Обработка ошибок API
        log.Printf("Ошибка API (статус %d): %v", e.StatusCode, e)
    default:
        // Обработка других ошибок
        log.Printf("Ошибка: %v", err)
    }
}
```

---

## 🧪 Тестирование

```bash
# Запуск тестов
go test ./...

# Запуск тестов с покрытием
go test -cover ./...

# Запуск тестов с подробным выводом
go test -v ./...
```

---

## 🤝 Участие в разработке

Мы приветствуем вклад в проект! Вот как вы можете помочь:

### Настройка окружения разработки

```bash
# Клонировать репозиторий
git clone https://github.com/tigusigalpa/yandex-cloud-client-go.git
cd yandex-cloud-client-go

# Установить зависимости
go mod download

# Скопировать файл окружения
cp .env.example .env
```

### Рекомендации по участию

- ✅ **Следуйте соглашениям Go** и лучшим практикам
- ✅ **Пишите идиоматичный Go** код
- ✅ **Пишите тесты** для новых функций
- ✅ **Обновляйте документацию** при необходимости
- ✅ **Одна функция на PR** - держите фокус

### Процесс Pull Request

1. Форкните репозиторий
2. Создайте ветку функции (`git checkout -b feature/amazing-feature`)
3. Внесите изменения
4. Запустите тесты и убедитесь, что они проходят
5. Закоммитьте изменения (`git commit -m 'Add amazing feature'`)
6. Запушьте в ветку (`git push origin feature/amazing-feature`)
7. Откройте Pull Request

---

## 🔒 Безопасность

Если вы обнаружили уязвимости безопасности, пожалуйста, напишите на **sovletig@gmail.com** вместо использования issue tracker.

Мы серьезно относимся к безопасности и оперативно реагируем.

---

## 👨‍💻 Автор и участники

**Создано с ❤️ [Igor Sazonov](https://github.com/tigusigalpa)**

- 📧 Email: sovletig@gmail.com
- 🐙 GitHub: [@tigusigalpa](https://github.com/tigusigalpa)

### Участники

Спасибо [всем участникам](../../contributors), которые помогают улучшать этот пакет!

---

## 📄 Лицензия

Этот проект лицензирован под **MIT License** - см. файл [LICENSE](LICENSE) для деталей.

Свободен для использования в личных и коммерческих проектах. ✨

---

## 🔗 Связанные пакеты

Изучите наши другие пакеты для Yandex Cloud:

| Пакет                    | Описание              | Ссылки                                                                                                                                             |
|--------------------------|-----------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| **YandexGPT Go**         | SDK для YandexGPT API | [GitHub](https://github.com/tigusigalpa/yandexgpt-go)                                                                                             |
| **YandexGPT PHP**        | SDK для YandexGPT API | [GitHub](https://github.com/tigusigalpa/yandexgpt-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandexgpt-php)                     |
| **Yandex Cloud Client PHP** | PHP SDK для Yandex Cloud | [GitHub](https://github.com/tigusigalpa/yandex-cloud-client-php) • [Packagist](https://packagist.org/packages/tigusigalpa/yandex-cloud-client-php) |

---

## 🔗 Полезные ссылки

### Официальная документация

- 📖 [Документация Yandex Cloud](https://yandex.cloud/docs)
- 🏢 [Справочник API организаций](https://yandex.cloud/ru/docs/organization/api-ref/)
- ☁️ [Справочник API Resource Manager](https://yandex.cloud/ru/docs/resource-manager/api-ref/)
- 🔐 [Справочник IAM API](https://yandex.cloud/ru/docs/iam/api-ref/)

### Ресурсы пакета

- 🐙 [GitHub репозиторий](https://github.com/tigusigalpa/yandex-cloud-client-go)
- 🐛 [Issue Tracker](https://github.com/tigusigalpa/yandex-cloud-client-go/issues)
- 💬 [Обсуждения](https://github.com/tigusigalpa/yandex-cloud-client-go/discussions)

---

<div align="center">

### ⭐ Поставьте звезду на GitHub!

Если этот пакет помог вам, пожалуйста, поставьте звезду ⭐

**Сделано с ❤️ для Go сообщества**

[Сообщить об ошибке](https://github.com/tigusigalpa/yandex-cloud-client-go/issues) • [Запросить функцию](https://github.com/tigusigalpa/yandex-cloud-client-go/issues) • [Участвовать](https://github.com/tigusigalpa/yandex-cloud-client-go/pulls)

</div>
