# Gos - Go项目脚手架生成器

Gos 是一个轻量级的Go项目脚手架生成工具，帮助你快速创建标准化的Go项目结构。

## 功能特点

- 生成标准的Go项目目录结构
- 包含基础的项目配置文件
- 提供常用的工具包和中间件
- 集成日志、配置管理等基础组件
- 包含完整的Makefile构建脚本

## 安装

```bash
go install github.com/wtpevxcrqw/gos@latest
```

## 快速开始

### 1. 创建新项目
```bash
# 创建一个名为myproject的新项目
gos new myproject

# 进入新项目目录
cd myproject

# 初始化依赖
go mod tidy

# 初始化配置
cp config.yaml.example config.yaml

# 运行项目
go run main.go

# 测试
curl http://localhost:8080/ping
```
### 2. 项目结构说明
```bash
myproject/
├── cmd/ # 命令行工具目录
│ └── server/ # 服务器启动命令
├── internal/ # 私有应用代码
│ ├── config/ # 配置管理
│ │ ├── config.go # 配置结构定义
│ │ └── loader.go # 配置加载器
│ ├── handler/ # HTTP处理器
│ │ └── example.go # 示例处理器
│ └── router/ # 路由管理
│ └── router.go # 路由注册
├── pkg/ # 公共包
│ └── logger/ # 日志包
│ └── logger.go # 日志实现
├── config.yml # 配置文件
├── config.yml.example # 配置文件示例
├── go.mod # Go模块文件
├── main.go # 主程序入口
├── Makefile # 构建脚本
└── README.md # 项目说明文档
```

## 贡献指南

1. Fork 项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情

## 联系方式

如有问题或建议，请提交 [Issue](https://github.com/yourusername/gos/issues)

## 注意
本项目仅为示例，实际项目开发中请根据自身需求进行调整和扩展。