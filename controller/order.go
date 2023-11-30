package controller

import (
	"assignment-two/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
)

func Create(c echo.Context) error {
	var order model.Order
	if err := c.Bind(&order); err != nil {
		return err
	}

	err := DB.Create(&order).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status": "error create data",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "success create data",
	})
}

func Delete(c echo.Context) error {
	orderId := c.Param("id")

	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Where("id = ?", orderId).Delete(&model.Order{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where("order_id = ?", orderId).Delete(&model.Item{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success":  true,
		"messages": "Delete data success",
	})
}

func Get(c echo.Context) error {
	var order model.Order
	orderId := c.Param("id")

	// err := DB.Where("id = ?", orderId).Take(&order).Error
	err := DB.Preload(clause.Associations).Find(&order, "id = ?", orderId).Error

	if err != nil {
		return err
	}

	if order.ID == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"success":  false,
			"messages": "Not found",
		})

	}

	return c.JSON(http.StatusOK, order)
}

func Update(c echo.Context) error {
	return nil
}
