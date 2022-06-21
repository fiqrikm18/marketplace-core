package repositories

import (
	"fmt"
	"github.com/cip8/autoname"
	"github.com/fiqrikm18/markerplace_core/pkg/config"
	const_ "github.com/fiqrikm18/markerplace_core/pkg/const"
	"github.com/fiqrikm18/markerplace_core/pkg/domain"
	"github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"os"
	"strings"
	"testing"
	"time"
)

var (
	db *gorm.DB
)

type UserRepositoryTestSuite struct {
	suite.Suite
	repo *UserRepository
}

func init() {
	viper.Set("app_environment", "test")
	err := os.Setenv("MARKETPLACE_CORE_CONFIG", "/Users/e180/Projects/personal/marketplace_core")
	if err != nil {
		panic(err)
	}

	dbConf, err := config.NewDBConnection()
	if err != nil {
		panic(err)
	}

	db = dbConf.DB
}

func (s *UserRepositoryTestSuite) SetupConnection() {
	s.repo = NewUserRepository(db)
}

func (s *UserRepositoryTestSuite) TestNewRepository() {
	s.repo = nil
	s.Nil(s.repo)

	s.repo = NewUserRepository(db)
	s.NotNil(s.repo)
}

func (s *UserRepositoryTestSuite) TestSaveUser_failed() {
	s.SetupConnection()
	user := &domain.User{
		ID:          uuid.UUID{},
		Name:        "",
		Username:    "",
		Email:       "",
		PhoneNumber: "",
		Gender:      0,
		BirthDay:    time.Time{},
		BirthPlace:  "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}

	err := s.repo.Save(user)
	s.Error(err)
}

func (s *UserRepositoryTestSuite) TestSaveUser_success() {
	s.SetupConnection()

	name := autoname.Generate("")
	username := strings.Replace(name, " ", "_", -1)
	email := fmt.Sprintf("%s@email.com", username)

	user := &domain.User{
		ID:          uuid.UUID{},
		Name:        name,
		Username:    username,
		Email:       email,
		PhoneNumber: "01923999293893",
		Gender:      const_.Male,
		BirthDay:    time.Time{},
		BirthPlace:  "",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
		DeletedAt:   gorm.DeletedAt{},
	}

	err := s.repo.Save(user)
	s.NoError(err)
}

func (s *UserRepositoryTestSuite) TestGetUser_notfound() {
	s.SetupConnection()

	uid := uuid.NewV4()
	_, err := s.repo.GetUserById(uid.String())
	s.Error(err)
}

func (s *UserRepositoryTestSuite) TestGetUser_found() {
	s.SetupConnection()

	user, err := s.repo.GetUserById("b4dc7945-1921-4b1d-853a-e262a9aa465c")
	s.NoError(err)

	s.NotNil(user)
	s.Equal(user.ID.String(), "b4dc7945-1921-4b1d-853a-e262a9aa465c")
}

func TestUserRepository(t *testing.T) {
	suite.Run(t, &UserRepositoryTestSuite{})
}
