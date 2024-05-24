package user

import (
	"ecommerce/src/config"
	"ecommerce/src/constants"
	"ecommerce/src/daos"
	"ecommerce/src/database/models"
	"ecommerce/src/dtos"
	"ecommerce/src/dtos/response"
	"ecommerce/src/utils/context"
	"ecommerce/src/utils/token"
	"errors"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	UserDAO      daos.UserDAO
	accessToken  daos.AccessTokenDAO
	refreshToken daos.RefreshTokenDAO 
}

func NewUsers() *Users {
	return &Users{
		UserDAO:      daos.NewUser(),
		accessToken:  daos.NewAccessToken(),
		refreshToken: daos.NewRefreshToken(),
	}
}

func (u *Users) UserFromRegisterReq(req *dtos.RegisterReq) *models.Users {
	return &models.Users{
		ID:        uuid.New().String(),
		Email:     req.Email,
		Password:  req.Password,
		Username:  req.Username,
		RoleID:    "34e9ec0f-aa20-42ab-b70d-4ba456c3bc48",
		CreatedAt: time.Now(),
	}
}

func (u *Users) Register(ctx *context.Context, req *dtos.RegisterReq) error {
	if req.Email == "" && req.Username == "" && req.Password == "" && req.RoleID == "" {
		return errors.New("all fields cannot be empty")
	}
	user := u.UserFromRegisterReq(req)

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		log.Println("Unable to create password hash. Err:", err)
		return err
	}

	if exists, _ := u.UserDAO.CheckEmailExists(ctx, user.Email); exists {
		return constants.ErrEmailTaken
	}
	if exists, _ := u.UserDAO.CheckUsernameExists(ctx, user.Username); exists {
		return constants.ErrUsernameTaken
	}

	user.Password = string(hash)

	return u.UserDAO.Create(ctx, user)
}

func (u *Users) GetUsers(ctx *context.Context) ([]models.Users, error) {
	var users []*models.Users
	err := ctx.DB.Table("users").Find(&users).Error
	if err != nil {
		return nil, err
	}

	var userList []models.Users
	for _, user := range users {
		userList = append(userList, *user)
	}
	return userList, nil
}

func (u *Users) Login(ctx *context.Context, req *dtos.LoginReq) (*response.LoginRes, error) {
    if req.Email == "" && req.Username == "" {
        return nil, errors.New("email or username must be provided")
    }
    if req.Password == "" {
        return nil, errors.New("password cannot be empty")
    }
    if req.RoleID == "" {
        return nil, errors.New("role_id cannot be empty")
    }

    user, err := u.UserDAO.GetAccountFromEmailOrMobile(ctx, req.Email, req.Username, req.RoleID)
    if err != nil {
        log.Println("NO record found. Err: ", err)
        return nil, constants.ErrInvalidCredentials
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
    if err != nil {
        log.Println("Unable to hash the password. Err: ", err)
        return nil, constants.ErrInvalidCredentials
    }

    accessToken, refreshToken := token.GetAccessAndRefreshToken(config.Conf.TokenSize)

    err = u.refreshToken.Create(ctx, &models.RefreshToken{
        Token:     refreshToken,
        UserID:    user.ID,
        ExpiresAt: time.Now().Add(time.Duration(config.Conf.RefreshTokenExpiry) * time.Hour),
    })
    if err != nil {
        log.Println("Unable to create refresh token. Err: ", err)
        return nil, err
    }

    err = u.accessToken.Create(ctx, &models.AccessToken{
        Token:         accessToken,
        RefreshTokens: refreshToken,
        UserID:        user.ID,
        ExpiresAt:     time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
    })
    if err != nil {
        log.Println("Unable to create access token. Err: ", err)
        return nil, err
    }

    return &response.LoginRes{
        Token: accessToken,
    }, nil
}


func (u *Users) GetAccountWithAccessToken(ctx *context.Context, token string) (*models.Users, error) {
	at, err := u.accessToken.Get(ctx, token)
	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return nil, err
	}

	if at.ExpiresAt.Before(time.Now()) {
		return nil, constants.ErrAccessTokenExpired
	}

	user, err := u.UserDAO.Get(ctx, at.UserID)
	if err != nil {
		log.Println("Unable to get users. Err: ", err)
		return nil, err
	}
	return user, nil
}

func (u *Users) GetAccessFromRefreshToken(ctx *context.Context, tkn string) (string, error) {
	rt, err := u.refreshToken.Get(ctx, tkn)
	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return "", err
	}

	if rt.ExpiresAt.Before(time.Now()) {
		return "", constants.ErrAccessTokenExpired
	}

	accessToken, _ := token.GetAccessAndRefreshToken(config.Conf.TokenSize)
	err = u.accessToken.Create(ctx, &models.AccessToken{
		Token:         accessToken,
		UserID:        rt.UserID,
		RefreshTokens: rt.Token,
		ExpiresAt:     time.Now().Add(time.Duration(config.Conf.AccessTokenExpiry) * time.Hour),
	})

	if err != nil {
		log.Println("Unable to get access token. Err: ", err)
		return "", err
	}
	return accessToken, nil
}
