package handlers

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/pquerna/otp/totp"

	"task_management_service/config"
	"task_management_service/database"
	"task_management_service/models"
	"task_management_service/types"
	"task_management_service/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"github.com/golang-jwt/jwt/v5"
)

var issuer2Fa = "TaskManagementService"

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}
	return string(bytes), err
}

func SignUp(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusBadRequest, "NOK", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", err.Error())
		return response
	}

	// ตรวจสอบว่ามีการลงทะเบียนของ uername อยู่แล้วหรือไม่
	db := database.DBConn
	var userCount int64
	_ = db.First(&user, "username = ?", user.Username).Count(&userCount).Error
	if userCount > 0 {
		log.Print("มีชื่อบัญชีผู้ใช้งานนี้เเล้ว")
		response := utils.ResponseError(c, fiber.StatusBadRequest, "NOK", "มีชื่อบัญชีผู้ใช้งานนี้เเล้ว")
		return response
	}

	user.Password, _ = hashPassword(user.Password)
	user.Status = true
	err = db.Create(&user).Error
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", err.Error())
		return response
	}

	user.Password = ""
	response := utils.ResponseSuccess(c, user, fiber.StatusOK, "OK", "User was signed up successfully!")
	return response
}

func SignIn(c *fiber.Ctx) error {
	user := models.User{}
	err := c.BodyParser(&user)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusBadRequest, "NOK", err.Error())
		return response
	}

	// Validate the request body
	if user.Username == "" || user.Password == "" {
		log.Print(err)
		log.Print("Username or Password is empty.")
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", "Username or Password is empty.")
		return response
	}

	passwordInput := user.Password
	user.Password, _ = hashPassword(user.Password)
	db := database.DBConn
	err = db.First(&user, "username = ?", user.Username).Error
	if err != nil {
		log.Print(err)
		log.Print("User Not found.")
		response := utils.ResponseError(c, fiber.StatusNotFound, "NOK", "User Not found.")
		return response
	}

	// ตรวจสอบสถานะของผู้ใช้
	if !user.Status {
		log.Print("Account is locked due to multiple invalid sign in attempts")
		response := utils.ResponseError(c, fiber.StatusForbidden, "NOK", "Account is locked.")
		return response
	}

	if !checkPasswordHash(passwordInput, user.Password) {
		log.Print("Invalid Password")

		// เพิ่มจำนวน InvalidSignInCount
		user.InvalidSignInCount++
		// ถ้า sign in ไม่ถูกต้อง 5 ครั้ง ให้ lock account
		if user.InvalidSignInCount >= 5 {
			user.Status = false

			db.Save(&user)
			log.Print("Account is locked due to multiple invalid sign in attempts")
			response := utils.ResponseError(c, fiber.StatusForbidden, "NOK", "Account is locked.")
			return response
		}

		db.Save(&user)

		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", "Invalid Password!")
		return response
	}

	// Create the Claims
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"admin":    false,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	var secretKey = config.Config("SECRET_KEY")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
		return response
	}

	//	create result
	result := fiber.Map{
		"token":    t,
		"username": user.Username,
		"user_id":  user.ID,
	}

	// รีเซ็ต InvalidSignInCount หลังจากเข้าสู่ระบบสำเร็จ
	user.InvalidSignInCount = 0
	db.Save(&user)

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "User was signed in successfully!")
	return response

}

func SignUp2Fa(c *fiber.Ctx) error {
	user := models.User{}

	err := c.BodyParser(&user)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusBadRequest, "NOK", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(response)
	}

	validate := validator.New()
	err = validate.Struct(user)
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", err.Error())
		return response
	}

	// ตรวจสอบว่ามีการลงทะเบียนของ uername อยู่แล้วหรือไม่
	db := database.DBConn
	var userCount int64
	_ = db.First(&user, "username = ?", user.Username).Count(&userCount).Error

	if userCount > 0 {
		log.Print("มีชื่อบัญชีผู้ใช้งานนี้เเล้ว")
		response := utils.ResponseError(c, fiber.StatusBadRequest, "NOK", "มีชื่อบัญชีผู้ใช้งานนี้เเล้ว")
		return response
	}

	user.Password, _ = hashPassword(user.Password)
	user.IsTwoFactorEnabled = true
	err = db.Create(&user).Error
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", err.Error())
		return response
	}

	// สร้าง Secret Key สำหรับ TOTP
	secret, err := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer2Fa,
		AccountName: user.Username,
	})

	if err != nil {
		log.Print(err)
		_ = db.Delete(&user).Where("username = ?", user.Username)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", "Failed to generate 2FA secret.")
		return response
	}

	// เก็บ Secret Key ไว้ในฟิลด์ TwoFactorToken และบันทึกในฐานข้อมูล
	user.TwoFactorToken = secret.Secret()
	user.Status = true
	log.Println(user)
	err = db.Model(&user).Where("username = ?", user.Username).Updates(&user).Error
	if err != nil {
		log.Print(err)
		log.Println("Failed to enable 2FA.")
		_ = db.Delete(&user).Where("username = ?", user.Username)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", "Failed to enable 2FA.")
		return response
	}

	// ส่ง QR Code ให้ผู้ใช้สแกน
	result := fiber.Map{
		"id":            user.ID,
		"user_id":       user.ID,
		"username":      user.Username,
		"full_name":     user.FullName,
		"profile_image": user.ProfileImage,
		"qr_url":        secret.URL(), // URL สำหรับให้สแกน QR ใน Google Authenticator
	}

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "User was signed up successfully!")
	return response
}

