package Models

import (
	"fmt"
	"gin/config"

	_ "github.com/go-sql-driver/mysql"
)

func getTodos(todo *[]TODO) (err error) {
	if err = config.DB.Find(todo).Error; err != nil {
		return err
	}
	return nil
}
func CreateTodos(todo *[]TODO) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
func Deletetodo(todo *[]TODO, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(todo)
	return nil
}

func Update(todo *[]TODO, id string) (err error) {
	fmt.Println(todo)
	config.DB.Save(todo)
	return nil
}

// func gettodo(todo *[]TODO, id string) (err error) {
// 	if err = config.DB.Where("id = ?", id).First(todo).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
