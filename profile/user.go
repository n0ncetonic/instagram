package profile

// User contains all available data on a user
type User struct {
	Biography                     string               `json:"biography"`
	BlockedByViewer               bool                 `json:"blocked_by_viewer"`
	CountryBlock                  bool                 `json:"country_block"`
	ExternalURL                   *string              `json:"external_url"`
	ExternalURLLinkShimmed        *string              `json:"external_url_link_shimmed"`
	EdgeFollowedBy                Count                `json:"edge_followed_by"`
	FollowedByViewer              bool                 `json:"followed_by_viewer"`
	EdgeFollow                    Count                `json:"edge_follow"`
	FollowsViewer                 bool                 `json:"follows_viewer"`
	FullName                      string               `json:"full_name"`
	HasChannel                    bool                 `json:"has_channel"`
	HasBlockedViewer              bool                 `json:"has_blocked_viewer"`
	HighlightReelCount            int                  `json:"highlight_reel_count"`
	HasRequestedViewer            bool                 `json:"has_requested_viewer"`
	ID                            string               `json:"id"`
	IsPrivate                     bool                 `json:"is_private"`
	IsVerified                    bool                 `json:"is_verified"`
	MutualFollowers               *string              `json:"mutual_followers"`
	ProfilePicURL                 string               `json:"profile_pic_url"`
	ProfilePicURLHd               string               `json:"profile_pic_url_hd"`
	RequestedByViewer             bool                 `json:"requested_by_viewer"`
	Username                      string               `json:"username"`
	ConnectedFbPage               *string              `json:"connected_fb_page"`
	EdgeFelixCombinedPostUploads  CombinedPostUploads  `json:"edge_felix_combined_post_uploads"`
	EdgeFelixCombinedDraftUploads CombinedDraftUploads `json:"edge_felix_combined_draft_uploads"`
	EdgeFelixVideoTimeline        VideoTimeline        `json:"edge_felix_video_timeline"`
	EdgeFelixDrafts               Drafts               `json:"edge_felix_drafts"`
	EdgeFelixPendingPostUploads   PendingPostUploads   `json:"edge_felix_pending_post_uploads"`
	EdgeFelixPendingDraftUploads  PendingDraftUploads  `json:"edge_felix_pending_draft_uploads"`
	EdgeOwnerToTimelineMedia      OwnerToTimelineMedia `json:"edge_felix_owner_to_timeline_media"`
	EdgeSavedMedia                SavedMedia           `json:"edge_felix_saved_media"`
	EdgeMediaCollections          MediaCollections     `json:"edge_felix_media_collections"`
}

// GetBiography gets biography of a User
//
// Returns
//
// string
func (u User) GetBiography() string {
	return u.Biography
}

// GetExternalURL gets External URL of a User
//
// Returns
//
// *string (can be nil)
func (u User) GetExternalURL() *string {
	return u.ExternalURL
}

// GetExternalURLShim gets External URL Link Shimmed of a User
//
// Returns
//
// *string (can be nil)
func (u User) GetExternalURLShim() *string {
	return u.ExternalURLLinkShimmed
}

// GetFollowerCount gets number of followers of a User
//
// Returns
//
// int
func (u User) GetFollowerCount() int {
	return u.EdgeFollowedBy.Count
}

// GetFollowingCount gets number of accounts followed by a User
//
// Returns
//
// int
func (u User) GetFollowingCount() int {
	return u.EdgeFollow.Count
}

// GetFullname gets full name of a User
//
// Returns
//
// string
func (u User) GetFullname() string {
	return u.FullName
}

// GetID gets ID of a User
//
// Returns
//
// int
func (u User) GetID() string {
	return u.ID
}

// CheckPrivate checks privacy status of a User
//
// Returns
//
// string
func (u User) CheckPrivate() bool {
	return u.IsPrivate
}

// CheckVerified checks verified status of a User
//
// Returns
//
// string
func (u User) CheckVerified() bool {
	return u.IsVerified
}

// GetProfilePicURL gets url to profile pic of a User
//
// Returns
//
// string
func (u User) GetProfilePicURL() string {
	return u.ProfilePicURL
}

// GetProfilePicURLHD gets url to HD profile pic of a User
//
// Returns
//
// string
func (u User) GetProfilePicURLHD() string {
	return u.ProfilePicURLHd
}

// GetUsername gets username of a User
//
// Returns
//
// string
func (u User) GetUsername() string {
	return u.Username
}

// GetTimelineMedia gets timeline media of a User
//
// Returns
//
// string
func (u User) GetTimelineMedia() OwnerToTimelineMedia {
	return u.EdgeOwnerToTimelineMedia
}
