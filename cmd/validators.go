package main

import (
	"slices"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"github.com/jangidRkt08/pizza-tracker-go/internal/models"
)


func RegistorCustomValidators(){
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok{
		v.RegisterValidation("valid_pizza_type", createSliceValidator(models.PizzaTypes))
				v.RegisterValidation("valid_pizza_size", createSliceValidator(models.PizzaSizes))
		
	}
}

func createSliceValidator(allowedValues []string) validator.Func {
	return func(fl validator.FieldLevel) bool {
 // Validate/Check if the value(fl.Field().String()) is in the allowedValues(pizzaTypes or pizzaSizes)
 		return slices.Contains(allowedValues,fl.Field().String()) 
	}
}