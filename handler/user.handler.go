package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/model/entity"
	"go-fiber-gorm/model/request"
	"go-fiber-gorm/utils"
	"log"
)

func UserHandlerGetAll(ctx *fiber.Ctx) error {
	var users []entity.User

	//userInfo := ctx.Locals("userInfo").(jwt.MapClaims)
	//log.Println("email :: ", userInfo["email"])

	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}
	//err := database.DB.Find(&users).Error
	//if err != nil {
	//	log.Println(err)
	//}
	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)
	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	// VALIDASI REQUEST
	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashingPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed to store data",
		})
	}

	return ctx.JSON(fiber.Map{
		"messaage": "success",
		"data":     newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")

	var user entity.User
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//userResponse := response.UserResponse{
	//	ID:        user.ID,
	//	Name:      user.Name,
	//	Email:     user.Email,
	//	Address:   user.Address,
	//	Phone:     user.Phone,
	//	CreatedAt: user.CreatedAt,
	//	UpdatedAt: user.UpdatedAt,
	//}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}

	var user entity.User

	userId := ctx.Params("id")
	// CHECK AVAILABLE USER
	err := database.DB.First(&user, "id = ?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// UPDATE USER DATA
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone

	errUpdate := database.DB.Save(&user).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	//	CHECK AVAILABLE USER
	err := database.DB.Debug().First(&user, "id=?", userId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error
	if errDelete != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user was deleted",
	})
}
