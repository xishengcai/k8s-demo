# 服务网格测试镜像

| 项目角色 | 项目名称 | 类型 |功能|
| ------ | ------ | ------ | -----|
| 前端 | mesh-front|  angular | show canary_image version and show mysql user data|
| 后端 | canary_image  |  golang  | 灰度发布测试镜像，|
| 后端 ｜｜mesh-backend|query user|
｜数据库｜｜mysql｜|


# 测试用例
- 网关 nginx-ingress
- 前端 mesh-front --> canary_image
