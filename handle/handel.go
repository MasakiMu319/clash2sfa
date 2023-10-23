package handle

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"log/slog"

	"github.com/xmdhs/clash2sfa/db"
	"github.com/xmdhs/clash2sfa/model"
	"github.com/xmdhs/clash2sfa/service"
)

func PutArg(db db.DB, l *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cxt := r.Context()

		arg := model.ConvertArg{}
		err := json.NewDecoder(r.Body).Decode(&arg)
		if err != nil {
			l.DebugContext(cxt, err.Error())
			http.Error(w, err.Error(), 400)
			return
		}
		if arg.Sub == "" {
			l.DebugContext(cxt, "订阅链接不得为空")
			http.Error(w, "订阅链接不得为空", 400)
			return
		}
		s, err := service.PutArg(cxt, arg, db)
		if err != nil {
			l.WarnContext(cxt, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write([]byte(s))
	}
}

func Frontend(frontendByte []byte, age int) http.HandlerFunc {
	sage := strconv.Itoa(age)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age="+sage)
		w.Write(frontendByte)
	}
}

func Sub(c *http.Client, db db.DB, frontendByte []byte, l *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		id := r.FormValue("id")
		config := r.FormValue("config")
		curl := r.FormValue("configurl")
		sub := r.FormValue("sub")
		include := r.FormValue("include")
		exclude := r.FormValue("exclude")
		urltest := r.FormValue("urltest")
		addTag := r.FormValue("addTag")
		addTagb := false

		if id == "" && sub == "" {
			l.DebugContext(ctx, "id 不得为空")
			http.Error(w, "id 不得为空", 400)
			return
		}
		if addTag == "true" {
			addTagb = true
		}

		rc := http.NewResponseController(w)
		rc.SetWriteDeadline(time.Now().Add(1 * time.Minute))

		b, err := func() ([]byte, error) {
			if sub != "" {
				if config != "" {
					b, err := zlibDecode(config)
					if err != nil {
						return nil, err
					}
					config = string(b)
				}

				a := model.ConvertArg{
					Sub:       sub,
					Include:   include,
					Exclude:   exclude,
					Config:    config,
					ConfigUrl: curl,
					AddTag:    addTagb,
					UrlTest: []model.UrlTestArg{
						{
							Tag:     "HK",
							Include: "HK|HongKong|🇭🇰|香港",
							Type:    "selector",
						},
						{
							Tag:     "TW",
							Include: "TW|Taiwan|🇹🇼|台湾",
							Type:    "selector",
						},
						{
							Tag:     "JP",
							Include: "JP|Japan|🇯🇵|日本",
							Type:    "selector",
						},
						{
							Tag:     "SG",
							Include: "SG|Singapore|🇸🇬|新加坡",
							Type:    "selector",
						},
						{
							Tag:     "US",
							Include: "US|United States|🇺🇸|美国",
							Type:    "selector",
						},
					},
				}
				if urltest != "" {
					b, err := zlibDecode(urltest)
					if err != nil {
						return nil, err
					}
					var u []model.UrlTestArg
					err = json.Unmarshal(b, &u)
					if err != nil {
						return nil, err
					}
					a.UrlTest = append(a.UrlTest, u...)
				}
				return service.MakeConfig(ctx, c, frontendByte, l, a)
			}
			return service.GetSub(ctx, c, db, id, frontendByte, l)
		}()
		if err != nil {
			l.WarnContext(ctx, err.Error())
			http.Error(w, err.Error(), 500)
			return
		}
		w.Write(b)
	}
}

func zlibDecode(s string) ([]byte, error) {
	b, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return nil, err
	}
	r, err := zlib.NewReader(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	b, err = io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}
