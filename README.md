# kklogTUI

kklog 的终端用户界面，可视化生成查看日志的命令，配合 [kklog](https://gitlab.weike.fm/zk/kklog) 一起食用

通过界面操作选择：

- 1 环境
  - 测试 dev
  - 生产 prod
- 2 项目
  - 即项目名，注意项目名的分割符都是中划线 比如：wk_tag_manage 在集群中的 Deployment 是 wk-tag-manage
- 3 服务
  - 项目的 API 服务或者起的脚本
- 4 [命名空间] - 仅当选测试环境时需要
  - 测试环境下 sit dev1~n 等，手动选择
  - 生产环境下 production core iprod 等，已与项目绑定，不需要手动选择

然后将已选项传递给 kklog，命令格式如下：

```sh
/usr/local/bin/kklog -d wk-miniprogram-cms -e dev -n wk-miniprogram-cms -ns sit -t api -l 500 2> /tmp/kklog_grep_buf_`date +%s` | tail -f /tmp/kklog_grep_buf_`date +%s`
```

## 使用

> 请确保已安装 `kklog`，检查方式 `kklog -h`

- 拉项目
- 安装

    ```sh
    make
    make install
    ```

- 执行

    ```sh
    kklogTUI
    ```

## 新增项目

配置于 `constant/constant.go`，修改请提 PR

## 开发

- 拉项目
- 装环境 [当前版本 go==1.18]

    ```sh
    go mod tidy
    ```

- 本地启动

    ```sh
    make run
    ```
