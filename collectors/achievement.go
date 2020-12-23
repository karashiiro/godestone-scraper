package collectors

import (
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/karashiiro/godestone/models"
	"github.com/karashiiro/godestone/selectors"
)

// BuildAchievementCollector builds the collector used for processing the page.
func BuildAchievementCollector(meta *models.Meta, profSelectors *selectors.ProfileSelectors, output chan *models.AchievementInfo) *colly.Collector {
	c := colly.NewCollector(
		colly.UserAgent(meta.UserAgentDesktop),
		colly.IgnoreRobotsTxt(),
		colly.MaxDepth(100), // Should be set to ceil(nAchievements / 50) + 1
	)

	achievementSelectors := profSelectors.Achievements

	totalAchievementInfo := &models.TotalAchievementInfo{}
	c.OnHTML(achievementSelectors.TotalAchievements.Selector, func(e *colly.HTMLElement) {
		taStr := achievementSelectors.TotalAchievements.Parse(e)[0]
		ta, err := strconv.ParseUint(taStr, 10, 32)
		if err == nil {
			totalAchievementInfo.TotalAchievements = uint32(ta)
		}
	})
	c.OnHTML(achievementSelectors.AchievementPoints.Selector, func(e *colly.HTMLElement) {
		apStr := achievementSelectors.AchievementPoints.Parse(e)[0]
		ap, err := strconv.ParseUint(apStr, 10, 32)
		if err == nil {
			totalAchievementInfo.TotalAchievementPoints = uint32(ap)
		}
	})

	nextURI := ""
	c.OnHTML(achievementSelectors.ListNextButton.Selector, func(e *colly.HTMLElement) {
		nextURI = achievementSelectors.ListNextButton.Parse(e)[0]
	})

	c.OnHTML(achievementSelectors.List.Selector, func(e1 *colly.HTMLElement) {
		e1.ForEach(achievementSelectors.Entry.Selector, func(i int, e2 *colly.HTMLElement) {
			nextAchievement := &models.AchievementInfo{TotalAchievementInfo: totalAchievementInfo}

			idStr := achievementSelectors.ID.ParseThroughChildren(e2)[0]
			id, err := strconv.ParseUint(idStr, 10, 32)
			if err == nil {
				nextAchievement.ID = uint32(id)
			}

			datetimeSecondsStr := achievementSelectors.Time.ParseThroughChildren(e2)[0]
			datetimeSeconds, err := strconv.ParseInt(datetimeSecondsStr, 10, 32)
			if err == nil {
				nextAchievement.Date = time.Unix(0, datetimeSeconds*1000*int64(time.Millisecond))
			}

			output <- nextAchievement
		})

		if nextURI != "javascript:void(0);" {
			err := e1.Request.Visit(nextURI)
			if err != nil {
				output <- &models.AchievementInfo{
					Error: err,
				}
			}
		}
	})

	return c
}