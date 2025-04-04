# GXX - 新一代基于YAML的指纹识别工具

GXX是一款强大的指纹识别工具，基于YAML配置的规则进行目标系统识别。
本工具支持多种协议（HTTP/HTTPS、TCP、UDP），可进行高效的批量目标扫描和精准识别。

## 功能特点

- **基于YAML配置**：使用简洁明了的YAML格式定义指纹识别规则
- **多协议支持**：支持HTTP/HTTPS、TCP、UDP协议
- **代理功能**：支持配置HTTP/SOCKS5代理
- **批量扫描**：支持从文件读取多个目标进行批量扫描
- **多格式输出**：支持TXT/CSV等多种输出格式
- **自定义规则**：可根据需要自定义指纹识别规则
- **调试模式**：内置调试功能，便于排查问题

## 快速开始

### 安装

```bash
# 使用 go install 安装
go install github.com/yourusername/gxx@latest

# 或下载预编译版本
# 从 releases 页面下载对应平台的二进制文件
```

### 基本使用

```bash
# 扫描单个目标
gxx -u https://example.com

# 从文件读取目标列表
gxx -f targets.txt

# 使用代理
gxx -u https://example.com --proxy http://127.0.0.1:8080

# 指定输出文件
gxx -u https://example.com -o results.txt

# 开启调试模式
gxx -u https://example.com --debug
```

## 命令行参数

### 输入选项
- `-u, --url`: 要扫描的目标URL/主机（可指定多个）
- `-f, --file`: 包含目标URL/主机列表的文件（每行一个）
- `-t, --threads`: 并发线程数（默认：10）

### 输出选项
- `-o, --output`: 输出文件路径
- `--format`: 输出文件格式（支持 txt/csv，默认：txt）

### 调试选项
- `--proxy`: HTTP/SOCKS5代理（支持逗号分隔的列表或文件输入）
- `-p, --poc`: 测试单个YAML文件
- `-pf, --poc-file`: 测试指定目录下的所有YAML文件
- `--debug`: 开启调试模式

## 项目目录结构

```
gxx/
├── cmd/                    # 命令行应用程序入口点
├── utils/                  # 工具和核心功能代码
│   ├── config/             # 配置管理
│   ├── logger/             # 日志管理
│   ├── common/             # 通用工具函数
│   ├── finger/             # 核心指纹识别功能
│   ├── proto/              # 协议相关代码
│   ├── cel/                # CEL表达式处理
│   ├── request/            # 请求处理
│   └── pkg/                # 可以被外部应用程序使用的库代码
├── finger/                 # 指纹文件目录
├── test/                   # 测试目录
├── logs/                   # 日志输出目录
├── go.mod                  # Go模块定义
├── go.sum                  # Go模块依赖校验和
└── README.md               # 项目说明文档
```

## 编译

### 使用 Makefile (推荐)

```bash
# 构建项目（不嵌入指纹库）
make build

# 构建项目（嵌入指纹库）
make build-embed

# 构建发布包（不嵌入指纹库）
make release

# 构建发布包（嵌入指纹库）
make release-embed

# 查看所有可用命令
make help
```

### 手动编译

```bash
# 基本编译
CGO_ENABLED=0 GOARCH=arm64 GOOS=darwin go build -ldflags "-w -s" -o gxx main.go

# 使用构建脚本
chmod +x build.sh
./build.sh
```

### 使用goreleaser编译

```bash
goreleaser build --snapshot --clean --snapshot
```

## API使用

- `FingerScan`: 指纹扫描API函数
- `NewFingerOptions`: 创建新的指纹扫描选项

## 指纹规则格式

详细的指纹规则格式说明请参考 [指纹规则格式说明.md](docs/指纹规则格式说明.md) 和 [指纹开发快速参考.md](docs/指纹开发快速参考.md)。

### 基本结构示例

```yaml
id: web-application
    
info:
  name: Web应用识别
  author: 作者名
  description: 识别特定Web应用
  reference:
    - https://example.com
  created: 2025/04/01
    
rules:
  r0:
    request:
      method: GET
      path: /
    expression: response.status == 200 && response.body.bcontains(b"特征字符串")
    
expression: r0()
```

## 贡献方式

- **规则贡献**：通过添加新的YAML格式指纹规则文件扩展指纹库
- **代码贡献**：遵循项目代码结构进行功能开发

## 许可证

[许可证信息](LICENSE)