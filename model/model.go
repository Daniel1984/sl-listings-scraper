package model

type Listings []struct {
	Listing struct {
		Bathrooms             float32 `json:"bathrooms"`
		Bedrooms              float32 `json:"bedrooms"`
		Beds                  float32 `json:"beds"`
		City                  string  `json:"city"`
		ID                    int64   `json:"id"`
		IsNewListing          bool    `json:"is_new_listing"`
		IsSuperhost           bool    `json:"is_superhost"`
		Lat                   float64 `json:"lat"`
		Lng                   float64 `json:"lng"`
		LocalizedCity         string  `json:"localized_city"`
		LocalizedNeighborhood string  `json:"localized_neighborhood"`
		Name                  string  `json:"name"`
		Neighborhood          string  `json:"neighborhood"`
		PersonCapacity        int     `json:"person_capacity"`
		PictureCount          int     `json:"picture_count"`
		PictureURL            string  `json:"picture_url"`
	} `json:"listing"`
}

type Address struct {
	ExploreTabs []struct {
		TabId    string `json:"tab_id"`
		TabName  string `json:"tab_name"`
		Sections []struct {
			Listings Listings `json:"listings"`
		} `json:"sections"`
		PaginationMetadata struct {
			HasNextPage   bool `json:"has_next_page"`
			SectionOffset int  `json:"section_offset"`
			ItemsOffset   int  `json:"items_offset"`
		} `json:"pagination_metadata"`
	} `json:"explore_tabs"`
}