func SignIn2Fa(c *fiber.Ctx) error {
	user := models.User{}
	inputBody := types.AuthSignIn{}
	err := c.BodyParser(&user)

	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
		return response
	}

	if user.Username == "" || user.Password == "" {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", "Username or Password is empty.")
		return response
	}

	// ตรวจสอบว่าชื่อผู้ใช้และรหัสผ่านถูกต้องหรือไม่
	passwordInput := user.Password
	db := database.DBConn
	err = db.First(&user, "username = ?", user.Username).Error
	if err != nil {
		log.Print(err)
		log.Print("User Not found.")
		response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", "User Not found.")
		return response
	}

	// ตรวจสอบสถานะของผู้ใช้
	if !user.IsTwoFactorEnabled {
		log.Print("ยันไม่ยืนยันตัวตน 2FA")
		response := utils.ResponseError(c, fiber.StatusForbidden, "NOK", "ยันไม่ยืนยันตัวตน 2FA")
		return response
	}

	// ตรวจสอบสถานะของผู้ใช้
	if !user.Status {
		log.Print("Account is locked due to multiple invalid sign in attempts")
		response := utils.ResponseError(c, fiber.StatusForbidden, "NOK", "Account is locked.")
		return response
	}

	if !checkPasswordHash(passwordInput, user.Password) {
		log.Print(err)
		log.Print("Invalid Password")

		// เพิ่มจำนวน InvalidSignInCount
		user.InvalidSignInCount++
		// ถ้า sign in ไม่ถูกต้อง 5 ครั้ง ให้ lock account
		if user.InvalidSignInCount >= 5 {
			user.Status = false

			db.Save(&user)
			log.Print("Account is locked due to multiple invalid sign in attempts")
			response := utils.ResponseError(c, fiber.StatusForbidden, "NOK", "Account is locked.")
			return response
		}

		db.Save(&user)
		response := utils.ResponseError(c, fiber.StatusUnauthorized, "NOK", "Invalid Password!")
		return response
	}

	// ตรวจสอบว่าผู้ใช้เปิดใช้งาน 2FA หรือไม่
	if user.IsTwoFactorEnabled {
		// ดึงรหัส OTP จาก request
		// otpCode := c.FormValue("otp_code")
		_ = c.BodyParser(&inputBody)
		otpCode := inputBody.OtpCode
		if otpCode == "" {
			log.Print(err)
			response := utils.ResponseError(c, fiber.StatusUnprocessableEntity, "NOK", "OTP code is required.")
			return response
		}

		// ตรวจสอบ OTP
		valid := totp.Validate(otpCode, user.TwoFactorToken)
		if !valid {
			log.Print(err)
			response := utils.ResponseError(c, fiber.StatusUnauthorized, "NOK", "Invalid OTP code.")
			return response
		}
	}

	// สร้างโทเค็น JWT
	claims := jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"admin":    false,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secretKey := config.Config("SECRET_KEY")
	t, err := token.SignedString([]byte(secretKey))
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", err.Error())
		return response
	}

	// ส่งตอบกลับให้ผู้ใช้
	result := fiber.Map{
		"token":    t,
		"username": user.Username,
		"user_id":  user.ID,
	}

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "Sign in successfully.")

	// รีเซ็ต InvalidSignInCount หลังจากเข้าสู่ระบบสำเร็จ
	user.InvalidSignInCount = 0
	db.Save(&user)

	return response
}

// EnableTwoFactor เปิดใช้งาน 2FA สำหรับผู้ใช้
func EnableTwoFactor(c *fiber.Ctx) error {
	username := c.Params("username")
	db := database.DBConn

	user := models.User{}
	err := db.First(&user, "username = ? AND status = true", username).Error
	if err != nil {
		log.Print(err)
		response := utils.ResponseError(c, fiber.StatusNotFound, "NOK", "User not found.")
		return response
	}

	var qrURL string

	// ตรวจสอบว่ามี Secret Key อยู่แล้วหรือไม่
	if user.TwoFactorToken != "" && user.IsTwoFactorEnabled {
		// ถ้ามี Secret Key อยู่แล้ว สร้าง URL สำหรับ QR Code โดยตรง
		otp := &url.URL{
			Scheme: "otpauth",
			Host:   "totp",
			Path:   fmt.Sprintf("/%s:%s", issuer2Fa, user.Username),
		}

		q := otp.Query()
		q.Set("secret", user.TwoFactorToken)
		q.Set("issuer", issuer2Fa)
		otp.RawQuery = q.Encode()

		qrURL = otp.String()
	} else {
		// ถ้ายังไม่มี Secret Key ให้สร้างใหม่
		secret, err := totp.Generate(totp.GenerateOpts{
			Issuer:      issuer2Fa,
			AccountName: user.Username,
		})
		if err != nil {
			log.Print(err)
			response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", "Failed to generate new 2FA secret.")
			return response
		}

		// บันทึก Secret Key ใหม่
		user.TwoFactorToken = secret.Secret()
		qrURL = secret.URL()
	}

	// เปิดใช้งาน 2FA
	user.IsTwoFactorEnabled = true
	if err := db.Save(&user).Error; err != nil {
		response := utils.ResponseError(c, fiber.StatusInternalServerError, "NOK", "Failed to enable 2FA.")
		return response
	}

	result := fiber.Map{
		"user_id":  user.ID,
		"username": user.Username,
		"message":  "2FA enabled. Scan this QR code with your authenticator app.",
		"qr_url":   qrURL,
	}

	response := utils.ResponseSuccess(c, result, fiber.StatusOK, "OK", "2FA enabled successfully.")
	return response
}

func DecodeUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userMap := fiber.Map{"user": fiber.Map{
		"user_id":  claims["id"],
		"username": claims["username"],
	}}

	return c.JSON(userMap)

}
