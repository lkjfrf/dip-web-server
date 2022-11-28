package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	db "github.com/DW-inc/LauncherFileServer/DB"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/logger"
<<<<<<< HEAD
	"github.com/gofiber/fiber/v2/middleware/session"

	cryptoRand "crypto/rand"
=======
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/template/html"
>>>>>>> parent of 6aa176f (세션으로 id전달)
)

var Charset string

var Tokens sync.Map

type Player struct {
	Name   string
	SSO_ID string
	Grade  string
	Team   string
}
type SaveLogin struct {
	SsoId string
}

func main() {
	Tokens = sync.Map{}

	db.GetDBManager().Init()
<<<<<<< HEAD
	logm.GetLogManager().SetLogFile()
	app := fiber.New(fiber.Config{})
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New(logger.ConfigDefault))
	SessionStorage = session.New(session.Config{
		CookieName: "dip_session",
=======

	engine := html.New("./WebServer", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	//app := fiber.New()
	app.Use(cors.New(cors.ConfigDefault))
	app.Use(logger.New(logger.ConfigDefault))

	app.Static("/", "./WebServer")
	app.Static("/empty", "./Test")

	app.Use("/files", filesystem.New(filesystem.Config{
		Root: http.Dir("./file"),
	}))

	app.Use("/launcher", func(c *fiber.Ctx) error {
		c.Context().RemoteAddr()
		return c.Download("./Launcher/Setup.zip", "Setup.zip")
>>>>>>> parent of 6aa176f (세션으로 id전달)
	})

	app.Get("/monit", monitor.New(monitor.Config{
		Title: "DIP Web Monitor",
	}))

	app.Get("/sso/autologin", func(c *fiber.Ctx) error {
<<<<<<< HEAD
		//------------ Get SSOID to Apache  ------------//
		sso_id := c.Get("SSO_ID")
		if sso_id == "" {
			log.Println("AUTOLOGIN: SSO_ID is NULL")
			return nil
		}
		log.Println("AUTOLOGIN RECV:", sso_id)
		//------------ Get SSOID to Apache  ------------//
		//------------ Set Session ID  ------------//
		sess, err := SessionStorage.Get(c)
		if err != nil {
			log.Println("SessionGet Err", err)
		}
		sess.Set("ssoid", sso_id)
		if err := sess.Save(); err != nil {
			log.Println("SessionSave Err", err)
		} else {
			log.Println("SessionSave Success", sso_id)
		}
		//------------ Set Session ID  ------------//
		return c.Next()
	})
	app.Static("/sso/autologin", "./WebServer")
	app.Get("/sso/getinfo", func(c *fiber.Ctx) error {
		sso_id := GetSsoidFromSession(c)
		log.Println("sso/getinfo recv ssoid :", sso_id)
=======
		//sso_id := c.Get("SSO_ID")
		sso_id := "1000001"
		log.Println("SSO/AUTOLOGIN Recv : SSO_ID =", sso_id)
>>>>>>> parent of 6aa176f (세션으로 id전달)
		dbData := db.ZCMUSER{}
		err := db.GetDBManager().DBMS.Table("zcmuser").Where("itg_user_id = ?", sso_id).Select("itg_user_nm", "itg_org_nm", "user_poa_nm").First(&dbData).Error

		var data Player

		if sso_id == "" || err != nil {
<<<<<<< HEAD
			if err != nil {
				log.Println("DBError:", err)
			}
			data = SinginData{
				Name:     "",
				SSO_ID:   "",
				Grade:    "",
				Team:     "",
				Security: GetSecurity(),
			}
		} else {
			data = SinginData{
				Name:     dbData.Itg_user_nm,
				SSO_ID:   sso_id,
				Grade:    dbData.UserPoaNm,
				Team:     dbData.Itg_org_nm,
				Security: GetSecurity(),
=======
			//if sso_id != "" || err == nil {
			log.Println(err)
			data = Player{
				//Name:   "tenotsr",
				//SSO_ID: "1000001",
				//Grade:  "idk",
				//Team:   "idkdikdikdidik",
				Name:   "CannotAceess",
				SSO_ID: "CannotAceess",
				Grade:  "CannotAceess",
				Team:   "CannotAceess",
			}
		} else {
			data = Player{
				Name:   dbData.Itg_user_nm,
				SSO_ID: sso_id,
				Grade:  dbData.UserPoaNm,
				Team:   dbData.Itg_org_nm,
				//Name:   "names",
				//SSO_ID: "ssoids",
				//Grade:  "grades",
				//Team:   "teams",
>>>>>>> parent of 6aa176f (세션으로 id전달)
			}
		}
		js, err := json.Marshal(data)
		if err != nil {
			log.Println(err)
		}
<<<<<<< HEAD
		return c.JSON(string(js))
	})
	app.Use("/files", filesystem.New(filesystem.Config{
		Root: http.Dir("./file"),
	}))
	app.Post("/savelogin", func(c *fiber.Ctx) error {
		data := SsoIdStruct{}
=======
		c.App().Config()
		//log.Println(js)

		c.Append("SSO_ID", string(js))

		//return c.Redirect("http://192.168.0.9:3000")
		return c.Render("index", fiber.Map{
			"SSOID": "testerssoid",
		})
	})

	app.Get("/sso/matchlogin", func(c *fiber.Ctx) error {
		IP := strings.Split(c.Context().RemoteAddr().String(), ":")[0]
		if t, ok := Tokens.Load(IP); ok {
			js, err := json.Marshal(t.(string))
			if err != nil {
				log.Println(err)
			}
			c.App().Config()
			return c.JSON(string(js))
		} else {
			return c.JSON(string(""))
		}
	})

	app.Post("/sso/savelogin", func(c *fiber.Ctx) error {
		data := SaveLogin{}
>>>>>>> parent of 6aa176f (세션으로 id전달)
		log.Println(string(c.Body()))
		err := json.Unmarshal(c.Body(), &data)
		if err != nil {
			log.Println(err)
		}
<<<<<<< HEAD
		sso_id := data.SSO_ID
		//------------ IP Store ------------//
		IP := strings.Split(c.Context().RemoteAddr().String(), ":")[0]
		log.Println(IP, "login and store:", sso_id)
		Tokens.Store(IP, sso_id)
		//------------ IP Store ------------//
		player := db.WebLogin{}
		if r := db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", sso_id).First(&player); r.RowsAffected == 0 {
			player = db.WebLogin{SsoId: sso_id, LastLoginTime: time.Now()}
			db.GetDBManager().DBMS.Create(&player)
=======
		log.Println(data.SsoId)

		IP := strings.Split(c.Context().RemoteAddr().String(), ":")[0]
		log.Println(IP, "login and store:", data.SsoId)
		Tokens.Store(IP, data.SsoId)

		return nil
	})

	app.Get("/security", func(c *fiber.Ctx) error {
		log.Println("Recv : Security")

		dbData := db.SecurityPhrase{}
		err := db.GetDBManager().DBMS.Table("security_phrase").Where("status = ? AND country = ?", "webpage", "korean").Select("phrases").First(&dbData).Error

		SendJson := ""
		if dbData.Phrases == "" || err != nil {
			log.Println(err)
			SendJson = "CannotAceess"
>>>>>>> parent of 6aa176f (세션으로 id전달)
		} else {
			SendJson = dbData.Phrases
		}
<<<<<<< HEAD
		return nil
	})
	app.Get("/requestkey", func(c *fiber.Ctx) error {
		IP := strings.Split(c.Context().RemoteAddr().String(), ":")[0]
		if t, ok := Tokens.Load(IP); ok {
			key := GetUnicRandomKey(6)
			player := db.WebLogin{}
			if r := db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", t.(string)).First(&player); r.RowsAffected == 0 {
				player = db.WebLogin{SsoId: t.(string), KeyValue: key, KeyStoreTime: time.Now()}
				db.GetDBManager().DBMS.Create(&player)
			} else {
				db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", t.(string)).Update("key_value", key)
				updateLogin := map[string]interface{}{
					"key_value":      key,
					"key_store_time": time.Now(),
				}
				db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", t).Updates(&updateLogin)
			}
			log.Println("IP:", IP, "/id:", t.(string), "/key:", key)
		}
		return nil
	})
	app.Get("/getkey", func(c *fiber.Ctx) error {
		sso_id := GetSsoidFromSession(c)
		if sso_id == "" {
			return c.JSON(string("error"))
		}
		data := db.WebLogin{}
		curTime := time.Now()
		err := db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", sso_id).Scan(&data).Error
		if err != nil || data.KeyValue == "" || curTime.Sub(data.LastLoginTime).Minutes() > 60 || curTime.Sub(data.KeyStoreTime).Minutes() > 5 {
			log.Println("getkey err:", err)
			return c.JSON(string("error"))
		} else {
			return c.JSON(string(data.KeyValue))
		}
	})
	app.Use("/launcher", func(c *fiber.Ctx) error {
		c.Context().RemoteAddr()
		return c.Download("./Launcher/Setup.zip", "Setup.zip")
	})
	app.Get("/logout", func(c *fiber.Ctx) error {
		sso_id := GetSsoidFromSession(c)
		db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", sso_id).Delete(&db.WebLogin{})
		log.Println(sso_id, " has logout")
		return nil
	})
	app.Get("/timecontinuation", func(c *fiber.Ctx) error {
		sso_id := GetSsoidFromSession(c)
		db.GetDBManager().DBMS.Table("web_login").Where("sso_id = ?", sso_id).Update("last_login_time", time.Now())
		return nil
=======

		js, err := json.Marshal(SendJson)
		if err != nil {
			log.Println(err)
		}
		c.App().Config()
		return c.JSON(string(js))
>>>>>>> parent of 6aa176f (세션으로 id전달)
	})

	app.Listen(":3000")
}
<<<<<<< HEAD

func GetSecurity() string {
	dbData := db.SecurityPhrase{}
	err := db.GetDBManager().DBMS.Table("security_phrase").Where("status = ? AND country = ?", "webpage", "korean").Select("phrases").First(&dbData).Error
	if dbData.Phrases == "" || err != nil {
		log.Println(err)
		return "Security Phrases Get Fail"
	} else {
		return dbData.Phrases
	}
}
func GetUnicRandomKey(length int) string {
	seed, _ := cryptoRand.Int(cryptoRand.Reader, big.NewInt(math.MaxInt64))
	rand.Seed(seed.Int64())
	b := make([]byte, length)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(b)
}
func GetSsoidFromSession(c *fiber.Ctx) string {
	sess, err := SessionStorage.Get(c)
	if err != nil {
		log.Println("SessionGet Err", err)
	}
	raw := sess.Get("ssoid")
	if raw == nil {
		log.Println("Session not logged in", err)
	}
	sso_id, ok := raw.(string)
	if !ok {
		log.Println("Session Convert Err :", sso_id)
	}
	return sso_id
}
=======
>>>>>>> parent of 6aa176f (세션으로 id전달)
