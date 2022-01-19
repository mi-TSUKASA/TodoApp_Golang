package main

import "models"

func main() {
	/*
		fmt.Println(config.Config.Port)
		fmt.Println(config.Config.SQLDriver)
		fmt.Println(config.Config.DbName)
		fmt.Println(config.Config.LogFile)

		log.Println("test")
	*/
	// fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "test"
	// u.Email = "test@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@example.com"
	// u.PassWord = "test3"
	// fmt.Println(u)
	// u.CreateUser()

	// user, _ := models.GetUser(3)
	// user.CreateTodo("Forth Todo")

	// t, _ := user.GetTodosByUser()
	// for _, v := range t {
	// 	fmt.Println(v)
	// }

	// t, _ := models.gettodo(1)
	// t.content = "update todo"
	// t.UpdateTodo()
	t, _ := models.GetTodo(3)
	t.DeleteTodo()
}
