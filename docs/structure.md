### 目录结构

```
crm
  ├── doc               // 项目文档
    ├── structure.md        // 目录结构
    ├── techstack.md        // 技术栈
  ├── server            // 服务端
    ├── api                 // 接口层
    ├── common              // 公共的方法
    ├── config              // 配置文件映射的结构体
    ├── db                  // 数据库文件
    ├── global              // 全局变量
    ├── initialize          // 初始化运行环境
    ├── middleware          // 中间件
    ├── models              // 结构体
    ├── response            // 响应结果封装
    ├── service             // 服务层
    ├── config.yaml         // 配置文件
    ├── go.mod              // Go模块管理
    ├── go.sum              // Go模块管理
    ├── main.go             // 程序执行入口
  ├── web               // Web端
    ├── public              // 默认生成的
    ├── src                 // 源代码
        ├── api             // 请求接口
        ├── assets          // 静态资源
        ├── axios           // 请求库封装
        ├── components      // 组件
        ├── router          // 页面路由
        ├── store           // 状态管理
        ├── views           // 页面
        ├── App.vue         // 全局组件
        ├── main.js         // 程序执行入口
    ├── index.html          // 首页
    ├── package-lock.json   // 依赖包
    ├── package.json        // 依赖包
    ├── vite.config.js      // Vite配置文件
  ├── .gitignore        // Git文件忽略
  ├── LICENSE           // 开源许可证
  ├── README.md         // 项目文档
```