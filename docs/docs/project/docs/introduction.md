# 简介

::: tip 提示
你正在阅读的是 Crm 最新版的[项目文档](https://docs.zocrm.cloud)。
:::

## 什么是 Crm ？

Crm（英文全称 Customer relationship management ）是一个客户关系管理系统，主要功能有仪表盘、客户管理、合同管理、产品管理、配置、订阅等功能。

- 在线演示：[zocrm.cloud](https://zocrm.cloud)

- 项目文档：[docs.zocrm.cloud](https://docs.zocrm.cloud)

## 技术栈

Crm 系统主要采用 Vue3 和 Golang 实现。

### 前端技术

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Vue.js | 前端框架 | https://v3.cn.vuejs.org |
| Vue Router | 页面路由 | https://router.vuejs.org |
| Axios | 网络请求库 | https://axios-http.com |
| Pinia | 状态管理 | https://pinia.vuejs.org |
| Vite | 构建工具 | https://vitejs.dev |
| Ant Design Vue | 前端UI组件库 | https://www.antdv.com |
| Apache ECharts | 可视化图表库 | https://echarts.apache.org |
| Moment | 日期库 | https://momentjs.com |

### 后端技术

| 技术 | 说明 | 相关文档 |
|---|---|---|
| Gin | Web框架 | https://gin-gonic.com |
| Gorm | ORM框架 | https://gorm.io |
| Jwt | 用户认证 | https://github.com/golang-jwt/jwt |
| Viper | 配置管理 | https://github.com/spf13/viper |
| Redis | 数据缓存 | https://github.com/go-redis/redis |
| Gomail | 邮件服务SDK | https://github.com/go-gomail/gomail |
| Gopay | 支付服务SDK | https://github.com/go-pay/gopay |

## 目录结构

Crm 的目录结构如下：

```text
crm
 ├── .github/workflows       // 工作流
     ├── deploy.yaml         // 自动部署配置文件
 ├── docs                    // 项目文档
     ├── docs                // 项目文档
         ├── .vitepress      // VitePress配置
         ├── logs            // 更新日志
         ├── project         // Markdown文档
         ├── sponsor         // 赞赏
         ├── index.md        // 文档首页
     ├── package-lock.json   // Npm依赖管理
     ├── package.json        // Npm依赖管理
 ├── server                  // 服务端
     ├── api                 // API层
     ├── common              // 通用的工具
     ├── config              // 配置文件
     ├── dao                 // 数据访问层
     ├── db                  // 数据库 SQL 文件
     ├── global              // 全局对象
     ├── initialize          // 初始化
     ├── middleware          // 中间件
     ├── models              // 数据模型
     ├── response            // 数据响应封装
     ├── service             // service层
     ├── config.yaml         // Yaml配置文件
     ├── go.mod              // Go模块
     ├── go.sum              // Go模块相关
     ├── main.go             // 程序启动的入口
 ├── web                     // Web端
     ├── public              // 公共的资源
     ├── src                 // 源代码
         |── api             // API接口
         ├── assets          // 资源
         ├── axios           // 网络请求
         ├── components      // 自定义组件
         ├── router          // 页面路由
         ├── store           // 状态管理
         ├── views           // 页面
         ├── App.vue         // 组件入口
         ├── main.js         // 程序启动的入口
     ├── .env.dev            // 开发模式环境变量
     ├── .env.prod           // 生产模式环境变量
     ├── index.html          // 首页
     ├── package-lock.json   // Npm依赖管理
     ├── package.json        // Npm依赖管理
     ├── vite.config.js      // Vite配置文件
 ├── .gitattributes          // Git描述
 ├── .gitignore              // Git文件忽略
 ├── LICENSE                 // 许可证
 ├── README.md               // 项目简介文档
```

## 系统架构

Crm 系统采用前后端分离架构，前端与后端分开部署，且部署到同一台服务器。

系统架构图：

<img src="/sys_structure.png" style="border-radius: 5px;" alt="system tructure"/>

## 许可证

Crm 是采用 MIT 许可的开源项目，使用完全免费。想要了解有关 MIT 许可证的更多信息，请访问[MIT License](https://github.com/zchengo/crm/blob/main/LICENSE)。