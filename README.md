# STU Online Judge System

## 项目简介

STUOJ 是汕头大学疾风算法协会的 ACM-ICPC 在线代码评测系统，基于 Go 语言和 Gin 框架开发。

用户可以在平台上阅读算法题目，并可提交代码到代码沙箱进行评测，评测完成后系统将返回评测结果。管理员可以管理用户、导入题目、修改评测点数据、管理提交记录、查询系统统计数据和修改系统设置。

## 项目仓库

- 后端仓库：[https://github.com/STUOJ/STUOJ](https://github.com/STUOJ/STUOJ)
- 前端仓库：[https://github.com/STUOJ/stuoj-web](https://github.com/STUOJ/stuoj-web)
- 数据库仓库：[https://github.com/STUOJ/stuoj-database](https://github.com/STUOJ/stuoj-database)

## API 文档

- Apifox：[https://apifox.com/apidoc/shared-431b8879-dd14-41f8-a011-613050aee4f4](https://apifox.com/apidoc/shared-431b8879-dd14-41f8-a011-613050aee4f4)

## 系统架构

- 后端架构：Gin + Gorm
- 前端架构：Vue
- 数据库：MySQL
- 代码沙箱: Judge0
- 图床服务: yuki-image

## 系统功能

![](https://github.com/user-attachments/assets/29b370aa-5a02-4e4d-ab37-3949c5dce19e)

### 功能简述

| 功能名称       | 请求方法 | 路由路径                  | 操作者   | 功能简述                                           |
|------------|--------|-----------------------|----------|------------------------------------------------|
| 用户注册       | POST   | /user/register         | 未登录用户 | 用户输入用户名、邮箱和密码，创建账号。                            |
| 用户登录       | POST   | /user/login            | 未登录用户 | 用户输入邮箱和密码，登录账号。                                |
| 获取用户头像     | GET    | /user/avatar/:id       | 用户     | 用户可以获取某个用户的头像。                                 |
| 获取用户信息     | GET    | /user/:id              | 用户     | 用户可以获取某个用户的详细信息。                               |
| 获取当前用户ID   | GET    | /user/current          | 用户     | 用户可以获取当前登录用户的ID。                                |
| 获取当前用户头像 | GET    | /user/avatar           | 用户     | 用户可以获取当前登录用户的头像。                               |
| 修改用户信息     | PUT    | /user/modify           | 用户     | 用户可以修改自己的用户信息。                                 |
| 修改用户密码     | PUT    | /user/password         | 用户     | 用户可以修改自己的密码。                                   |
| 更新用户头像     | POST   | /user/avatar           | 用户     | 用户可以更新自己的头像。                                   |
| 获取题目列表     | GET    | /problem               | 用户     | 用户可以获取题目列表。                                    |
| 根据难度获取题目列表 | GET    | /problem/difficulty/:id | 用户     | 用户可以根据难度获取题目列表。                                 |
| 根据标签获取题目列表 | GET    | /problem/tag/:id       | 用户     | 用户可以根据标签获取题目列表。                                 |
| 获取题目详情     | GET    | /problem/:id           | 用户     | 用户可以获取题目的详细信息。                                 |
| 获取标签列表     | GET    | /problem/tag           | 用户     | 用户可以获取标签列表。                                    |
| 获取编程语言列表   | GET    | /judge/language        | 用户     | 用户可以获取支持的编程语言列表。                               |
| 提交代码       | POST   | /judge/submit          | 用户     | 用户可以提交代码进行评测。                                  |
| 获取提交记录列表   | GET    | /record               | 用户     | 用户可以获取提交记录列表。                                  |
| 获取提交记录详情   | GET    | /record/:id           | 用户     | 用户可以获取提交记录的详细信息。                               |
| 获取用户提交记录   | GET    | /record/user/:id      | 用户     | 用户可以获取某个用户的提交记录。                               |
| 获取题目提交记录   | GET    | /record/problem/:id   | 用户     | 用户可以获取某个题目的提交记录。                               |
| 获取用户列表     | GET    | /admin/user            | 管理员   | 管理员可以获取用户列表。                                   |
| 获取用户详情     | GET    | /admin/user/:id        | 管理员   | 管理员可以获取用户的详细信息。                               |
| 根据角色获取用户列表 | GET    | /admin/user/role/:id   | 管理员   | 管理员可以根据角色获取用户列表。                                |
| 添加用户       | POST   | /admin/user            | 管理员   | 管理员可以添加用户。                                     |
| 修改用户信息     | PUT    | /admin/user            | 管理员   | 管理员可以修改用户信息。                                   |
| 删除用户       | DELETE | /admin/user/:id        | 管理员   | 管理员可以删除用户。                                     |
| 获取题目列表     | GET    | /admin/problem         | 管理员   | 管理员可以获取题目列表。                                   |
| 根据状态获取题目列表 | GET    | /admin/problem/status/:id | 管理员   | 管理员可以根据状态获取题目列表。                                |
| 获取题目详情     | GET    | /admin/problem/:id     | 管理员   | 管理员可以获取题目的详细信息。                               |
| 添加题目       | POST   | /admin/problem         | 管理员   | 管理员可以添加题目。                                     |
| 修改题目       | PUT    | /admin/problem         | 管理员   | 管理员可以修改题目。                                     |
| 删除题目       | DELETE | /admin/problem/:id     | 管理员   | 管理员可以删除题目。                                     |
| 获取题目历史记录   | GET    | /admin/problem/history/:id | 管理员   | 管理员可以获取题目的历史记录。                               |
| 添加题目标签     | POST   | /admin/problem/tag     | 管理员   | 管理员可以给题目添加标签。                                  |
| 删除题目标签     | DELETE | /admin/problem/tag     | 管理员   | 管理员可以删除题目的标签。                                  |
| 获取评测点详情    | GET    | /admin/testcase/:id    | 管理员   | 管理员可以获取评测点的详细信息。                               |
| 添加评测点      | POST   | /admin/testcase        | 管理员   | 管理员可以添加评测点。                                    |
| 修改评测点      | PUT    | /admin/testcase        | 管理员   | 管理员可以修改评测点。                                    |
| 删除评测点      | DELETE | /admin/testcase/:id    | 管理员   | 管理员可以删除评测点。                                    |
| 获取标签列表     | GET    | /admin/tag             | 管理员   | 管理员可以获取标签列表。                                   |
| 添加标签       | POST   | /admin/tag             | 管理员   | 管理员可以添加标签。                                     |
| 修改标签       | PUT    | /admin/tag             | 管理员   | 管理员可以修改标签。                                     |
| 删除标签       | DELETE | /admin/tag/:id         | 管理员   | 管理员可以删除标签。                                     |
| 获取提交记录列表   | GET    | /admin/record          | 管理员   | 管理员可以获取提交记录列表。                                 |
| 获取提交记录详情   | GET    | /admin/record/:id      | 管理员   | 管理员可以获取提交记录的详细信息。                             |
| 删除提交记录     | DELETE | /admin/record/:id      | 管理员   | 管理员可以删除提交记录。                                   |
| 获取统计数据     | GET    | /admin/statistics      | 管理员   | 管理员可以获取统计数据。                                   |
| 修改用户角色     | PUT    | /admin/user/role       | 超级管理员 | 超级管理员可以修改用户角色。                                 |
| 获取系统设置     | GET    | /admin/config          | 超级管理员 | 超级管理员可以获取系统设置。                                 |
| 修改系统设置     | PUT    | /admin/config          | 超级管理员 | 超级管理员可以修改系统设置。                                 |

## UML

### 用例图

![image](https://github.com/user-attachments/assets/d27bc6a6-bcdd-422b-baa5-8a85ba05b79b)

### 状态图

#### 用户
![image](https://github.com/user-attachments/assets/ebcff400-f7ba-433c-93be-47d211beca01)

#### 题目
![image](https://github.com/user-attachments/assets/31519b73-9180-4e6b-bb60-28e562e8e218)

#### 提交记录
![image](https://github.com/user-attachments/assets/4538e435-d39c-4851-90a3-57317acd2be7)

#### 博客
![image](https://github.com/user-attachments/assets/a48cf23f-ba6b-42fa-99be-3d32abde7a9f)

### 活动图

#### 用户

##### 注册
![image](https://github.com/user-attachments/assets/10867d10-bae6-42d8-a613-bf6aed90e071)

##### 登录
![image](https://github.com/user-attachments/assets/cda37df8-469b-46f4-90b6-a74d1c097458)

##### 修改个人信息
![image](https://github.com/user-attachments/assets/cb85d84e-11ce-4d43-b6d2-c85a799276ad)

##### 修改密码
![image](https://github.com/user-attachments/assets/f98ad919-83bb-4543-bd34-01643962498f)

#### 题目
![image](https://github.com/user-attachments/assets/53bdd18b-8498-45a0-af7a-29253d5c0109)

#### 提交代码
![image](https://github.com/user-attachments/assets/f910a74f-1c15-4a83-aa79-f8b454671f28)

#### 提交记录
![image](https://github.com/user-attachments/assets/e734151a-a403-46da-af01-1a9620f3049c)

