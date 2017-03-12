package tindergo

import (
	"encoding/json"
	"errors"
	"time"
)

type Meta struct {
	Status          int `json:"status"`
	ClientResources struct {
		RateCard struct {
			Carousel []struct {
				Slug string `json:"slug"`
			} `json:"carousel"`
		} `json:"rate_card"`
		PlusScreen []string `json:"plus_screen"`
	} `json:"client_resources"`
	Notifications []interface{} `json:"notifications"`
	Groups        []struct {
		Type        string `json:"type"`
		SubType     string `json:"sub_type"`
		Key         string `json:"key"`
		GroupID     string `json:"group_id"`
		Version     int    `json:"version"`
		IsPrimary   bool   `json:"is_primary,omitempty"`
		IsBaseGroup bool   `json:"is_base_group,omitempty"`
	} `json:"groups"`
	Rating struct {
		LikesRemaining int `json:"likes_remaining"`
		SuperLikes     struct {
			Remaining                    int         `json:"remaining"`
			AlcRemaining                 int         `json:"alc_remaining"`
			NewAlcRemaining              int         `json:"new_alc_remaining"`
			Allotment                    int         `json:"allotment"`
			SuperlikeRefreshAmount       int         `json:"superlike_refresh_amount"`
			SuperlikeRefreshInterval     int         `json:"superlike_refresh_interval"`
			SuperlikeRefreshIntervalUnit string      `json:"superlike_refresh_interval_unit"`
			ResetsAt                     interface{} `json:"resets_at"`
		} `json:"super_likes"`
	} `json:"rating"`
	Travel struct {
		IsTraveling bool `json:"is_traveling"`
	} `json:"travel"`
	Purchases []interface{} `json:"purchases"`
	Versions  struct {
		ActiveText         string `json:"active_text"`
		AgeFilter          string `json:"age_filter"`
		Matchmaker         string `json:"matchmaker"`
		Trending           string `json:"trending"`
		TrendingActiveText string `json:"trending_active_text"`
	} `json:"versions"`
	Globals struct {
		BoostIntroMultiplier     int    `json:"boost_intro_multiplier"`
		InviteType               string `json:"invite_type"`
		RecsInterval             int    `json:"recs_interval"`
		UpdatesInterval          int    `json:"updates_interval"`
		RecsSize                 int    `json:"recs_size"`
		MatchmakerDefaultMessage string `json:"matchmaker_default_message"`
		ShareDefaultText         string `json:"share_default_text"`
		BoostDecay               int    `json:"boost_decay"`
		BoostUp                  int    `json:"boost_up"`
		BoostDown                int    `json:"boost_down"`
		Sparks                   bool   `json:"sparks"`
		Kontagent                bool   `json:"kontagent"`
		SparksEnabled            bool   `json:"sparks_enabled"`
		KontagentEnabled         bool   `json:"kontagent_enabled"`
		Mqtt                     bool   `json:"mqtt"`
		TinderSparks             bool   `json:"tinder_sparks"`
		AdSwipeInterval          int    `json:"ad_swipe_interval"`
		FetchConnections         bool   `json:"fetch_connections"`
		RateApp                  bool   `json:"rate_app"`
		AppBoy                   bool   `json:"app_boy"`
		SuperLikeAlcMode         int    `json:"super_like_alc_mode"`
		Plus                     bool   `json:"plus"`
		SuperLike                bool   `json:"super_like"`
		RecsBlend                bool   `json:"recs_blend"`
		SquadsEnabled            bool   `json:"squads_enabled"`
		SquadsExtensionLength    int    `json:"squads_extension_length"`
		SquadsExpirationNotice   int    `json:"squads_expiration_notice"`
		PhotoPreviewEnabled      bool   `json:"photo_preview_enabled"`
		Discount                 bool   `json:"discount"`
		Boost                    bool   `json:"boost"`
		BoostDuration            int    `json:"boost_duration"`
		CanEditJobs              bool   `json:"can_edit_jobs"`
		CanEditSchools           bool   `json:"can_edit_schools"`
		CanAddPhotosFromFacebook bool   `json:"can_add_photos_from_facebook"`
		CanShowCommonConnections bool   `json:"can_show_common_connections"`
	} `json:"globals"`
	Tutorials []string `json:"tutorials"`
	Products  struct {
		Plus struct {
			Regular struct {
				Skus []struct {
					ProductType  string `json:"product_type"`
					PurchaseType string `json:"purchase_type"`
					ProductID    string `json:"product_id"`
					IsPrimary    bool   `json:"is_primary"`
					Terms        int    `json:"terms"`
					IsBaseGroup  bool   `json:"is_base_group,omitempty"`
				} `json:"skus"`
			} `json:"regular"`
		} `json:"plus"`
		Superlike struct {
			Regular struct {
				Skus []struct {
					ProductType  string `json:"product_type"`
					PurchaseType string `json:"purchase_type"`
					ProductID    string `json:"product_id"`
					Amount       int    `json:"amount"`
					IsBaseGroup  bool   `json:"is_base_group,omitempty"`
					IsPrimary    bool   `json:"is_primary,omitempty"`
				} `json:"skus"`
			} `json:"regular"`
		} `json:"superlike"`
		Boost struct {
			Regular struct {
				Skus []struct {
					ProductType  string `json:"product_type"`
					PurchaseType string `json:"purchase_type"`
					ProductID    string `json:"product_id"`
					Amount       int    `json:"amount"`
					IsPrimary    bool   `json:"is_primary,omitempty"`
					IsBaseGroup  bool   `json:"is_base_group,omitempty"`
				} `json:"skus"`
			} `json:"regular"`
		} `json:"boost"`
	} `json:"products"`
	User struct {
		ID              string    `json:"_id"`
		ActiveTime      string    `json:"active_time"`
		AgeFilterMax    int       `json:"age_filter_max"`
		AgeFilterMin    int       `json:"age_filter_min"`
		APIToken        string    `json:"api_token"`
		Bio             string    `json:"bio"`
		BirthDate       time.Time `json:"birth_date"`
		Blend           string    `json:"blend"`
		CreateDate      time.Time `json:"create_date"`
		Discoverable    bool      `json:"discoverable"`
		ConnectionCount int       `json:"connection_count"`
		Interests       []struct {
			CreatedTime string `json:"created_time"`
			ID          string `json:"id"`
			Name        string `json:"name"`
		} `json:"interests"`
		DistanceFilter int    `json:"distance_filter"`
		FullName       string `json:"full_name"`
		Gender         int    `json:"gender"`
		GenderFilter   int    `json:"gender_filter"`
		Name           string `json:"name"`
		Photos         []struct {
			ID             string `json:"id"`
			URL            string `json:"url"`
			Main           bool   `json:"main,omitempty"`
			FileName       string `json:"fileName"`
			FbID           string `json:"fbId"`
			Extension      string `json:"extension"`
			ProcessedFiles []struct {
				URL    string `json:"url"`
				Height int    `json:"height"`
				Width  int    `json:"width"`
			} `json:"processedFiles"`
		} `json:"photos"`
		PhotosProcessing        bool          `json:"photos_processing"`
		PhotoOptimizerEnabled   bool          `json:"photo_optimizer_enabled"`
		PhotoOptimizerHasResult bool          `json:"photo_optimizer_has_result"`
		PingTime                time.Time     `json:"ping_time"`
		SpotifyConnected        bool          `json:"spotify_connected"`
		Badges                  []interface{} `json:"badges"`
		Jobs                    []struct {
			Company struct {
				Displayed bool   `json:"displayed"`
				Name      string `json:"name"`
				ID        string `json:"id"`
			} `json:"company"`
			Title struct {
				Displayed bool   `json:"displayed"`
				Name      string `json:"name"`
				ID        string `json:"id"`
			} `json:"title"`
		} `json:"jobs"`
		Schools            []interface{} `json:"schools"`
		Username           string        `json:"username"`
		SquadsDiscoverable bool          `json:"squads_discoverable"`
		SquadsOnly         bool          `json:"squads_only"`
		Squads             []interface{} `json:"squads"`
		CanCreateSquad     bool          `json:"can_create_squad"`
		SquadAdsShown      bool          `json:"squad_ads_shown"`
	} `json:"user"`
}

func (t *TinderGo) Meta() (Meta, error) {
	meta := Meta{}
	url := "https://api.gotinder.com/meta"
	b, errs := t.requester.Get(url)
	if errs != nil {
		return meta, errs[0]
	}

	err := json.Unmarshal([]byte(b), &meta)
	if err != nil {
		return meta, err
	}

	if meta.Status != 200 {
		return meta, errors.New("Error getting Meta.")
	}

	return meta, nil
}
