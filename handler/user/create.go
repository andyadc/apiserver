package user

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"fmt"
	"github.com/gin-gonic/gin"
	logger "github.com/lexkong/log"
)

// Create creates a new user account.
func Create(c *gin.Context) {
	var req CreateRequest
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	username := c.Param("username")
	logger.Infof("URL username: %s", username)

	desc := c.Query("desc")
	logger.Infof("URL key param desc: %s", desc)

	contentType := c.GetHeader("Content-Type")
	logger.Infof("Header Content-Type: %s", contentType)

	logger.Debugf("username is: [%s], password is [%s]", req.Username, req.Password)
	if req.Username == "" {
		SendResponse(c, errno.New(errno.ErrUserNotFound, fmt.Errorf("username can not found in db: xx.xx.xx.xx")), nil)
		return
	}

	if req.Password == "" {
		SendResponse(c, fmt.Errorf("password is empty"), nil)
	}

	resp := CreateResponse{
		Username: req.Username,
	}

	// Show the user information.
	SendResponse(c, nil, resp)
}
