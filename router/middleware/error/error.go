package error

import (
	"github.com/Sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

// SetError initializes the error middleware.
func SetError() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// At this point, all the handlers finished. Let's read the errors,
		// in this example we only will use the **last error typed as public**
		// but you could iterate over all them since c.Errors is a slice!
		err := c.Errors.ByType(
			gin.ErrorTypePublic,
		).Last()

		if err != nil {
			status, ok := err.Meta.(int)

			if !ok {
				status = 500
			}

			logrus.Warn(err.Error())

			c.IndentedJSON(
				status,
				gin.H{
					"status":  status,
					"message": err.Error(),
				},
			)
		}
	}
}
