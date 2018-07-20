package profile

// TODO: I'm sure all this repeated code can be condensed

// Count holds an integer count
type Count struct {
	Count int `json:"count"`
}

// CombinedPostUploads holds pagination and edge information for post uploads
type CombinedPostUploads struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// CombinedDraftUploads holds pagination and edge information for draft uploads
type CombinedDraftUploads struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// PageInfo stores a cursor for pagination
type PageInfo struct {
	HasNextPage bool    `json:"has_next_page"`
	EndCursor   *string `json:"end_cursor"`
}

// Edge describes an edge.
// TODO: Implement Edge node structure
type Edge struct {
}

// VideoTimeline holds pagination and edge information for video timelines
type VideoTimeline struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// Drafts holds pagination and edge information for drafts
type Drafts struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// PendingPostUploads holds pagination and edge information for pending uploads
type PendingPostUploads struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// PendingDraftUploads holds pagination and edge information for pending drafts
type PendingDraftUploads struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// OwnerToTimelineMedia holds pagination and edge information for timeline media
type OwnerToTimelineMedia struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// SavedMedia holds pagination and edge information for saved media
type SavedMedia struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}

// MediaCollections holds pagination and edge information for media collections
type MediaCollections struct {
	Count
	PageInfo PageInfo `json:"page_info"`
	Edges    []Edge   `json:"edges"`
}
