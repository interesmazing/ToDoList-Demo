# ToDoList-Demo

这是一个由 AI Agent (GitHub Copilot) 创建的全栈待办事项列表（ToDo List）演示项目。

该项目旨在展示如何使用现代技术栈快速构建一个功能完整的、容器化的 Web 应用程序。

## ✨ 技术栈

*   **后端**: Go (使用 [Gin](https://gin-gonic.com/) 框架)
*   **前端**: Vue.js (使用 [Vite](https://vitejs.dev/) 作为构建工具)
*   **数据库**: PostgreSQL
*   **容器化**: Docker & Docker Compose

## 🚀 如何开始

请确保你的电脑上已经安装了 [Docker](https://www.docker.com/get-started) 和 [Docker Compose](https://docs.docker.com/compose/install/)。

### 1. 克隆仓库

```bash
git clone https://github.com/interesmazing/ToDoList-Demo.git
cd ToDoList-Demo
```

### 2. 启动应用

在项目根目录下，运行以下命令：

```bash
docker-compose up --build
```

`--build` 标志会强制 Docker Compose 在启动前重新构建镜像，确保应用使用的是最新的代码。

应用启动后，你可以通过以下地址访问：

*   **前端应用**: [http://localhost:8080](http://localhost:8080)
*   **后端 API**: [http://localhost:8081/api/todos](http://localhost:8081/api/todos)

### 3. 停止应用

要停止并移除所有相关的容器和网络，请在项目根目录下运行：

```bash
docker-compose down
```

## 📁 项目结构

```
.
├── backend/            # Go 后端服务
│   ├── Dockerfile
│   ├── go.mod
│   └── main.go
├── db/                 # 数据库初始化脚本
│   └── init.sql
├── frontend/           # Vue 前端应用
│   ├── Dockerfile
│   ├── index.html
│   ├── package.json
│   ├── vite.config.js
│   └── src/
├── docker-compose.yml  # Docker Compose 编排文件
├── LICENSE             # MIT 许可证
└── README.md           # 项目说明
```

## ⚙️ API 端点

所有 API 端点都以 `/api` 为前缀。

| 方法   | 路径               | 描述                 |
|--------|--------------------|----------------------|
| `GET`  | `/todos`           | 获取所有待办事项     |
| `POST` | `/todos`           | 创建一个新的待办事项 |
| `PUT`  | `/todos/:id`       | 更新一个待办事项     |
| `DELETE`| `/todos/:id`      | 删除一个待办事项     |

---

该项目由 GitHub Copilot 自动生成。
