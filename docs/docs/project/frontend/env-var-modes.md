# 环境变量和模式

Crm 采用 Vite 提供的环境变量与模式，来设置应用的环境变量与运行模式。

了解有关 Vite 环境变量与模式的更多信息，请参考文档[Vite 环境变量与模式](https://vitejs.dev/guide/env-and-mode.html#env-files)。

## 创建 ```.env``` 文件

在 ```crm/web/``` 目录下创建 ```.env``` 文件，Vite 会从你的 ```环境目录``` 中的下列文件加载额外的环境变量：

```txt
.env                # 所有情况下都会加载
.env.local          # 所有情况下都会加载，但会被 git 忽略
.env.[mode]         # 只在指定模式下加载
.env.[mode].local   # 只在指定模式下加载，但会被 git 忽略
```

在 Crm 系统中，只定义了开发模式与生产模式：
```txt
.env.dev          # 开发模式的环境变量
.env.prod         # 生产模式的环境变量
```

## 定义环境变量

分别在开发模式和生产模式中定义开发环境和生产环境所需要环境变量。

请参考 ```crm/web/.env.dev``` 文件如下：

```txt
# .env.development
VITE_API_BASE_URL=http://127.0.0.1:8000/api
VITE_FILE_UPLOAD_URL=http://127.0.0.1:8000/api/common/file/upload
```

请参考 ```crm/web/.env.prod``` 文件如下：

```txt
# .env.production
VITE_API_BASE_URL=https://zocrm.cloud/api
VITE_FILE_UPLOAD_URL=https://zocrm.cloud/api/common/file/upload
```

## 加载环境变量

通过 ```import.meta.env``` 来获得环境变量：

```js
console.log(import.meta.env.VITE_API_BASE_URL) // http://127.0.0.1:8000/api
console.log(import.meta.env.VITE_FILE_UPLOAD_URL) // http://127.0.0.1:8000/api/common/file/upload
```

请注意，Vite 只会加载以 ```VITE_``` 为前缀的变量。

在 ```crm/web/src/axios/index.js``` 中加载环境变量：

```js
import axios from 'axios';
import { message } from 'ant-design-vue';

axios.defaults.baseURL = import.meta.env.VITE_API_BASE_URL

const request = axios.create({
    // timeout: 5000,`
    headers: {
        'Content-Type': 'application/json;charset=UTF-8'
    }
})
...
```

## 运行模式

在 ```crm/web/package.json``` 中添加运行模式，请参考如下：

```json
{
  "name": "web",
  "private": true,
  "version": "0.0.0",
  "type": "module",
  "scripts": {
    "dev": "vite --mode dev",
    "build": "vite build --mode prod",
    "preview": "vite preview"
  },
  "dependencies": {
    ...
  },
  "devDependencies": {
    ...
  }
}
```

