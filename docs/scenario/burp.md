# 使用Burp与Xray进行联动
  
在实际测试过程中，除了被动扫描，也时常需要手工测试。这里使用Burp的原生功能与xray建立起一个多层代理，让流量从Burp转发到Xray中。  

首先 xray 建立起 webscan 的监听：  
![](https://raw.githubusercontent.com/Lz1y/imggo/master/20191024165621.png)
  
进入Burp后，打开`User options`标签页，然后找到`Upstream Proxy Servers`设置。  
点击Add添加上游代理以及作用域，`Destination host`处可以使用`*`匹配多个任意字符串，`?`匹配单一任意字符串，而上游代理的地址则填写Xray的监听地址。
  
![](https://raw.githubusercontent.com/Lz1y/imggo/master/20191024165734.png)  
接下来，在浏览器端代理上Burp的代理地址。
![](https://raw.githubusercontent.com/Lz1y/imggo/master/20191024170238.png)  

此时，请求已经通过了Burp  
![](https://raw.githubusercontent.com/Lz1y/imggo/master/20191024170348.png)  

转发到了Xray中  
![](https://raw.githubusercontent.com/Lz1y/imggo/master/20191024170458.png)  

至此，联动成功。
