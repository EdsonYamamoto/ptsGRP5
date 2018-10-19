package model

//FacebookCallback ...
type FacebookCallback struct {
	Object string `json:"object,omitempty"`
	Entry  []struct {
		ID        string              `json:"id,omitempty"`
		Time      int                 `json:"time,omitempty"`
		Messaging []FacebookMessaging `json:"messaging,omitempty"`
	} `json:"entry,omitempty"`
}

//FacebookMessaging ...
type FacebookMessaging struct {
	Sender    FacebookUser    `json:"sender,omitempty"`
	Recipient FacebookUser    `json:"recipient,omitempty"`
	Timestamp int             `json:"timestamp,omitempty"`
	Message   FacebookMessage `json:"message,omitempty"`
}

//FacebookUser ...
type FacebookUser struct {
	ID string `json:"id,omitempty"`
}

//FacebookMessage ...
type FacebookMessage struct {
	MID        string `json:"mid,omitempty"`
	Text       string `json:"text,omitempty"`
	QuickReply *struct {
		Payload string `json:"payload,omitempty"`
	} `json:"quick_reply,omitempty"`
	Attachments *[]FacebookAttachment `json:"attachments,omitempty"`
	Attachment  *FacebookAttachment   `json:"attachment,omitempty"`
}

//FacebookAttachment ...
type FacebookAttachment struct {
	Type    string          `json:"type,omitempty"`
	Payload FacebookPayload `json:"payload,omitempty"`
}

//FacebookResponse ...
type FacebookResponse struct {
	Recipient FacebookUser    `json:"recipient,omitempty"`
	Message   FacebookMessage `json:"message,omitempty"`
}

//FacebookPayload ...
type FacebookPayload struct {
	URL string `json:"url,omitempty"`
}
