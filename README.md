# Asashishi Agent - 智能编程助手 CLI

[![Go Version](https://img.shields.io/badge/Go-1.25.5+-00ADD8?style=flat-square&logo=go)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![OpenAI API](https://img.shields.io/badge/OpenAI-API-412991?style=flat-square&logo=openai)](https://platform.openai.com/)

<p align="center">
  <img src="https://raw.githubusercontent.com/Asashishi/asashishi-agent/refs/heads/main/resources/app.ico" alt="Logo" />
</p>

**Asashishi Agent** 是一个基于 Go 语言开发的智能编程助手命令行工具，通过自然语言交互帮助开发者进行项目开发、代码生成和文件管理。支持 OpenAI 兼容 API（如 DeepSeek），提供完整的文件操作、Shell 命令执行和网络搜索能力。

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
```bash
# 1. 克隆仓库
git clone https://github.com/asashishi/asashishi-agent.git
cd asashishi-agent

# 2. 安装依赖工具
go install github.com/akavel/rsrc@latest  # Windows 图标资源工具
# 下载 UPX：https://upx.github.io/（可选，用于压缩可执行文件）

# 3. 构建项目
./build.bat  # Windows
# 或手动执行构建命令
```

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
        "version": "3.2.7",
        "owner": "YourName",
        "name": "asashishi-agent.exe"
    },
    "proc": {
        "backup": true,
        "backup_excepts": ["log", "backup\\files", "node_modules"],
        "tick_per_sec": 90
    },
    "llm": {
        "api_key": "sk-your-api-key-here",
        "base_url": "https://api.deepseek.com/v1",
        "model_name": "deepseek-chat",
        "temperature": 0.5,
        "use_web_search": true,
        "show_toolcall_args": false,
        "context_length": 128,
        "max_response_token_length": 8192,
        "files_excepts": ["node_modules", "backup\\files"]
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

## 🛠️ 工具功能详解

### 📁 文件操作 (FileOps)
- `GetFileList(path)` - 获取目录文件列表
- `CreateFile(path)` / `RemoveFile(path)` - 创建/删除文件
- `CreateDir(path)` / `RemoveDir(path)` - 创建/删除目录
- `MoveContent(opath, npath)` - 重命名文件/目录
- `ReadFileContent(path)` - 读取文件内容
- `AppendContentAtTail(path, content)` - 在文件末尾追加内容
- `AppendContentAtMiddle(path, stp, content)` - 在文件中间插入内容
- `DeleteFileContent(path)` - 清空文件内容
- `FileContentRollBack(path)` - 回滚到上次缓存状态
- `RenewFileCache(path)` - 更新文件缓存

### 💻 Shell 操作 (ShellOps)
- `GetCommands()` - 获取当前命令组
- `AddCommands(command)` - 添加命令到命令组
- `PopCommands(num)` - 移除末尾命令
- `ClearCommands()` - 清空命令组
- `Excute()` - 执行命令组并返回结果

### 🌐 网络操作 (NetOps)
- `WebContentSearch(url)` - 获取网页文本内容（需启用 `use_web_search`）

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
├── main.go                    # 程序入口
├── config.json               # 用户配置文件
├── build.bat                 # Windows 构建脚本
├── run-test.bat              # 测试运行脚本
├── go.mod                    # Go 模块定义
├── go.sum                    # 依赖校验文件
├── .gitignore               # Git 忽略配置
├── LICENSE                   # MIT 许可证
│
├── agent/                    # AI 代理核心模块
│   ├── agent-client.go      # AI 客户端实现
│   ├── consts.go            # 常量定义
│   ├── tool-switch.go       # 工具调用路由
│   ├── types.go             # 类型定义
│   ├── use-tool.go          # 工具使用逻辑
│   └── utils.go             # 工具函数
│
├── tools/                    # 工具实现模块
│   ├── init-desc.go         # 工具描述定义
│   ├── flie-ops.go          # 文件操作实现
│   ├── shell-ops.go         # Shell 操作实现
│   ├── net-ops.go           # 网络操作实现
│   ├── time-ops.go          # 时间工具实现
│   └── consts.go            # 工具常量定义
│
├── conf/                     # 配置管理模块
│   ├── init-conf.go         # 配置初始化
│   ├── types.go             # 配置类型定义
│   └── consts.go            # 配置常量定义
│
├── backup/                   # 备份系统模块
│   ├── backup.go            # 备份逻辑实现
│   └── consts.go            # 备份常量定义
│
├── global/                   # 全局功能模块
│   ├── consts.go            # 全局常量定义
│   └── utils.go             # 全局工具函数
│
├── test/                     # 测试模块
│   ├── init-test.go         # 测试初始化
│   └── consts.go            # 测试常量定义
│
├── log/                      # 操作日志目录
│   └── *.md                 # 时间戳命名的日志文件
│
└── resources/                # 资源文件目录
    └── app.ico              # 程序图标
```

### 📁 模块说明

| 模块 | 功能描述 | 关键文件 |
|------|----------|----------|
| **agent/** | AI 代理核心，处理与 LLM 的交互和工具调用 | `agent-client.go`, `tool-switch.go` |
| **tools/** | 工具实现，提供文件、Shell、网络等操作能力 | `flie-ops.go`, `shell-ops.go`, `net-ops.go` |
| **conf/** | 配置管理，读取和验证用户配置 | `init-conf.go`, `types.go` |
| **backup/** | 备份系统，在重要操作前自动备份文件 | `backup.go` |
| **global/** | 全局功能，提供跨模块使用的常量和工具 | `consts.go`, `utils.go` |
| **test/** | 测试模块，提供测试初始化功能 | `init-test.go` |
| **log/** | 日志记录，保存所有操作的详细日志 | `YYYYMMDDhhmmss.md` 格式文件 |

### 🔄 数据流架构

```
用户输入 → agent/ → 工具调用 → 操作执行 → 结果返回 → 日志记录
    ↓           ↓           ↓           ↓           ↓
配置验证 → 智能规划 → 安全检查 → 备份保护 → 状态更新
```

### 🎯 设计原则

1. **模块化设计** - 每个模块职责单一，便于维护和扩展
2. **单向数据流** - 严格遵循 DDD 原则，避免循环依赖
3. **错误隔离** - 模块间错误不传播，确保系统稳定性
4. **日志驱动** - 所有操作都有详细日志，便于调试和审计## 🔧 开发指南

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
- 编写单元测试（TODO）

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
**A:** 网页内容会经过摘要处理，但仍会消耗一定 token。建议仅在需要最新信息时启用，或使用 `"use_web_search": false` 关闭。

### Q: 支持哪些 AI 模型？
**A:** 支持所有 OpenAI 兼容 API，包括：
- DeepSeek Chat
- OpenAI GPT 系列
- 其他兼容 OpenAI API 的服务

### Q: 可以在 Linux/macOS 上运行吗？
**A:** 虽然可以，但需要调整构建脚本和部分Shell操作代码以及系统提示词。核心 Go 代码是跨平台的。
- 短期内没有进一步支持 linux 平台的计划

## 📄 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 🙏 致谢

- 感谢 [OpenAI](https://openai.com/) 提供的 API 接口
- 感谢 [DeepSeek](https://www.deepseek.com/) 提供优质的 AI 服务
- 感谢所有贡献者和用户的支持

## 📞 支持与反馈

- **GitHub Issues**: [报告问题或请求功能](https://github.com/asashishi/asashishi-agent/issues)
- **文档**: 查看本 README 和代码注释
- **日志**: 程序运行日志保存在 `log/` 目录

---

**✨ 开始您的智能编程之旅！** 
- 如有任何问题或建议，欢迎参与项目贡献或提交反馈。
