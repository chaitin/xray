package evil

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"time"

	"github.com/dengsgo/math-engine/engine"
)

func ServeEvilServer(addr string, hard bool) error {
	return http.ListenAndServe(addr, NewEvilServeMux(hard))
}

func NewEvilServeMux(hard bool) *http.ServeMux {
	s := http.NewServeMux()
	mathRe := regexp.MustCompile(`\d+\s*[-+*/]\s*\d+`)
	sleepRe := regexp.MustCompile(`(?i)sleep\((\d+)\)`)
	waitForRe := regexp.MustCompile(`(?i)waitfor\s+delay\s+'0:0:(\d+)'`)
	s.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		buf := bufPool.Get().(*bytes.Buffer)
		defer func() {
			buf.Reset()
			bufPool.Put(buf)
		}()

		buf.Write(CommonEvilResponse)

		data, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Println(err)
		}
		buf.Write(data)

		if hard {
			Split(data, SepFunc, func(bytes []byte) bool {
				GenEvilContent(buf, bytes)
				return true
			})
		}

		unescape, _ := url.PathUnescape(string(data))
		unescape, _ = url.QueryUnescape(unescape)
		if hard {
			Split([]byte(unescape), SepFunc, func(bytes []byte) bool {
				GenEvilContent(buf, bytes)
				return true
			})
		}

		if hard {
			buf.WriteString("\nroot:x:0:0:root:/root:/bin/bash\n")
			buf.WriteString(`
			; for 16-bit app support
		    [fonts]
			[extensions]
			[mci extensions]
			[files]
			[Mail]
			MAPI=1`)
		}

		// 处理 sleep 和 WAITFOR DELAY
		sleepMatches := sleepRe.FindAllStringSubmatch(unescape, -1)
		for _, match := range sleepMatches {
			if len(match) > 1 {
				sleepTime, _ := strconv.Atoi(match[1])
				if sleepTime > 50 {
					time.Sleep(time.Millisecond * time.Duration(sleepTime))
				} else {
					time.Sleep(time.Second * time.Duration(sleepTime))
				}
			}
		}

		waitForMatches := waitForRe.FindAllStringSubmatch(unescape, -1)
		for _, match := range waitForMatches {
			if len(match) > 1 {
				waitTime, _ := strconv.Atoi(match[1])
				if waitTime > 50 {
					time.Sleep(time.Millisecond * time.Duration(waitTime))
				} else {
					time.Sleep(time.Second * time.Duration(waitTime))
				}
			}
		}

		for _, expr := range mathRe.FindAllString(unescape, -1) {
			r, err := engine.ParseAndExec(expr)
			if err != nil {
				log.Println(err)
				continue
			}
			GenEvilContent(buf, []byte(strconv.Itoa(int(r))))
		}

		_, _ = writer.Write(buf.Bytes())
	})
	return s
}

var bufPool = sync.Pool{New: func() any {
	return bytes.NewBuffer(nil)
}}

func GenEvilContent(dst *bytes.Buffer, data []byte) {
	dst.Write(data)
	hashMD5 := md5.Sum(data)
	dst.Write(hashMD5[:])
	dst.WriteString(" ")
	dst.WriteString(hex.EncodeToString(hashMD5[:]))
	dst.WriteString(" ")
	dst.WriteString(base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hashMD5[:]))))
	dst.WriteString(" ")
	hashSha1 := sha1.Sum(data)
	dst.Write(hashSha1[:])
	dst.WriteString(hex.EncodeToString(hashSha1[:]))
	dst.WriteString(" ")
	dst.WriteString(base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hashSha1[:]))))
	dst.WriteString(" ")
	dst.WriteString(base64.StdEncoding.EncodeToString(data))
	dst.WriteString(" ")
}

// CommonEvilResponse
// 常见md5/sha1/base64 (1-1000的数字)
// 常见登录表单
// 常见错误信息
var CommonEvilResponse = []byte(`<html>
		<head>
		<title>Level1</title>
		<title>OA系统</title>
		</head>
		
		<center><h1>Level1</h1></center>
		
		<header>
		    <h1>OA系统</h1>
		    <nav>
		      <ul>
		        <li><a href="#">首页</a></li>
		        <li><a href="#">通知公告</a></li>
		        <li><a href="#">个人中心</a></li>
		        <li><a href="#">退出登录</a></li>
		      </ul>
		    </nav>
		  </header>
		
		  <main>
		    <section class="banner">
		      <h2>欢迎使用OA系统</h2>
		      <p>这是一款集办公、管理、协同于一体的企业级应用软件。</p>
		    </section>
		
			<form action="/login" method="POST">
			  <div>
				<label for="username">用户名:</label>
				<input type="text" id="username" name="username" required>
			  </div>
			  <div>
				<label for="password">密码:</label>
				<input type="password" id="password" name="password" required>
			  </div>
			  <button type="submit">登录</button>
			</form>
		
		    <section class="news">
		      <h2>新闻动态</h2>
		      <ul>
		        <li><a href="#">公司年会隆重举行</a></li>
		        <li><a href="#">财务部门完成年度结算</a></li>
		        <li><a href="#">研发部门推出新产品</a></li>
		      </ul>
		    </section>
		
		    <section class="notice">
		      <h2>通知公告</h2>
		      <ul>
		        <li><a href="#">关于2022年春节放假的通知</a></li>
		        <li><a href="#">办公室搬迁通知</a></li>
		      </ul>
		    </section>
		  </main>
		
		  <footer>
		    <p>&copy; 2023 OA系统 版权所有</p>
		  </footer>
		
		<h1>Oops! Something went wrong.</h1>
		<p>We're sorry, but an error has occurred while processing your request. Please try again later.</p>
		<p>Error Code: 400</p>
		<p>Error Code: 401</p>
		<p>Error Code: 403</p>
		<p>Error Code: 404</p>
		<p>Error Code: 500</p>
		<p>Error Code: 501</p>
		<p>Error Code: 502</p>
		<p>Error Code: 503</p>
		{
		  "error": {
		    "code": 404,
		    "message": "未找到请求的资源",
		    "details": "请检查您的请求URL是否正确，并确保所请求的资源存在。"
		  }
		}
		</body>
</html>
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<!-- a padding to disable MSIE and Chrome friendly error page -->
<?xml version="1.0" encoding="UTF-8"?>
<?xml version="1.1" encoding="UTF-8"?>`)

func init() {
	buf := bytes.NewBuffer(nil)
	buf.Write(CommonEvilResponse)
	for i := 0; i < 1000; i++ {
		GenEvilContent(buf, []byte(strconv.Itoa(i)))
	}
	CommonEvilResponse = buf.Bytes()
}
