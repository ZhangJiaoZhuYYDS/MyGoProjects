package main

import (
	"fmt"
	"gin_test/middlewares"
	"gin_test/modles"
	"github.com/gin-contrib/sessions"
	redis2 "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Topic struct {
	Title   string   `json:"title"`
	Options []Option `json:"options"`
}

type Option struct {
	Name string `json:"name"`
}

func main() {
	//http.Handle("/", http.FileServer(http.Dir(".")))
	//http.ListenAndServe(":8080", nil)
	g := gin.Default()



	// 设置跨域
	//config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://localhost:63342"}
	//config.AllowCredentials = true
	//g.Use(cors.New(config))


	// Set up Redis client
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})

	// Set up Redis store
	store, err := redis2.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		panic(err)
	}
	// Set up session middleware
	g.Use(sessions.Sessions("mysession", store))

	// Define routes
	g.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Set("foo", "bar")
		session.Save()
		c.String(200, "Session value set")
	})

	g.GET("/getCookie", func(c *gin.Context) {
		session := sessions.Default(c)
		foo := session.Get("foo")
		c.String(200, "Session value: %v", foo)
	})



	g.LoadHTMLGlob("./static/*")

	g.GET("upIndex", func(c *gin.Context) {
		c.HTML(200,"demo.html",nil)
	})
	g.POST("/up", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
			return
		}

		content, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "文件读取失败"})
			return
		}

		fmt.Println(header.Filename)
		fmt.Println(string(content))
		c.JSON(http.StatusOK, gin.H{"msg": "上传成功"})
	})

	// 断点续传
	g.POST("/upload", func(c *gin.Context) {
			// get the file from the request
			file, err := c.FormFile("file")
			if err != nil {
				c.String(http.StatusBadRequest, "Bad request")
				return
			}

			// check if the file already exists on disk
			fileInfo, err := os.Stat("path/to/save/file")
			if err == nil {
				// file exists, resume upload
				// OpenFile函数用于打开磁盘上的文件。os的第一个参数。OpenFile是我们要打开的文件的路径。在本例中，路径为“path/to/save/file”。
				//		第二个参数是一组标志，决定如何打开文件。操作系统。O_APPEND标志用于在文件已经存在的情况下向文件追加。O_WRONLY标志用于以只写方式打开文件。
				//		第三个参数是文件模式，它决定了文件的权限。在本例中，文件模式是0644，这意味着文件所有者可读可写，其他人也可读。
				file, err = os.OpenFile("path/to/save/file", os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					c.String(http.StatusInternalServerError, "Internal server error")
					return
				}
				// 文件光标移动到文件末尾
				_, err = file.Seek(fileInfo.Size(), 0)
				if err != nil {
					c.String(http.StatusInternalServerError, "Internal server error")
					return
				}
			} else {
				// file does not exist, create new file
				file, err = os.Create("path/to/save/file")
				if err != nil {
					c.String(http.StatusInternalServerError, "Internal server error")
					return
				}
			}

			// append the contents of the uploaded file to the file on disk
			_, err = io.Copy(file, file)   // 1 目标文件，要写入的文件 2 源文件
			if err != nil {
				c.String(http.StatusInternalServerError, "Internal server error")
				return
			}

			c.String(http.StatusOK, "File uploaded successfully")
		})































	g.GET("/te", func(context *gin.Context) {
		context.JSON(200, "6666")
		return
		fmt.Println(666)
	})
	g.POST("/topics", func(c *gin.Context) {
		var topic Topic
		if err := c.ShouldBindJSON(&topic); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// Do something with the topic data, e.g. save it to a database

		c.JSON(200, gin.H{"message": "投票主题创建成功！"})
	})

	g.GET("/index", func(c *gin.Context) {
		c.HTML(200,"form.html",nil)
	})

	g.POST("/test", func(c *gin.Context) {
		type VoteForm struct {
			Title string   `form:"title" json:"title"`
			Description string `form:"description" json:"description"`
			Type string`form:"type" json:"type"`
			Sex string `form:"sex" json:"sex"`
			Hobbies []string `form:"hobbies" json:"hobbies"`
		}
		 vote := VoteForm{}
		err := c.ShouldBindJSON(&vote)
		if err != nil {
			c.JSON(400,"cuole")
			return
		}
		c.JSON(200,gin.H{
			"data":vote,
		})

	})

	g.POST("/login", func(c *gin.Context) {
		type User struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Do something with the user data...

		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})
	
	
	g.Use(middlewares.Logger())

	//登录判断是否成功，然后下发token
	g.POST("/logins", func(context *gin.Context) {

		var s modles.ParamLogin
		err := context.Bind(&s)
		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code":    "101",
				"message": "获取账号密码失败",
			})
		}
		token, err := middlewares.Login(&s)
		if err != nil {
			fmt.Println("获取token失败la", err)
		}
		fmt.Println(token)
	})
	// 访问“/ping”时，调用JWT中间件
	g.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		value, _ := c.Get("example")
		fmt.Println(value)
		//如果是登陆的用户，判断请求头中是否含有有效JWT，
		value, exists := c.Get("userID")
		fmt.Println(value, exists)
		c.String(http.StatusOK, "pong")
	})

	g.POST("/post", func(context *gin.Context) {
		dicts := context.PostFormArray("hobbies")
		fmt.Println(dicts)
	})

	g.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)

		// 打印："12345"
		log.Println(example)
	})

	g.GET("/get", func(context *gin.Context) {
		context.HTML(http.StatusOK, "demo.html", gin.H{
			"message": "123",
		})
	})
	g.GET("/get1", func(context *gin.Context) {
		// 外重定向
		context.Redirect(http.StatusTemporaryRedirect, "https://www.baidu.com")
		// 内重定向
		context.Request.URL.Path = "/get2"
		g.HandleContext(context)
	})
	g.GET("/get2", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "get2",
		})
	})
	v1 := g.Group("/v1")
	{
		v1.GET("/v2", func(c *gin.Context) {

			response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
			if err != nil || response.StatusCode != http.StatusOK {
				c.Status(http.StatusServiceUnavailable)
				return
			}

			reader := response.Body
			contentLength := response.ContentLength
			contentType := response.Header.Get("Content-Type")

			extraHeaders := map[string]string{
				"Content-Disposition": `attachment; filename="gopher.png"`,
			}

			c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

			c.String(http.StatusOK, "6666")
		})
	}
	// 设置和获取cookie
	g.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
			c.JSON(200,gin.H{
				"msg":666,
			})
		}
		fmt.Printf("Cookie value: %s \n", cookie)
		g , _ :=c.Cookie("Goland-5129773a")
		fmt.Println(g)
	})
	// 从reader中读取数据
	g.Run()

}
