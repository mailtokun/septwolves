# YuTu Gitops Engine
[玉兔Gitops引擎](http://yutu.pub)
## 目标用户
个人开发者或微小企业.
## 优势
只需要一条命令即完成CICD的配置. 

## 场景
### 场景 1: 将代码部署到 Docker 中
在目标机器上运行以下示例代码即可完成 CICD 的配置. 同时提供了两个示例
- [golang 示例](https://github.com/mailtokun/yutu-go-example)
- [vue 示例](https://github.com/mailtokun/yutu-vue-example)
```
mkdir ~/.yutu || true
cat >~/.yutu/projects.json <<EOL
[
    {
        "githubRepo": "https://github.com/xxxx/xxxx",
        "githubBranch": "main",
        "githubToken": "xxxxxxx-xxxxx-xxxxx-xxxxx"
    }
]
EOL
docker run -d --network="host" --name=yutu \
-v /var/run/docker.sock:/var/run/docker.sock \
-v ~/.yutu/projects.json:/yutu/projects.json \
mailtokun/yutu /yutu/main
```

### 场景 2: 将代码部署到 kubernetes 中
暂不支持


[支持一下作者](https://www.buymeacoffee.com/coffeefree)
