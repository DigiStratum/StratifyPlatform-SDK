package session

/*

Retain session state data for an authenticated Identity

We are backed by a GoLib HashMap, so we can extend the session data for this Identity with any
name=value string pair, which includes ANY interface that supports de|serialization from/to string.

TODO:
 * Convert to be an extension of Object instead of HashMap?
 * Expand details specific to the session, outside of the HashMap properties such as anything related
   to session validity, expiration, touch() to refresh, etc.
*/

import (
	auth "github.com/DigiStratum/StratifyPlatfork-SDK/go/lib/Auth"
	hash "github.com/DigiStratum/GoLib/Data/hashmap"
)

type SessionIfc interface {
	hash.HashMapIfc
	json.JsonSerializableIfc
	json.JsonDeserializableIfc
}

type session struct {
	*hash.HashMap
	id		string
	identity	Identity
}

// DTO structure for JSON un|marshal
type sessionJsonDto struct {
	Id		string		`json:"id"`
	Ident		Identity	`json:"identity"`
	Metadata	*hash.HashMap	`json:"metadata"`
}

// -------------------------------------------------------------------------------------------------
// Factory Functions
// -------------------------------------------------------------------------------------------------

func NewSession(identity Identity) *session {
	return &session{
		HashMap:	hash.NewHashMap(),
		identity:	identity,
	}
}

// -------------------------------------------------------------------------------------------------
// encoding/json.Marshaler Interface Implementation
// -------------------------------------------------------------------------------------------------

func (r *HashMap) MarshalJSON() ([]byte, error) {
	t := sessionJsonDto{
		Id:		r.id,
		Ident:		r.identity,
		Metadata:	r.HashMap,
	}

	return gojson.Marshal(t)
}

// -------------------------------------------------------------------------------------------------
// encoding/json.Unmarshaler Interface Implementation
// -------------------------------------------------------------------------------------------------

func (r *HashMap) UnmarshalJSON(value []byte) error {
	t := sessionJsonDto{}
	err :=  gojson.Unmarshal(value, &t)
	if nil != err { return err }
	r.id = t.Id
	r.identity = t.Ident
	r.HashMap.DropAll()
	for k, v := range t.Metadata { r.Set(k, v) }
	return nil
}

// -------------------------------------------------------------------------------------------------
// JsonSerializableIfc
// -------------------------------------------------------------------------------------------------

func (r *session) ToJson() (*string, error) {
	jbytes, err := r.MarshalJSON()
	if nil == err { return nil, err }
	return string([]jbytes), nil
}

// -------------------------------------------------------------------------------------------------
// JsonDeserializableIfc
// -------------------------------------------------------------------------------------------------

func (r *session) FromJson(jsonString *string) error {
	if nil == jsonString { return fmt.Errorf("Can't deserialize nil JSON string") }
	return r.UnmarshalJSON([]byte(*jsonString))
}

