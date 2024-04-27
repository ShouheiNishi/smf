/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package sbi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getEventExposureRoutes() []Route {
	return []Route{
		{
			Method:  http.MethodGet,
			Pattern: "/",
			APIFunc: func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{"status": "Service Available"})
			},
		},
		{
			Method:  http.MethodPost,
			Pattern: "subscriptions",
			APIFunc: s.SubscriptionsPost,
		},
		{
			Method:  http.MethodDelete,
			Pattern: "subscriptions/:subId",
			APIFunc: s.SubscriptionsSubIdDelete,
		},
		{
			Method:  http.MethodGet,
			Pattern: "subscriptions/:subId",
			APIFunc: s.SubscriptionsSubIdGet,
		},
		{
			Method:  http.MethodPut,
			Pattern: "subscriptions/:subId",
			APIFunc: s.SubscriptionsSubIdPut,
		},
	}
}

// SubscriptionsPost -
func (s *Server) SubscriptionsPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdDelete -
func (s *Server) SubscriptionsSubIdDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdGet -
func (s *Server) SubscriptionsSubIdGet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdPut -
func (s *Server) SubscriptionsSubIdPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}
