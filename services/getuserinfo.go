package services

import (
	"Projects/config"
	"Projects/models"
	"strconv"

	"github.com/gocolly/colly/v2"
)

func GetUserInfo(UserName string) (*models.User, error) {
	var user models.User

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		// fmt.Println(r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		// fmt.Println(r.StatusCode)
	})

	c.OnHTML("span.mainRating", func(h *colly.HTMLElement) {

		Rating, err := strconv.ParseFloat(h.ChildText("b"), 32)
		if err != nil {
			panic(err)
		}
		user.UserRating = float32(Rating)
	})
	c.OnHTML("span.offline_username", func(h *colly.HTMLElement) {
		// fmt.Println(h.Text)
		user.Name = h.Text
	})
	c.Visit("https://" + config.Conf.URL + "user/" + UserName)
	user.Id = 1
	return &user, nil

}
