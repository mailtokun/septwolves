# YuTu Gitops Engine
玉兔Gitops引擎
## 目标用户
个人开发者或微小企业.
## 优势
只需要一条命令即完成CICD的配置. 

## 场景
### 场景 1: 将代码部署到 Docker 中
```
docker run -d --network="host" --name=yutu \
-v /var/run/docker.sock:/var/run/docker.sock \
--env GITHUB_REPO=https://github.com/xxxx/xxxx \
--env GITHUB_BRANCH=main \
--env GITHUB_TOKEN=xxxxxxx-xxxxx-xxxxx-xx \
mailtokun/yutu /yutu/main
```

### 场景 2: 将代码部署到 kubernetes 中
