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
class JsonRequest {
    public String type;
    public String id;
    public String username;
    public String messageBody;
    public String recipientNumber;
    public String recipientGroupId;
    public Boolean voice;
    public String code;
    public String deviceName;
    public List<String> attachmentFilenames;
    public String uri;
    public String groupName;
    public List<String> members;
    public String avatar;

    JsonRequest() {}
}
*/

// Request represents a message sent to signald
type Request struct {
	Type                string       `json:"type"`
	ID                  string       `json:"id,omitempty"`
	Username            string       `json:"username,omitempty"`
	MessageBody         string       `json:"messageBody,omitempty"`
	RecipientNumber     string       `json:"recipientNumber,omitempty"`
	RecipientGroupID    string       `json:"recipientGroupId,omitempty"`
	Voice               bool         `json:"voice,omitempty"`
	Code                string       `json:"code,omitempty"`
	DeviceName          string       `json:"deviceName,omitempty"`
	AttachmentFilenames []string     `json:"attachmentFilenames,omitempty"`
	URI                 string       `json:"uri,omitempty"`
	Attachments         []Attachment `json:"attachments,omitempty"`
	GroupName           string       `json:"groupName,omitempty"`
	Members             []string     `json:"members,omitempty"`
	Avatar              string       `json:"avatar,omitempty"`
}

type Attachment struct {
	Filename  string `json:"filename"`
	Caption   string `json:"caption"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	VoiceNote bool   `json:"voiceNote"`
	Preview   bool   `json:"preview"`
}
