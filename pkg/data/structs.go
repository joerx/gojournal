package data

// This package defines the business data model used by the applications

// Entry represents a journal entry
type Entry struct {
	ID          uint64
	Title       string
	Content     string
	Attachments []*Attachment
}

// Attachment represents an arbitrary media file or piece of information that is somehow related to an entry,
// e.g. images, videos, weblinks. They stored somewhere in the cloud and are referenced by their URL
// As far as the backend is concerned they are more or less opaque, it's up to the client to represent them
// according to their type.
type Attachment struct {
	ID       uint64
	URL      string
	Type     string
	Metadata AttachmentMeta
}

// AttachmentMeta abstractly represents a JSON blob containing arbitrary metadata extracted from an attachment,
// specific to its media type
type AttachmentMeta interface{}

// ImageMeta represents metadata for an image attachment
type ImageMeta struct {
	Filename    string
	Orientation string
	Dimensions  string
}
