package static

import (
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/solderapp/solder/router/middleware/context"
	"github.com/solderapp/solder/static"
)

func SetStatic() gin.HandlerFunc {
	return func(c *gin.Context) {
		path := strings.TrimPrefix(
			c.Request.URL.Path,
			context.Config(c).Server.Root,
		)

		file, err := static.FSString(
			false,
			"/"+path,
		)

		logrus.Debugf(path)
		logrus.Debugf(file)
		logrus.Debugf("%q", err)

		if err != nil {
			c.Next()
			return
		}

		logrus.Debugf(file)
	}
}
