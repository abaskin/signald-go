// Copyright © 2018 Finn Herzfeld <finn@janky.solutions>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package signald

/* The class in signald:

public class JsonRequest {
    public String type;
    public String id;
    public String username;
    public String messageBody;
    public String recipientGroupId;
    public JsonAddress recipientAddress;
    public Boolean voice;
    public String code;
    public String deviceName;
    public List<JsonAttachment> attachments;
    public String uri;
    public String groupName;
    public List<String> members;
    public String avatar;
    public JsonQuote quote;
    public int expiresInSeconds;
    public String fingerprint;
    public String trustLevel;
    public ContactStore.ContactInfo contact;
    public String captcha;
    public String name;
    public List<Long> timestamps;
    public long when;
    public JsonReaction reaction;

    JsonRequest() {}
}

*/

// Request represents a message sent to signald
type Request struct {
	Type             string              `json:"type"`
	ID               string              `json:"id,omitempty"`
	Username         string              `json:"username,omitempty"`
	MessageBody      string              `json:"messageBody,omitempty"`
	RecipientAddress *RequestAddress     `json:"recipientAddress,omitempty"`
	RecipientGroupID string              `json:"recipientGroupId,omitempty"`
	Voice            bool                `json:"voice,omitempty"`
	Code             string              `json:"code,omitempty"`
	DeviceName       string              `json:"deviceName,omitempty"`
	Attachments      []RequestAttachment `json:"attachments,omitempty"`
	URI              string              `json:"uri,omitempty"`
	GroupName        string              `json:"groupName,omitempty"`
	Members          []string            `json:"members,omitempty"`
	Avatar           string              `json:"avatar,omitempty"`
	Quote            *RequestQuote       `json:"quote,omitempty"`
	ExpiresInSeconds int                 `json:"expiresInSeconds,omitempty"`
	Fingerprint      string              `json:"fingerprint,omitempty"`
	TrustLevel       string              `json:"trustLevel,omitempty"`
	Contact          *RequestContact     `json:"contact,omitempty"`
	Captcha          string              `json:"captcha,omitempty"`
	Name             string              `json:"name,omitempty"`
	Timestamps       []int64             `json:"timestamps,omitempty"`
	When             int64               `json:"when,omitempty"`
	Reaction         *RequestReaction    `json:"reaction,omitempty"`
	Pin              string              `json:"pin,omitempty"`
}

// RequestAttachment to a message attachhment sent to signald
type RequestAttachment struct {
	Filename  string `json:"filename,omitempty"`
	Caption   string `json:"caption,omitempty"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
	VoiceNote bool   `json:"voiceNote,omitempty"`
	Preview   bool   `json:"preview,omitempty"`
}

// RequestQuote to a message quote sent to signald
type RequestQuote struct {
	ID          int64                    `json:"id,omitempty"`
	Author      RequestAddress           `json:"author,omitempty"`
	Text        string                   `json:"text,omitempty"`
	Attachments []RequestQuoteAttachment `json:"attachments,omitempty"`
}

// RequestQuoteAttachment to a message quote attachment sent to signald
type RequestQuoteAttachment struct {
	ContentType string            `json:"contentType,omitempty"`
	FileName    string            `json:"fileName,omitempty"`
	Thumbnail   RequestAttachment `json:"thumbnail,omitempty"`
}

// RequestContact contact info for a message sent to signald
type RequestContact struct {
	Number string `json:"number,omitempty"`
	Name   string `json:"name,omitempty"`
	Color  string `json:"color,omitempty"`
}

// RequestAddress address info for a message sent to signald
type RequestAddress struct {
	Number string `json:"number,omitempty"`
	UUID   string `json:"uuid,omitempty"`
}

// Empty check if the RequestAddress is empty
func (r *RequestAddress) Empty() bool {
	return r.Number == "" && r.UUID == ""
}

// RequestReaction reaction info
type RequestReaction struct {
	Emoji               string         `json:"emoji,omitempty"`
	Remove              bool           `json:"remove,omitempty"`
	TargetAuthor        RequestAddress `json:"targetAuthor,omitempty"`
	TargetSentTimestamp int64          `json:"TargetSentTimestamp,omitempty"`
}
