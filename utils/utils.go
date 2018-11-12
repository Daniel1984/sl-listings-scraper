package utils

import (
	"fmt"
)

func GetListingsURL(location string) string {
	return fmt.Sprintf("https://www.airbnb.com/api/v2/explore_tabs"+
		"?version=1.3.9"+
		"&_format=for_explore_search_web"+
		"&experiences_per_grid=20"+
		"&items_per_grid=18"+
		"&guidebooks_per_grid=20"+
		"&auto_ib=false"+
		"&fetch_filters=true"+
		"&has_zero_guest_treatment=false"+
		"&is_guided_search=true"+
		"&is_new_cards_experiment=true"+
		"&luxury_pre_launch=false"+
		"&query_understanding_enabled=true"+
		"&show_groupings=true"+
		"&supports_for_you_v3=true"+
		"&timezone_offset=120"+
		"&client_session_id=c2102072-77fe-4663-8006-97eb739901ae"+
		"&metadata_only=false"+
		"&is_standard_search=true"+
		"&selected_tab_id=home_tab"+
		"&place_id=ChIJa3z2sROU3UYRQUVFTI3RACY"+
		"&screen_size=medium"+
		"&query=%v"+
		"&_intents=p1"+
		"&key=d306zoyjsyarp7ifhu67rjxn52tv0t20"+
		"&currency=EUR"+
		"&locale=en", location)
}
