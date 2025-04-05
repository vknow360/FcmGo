package handlers

import (
	"context"
	"net/http"

	types "github.com/vknow360/FcmGo/interfaces"

	"firebase.google.com/go/v4/messaging"
	"github.com/gin-gonic/gin"
	"github.com/vknow360/FcmGo/firebase"
)

func Status(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func SendNotificationToTopic(c *gin.Context) {
	var req types.NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := firebase.App.Messaging(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := &messaging.Message{
		Topic: req.Topic,
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:    req.Title,
				Body:     req.Body,
				Sound:    req.Sound,
				ImageURL: req.Image,
			},
		},
		Data: req.Data,
	}

	resp, err := client.Send(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"messageId": resp})
	}
}

func SendMessageToTopic(c *gin.Context) {
	var req types.NotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := firebase.App.Messaging(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := &messaging.Message{
		Topic: req.Topic,
		Data:  req.Data,
	}

	resp, err := client.Send(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"messageId": resp})
	}
}

func SendNotificationToTokens(c *gin.Context) {
	var req types.TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := firebase.App.Messaging(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := &messaging.MulticastMessage{
		Tokens: req.Tokens,
		Android: &messaging.AndroidConfig{
			Notification: &messaging.AndroidNotification{
				Title:    req.Title,
				Body:     req.Body,
				Sound:    req.Sound,
				ImageURL: req.Image,
			},
		},
		Data: req.Data,
	}

	resp, err := client.SendEachForMulticast(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"successCount": resp.SuccessCount,
			"failureCount": resp.FailureCount,
		})
	}
}

func SendMessageToTokens(c *gin.Context) {
	var req types.TokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	client, err := firebase.App.Messaging(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg := &messaging.MulticastMessage{
		Tokens: req.Tokens,
		Data:   req.Data,
	}

	resp, err := client.SendEachForMulticast(context.Background(), msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"successCount": resp.SuccessCount,
			"failureCount": resp.FailureCount,
		})
	}
}
