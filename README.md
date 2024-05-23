# Go 书店项目文档

### 项目简介

Go 书店项目是一个基于 Go 语言的 CRUD 应用，使用了 GORM 和 Gorilla Mux 库来实现数据的增删改查和路由功能。该项目适用于学习和实践如何在
Go 中使用数据库和 Web 服务。

### 开发环境

- **Go 版本:** 1.22.3
- **模块路径:** github.com/HarleyQiu/go-bookstore

### 依赖库

- **Gorilla Mux:** 用于路由的处理，版本 `1.8.1`
- **MySQL 驱动:** 用于数据库连接和操作，版本 `1.5.0`（间接依赖）
- **GORM:** ORM 库用于简化数据库操作，版本 `1.9.16`（间接依赖）
- **Inflection:** 用于字符串的复数和单数转换，版本 `1.0.0`（间接依赖）

### 配置文件的使用

为了保护敏感信息，如数据库连接信息，应将其存放在配置文件中（例如 `config.json`），并通过代码从配置文件中动态读取。
创建配置文件:

- 在项目的根目录下创建一个名为`config.json`的文件，并填入以下内容（使用实际的数据库连接信息替换示例内容）：

   ```json
   {
      "database": "用户名:密码@tcp(服务器地址:端口)/数据库名?charset=utf8&parseTime=True&loc=Local"
   }
   ```

### 安装方法

首先确保你的机器上已经安装了 Go，并设置了正确的环境变量。然后执行以下步骤：

1. 克隆项目到本地:
   ```shell
   git clone https://github.com/HarleyQiu/go-bookstore.git
   cd go-bookstore
   ```

2. 下载依赖:
   ```shell
   go mod tidy
   ```

3. 编译并运行项目:
   ```shell
    go run main.go
    ```

### 使用说明

项目运行后，将启动 HTTP 服务，默认端口为 `9010`（如果代码中设置了不同的端口，请以代码为准）。你可以通过 REST API
对书籍数据进行增删改查操作。具体的 API 路径和方法请参考项目的 `routes` 配置。