# STU Online Judge System

## 项目简介

STUOJ 是汕头大学疾风算法协会的 ACM-ICPC 在线代码评测系统，基于 Go 语言和 Gin 框架开发。

用户可以在平台上阅读算法题目，并可提交代码到代码沙箱进行评测，评测完成后系统将返回评测结果。管理员可以管理用户、导入题目、修改评测点数据、管理提交记录、查询系统统计数据和修改系统设置。

## 项目仓库

- 后端仓库：[https://github.com/HEX9CF/STUOJ](https://github.com/HEX9CF/STUOJ)
- 前端仓库：[https://github.com/HEX9CF/stuoj-frontend](https://github.com/HEX9CF/stuoj-frontend)
- 数据库仓库：[https://github.com/HEX9CF/stuoj-database](https://github.com/HEX9CF/stuoj-database)

## API 文档

- Apifox：[https://apifox.com/apidoc/shared-431b8879-dd14-41f8-a011-613050aee4f4](https://apifox.com/apidoc/shared-431b8879-dd14-41f8-a011-613050aee4f4)

## 系统架构

- 后端架构：Gin
- 前端架构：Vue
- 数据库：MySQL
- 代码沙箱: Judge0
- 图床服务: yuki-image

## 功能简述

| 功能名称       | 请求方法 | 路由路径                  | 操作者   | 功能简述                                           |
|------------|--------|-----------------------|----------|------------------------------------------------|
| 用户注册       | POST   | /user/register         | 未登录用户 | 用户输入用户名、邮箱和密码，创建账号。                            |
| 用户登录       | POST   | /user/login            | 未登录用户 | 用户输入邮箱和密码，登录账号。                                |
| 用户注销       | POST   | /user/logout           | 用户     | 用户可以退出账号。                                      |
| 找回密码       | POST   | /user/forget           | 用户     | 用户可以发送验证码到自己的邮箱，通过验证码重置密码。                     |
| 查看用户个人空间   | GET    | /user/:id             | 用户     | 用户可以在某个用户的个人空间页面查看该用户的用户名、邮箱、头像和个性签名等信息。       |
| 修改用户信息     | PUT    | /user/modify           | 用户     | 用户可以对自己的用户名、邮箱地址、头像和个性签名进行修改。                  |
| 用户修改密码     | PUT    | /user/password         | 用户     | 用户可以对自己的密码进行修改。                                |
| 提交代码       | POST   | /judge/submit          | 用户     | 用户可以在代码提交页面提交代码到代码沙箱进行评测，评测完成后系统将返回评测结果。       |
| 选择编程语言     | GET    | /judge/language        | 用户     | 用户可以在代码提交页面设置代码使用的编程语言。                        |
| 查看题目列表     | GET    | /problem               | 用户     | 用户可以查看题目列表。                                    |
| 查看题目内容     | GET    | /problem/:id           | 用户     | 用户可以在题目页面来查看单个题目的详细信息，包括标题、题目来源、难度、时间限制、内存限制等。 |
| 查看用户提交记录列表 | GET    | /record/user/:id      | 用户     | 用户可以查看某个用户的所有提交记录，包括用户ID、题目ID、评测状态、分数、语言等。     |
| 查看单题提交记录列表 | GET    | /record/problem/:id   | 用户     | 用户可以查看某个题目的所有提交记录以及其对应的评测状态、运行耗时、内存。           |
| 查看用户列表     | GET    | /admin/user            | 管理员   | 管理员可以查看用户列表。                                   |
| 查看单个用户信息   | GET    | /admin/user/:id        | 管理员   | 管理员可以查看某个用户的完整信息。                              |
| 添加用户       | POST   | /admin/user            | 管理员   | 管理员可以直接添加一个用户。                                 |
| 修改用户信息     | PUT    | /admin/user            | 管理员   | 管理员可以修改某个用户的信息。                                |
| 删除用户       | DELETE | /admin/user/:id        | 管理员   | 管理员可以删除某个用户。                                   |
| 查看题目列表     | GET    | /admin/problem         | 管理员   | 管理员可以查看题目列表。                                   |
| 查看单个题目信息   | GET    | /admin/problem/:id     | 管理员   | 管理员可以查看某个题目的标题、题目来源、难度、时间限制、内存限制、创建时间、更新时间等。   |
| 添加题目       | POST   | /admin/problem         | 管理员   | 管理员可以添加一个题目。                                   |
| 导入题目       | POST   | /admin/problem/import  | 管理员   | 管理员可以批量导入题目。                                   |
| 修改题目       | PUT    | /admin/problem         | 管理员   | 管理员可以修改某个题目的信息。                                |
| 删除题目       | DELETE | /admin/problem/:id     | 管理员   | 管理员可以删除某个题目。                                   |
| 添加评测点数据    | POST   | /admin/testcase        | 管理员   | 管理员可以给某个题目添加新的评测点。                             |
| 修改评测点数据    | PUT    | /admin/testcase        | 管理员   | 管理员可以修改某个题目的某个评测点的数据。                          |
| 删除评测点数据    | DELETE | /admin/testcase/:id    | 管理员   | 管理员可以删除某个题目的某个评测点。                             |
| 查看提交记录列表   | GET    | /admin/record          | 管理员   | 管理员可以查看所有题目的所有用户的提交记录列表。                       |
| 查看提交记录     | GET    | /admin/record/:id      | 管理员   | 管理员可以查看某个提交记录的提交信息以及评测结果。                      |
| 删除提交记录     | DELETE | /admin/record/:id      | 管理员   | 管理员可以删除某个提交记录的提交信息以及评测结果。                      |
| 查询统计数据     | GET    | /admin/statistics      | 管理员   | 管理员可以查询各种统计数据。                                 |
| 查询系统设置     | GET    | /admin/config          | 超级管理员 | 管理员可以查询各种系统设置。                                 |
| 修改系统设置     | PUT    | /admin/config          | 超级管理员 | 管理员可以修改系统设置。                                   |
