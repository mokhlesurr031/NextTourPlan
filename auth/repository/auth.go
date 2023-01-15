package repository

import (
	"context"
	"github.com/NextTourPlan/domain"
	"github.com/NextTourPlan/domain/dto"
	"github.com/NextTourPlan/internal/config"
	"github.com/NextTourPlan/utils"
	"gorm.io/gorm"
	"log"
)

func New(db *gorm.DB) domain.AuthRepository {
	return &AuthSqlStorage{
		db: db,
	}
}

type AuthSqlStorage struct {
	db *gorm.DB
}

func (a *AuthSqlStorage) PostSignUP(ctx context.Context, ctr *domain.SignUpInput) string {
	db := a.db
	if ctr.Password != ctr.PasswordConfirm {
		log.Println("Password Doesnt Match")
	}
	hashedPassword, err := utils.HashPassword(ctr.Password)
	if err != nil {
		log.Println("Password hashing failed")
	}
	ctr.Password = hashedPassword
	ctr.PasswordConfirm = hashedPassword

	user := domain.SignUpInput{}

	if ctr.Email != "" && ctr.Contact != "" {
		//Check if email already exists
		mail := db.First(&user, "email=?", ctr.Email)
		cred := &domain.SignUpInput{}
		if err := mail.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email != "" {
			return "email already exists"
		}

		//Check if contact already exists
		contact := db.First(&user, "contact=?", ctr.Contact)
		if err := contact.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email != "" {
			return "contact already exists"
		}
	}

	db.Create(ctr)

	return "success"
}

func (a *AuthSqlStorage) PostSignIn(ctx context.Context, ctr *domain.SignInInput) (*dto.JWTToken, error) {
	qry := a.db
	jwt := config.JWT()
	user := domain.SignUpInput{}

	if ctr.Email != "" && ctr.Password != "" {
		qry := qry.Find(&user, "email=?", ctr.Email)
		cred := &domain.SignUpInput{}
		if err := qry.WithContext(ctx).Take(cred).Error; err != nil {
			log.Println(err)
		}
		if cred.Email == "" {
			reqJwt := &dto.JWTToken{Message: "invalid data"}
			return reqJwt, nil
		}

		if err := utils.VerifyPassword(cred.Password, ctr.Password); err != nil {
			reqJwt := &dto.JWTToken{Message: "invalid password"}
			return reqJwt, err
		}

		token, err := utils.GenerateToken(jwt.ExpiredIn, cred.ID, jwt.Secret)
		if err != nil {
			log.Println(err)
		}

		loggedInData := &dto.LoggerInUserData{}
		loggedInData.FullName = user.FullName
		loggedInData.Email = user.Email
		loggedInData.Address = user.Address
		loggedInData.Contact = user.Contact
		loggedInData.Rating = user.Rating
		loggedInData.ID = user.ID

		reqJwt := &dto.JWTToken{User: loggedInData, Secret: token, MaxAge: jwt.MaxAge, ExpiredIn: jwt.ExpiredIn, Message: "success"}
		return reqJwt, nil
	}

	reqJwt := &dto.JWTToken{Message: "invalid data"}
	return reqJwt, nil

}
