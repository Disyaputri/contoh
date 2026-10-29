package handlers

import (
	"bufio"
	"clean-arch/internal/domain"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type userHandler struct {
	uc domain.UserUsecase
}

func NewUserHandler(uc domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		uc: uc,
	}
}

func (h *userHandler) Run() {
	for {
		fmt.Println("\n===== USER MENU =====")
		fmt.Println("1. Create User")
		fmt.Println("2. List Users")
		fmt.Println("3. Get User By ID")
		fmt.Println("4. Update User")
		fmt.Println("5. Delete User")
		fmt.Println("0. Exit")

		fmt.Print("Choose: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			h.create()
		case 2:
			h.list()
		case 3:
			h.detail()
		case 4:
			h.update()
		case 5:
			h.delete()
		case 0:
			fmt.Println("Good Bye!")
			return
		default:
			fmt.Println("Invalid Menu")
		}
	}
}

func (h *userHandler) create() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Name  : ")
	name, _ := reader.ReadString('\n')

	fmt.Print("Email : ")
	email, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	err := h.uc.Create(name, email)
	if err != nil {
		fmt.Println("Error :", err)
		return
	}

	fmt.Println("User created successfully.")
}

func (h *userHandler) list() {

	users, err := h.uc.GetAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()

	for _, user := range users {
		fmt.Printf("%d | %s | %s\n",
			user.ID,
			user.Name,
			user.Email,
		)
	}
}

func (h *userHandler) detail() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("User ID : ")

	text, _ := reader.ReadString('\n')

	id, _ := strconv.Atoi(strings.TrimSpace(text))

	user, err := h.uc.GetByID(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println()
	fmt.Println("ID    :", user.ID)
	fmt.Println("Name  :", user.Name)
	fmt.Println("Email :", user.Email)
}

func (h *userHandler) update() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("ID : ")
	idText, _ := reader.ReadString('\n')

	id, _ := strconv.Atoi(strings.TrimSpace(idText))

	fmt.Print("New Name : ")
	name, _ := reader.ReadString('\n')

	fmt.Print("New Email : ")
	email, _ := reader.ReadString('\n')

	name = strings.TrimSpace(name)
	email = strings.TrimSpace(email)

	err := h.uc.Update(id, name, email)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User updated successfully.")
}

func (h *userHandler) delete() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("User ID : ")

	text, _ := reader.ReadString('\n')

	id, _ := strconv.Atoi(strings.TrimSpace(text))

	err := h.uc.Delete(id)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User deleted successfully.")
}