package controllers

import "github.com/gin-gonic/gin"

import "github.com/redmeros/htrade/config"

import "errors"

import "github.com/jinzhu/gorm"

// GetConfig wyluskuje config z contextu
func GetConfig(c *gin.Context) (*config.Config, error) {
	val, exists := c.Get("config")
	if exists != true {
		return nil, errors.New("Config not injected to context")
	}
	cfg, ok := val.(*config.Config)
	if ok != true {
		return nil, errors.New("config key should contain config")
	}
	return cfg, nil
}

// GetDB wylusuje baze danych z contextu
func GetDB(c *gin.Context) (*gorm.DB, error) {
	val, exists := c.Get("DB")
	if exists != true {
		return nil, errors.New("Database not injected")
	}
	db, ok := val.(*gorm.DB)
	if ok != true {
		return nil, errors.New("DB key should contain db")
	}
	return db, nil
}
