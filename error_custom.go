package main

import "fmt"

type validationError struct {
	Message string
}

func (err *validationError) Error() string {
	return err.Message
}

type notFoundError struct {
	Message string
}

func (err *notFoundError) Error() string {
	return err.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"id is empty"}
	}

	if id != "angga" {
		return &notFoundError{"data not found"}
	}

	return nil
}

func main() {
	err := SaveData("asd", nil)
	// if else
	if err != nil {
		if validateErr, ok := err.(*validationError); ok {
			fmt.Println("Validation error:", validateErr.Error())
		} else if notFoundE, ok := err.(*notFoundError); ok {
			fmt.Println("Not found error:", notFoundE.Error(), ok)
		} else {
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Data saved successfully")
	}

	// switch
	if err != nil {
		switch finalErr := err.(type) {
		case *notFoundError:
			fmt.Println("Not found error:", finalErr.Error())
		case *validationError:
			fmt.Println("Validation error:", finalErr.Error())
		default:
			fmt.Println("Unknown error:", err)
		}
	}
}
