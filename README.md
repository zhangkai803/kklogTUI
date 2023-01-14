# kklogTUI

kklog 的终端用户界面，可视化生成查看日志的命令，配合 [kklog](https://gitlab.weike.fm/zk/kklog) 一起食用

通过界面操作选择：

- 环境
  - 测试 dev or 生产 prod
- 命名空间
  - 测试环境下 sit dev1~n 等
  - 生产环境下 production core iprod 等
- 项目
  - 即项目名，注意项目名的分割符都是中划线 比如：wk_tag_manage 在集群中的 Deployment 是 wk-tag-manage
- 服务
  - 项目的 web server 或者起的脚本
- 服务类型
  - 目前有 api / script

然后将已选项传递给 kklog，命令格式如下：

```sh
/usr/local/bin/kklog -d wk-miniprogram-cms -e dev -n wk-miniprogram-cms -ns sit -t api -l 500 2> /tmp/kklog_grep_buf_`date +%s` | tail -f /tmp/kklog_grep_buf_`date +%s`
```
