/*
 *  Copyright yuncun-leping Author(https://houseme.github.io/yuncun-leping/). All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *  You can obtain one at https://github.com/houseme/yuncun-leping.
 *
 */

package colly

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/houseme/yuncun-leping/internal/shared/domain"
	"github.com/houseme/yuncun-leping/utility/helper"
)

// GetList .
func GetList(ctx context.Context, visit string) (list []*domain.SongItem, err error) {
	logger := g.Log(helper.Helper().Logger(ctx))
	c := colly.NewCollector(
		colly.UserAgent(`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36`),
		colly.Headers(
			map[string]string{
				"Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
			}),
		colly.MaxDepth(1),
		colly.Async(true),
		colly.AllowedDomains("music.163.com"),
		colly.AllowURLRevisit(),
	)
	c.OnRequest(func(r *colly.Request) {
		logger.Debug(ctx, "Visiting", r.URL)
	})
	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		logger.Debug(ctx, "Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		logger.Debug(ctx, "Visited", r.Request.URL)
	})

	c.OnHTML(".f-hide li", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(_ int, el *colly.HTMLElement) {
			var (
				href = el.Attr("href")
				id   uint64
			)
			if hrefArr := strings.Split(href, "="); len(hrefArr) > 0 {
				id = gconv.Uint64(hrefArr[len(hrefArr)-1])
			}
			logger.Debug(ctx, "Visited id: ", el.Attr("href"), " text:", el.Text, " id:", id)
			list = append(list, &domain.SongItem{
				SID:   id,
				Title: el.Text,
			})
		})
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	if err = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	}); err != nil {
		return
	}
	if err = c.Visit(visit); err != nil {
		return
	}
	c.Wait()
	logger.Debug(ctx, "Finished", visit)
	return
}
