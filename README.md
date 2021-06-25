# - 区块链数据真实性检测算法

# 本算法主要包括针对文字的Simhash算法，针对图片的dhash、ahash、phash算法，以及对于数据库的相关使用


# 代码上传命令
#（1）（ git add .），若遇到如图3中的问题，是因为没有github库的地址，可以采用（git remote add origin https://github.com/hudiankai/Golearn.git）的命令重新设定相关github库的地址，后面的https是相关的地址。（注释：网上有说在git clone + https地址时可能会超时，相关办法是把https变为git）
#（2）（git commit -m "相关备注"），
#（3）（git pull），拉取github上的最新代码，遇到图3中的超时问题，
#解决方法：如下图所示，与（git push）超时一样，采用命令GOPROXY设定代理（set GOPROXY=https://goproxy.io,https://mirrors.aliyun.com/goproxy/,https://goproxy.cn,direct）
#（4）（git push）会弹出相关git界面，需要输入账号和密码，