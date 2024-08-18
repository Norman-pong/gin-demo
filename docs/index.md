# Gin 入门

1. 最佳实践
```
app/
├── cmd/
│   └── api/
│       └── main.go
├── internal/
│   ├── handlers/
│   │   └── user.go
│   ├── models/
│   │   └── user.go
│   ├── routes/
│   │   └── routes.go
│   └── services/
│       └── user_service.go
├── pkg/
│   └── utils/
│       └── utils.go
└── go.mod
```

## 风格

RESTful 和 CRUD 是两个不同的概念，尽管它们在 Web 开发中经常一起使用。以下是它们的区别：

### CRUD

CRUD 是指四种基本的数据库操作：

1. **Create**：创建数据
2. **Read**：读取数据
3. **Update**：更新数据
4. **Delete**：删除数据

CRUD 操作是数据库管理的基本操作，几乎所有的数据库系统都支持这些操作。

### RESTful

RESTful 是一种基于 REST（Representational State Transfer）架构风格的 Web 服务设计原则。RESTful API 使用 HTTP 协议的方法来实现 CRUD 操作，并且遵循以下原则：

1. **资源（Resources）**：一切皆资源。每个资源都有一个唯一的 URI。
2. **HTTP 方法**：使用标准的 HTTP 方法来执行操作。
   - **GET**：读取资源（对应 CRUD 中的 Read）
   - **POST**：创建资源（对应 CRUD 中的 Create）
   - **PUT**：更新资源（对应 CRUD 中的 Update）
   - **DELETE**：删除资源（对应 CRUD 中的 Delete）
3. **无状态（Stateless）**：每个请求都是独立的，服务器不保存客户端的状态。
4. **统一接口（Uniform Interface）**：通过统一的接口来操作资源。
5. **表示（Representation）**：资源可以有多种表示形式（如 JSON、XML）。

### 示例

假设我们有一个用户资源，以下是 RESTful API 和 CRUD 操作的对应关系：

- **Create（创建用户）**
  - CRUD：INSERT INTO users (name, email) VALUES ('John Doe', 'john@example.com');
  - RESTful：POST /users

- **Read（读取用户信息）**
  - CRUD：SELECT * FROM users WHERE id = 1;
  - RESTful：GET /users/1

- **Update（更新用户信息）**
  - CRUD：UPDATE users SET email = 'john.doe@example.com' WHERE id = 1;
  - RESTful：PUT /users/1

- **Delete（删除用户）**
  - CRUD：DELETE FROM users WHERE id = 1;
  - RESTful：DELETE /users/1

### 总结

- **CRUD** 是数据库操作的基本概念，专注于数据的创建、读取、更新和删除。
- **RESTful** 是一种 Web 服务设计原则，使用 HTTP 方法来实现 CRUD 操作，并遵循一系列设计原则以确保 API 的一致性和可扩展性。

通过结合使用 CRUD 和 RESTful，你可以设计出符合 REST 架构风格的 Web API，使其易于理解和使用。