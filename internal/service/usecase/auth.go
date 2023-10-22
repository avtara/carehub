package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/avtara/carehub/internal/models"
	"github.com/avtara/carehub/internal/service"
	"github.com/avtara/carehub/utils"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"io"
	"net/http"
)

type authenticationUsecase struct {
	userRepository  service.UserRepository
	asyncRepository service.BrokerRepository
	oauthConfig     *oauth2.Config
	db              *sqlx.DB
}

func NewAuthenticationUseCase(userRepository service.UserRepository, asyncRepository service.BrokerRepository, oauthConfig *oauth2.Config, db *sqlx.DB) service.AuthenticationUseCase {
	return &authenticationUsecase{
		userRepository:  userRepository,
		asyncRepository: asyncRepository,
		oauthConfig:     oauthConfig,
		db:              db,
	}
}

func (u *authenticationUsecase) AuthenticationGoogle(ctx context.Context) (redirectURL string, err error) {
	redirectURL = u.oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return
}

func (u *authenticationUsecase) GoogleCallback(ctx context.Context, code string) (response models.AuthenticationResponse, err error) {
	var (
		tx *sqlx.Tx
	)
	tx, err = u.db.Beginx()

	defer func() {
		if tx == nil {
			return
		}

		if err != nil {
			serr := tx.Rollback()
			if serr != nil {
				log.Errorf("[Usecase][GoogleCallback] failed to rollback: %s", err.Error())
			}
		}
	}()

	token, err := u.oauthConfig.Exchange(ctx, code)
	if err != nil {
		err = fmt.Errorf("[Usecase][GoogleCallback] while code exchange failed: %s", err.Error())
		return
	}

	var userInfoResponse *http.Response
	userInfoResponse, err = http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		err = fmt.Errorf("[Usecase][GoogleCallback] failed getting user info: %s", err.Error())
		return
	}
	defer userInfoResponse.Body.Close()

	userInfoByte, err := io.ReadAll(userInfoResponse.Body)
	if err != nil {
		err = fmt.Errorf("[Usecase][GoogleCallback] failed reading response body: %s", err.Error())
		return
	}

	var userInfo models.UserInfoGoogle
	err = json.Unmarshal(userInfoByte, &userInfo)
	if err != nil {
		err = fmt.Errorf("[Usecase][GoogleCallback] failed while unmarshal user info: %s", err.Error())
		return
	}

	var user models.User
	user, err = u.userRepository.GetByEmail(ctx, userInfo.Email)
	if err != nil && err != models.ErrorUserNotFound {
		err = fmt.Errorf("[Usecase][GoogleCallback] failed get user by email: %s", err.Error())
		return
	}

	if err == models.ErrorUserNotFound {
		password := utils.GenerateRandomString(20)

		var hashedPassword []byte
		hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			err = fmt.Errorf("[Usecase][GoogleCallback] failed while hash password: %s", err.Error())
			return
		}

		user = models.User{
			Name:     userInfo.Name,
			Password: string(hashedPassword),
			Email:    userInfo.Email,
			Photo:    userInfo.Picture,
			Role:     models.RoleTypeGuest,
		}
		user.ID, err = u.userRepository.Insert(ctx, tx, user)
		if err != nil {
			err = fmt.Errorf("[Usecase][GoogleCallback] failed while insert user: %s", err.Error())
			return
		}

		var encText string
		encText, err = utils.Encrypt(context.Background(), password, utils.GetEnv("encrypt.secret_key", "!@#SecretBgfast"))
		if err != nil {
			err = fmt.Errorf("[Usecase][GoogleCallback] failed while encrypt data user: %s", err.Error())
			return
		}

		user.Password = encText

		err = tx.Commit()
		if err != nil {
			err = fmt.Errorf("[Usecase][GoogleCallback] failed to commit: %s", err.Error())
			return
		}

		go func(user models.User) {
			_, err = u.asyncRepository.Publish(context.Background(), models.TaskSendEmailNewUser, user)
			if err != nil {
				log.Warnf("[Usecase][GoogleCallback] failed while publish task: %s", err.Error())
			}
		}(user)
	}

	response.AccessToken, err = utils.GenerateAccessToken(ctx, user.ID, user.Role)
	if err != nil {
		err = fmt.Errorf("[Usecase][GoogleCallback] failed generate access token: %s", err.Error())
		return
	}

	return
}

func (u *authenticationUsecase) Login(ctx context.Context, args models.AuthenticationParams) (response models.AuthenticationResponse, err error) {
	var user models.User
	user, err = u.userRepository.GetByEmail(ctx, args.Email)
	if err != nil && err == models.ErrorUserNotFound {
		return
	}

	if err != nil {
		err = fmt.Errorf("[Usecase][Login] failed get user by email: %s", err.Error())
		return
	}

	err = utils.VerifyPassword(ctx, args.Password, user.Password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = models.ErrorUserWrongPassword
		return
	}

	if err != nil {
		err = fmt.Errorf("[Usecase][Login] failed while verify password: %s", err.Error())
		return
	}

	response.AccessToken, err = utils.GenerateAccessToken(ctx, user.ID, user.Role)
	if err != nil {
		err = fmt.Errorf("[Usecase][Login] failed generate access token: %s", err.Error())
		return
	}

	return
}
