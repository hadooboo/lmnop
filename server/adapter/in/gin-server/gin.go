package gin_server

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"jaehonam.com/lmnop/application/port/in"
	"jaehonam.com/lmnop/domain"
)

type GinAPIServer struct {
	server  *gin.Engine
	service in.Query
	port    int
}

func NewGinAPIServer(service in.Query, port int) *GinAPIServer {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	server := &GinAPIServer{
		server:  r,
		service: service,
		port:    port,
	}
	server.initRouter()

	return server
}

func (r *GinAPIServer) initRouter() {
	v1 := r.server.Group("/api/v1")
	{
		v1.GET("/users/:handle", r.getUser)
		v1.GET("/optimum-problem", r.getOptimumProblem)
	}
}

func (r *GinAPIServer) Serve() error {
	return r.server.Run(fmt.Sprintf(":%v", r.port))
}

func (r *GinAPIServer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.server.ServeHTTP(w, req)
}

func (r *GinAPIServer) getUser(c *gin.Context) {
	user, err := r.service.GetUser(c.Param("handle"))
	if err != nil {
		switch err {
		case domain.ErrNotFoundUser:
			c.String(http.StatusNotFound, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, &getUserResp{user: user})
}

func (r *GinAPIServer) getOptimumProblem(c *gin.Context) {
	handle := c.Query("handle")
	if len(handle) == 0 {
		c.String(http.StatusBadRequest, "invalid handle")
		return
	}

	if len(c.Query("level")) == 0 {
		c.String(http.StatusBadRequest, "invalid level")
		return
	}
	level, err := strconv.ParseInt(c.Query("level"), 10, 32)
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("invalid level: %v", err))
		return
	}

	except := make([]int64, 0)
	if len(c.Query("except")) > 0 {
		for _, v := range strings.Split(c.Query("except"), ",") {
			problemID, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("invalid problem id in except: %v", err))
				return
			}
			except = append(except, problemID)
		}
	}

	problem, err := r.service.GetOptimumProblem(handle, domain.Level(level), except)
	if err != nil {
		switch err {
		case domain.ErrNotFoundProblem:
			c.String(http.StatusNotFound, err.Error())
		default:
			c.String(http.StatusInternalServerError, err.Error())
		}
		return
	}

	c.JSON(http.StatusOK, &getOptimumProblemResp{problem: problem})
}