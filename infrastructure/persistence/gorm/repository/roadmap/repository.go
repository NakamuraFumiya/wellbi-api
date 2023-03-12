package roadmap

import (
	"errors"
	"github.com/labstack/echo/v4"
	"wellbi-api/core/domain/model"
	"wellbi-api/core/domain/repository/roadmap"
	"wellbi-api/infrastructure/persistence/gorm/handler"
)

type RoadmapRepository struct {
	handler *handler.Handler
}

func NewRoadmapRepository(handler *handler.Handler) roadmap.RoadmapRepository {
	return &RoadmapRepository{handler}
}

func (r *RoadmapRepository) Fetch(c echo.Context) ([]model.Roadmap, error) {
	var roadmaps []model.Roadmap

	if err := r.handler.DB().Order("id desc").Find(&roadmaps).Error; err != nil {
		return nil, err
	}

	return roadmaps, nil
}

func (r *RoadmapRepository) FetchByID(c echo.Context) (*model.Roadmap, error) {
	var roadmap model.Roadmap

	if err := r.handler.DB().First(&roadmap, c.Param("id")).Error; err != nil {
		return nil, err
	}

	return &roadmap, nil
}

func (r *RoadmapRepository) Create(c echo.Context) (*model.Roadmap, error) {
	title := c.FormValue("title")
	if title == "" {
		return nil, errors.New("invalid to or title fields")
	}
	message := c.FormValue("message")
	// FIXME: 実際はlexicalのノードが送られてくるので、バリデーションが効いていないので修正必要
	if message == "" {
		return nil, errors.New("invalid to or message fields")
	}
	var imageURL *string
	formImageURL := c.FormValue("imageURL")
	if formImageURL == "" {
		imageURL = nil
	} else {
		imageURL = &formImageURL
	}
	roadmap := model.Roadmap{
		Title:    title,
		Message:  message,
		ImageURL: imageURL,
	}
	if err := r.handler.DB().Create(&roadmap).Error; err != nil {
		return nil, err
	}
	return &roadmap, nil
}

func (r *RoadmapRepository) Update(c echo.Context) error {
	title := c.FormValue("title")
	if title == "" {
		return errors.New("invalid to or title fields")
	}
	message := c.FormValue("message")
	// FIXME: 実際はlexicalのノードが送られてくるので、バリデーションが効いていないので修正必要
	if message == "" {
		return errors.New("invalid to or message fields")
	}
	var imageURL *string
	formImageURL := c.FormValue("imageURL")
	if formImageURL == "" {
		imageURL = nil
	} else {
		imageURL = &formImageURL
	}

	if err := r.handler.DB().Model(&model.Roadmap{}).
		Where("id = ?", c.Param("id")).
		Updates(map[string]interface{}{"title": title, "message": message, "image_url": imageURL}).
		Error; err != nil {
		return err
	}
	return nil
}

func (r *RoadmapRepository) Delete(c echo.Context) error {
	if err := r.handler.DB().Delete(&model.Roadmap{}, c.Param("id")).Error; err != nil {
		return err
	}
	return nil
}
