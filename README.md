## 验证码服务

目前在 mac centos6.x ubuntu14 上测试通过

![demo.gif](https://raw.githubusercontent.com/jianxinio/captcha/master/src/captcha/tmp/demo.gif)

- - -

#### 准备工作

依赖安装： `golang` `imagemagic`

验证成功：

`pkg-config --cflags --libs MagickWand`

`go version`

没有报错说明安装成功

#### 下载并编译

`git clone https://github.com/jianxinio/captcha/`

或者 [下载](https://github.com/jianxinio/captcha/archive/master.zip) 并解压


进入captcha文件夹后运行 `source install.sh`

`Downloading necessary files` 这步可能耗时很久，后面留意观察错误，如果下载完后持续报错可能是 imagemagic 安装有误

如果一切顺利，可以看到多了一个 build 文件夹，其中

1. bin 为可执行文件夹
2. assets 为使用到的静态资源文件夹
3. tmp 为缓存验证码文件

除了 build 文件夹外，其他文件不会再被使用

#### 启动

进入到 bin 文件夹，`./captcha` 即可启动。

见到 `Init success.` 的提示说明初始化缓存生成成功。此时可以在 tmp 文件夹中看到 100 张验证码。

线上环境请自备守护进程

#### 程序集成

访问 localhost:8001 即可见到验证码内容（格式为 base64(buffer)|result）。

使用时需要先 split('|') 然后将 base64 解密后给前端。

	get '/captcha' do
	  captcha = Faraday.get settings.captcha_server
	  captcha_arr = captcha.body.split('|')
	  @session['captcha_result'] = captcha_arr[1]
	  content_type 'image/gif'
	  Base64.decode64 captcha.body.split('|')[0]
	end

#### 高级配置

BriefDesign 中有详细的设计文档。

config.json 说明：

	initial_count 初始化时产生的验证码数量
	check_interval 检查使用量间隔
	threshold 更新验证码阀值
	update_count 更新的验证码数量

如果访问量较大，建议提高 initial_count 以及 update_count

更新配置文件后 reload 操作：

reload 只支持 initial_count 以及 update_count 的修改

操作方法： `kill -USR2 pid`


#### 开源协议

MIT