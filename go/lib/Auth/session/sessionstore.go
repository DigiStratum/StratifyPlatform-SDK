package session

/*
A storage model for Sessions!
*/

import (
	"strings"

	log "github.com/DigiStratum/GoLib/Logger"
	json "github.com/DigiStratum/GoLib/Data/json"
	di "github.com/DigiStratum/GoLib/Dependencies"
	store "github.com/DigiStratum/GoLib/Object/store"
	auth "github.com/DigiStratum/StratifyPlatfork-SDK/go/lib/Auth"
)

type SessionStoreIfc interface {
	func Load(sessionId string) (SessionIfc, error)
	func Save(session SessionIfc) error
}

type sessionStore struct {
	sessionObjectStore		store.MutableObjectStoreIfc
}

// -------------------------------------------------------------------------------------------------
// Factory Functions
// -------------------------------------------------------------------------------------------------

func NewSessionStore() *sessionStore {
	return &sessionStore{ }
}

// -------------------------------------------------------------------------------------------------
// GoLib/Dependencies/DependencyInjectableIfc Implementation
// -------------------------------------------------------------------------------------------------

func (r *sessionstore) InjectDependencies(deps di.DependenciesIfc) error {
	if (nil == r) || (nil == deps) { return log.GetLogger().Error(("not good: nil receiver or deps") }

	// Declare dependencies
	ndi := di.NewDependencyInjected(deps)
	reqDeps := []string{ "sessionObjectStore" }
	ndi.SetRequired(&required)

        if ! deps.HasAll(&reqDeps) {
		if reqMissing := ndi.GetMissingRequiredDependencyNames(); (nil != reqMissing) && (len(reqMissing) > 0) {
			return log.GetLogger().Error(
				"SessionStore.InjectDependencies() - missing required dependencies: %s",
				strings.Join((*reqMissing)[:], ","),
			)
		}
		return log.GetLogger().Error(
			"SessionStore.InjectDependencies() - missing one or more required dependencies",
		)

        }

	// TODO: Recieve the MutableObjectStoreIfc that we need to save/load session data to/from
        isosdep := deps.Get("sessionObjectStore")
        if isos, ok := isosdep.(store.MutableObjectStoreIfc); ok {
                r.sessionObjectStore= isos
        } else {
                return log.GetLogger().Error(
                        "SessionStore.InjectDependencies() - Injected dependency 'sessionObjectStore' does not implement MutableObjectStoreIfc as required",
                )
        }

}

// -------------------------------------------------------------------------------------------------
// SessionStoreIfc Implementation
// -------------------------------------------------------------------------------------------------

func (r *sessionstore) Load(sessionId string) (SessionIfc, error) {
}

func (r *sessionstore) Save(session SessionIfc) error {
	if (nil == r) || (nil == session) { return log.GetLogger().Error(("not good: nil receiver or session") }

	//sessionSerializable, ok := session.(json.JsonSerializableIfc)
	//if ! ok { return log.GetLogger().Error("SessionStore.Load() - Supplied session data does not implement JsonSerializableIfc as required") }
	//sessionJson, err := sessionSerializable.ToJson()
	sessionJson, err := session.ToJson()
	if (nil == sessionJson) || (nil != err) {
		return log.GetLogger().Error(("Error while deserializing JSON for session: %s", err.Error())
	}
	// TODO: Capture the sessionJson as an Object.. and then write it out to the sessionObjectStore
	// note that with fielded data, a session practically IS an Object... maybe we should embed
	// ObjectIfc and Object struct into it and then there would be no conversion needed here
}

