package dstate

import (
	"github.com/bwmarrin/discordgo"
	"time"
)

// MemberState represents the state of a member
type MemberState struct {
	Guild *GuildState

	// The ID of the member, safe to access without locking
	ID string `json:"id"`

	MemberSet bool

	// The time at which the member joined the guild, in ISO8601.
	// This may be zero if the member hasnt been updated
	JoinedAt time.Time `json:"joined_at"`

	// The nickname of the member, if they have one.
	Nick string `json:"nick"`

	// A list of IDs of the roles which are possessed by the member.
	Roles []string `json:"roles"`

	// Wether the presence Information was set
	PresenceSet    bool             `json:"presence_set"`
	PresenceStatus discordgo.Status `json:"presence_status"`
	PresenceGame   *discordgo.Game  `json:"presence_game"`

	// The users username.
	Username string `json:"username"`

	// The hash of the user's avatar. Use Session.UserAvatar
	// to retrieve the avatar itself.
	Avatar string `json:"avatar"`

	// The discriminator of the user (4 numbers after name).
	Discriminator string `json:"discriminator"`

	// Whether the user is a bot, safe to access without locking
	Bot bool `json:"bot"`
}

func (m *MemberState) UpdateMember(member *discordgo.Member) {
	// Patch
	if member.JoinedAt != "" {
		m.JoinedAt, _ = time.Parse(time.RFC3339, member.JoinedAt)
	}

	if member.Roles != nil {
		m.Roles = member.Roles
	}

	// Seems to always be provided
	m.Nick = member.Nick

	m.Username = member.User.Username
	m.Avatar = member.User.Avatar
	m.Discriminator = member.User.Discriminator
	m.MemberSet = true
}

func (m *MemberState) UpdatePresence(presence *discordgo.Presence) {
	m.PresenceSet = true
	m.PresenceStatus = presence.Status
	m.PresenceGame = presence.Game

	if !m.MemberSet {
		m.Nick = presence.Nick
	}

	if presence.User.Username != "" {
		m.Username = presence.User.Username
	}

	if presence.User.Avatar != "" {
		m.Avatar = presence.User.Avatar
	}
}

// Copy returns a copy of the state, this is not a deep copy so the slices will point to the same arrays, so they're only read safe, not write safe
func (m *MemberState) Copy() *MemberState {
	cop := new(MemberState)
	*cop = *m
	return cop
}
