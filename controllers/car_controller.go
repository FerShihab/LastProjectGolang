package controllers

import (
	"lastProject/configs"
	"lastProject/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddCar(c echo.Context) error {
	var addCar models.Car
	c.Bind(&addCar)

	// masukkan ke database
	result := configs.DB.Create(&addCar)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data cars",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    addCar,
	})

}

func GetDetailCarController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var car models.Car = models.Car{}
	result := configs.DB.First(&car, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Data cars not found",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    car,
	})
}

func DeleteDetailCarController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var car models.Car = models.Car{}
	result := configs.DB.Delete(&car, id)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed delete data cars",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil menghapus data",
	})
}

func GetCarsController(c echo.Context) error {

	var cars []models.Car

	result := configs.DB.Find(&cars)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, models.BaseResponse{
			Status:  false,
			Message: "Failed insert data cars",
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.BaseResponse{
		Status:  true,
		Message: "Berhasil",
		Data:    cars,
	})
}
