package match_service

import (
	"fmt"
	"github.com/gocolly/colly"
	"pp-backend/models"
	"time"
)

func (m *match) GetAll() ([]*models.Article, error) {
	var articles []*models.Article

	articles, err := models.GetArticles(a.getMaps())
	if err != nil {
		return nil, err
	}

	return articles, nil
}

type spiderMatch struct {
	P1      string
	P1Score string
	P2      string
	P2Score string
}

func spiderMatches() {
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (X11; Linux x86_64)    AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.100 Safari/537.36"),
	)

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	detailCollector := c.Clone()

	detailCollector.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})

	detailCollector.OnHTML("body", func(e *colly.HTMLElement) {
		n := spiderMatch{}
		n.P1 = e.ChildText("div[class=Content-sc-1morvta-0 EventCellstyles__WinIndicator-sc-4ti8ha-4 jnfaHp hieeMY]>div")
		n.P2 = e.ChildText("div[class=Content-sc-1morvta-0 EventCellstyles__WinIndicator-sc-4ti8ha-4 jnfaHp hieeMY]>div")
		n.P1Score = e.ChildText("div[class=Content-sc-1morvta-0 EventCellstyles__WinIndicator-sc-4ti8ha-4 jnfaHp hieeMY]>div")
		n.P2Score = e.ChildText("div[class=Content-sc-1morvta-0 EventCellstyles__WinIndicator-sc-4ti8ha-4 jnfaHp hieeMY]>div")
	})

	c.OnHTML("a[class=tt]", func(e *colly.HTMLElement) {
		detailCollector.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://www.sofascore.com/table-tennis/livescore")

	c.Wait()
	detailCollector.Wait()
}
