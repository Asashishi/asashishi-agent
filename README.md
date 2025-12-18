# Asashishi Agent - 智能编程助手

[![Go Version](https://img.shields.io/badge/Go-1.25.5+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![OpenAI API](https://img.shields.io/badge/OpenAI-API-412991?style=flat-square&logo=openai)](https://platform.openai.com/)

<p align="center">
  <img src="https://raw.githubusercontent.com/Asashishi/asashishi-agent/refs/heads/main/resources/app.ico" alt="Logo" />
</p>

**Asashishi Agent** 是一个基于 Go 语言开发的智能编程助手工具，通过自然语言交互帮助开发者进行项目开发、代码生成和文件管理。支持 OpenAI 兼容 API（如 DeepSeek），提供完整的文件操作、Shell 命令执行和网络搜索能力。

## ✨ 核心特性

### 🤖 智能交互

- **自然语言编程** - 使用自然语言描述需求，自动生成代码和项目结构
- **多语言支持** - 专精于 JavaScript/TypeScript、Python、HTML/CSS 项目开发
- **智能规划** - 自动分析需求并制定详细的操作计划

### 🛠️ 强大工具集

- **文件操作** - 完整的文件/目录创建、读取、修改、删除、重命名功能
- **Shell 集成** - 安全的命令组管理和执行，支持 PowerShell
- **网络搜索** - 联网获取最新库信息和 API 文档（可选）
- **自动备份** - 支持启动时备份，支持文件内容操作自动回滚

### ⚡ 高效开发

- **快速项目初始化** - 一键创建标准化项目结构
- **智能依赖管理** - 自动处理 npm/pip 依赖安装
- **代码规范检查** - 遵循最新的编码标准和最佳实践
- **完整日志记录** - 所有操作详细记录，便于追溯和调试

## 🚀 快速开始

### 系统要求

- **Go 1.25.5+** - [下载地址](https://golang.org/dl/)
- **Windows 10/11**（支持其他平台，但需要调整构建脚本）
- **OpenAI 兼容 API 密钥**（如 DeepSeek、OpenAI 等）

### 安装方式

#### 方式一：从源码构建（推荐）

##### 🖥️ Windows 系统构建

```bash
# 1. 克隆仓库
git clone https://github.com/asashishi/asashishi-agent.git
cd asashishi-agent

# 2. 安装依赖工具（Windows 图标资源工具）
go install github.com/akavel/rsrc@latest

# 3. 下载 UPX（可选，用于压缩可执行文件）
# 前往 https://upx.github.io/ 下载并添加到 PATH

# 4. 使用构建脚本（推荐）
./build.bat

# 或手动执行构建命令
go build -ldflags="-s -w -H=windowsgui" -trimpath -o asashishi-agent.exe

# 5. 使用 UPX 压缩（可选）
upx --best --lzma asashishi-agent.exe
```

##### 🐧 Linux/macOS 系统构建

```bash
# 1. 克隆仓库
git clone https://github.com/asashishi/asashishi-agent.git
cd asashishi-agent

# 2. 安装依赖工具
go mod tidy

# 3. 使用构建脚本
chmod +x build.sh
./build.sh

# 或手动执行构建命令
go build -ldflags="-s -w" -trimpath -o asashishi-agent

# 4. 使用 UPX 压缩（可选）
upx --best --lzma asashishi-agent
```

##### 📦 构建选项说明

| 构建选项                     | 说明                           | 推荐值                                                          |
| ---------------------------- | ------------------------------ | --------------------------------------------------------------- |
| `-ldflags="-s -w"`         | 移除调试信息，减小文件大小     | 推荐使用                                                        |
| `-ldflags="-H=windowsgui"` | Windows 隐藏控制台窗口         | Windows 专用                                                    |
| `-trimpath`                | 移除构建路径信息，提高可移植性 | 推荐使用                                                        |
| `-o`                       | 指定输出文件名                 | `asashishi-agent.exe` (Windows)`<br>asashishi-agent` (Unix) |

##### 🔧 构建脚本功能

**build.bat (Windows) / build.sh (Unix)** 提供以下功能：

1. **自动依赖检查** - 检查 Go 版本和必要工具
2. **图标资源嵌入** - 自动嵌入程序图标（Windows）
3. **优化构建** - 使用推荐的构建参数
4. **版本信息** - 嵌入版本和构建时间信息
5. **清理功能** - 可选的清理中间文件

##### 🚀 快速构建命令

```bash
# Windows 快速构建（使用默认配置）
go build -o asashishi-agent.exe

# Linux/macOS 快速构建
go build -o asashishi-agent

# 开发模式构建（保留调试信息）
go build -gcflags="all=-N -l" -o asashishi-agent-dev
```

##### 🛠️ 开发环境设置

```bash
# 1. 验证 Go 环境
go version  # 确保 >= 1.25.5

# 2. 获取项目依赖
go mod download
go mod verify

# 3. 运行测试（如有）
go test ./...

# 4. 格式化代码
gofmt -w .

# 5. 静态分析
go vet ./...
```

##### 📝 构建注意事项

1. **Go 版本要求**：必须使用 Go 1.25.5 或更高版本
2. **网络连接**：首次构建需要下载依赖，确保网络畅通
3. **磁盘空间**：构建过程需要约 100MB 临时空间
4. **权限要求**：需要写入当前目录的权限
5. **防病毒软件**：某些防病毒软件可能误报，请添加例外

##### 🔍 构建问题排查

如果构建失败，请检查：

1. **Go 版本**：`go version` 确认版本符合要求
2. **依赖工具**：确保 `rsrc` 工具已正确安装
3. **网络代理**：如有需要，设置 Go 代理 `go env -w GOPROXY=...`
4. **环境变量**：检查 `GOPATH` 和 `GOROOT` 设置
5. **查看错误**：仔细阅读构建错误信息，通常包含具体原因

#### 方式二：下载预编译版本

前往 [Releases](https://github.com/asashishi/asashishi-agent/releases) 页面下载最新版本的可执行文件。

### 基本配置

1. **编辑配置文件** - 项目已包含 `config.json` 文件，直接编辑即可：

```json
{
    "llm": {
        "api_key": "sk-your-api-key-here",
        "base_url": "https://api.deepseek.com/v1",
        "model_name": "deepseek-chat"
    }
}
```

2. **配置说明**：

   - `api_key`: **必填** - 您的 OpenAI 兼容 API 密钥
   - `base_url`: API 服务地址，默认为 DeepSeek
   - `model_name`: 使用的 AI 模型名称
3. **启动程序**

```bash
./asashishi-agent.exe
# 或
go run main.go
```

## ⚙️ 详细配置

### 配置文件结构

```json
{
    "info": {
        "version": "3.4.1",
        "name": "asashishi-agent.exe"
    },
    "proc": {
        "backup": false,
        "backup_excepts": [
            "log",
            ".git",
            "build",
            "backup\\files",
            "node_modules"
        ],
        "tick_per_sec": 90,
        "web": {
            "web_mode": false,
            "http_port": 3000,
            "websocket_route": "/ws",
            "server_root_path": "web"
        },
        "terminal_code_style": "monokai"
    },
    "llm": {
        "temperature": 0.5,
        "dir_excepts": [
            "build",
            ".git",
            "node_modules",
            "backup\\files"
        ],
        "context_length": 128,
        "use_web_search": true,
        "show_toolcall_args": false,
        "model_name": "deepseek-chat",
        "max_response_token_length": 8192,
        "base_url": "https://api.deepseek.com/v1",
        "api_key": "sk-your-api-key-here"
    }
}
```

### 关键配置说明

| 配置项 | 说明 | 推荐值 |
|--------|------|--------|
| `llm.api_key` | **必填** - OpenAI 兼容 API 密钥 | 从服务商获取 |
| `llm.base_url` | API 服务地址 | `https://api.deepseek.com/v1` |
| `llm.model_name` | 使用的 AI 模型 | `deepseek-chat` |
| `llm.temperature` | 创造力参数 (0.0-2.0) | `0.5`（平衡） |
| `proc.backup` | 是否启用启动时备份 | `true`（生产环境） |
| `llm.use_web_search` | 是否启用联网搜索 | `true`（需要最新信息时） |
| `proc.web.web_mode` | 是否启用 Web 模式 | `false`（默认 CLI 模式） |
| `proc.web.http_port` | Web 服务器端口 | `3000` |
| `proc.web.websocket_route` | WebSocket 路由 | `/ws` |
| `proc.web.server_root_path` | Web 静态文件目录 | `web` |
| `llm.dir_excepts` | 文件操作排除目录 | `["build", ".git", "node_modules", "backup\\files"]` |

## 🛠️ 工具功能详解

### 📁 文件操作 (FileOps)
- `GetFileList` - 获取目录文件列表
- `CreateFile` / `RemoveFile` - 创建/删除文件
- `CreateDir` / `RemoveDir` - 创建/删除目录
- `MoveContent` - 重命名文件/目录
- `ReadFileContent` - 读取文件内容
- `AppendContentAtTail` - 在文件末尾追加内容
- `AppendContentAtMiddle` - 在文件中间插入内容
- `DeleteFileContent` - 清空文件内容
- `FileContentRollBack` - 回滚到上次缓存状态
- `RenewFileCache` - 更新文件缓存

### 💻 Shell 操作 (ShellOps)
- `GetCommands` - 获取当前命令组
- `AddCommands` - 添加命令到命令组
- `PopCommands` - 移除末尾命令
- `ClearCommands` - 清空命令组
- `Excute` - 执行命令组并返回结果

### 🌐 网络操作 (NetOps)
- `HttpSearch` - 获取网页文本内容（需启用 `use_web_search`）

### ⏰ 时间工具
- `GetFormatedTime()` - 获取格式化时间戳 (YYYYMMDDhhmmss)

## 📖 使用示例

### 示例 1：创建 TypeScript 项目

```
用户：创建一个 TypeScript 项目，包含 Express 服务器和基础路由

助手：正在规划项目结构...
1. 创建项目目录和 package.json
2. 安装 TypeScript 和 Express 依赖
3. 配置 tsconfig.json
4. 创建服务器入口文件
5. 添加基础路由和中间件
...
```

### 示例 2：修改现有代码

```
用户：在现有的 user.service.ts 中添加用户验证功能

助手：正在分析现有代码结构...
1. 读取当前文件内容
2. 分析现有函数和接口
3. 添加验证函数和类型定义
4. 更新相关导入和导出
...
```

### 示例 3：执行系统命令

```
用户：安装项目依赖并启动开发服务器

助手：正在执行命令...
1. npm install
2. 配置环境变量
3. 启动开发服务器 (npm run dev)
...
```

## 🏗️ 项目结构

```
asashishi-agent/
├── main.go                   # 程序主入口文件
├── config.json               # 用户配置文件
├── build.bat                 # Windows 构建脚本
├── build.sh                  # Linux/macOS 构建脚本
├── run-test.bat              # Windows 测试运行脚本
├── run-test.sh               # Linux/macOS 测试运行脚本
├── go.mod                    # Go 模块定义文件
├── go.sum                    # Go 依赖校验文件
├── .gitignore                # Git 忽略配置文件
├── LICENSE                   # MIT 许可证文件
│
├── agent/                    # AI 代理核心模块
├── tools/                    # 工具实现模块
├── conf/                     # 配置管理模块
├── backup/                   # 备份系统模块
├── global/                   # 全局功能模块
├── test/                     # 测试模块
├── cmd/                      # 快捷命令模块
├── ui/                       # 用户界面模块
├── entry/                    # 程序入口模块（CLI/Web 模式）
├── websocket/                # WebSocket 通信模块
├── web/                      # Web 界面文件目录
│   ├── index.html            # Web 主页面
│   ├── example.html          # WebSocket 示例页面
│   └── public/               # 静态资源目录
├── log/                      # 操作日志目录
└── resources/                # 资源文件目录
```

### 📁 模块功能概览
| 模块                 | 功能描述     | 关键职责                               |
| -------------------- | ------------ | -------------------------------------- |
| **agent/**     | AI 代理核心  | 处理与 LLM 的交互、工具调用和智能规划  |
| **tools/**     | 工具实现     | 提供文件、Shell、网络、时间等操作能力  |
| **conf/**      | 配置管理     | 读取、验证和提供用户配置               |
| **backup/**    | 备份系统     | 在重要操作前自动备份文件，支持回滚     |
| **global/**    | 全局功能     | 提供跨模块使用的常量、工具和样式       |
| **test/**      | 测试模块     | 提供测试初始化和验证功能               |
| **cmd/**       | 命令行模块   | 处理命令行参数和用户交互               |
| **ui/**        | 用户界面模块 | 处理终端颜色、样式和输出格式化         |
| **entry/**     | 程序入口模块 | 提供 CLI 和 Web 两种启动模式           |
| **websocket/** | WebSocket 模块 | WebSocket 通信和消息传递的数据格式定义         |
| **web/**       | Web 界面模块 | 提供 Web 界面和静态文件服务           |
| **log/**       | 日志记录     | 保存所有操作的详细日志，便于调试和审计 |
| **resources/** | 资源文件     | 存储程序图标等静态资源                 |

### 🔄 系统架构与数据流
```
CLI 模式数据流：
用户输入 (命令行)
    ↓
cmd/ (命令行解析)
    ↓
agent/ (AI 代理核心)
    ├── 配置验证 (conf/)
    ├── 智能规划 (agent/)
    ├── 工具调用 (tools/)
    └── 安全检查 (global/)
    ↓
工具执行
    ├── 文件操作 (tools/)
    ├── Shell 命令 (tools/)
    ├── 网络搜索 (tools/)
    └── 时间工具 (tools/)
    ↓
结果处理
    ├── 备份保护 (backup/)
    ├── 状态更新 (agent/)
    ├── UI 渲染 (ui/)
    └── 日志记录 (log/)
    ↓
终端输出

Web 模式数据流：
用户输入 (Web 界面)
    ↓
websocket/ (WebSocket 通信)
    ↓
entry/ (Web 模式入口)
    ↓
agent/ (AI 代理核心)
    ├── 配置验证 (conf/)
    ├── 智能规划 (agent/)
    ├── 工具调用 (tools/)
    └── 安全检查 (global/)
    ↓
工具执行
    ├── 文件操作 (tools/)
    ├── Shell 命令 (tools/)
    ├── 网络搜索 (tools/)
    └── 时间工具 (tools/)
    ↓
结果处理
    ├── 备份保护 (backup/)
    ├── 状态更新 (agent/)
    ├── WebSocket 响应 (websocket/)
    └── 日志记录 (log/)
    ↓
Web 界面输出
```

### 🎯 核心设计原则

1. **模块化设计** 📦

   - 每个模块职责单一，高内聚低耦合
   - 清晰的接口定义，便于维护和扩展
   - 模块间通过明确定义的接口通信
2. **单向数据流** 🔄

   - 严格遵循领域驱动设计(DDD)原则
   - 避免循环依赖，确保数据流向清晰
   - 输入 → 处理 → 输出 的线性流程
3. **错误隔离与恢复** 🛡️

   - 模块间错误不传播，确保系统稳定性
   - 自动备份和回滚机制
   - 详细的错误日志记录
4. **日志驱动开发** 📝

   - 所有操作都有详细的时间戳日志
   - 日志格式统一，便于分析和审计
   - 支持操作追溯和问题排查
5. **UI 与业务逻辑分离** 🎨

   - 用户界面逻辑独立于业务逻辑
   - 支持不同的输出格式和样式
   - 便于定制和扩展用户界面
6. **安全第一** 🔒

   - 所有操作前进行安全检查
   - 文件操作限制在当前工作目录
   - 网络操作验证来源和完整性

### 🏆 技术亮点- **智能工具路由**：根据用户需求自动选择合适的工具
- **自动备份系统**：重要操作前自动备份，支持一键回滚
- **跨平台支持**：提供 Windows 和 Unix-like 系统的构建脚本
- **可扩展架构**：易于添加新的工具和功能模块
- **详细日志系统**：所有操作都有完整记录，便于调试
- **双模式运行**：支持 CLI 命令行和 Web 界面两种交互方式
- **实时 WebSocket 通信**：Web 模式下提供实时双向通信
- **完整测试套件**：内置全面的功能测试，确保系统稳定性
- **自动打包发布**：构建脚本支持自动打包为 zip/tar.gz 格式

### 环境设置

```bash
# 1. 安装 Go 1.25.5+
go version

# 2. 获取项目依赖
go mod tidy

# 3. 安装开发工具
go install github.com/akavel/rsrc@latest
```

### 构建说明

```bash
# Windows 构建（使用 build.bat）
./build.bat

# 手动构建
go build -ldflags="-s -w" -trimpath -o asashishi-agent.exe

# 使用 UPX 压缩（可选）
upx --best --lzma asashishi-agent.exe
```

### 代码规范

- 遵循 Go 官方代码规范
- 使用 `gofmt` 格式化代码
- 添加必要的注释和文档
- 编写单元测试

## 🤝 贡献指南

我们欢迎各种形式的贡献！请参考以下步骤：

### 报告问题

1. 在 [Issues](https://github.com/asashishi/asashishi-agent/issues) 页面搜索是否已有类似问题
2. 创建新 issue，详细描述问题、复现步骤和期望行为
3. 提供相关日志和系统信息

### 提交代码

1. Fork 本仓库
2. 创建功能分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'Add some amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

### 开发规范

- 保持代码简洁和可读性
- 添加适当的测试用例
- 更新相关文档
- 遵循现有的代码风格

## ❓ 常见问题

### Q: 程序无法启动或报错？

**A:** 检查以下事项：

1. `config.json` 中的 API 密钥是否正确
2. 网络连接是否正常（如需访问 API）
3. 是否有文件读写权限
4. 查看 `log/` 目录下的详细错误日志

### Q: 如何备份我的项目？

**A:** 在 `config.json` 中设置 `"backup": true`，程序会在启动时自动备份文件到 `backup/files/` 目录。

### Q: 联网搜索功能消耗多少 token？

**A:** 网页内容会经过摘要处理，会消耗一定 token (根据网页内容大小, 通常在 110k 左右/次, 摘要由独立上下文处理，主上下文消耗在 5-10k 左右/次)。建议仅在需要最新信息时启用，或使用 `"use_web_search": false` 关闭。

### Q: Linux 可以运行么？
**A:** 可以，项目已经完全支持 Linux 所有功能已经通过测试, 但需要编译安装, 暂不提供预编译版本

### Q: 支持哪些 AI 模型？

**A:** 支持所有 OpenAI 兼容 API，包括：

- DeepSeek Chat
- OpenAI GPT 系列
- 其他兼容 OpenAI API 的服务

### Q: 如何启用 Web 模式？
**A:** 
1. 在 `config.json` 中设置 `"proc.web.web_mode": true`，然后重启程序。程序将在指定端口（默认 3000）启动 Web 服务器，您可以通过浏览器访问 `http://localhost:3000` 使用 Web 界面
2. 注意! Web 模式尚未正式发布，如需提前使用，请自行按照 ./web/example.html 下的示例实现 socket 回调和页面样式

### Q: Web 模式支持哪些功能？
**A:** Web 将模式支持所有 CLI 模式的功能

### Q: 如何运行测试？

**A:** 有两种方式运行测试：
1. **使用测试脚本**：执行 `./run-test.bat` (Windows) 或 `./run-test.sh` (Linux/macOS)
2. **命令行参数**：运行 `./asashishi-agent.exe rt` 或 `./asashishi-agent rt`

### Q: 测试包含哪些内容？

**A:** 测试套件包含：
- 文件操作测试（创建、读取、修改、删除文件/目录）
- Shell 命令组管理测试
- 网络搜索测试
- 时间工具测试
- 所有工具功能的完整性验证

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🧩 三方库集成

[OpenAI](https://openai.com/)<br/>
[Chroma-V2](https://github.com/alecthomas/chroma)<br/>
[Coder-Websocket](https://github.com/coder/websocket)<br/>

## 🌟 支持与反馈

- **GitHub Issues**: [报告问题或请求功能](https://github.com/asashishi/asashishi-agent/issues)
- **文档**: 查看本 README 和代码注释
- **日志**: 程序运行日志保存在 `log/` 目录

---

**✨ 开始您的智能编程之旅！**

- 如有任何问题或建议，欢迎参与项目贡献或提交反馈。
