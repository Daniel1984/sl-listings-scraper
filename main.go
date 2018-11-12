package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/scraping-service/shared"
	"github.com/sl-listings-scraper/model"
	"github.com/sl-listings-scraper/utils"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

func GetPropertiesFromAirbnb(url string, ch chan []byte) {
	httpClient := &http.Client{}
	userAgentStr := shared.GetUserAgent()
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("authority", "www.airbnb.com")
	req.Header.Set("User-Agent", userAgentStr)
	req.Header.Set("x-csrf-token", "V4$.airbnb.com$HxMVGU-RyKM$1Zwcm1JOrU3Tn0Y8oRrvN3Hc67ZQSbOKVnMjCRtZPzQ=")

	res, err := httpClient.Do(req)

	if err != nil {
		fmt.Println("getting properties errror: ", err)
		return
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	ch <- body
}

func respondWithError(statusCode int, msg string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: statusCode,
		Body:       msg,
	}, nil
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	street := req.QueryStringParameters["street"]

	if street == "" {
		return respondWithError(500, "Street name must be present in query params")
	}

	streetName := &url.URL{Path: street}
	encodedStreetName := streetName.String()
	listingUrl := utils.GetListingsURL(encodedStreetName)
	channel := make(chan []byte)

	var sectionOffset int = 0
	var itemsOffset int = 0
	var scrapeInProgress bool = true
	var listings model.Listings

	for scrapeInProgress {
		offsetUrl := fmt.Sprintf("%s%s%d%s%d", listingUrl, "&section_offset=", sectionOffset, "&items_offset=", itemsOffset)
		go GetPropertiesFromAirbnb(offsetUrl, channel)
		responseBody := <-channel
		listing := model.Listing{}

		if err := json.Unmarshal(responseBody, &listing); err != nil {
			continue
		} else {
			hometabIndex := sort.Search(len(listing.ExploreTabs), func(i int) bool {
				tabId := listing.ExploreTabs[i].TabId
				tabName := listing.ExploreTabs[i].TabName
				return tabId == "home_tab" || tabId == "all_tab" || tabName == "Homes"
			})

			if hometabIndex < len(listing.ExploreTabs) {
				exploreTabs := listing.ExploreTabs[hometabIndex]
				sectionOffset = exploreTabs.PaginationMetadata.SectionOffset
				itemsOffset = exploreTabs.PaginationMetadata.ItemsOffset
				listings = append(listings, exploreTabs.Sections[0].Listings...)
				scrapeInProgress = exploreTabs.PaginationMetadata.HasNextPage
			}
		}
	}

	responsePayload, err := json.Marshal(listings)

	if err != nil {
		return respondWithError(500, err.Error())
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(responsePayload),
	}, nil
}

func main() {
	lambda.Start(handler)
}
