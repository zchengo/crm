# 持续集成与交付 (CI/CD)

Crm 系统使用 GitHub Actions 作为持续集成和持续交付 (CI/CD) 平台，自动执行生成、测试和部署，通过创建工作流程来自动构建和部署项目到生产环境。

想要了解有关 GitHub Actions 的更多信息，请访问[GitHub Actions](https://docs.github.com/en/actions)。

## 创建工作流

在您的 Github 存储库中创建 ```.github/workflows``` 目录，在该目录下创建一个 ```Ymal``` 文件（如 deploy.ymal）。

## 编写配置文件

配置文件 ```deploy.ymal```， 请参考如下：

```yaml
name: CRM CI

on:
  push:
    branches: [ "main" ]
  pull_request:
    branches: [ "main" ]
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    
    - name: Use Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18.12.0'    
        
    - name: Build Web
      run: cd web && npm install && npm run build

    - name: Build Docs
      run: cd docs && npm install && npm run docs:build
    
    - name: Use Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.19
    
    - name: Build Server
      run: cd server && go mod tidy && go build -o crmserver main.go
    
    - name: Deploy CRM
      env:
        KEY: ${{ secrets.SSH_PRIVATE_KEY }}
        HOST: ${{ secrets.REMOTE_HOST }}
      run: |
        mkdir -p ~/.ssh/ && echo "$KEY" > ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa
        scp -o StrictHostKeyChecking=no -r web/dist ubuntu@${HOST}:/usr/local/nginx/html/
        scp -o StrictHostKeyChecking=no -r docs/docs/.vitepress/dist ubuntu@${HOST}:/usr/local/nginx/html/docs/
        ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /usr/local/nginx/sbin/nginx -s reload"
        scp -o StrictHostKeyChecking=no -r server/crmserver ubuntu@${HOST}:/home/ubuntu/
        ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /home/ubuntu/crmapi/restart.sh > /dev/null 2>&1 &"
```

## 配置文件说明

### 使用 Ubuntu 系统构建应用

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
```

### 构建Web端与项目文档

```yaml
    steps:
    - uses: actions/checkout@v3
    
    - name: Use Node.js
      uses: actions/setup-node@v3
      with:
        node-version: '18.12.0'    
        
    - name: Build Web
      run: cd web && npm install && npm run build

    - name: Build Docs
      run: cd docs && npm install && npm run docs:build
```

### 构建服务端

```yaml
    - name: Use Go
      uses: actions/setup-go@v3
      with:
        go-version: 1.19
    
    - name: Build Server
      run: cd server && go mod tidy && go build -o crmserver main.go
```

### 部署到服务器

```yaml
    - name: Deploy CRM
      env:
        KEY: ${{ secrets.SSH_PRIVATE_KEY }}
        HOST: ${{ secrets.REMOTE_HOST }}
      run: |
        mkdir -p ~/.ssh/ && echo "$KEY" > ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa
        scp -o StrictHostKeyChecking=no -r web/dist ubuntu@${HOST}:/usr/local/nginx/html/
        scp -o StrictHostKeyChecking=no -r docs/docs/.vitepress/dist ubuntu@${HOST}:/usr/local/nginx/html/docs/
        ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /usr/local/nginx/sbin/nginx -s reload"
        scp -o StrictHostKeyChecking=no -r server/crmserver ubuntu@${HOST}:/home/ubuntu/
        ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /home/ubuntu/crmapi/restart.sh > /dev/null 2>&1 &"
```

设置 ```SSH_PRIVATE_KEY``` 和 ```REMOTE_HOST```：

```SSH_PRIVATE_KEY``` 指的是你在自己电脑上生成的 SSH 私钥（与服务器关联的那个秘钥），```REMOTE_HOST``` 指的是服务器的IP地址。

由于这两个变量属于敏感信息，不方便直接写在配置文件中，所以需要在您的 ```代码库``` -> ```Settings``` -> ```secrets``` -> ```Actions``` 中设置。设置后，可通过 ```secrets``` 来调用。

在虚拟机环境中配置 ```SSH```：

```bash
mkdir -p ~/.ssh/ && echo "$KEY" > ~/.ssh/id_rsa && chmod 600 ~/.ssh/id_rsa
```

将构建的页面资源上传到服务器的特定目录：

```bash
scp -o StrictHostKeyChecking=no -r web/dist ubuntu@${HOST}:/usr/local/nginx/html/
scp -o StrictHostKeyChecking=no -r docs/docs/.vitepress/dist ubuntu@${HOST}:/usr/local/nginx/html/docs/
```

连接到服务器、重启 Nginx 服务：
```bash
ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /usr/local/nginx/sbin/nginx -s reload"
```

将构建的Go可执行文件上传到服务器的特定目录：

```bash
scp -o StrictHostKeyChecking=no -r server/crmserver ubuntu@${HOST}:/home/ubuntu/
```

使用 Shell 脚本重启Go服务：

```bash
ssh -o StrictHostKeyChecking=no ubuntu@${HOST} "sudo /home/ubuntu/crmapi/restart.sh > /dev/null 2>&1 &"
```

Shell 重启Go服务脚本 ```restart.sh```，请参考如下：

```shell
#!/bin/bash
sudo pkill crmserver
sudo cp /home/ubuntu/crmserver /home/ubuntu/crmapi/
sudo nohup /home/ubuntu/crmapi/crmserver > /home/ubuntu/crmapi/crmserver.log 2>&1 &
```

## 常见问题

遇到问题，请访问[GitHub Actions](https://docs.github.com/en/actions)查看更多信息。




