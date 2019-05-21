// Auto-generated by avdl-compiler v1.3.29 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/favorite.avdl

package keybase1

import (
	"errors"
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type FolderType int

const (
	FolderType_UNKNOWN FolderType = 0
	FolderType_PRIVATE FolderType = 1
	FolderType_PUBLIC  FolderType = 2
	FolderType_TEAM    FolderType = 3
)

func (o FolderType) DeepCopy() FolderType { return o }

var FolderTypeMap = map[string]FolderType{
	"UNKNOWN": 0,
	"PRIVATE": 1,
	"PUBLIC":  2,
	"TEAM":    3,
}

var FolderTypeRevMap = map[FolderType]string{
	0: "UNKNOWN",
	1: "PRIVATE",
	2: "PUBLIC",
	3: "TEAM",
}

func (e FolderType) String() string {
	if v, ok := FolderTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type FolderConflictType int

const (
	FolderConflictType_NONE                  FolderConflictType = 0
	FolderConflictType_IN_CONFLICT           FolderConflictType = 1
	FolderConflictType_IN_CONFLICT_AND_STUCK FolderConflictType = 2
)

func (o FolderConflictType) DeepCopy() FolderConflictType { return o }

var FolderConflictTypeMap = map[string]FolderConflictType{
	"NONE":                  0,
	"IN_CONFLICT":           1,
	"IN_CONFLICT_AND_STUCK": 2,
}

var FolderConflictTypeRevMap = map[FolderConflictType]string{
	0: "NONE",
	1: "IN_CONFLICT",
	2: "IN_CONFLICT_AND_STUCK",
}

func (e FolderConflictType) String() string {
	if v, ok := FolderConflictTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type ConflictStateType int

const (
	ConflictStateType_AutomaticResolving        ConflictStateType = 0
	ConflictStateType_ManualResolvingServerView ConflictStateType = 1
	ConflictStateType_ManualResolvingLocalView  ConflictStateType = 2
)

func (o ConflictStateType) DeepCopy() ConflictStateType { return o }

var ConflictStateTypeMap = map[string]ConflictStateType{
	"AutomaticResolving":        0,
	"ManualResolvingServerView": 1,
	"ManualResolvingLocalView":  2,
}

var ConflictStateTypeRevMap = map[ConflictStateType]string{
	0: "AutomaticResolving",
	1: "ManualResolvingServerView",
	2: "ManualResolvingLocalView",
}

func (e ConflictStateType) String() string {
	if v, ok := ConflictStateTypeRevMap[e]; ok {
		return v
	}
	return ""
}

type ConflictAutomaticResolving struct {
	IsStuck bool `codec:"isStuck" json:"isStuck"`
}

func (o ConflictAutomaticResolving) DeepCopy() ConflictAutomaticResolving {
	return ConflictAutomaticResolving{
		IsStuck: o.IsStuck,
	}
}

type ConflictManualResolvingServerView struct {
	LocalViews []Path `codec:"localViews" json:"localViews"`
}

func (o ConflictManualResolvingServerView) DeepCopy() ConflictManualResolvingServerView {
	return ConflictManualResolvingServerView{
		LocalViews: (func(x []Path) []Path {
			if x == nil {
				return nil
			}
			ret := make([]Path, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.LocalViews),
	}
}

type ConflictManualResolvingLocalView struct {
	ServerView Path `codec:"serverView" json:"serverView"`
}

func (o ConflictManualResolvingLocalView) DeepCopy() ConflictManualResolvingLocalView {
	return ConflictManualResolvingLocalView{
		ServerView: o.ServerView.DeepCopy(),
	}
}

type ConflictState struct {
	ConflictStateType__         ConflictStateType                  `codec:"conflictStateType" json:"conflictStateType"`
	Automaticresolving__        *ConflictAutomaticResolving        `codec:"automaticresolving,omitempty" json:"automaticresolving,omitempty"`
	Manualresolvingserverview__ *ConflictManualResolvingServerView `codec:"manualresolvingserverview,omitempty" json:"manualresolvingserverview,omitempty"`
	Manualresolvinglocalview__  *ConflictManualResolvingLocalView  `codec:"manualresolvinglocalview,omitempty" json:"manualresolvinglocalview,omitempty"`
}

func (o *ConflictState) ConflictStateType() (ret ConflictStateType, err error) {
	switch o.ConflictStateType__ {
	case ConflictStateType_AutomaticResolving:
		if o.Automaticresolving__ == nil {
			err = errors.New("unexpected nil value for Automaticresolving__")
			return ret, err
		}
	case ConflictStateType_ManualResolvingServerView:
		if o.Manualresolvingserverview__ == nil {
			err = errors.New("unexpected nil value for Manualresolvingserverview__")
			return ret, err
		}
	case ConflictStateType_ManualResolvingLocalView:
		if o.Manualresolvinglocalview__ == nil {
			err = errors.New("unexpected nil value for Manualresolvinglocalview__")
			return ret, err
		}
	}
	return o.ConflictStateType__, nil
}

func (o ConflictState) Automaticresolving() (res ConflictAutomaticResolving) {
	if o.ConflictStateType__ != ConflictStateType_AutomaticResolving {
		panic("wrong case accessed")
	}
	if o.Automaticresolving__ == nil {
		return
	}
	return *o.Automaticresolving__
}

func (o ConflictState) Manualresolvingserverview() (res ConflictManualResolvingServerView) {
	if o.ConflictStateType__ != ConflictStateType_ManualResolvingServerView {
		panic("wrong case accessed")
	}
	if o.Manualresolvingserverview__ == nil {
		return
	}
	return *o.Manualresolvingserverview__
}

func (o ConflictState) Manualresolvinglocalview() (res ConflictManualResolvingLocalView) {
	if o.ConflictStateType__ != ConflictStateType_ManualResolvingLocalView {
		panic("wrong case accessed")
	}
	if o.Manualresolvinglocalview__ == nil {
		return
	}
	return *o.Manualresolvinglocalview__
}

func NewConflictStateWithAutomaticresolving(v ConflictAutomaticResolving) ConflictState {
	return ConflictState{
		ConflictStateType__:  ConflictStateType_AutomaticResolving,
		Automaticresolving__: &v,
	}
}

func NewConflictStateWithManualresolvingserverview(v ConflictManualResolvingServerView) ConflictState {
	return ConflictState{
		ConflictStateType__:         ConflictStateType_ManualResolvingServerView,
		Manualresolvingserverview__: &v,
	}
}

func NewConflictStateWithManualresolvinglocalview(v ConflictManualResolvingLocalView) ConflictState {
	return ConflictState{
		ConflictStateType__:        ConflictStateType_ManualResolvingLocalView,
		Manualresolvinglocalview__: &v,
	}
}

func (o ConflictState) DeepCopy() ConflictState {
	return ConflictState{
		ConflictStateType__: o.ConflictStateType__.DeepCopy(),
		Automaticresolving__: (func(x *ConflictAutomaticResolving) *ConflictAutomaticResolving {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Automaticresolving__),
		Manualresolvingserverview__: (func(x *ConflictManualResolvingServerView) *ConflictManualResolvingServerView {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Manualresolvingserverview__),
		Manualresolvinglocalview__: (func(x *ConflictManualResolvingLocalView) *ConflictManualResolvingLocalView {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Manualresolvinglocalview__),
	}
}

// Folder represents a favorite top-level folder in kbfs.
// This type is likely to change significantly as all the various parts are
// connected and tested.
type Folder struct {
	Name          string         `codec:"name" json:"name"`
	Private       bool           `codec:"private" json:"private"`
	Created       bool           `codec:"created" json:"created"`
	FolderType    FolderType     `codec:"folderType" json:"folderType"`
	TeamID        *TeamID        `codec:"team_id,omitempty" json:"team_id,omitempty"`
	ResetMembers  []User         `codec:"reset_members" json:"reset_members"`
	Mtime         *Time          `codec:"mtime,omitempty" json:"mtime,omitempty"`
	ConflictState *ConflictState `codec:"conflictState,omitempty" json:"conflictState,omitempty"`
}

func (o Folder) DeepCopy() Folder {
	return Folder{
		Name:       o.Name,
		Private:    o.Private,
		Created:    o.Created,
		FolderType: o.FolderType.DeepCopy(),
		TeamID: (func(x *TeamID) *TeamID {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.TeamID),
		ResetMembers: (func(x []User) []User {
			if x == nil {
				return nil
			}
			ret := make([]User, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.ResetMembers),
		Mtime: (func(x *Time) *Time {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.Mtime),
		ConflictState: (func(x *ConflictState) *ConflictState {
			if x == nil {
				return nil
			}
			tmp := (*x).DeepCopy()
			return &tmp
		})(o.ConflictState),
	}
}

type FavoritesResult struct {
	FavoriteFolders []Folder `codec:"favoriteFolders" json:"favoriteFolders"`
	IgnoredFolders  []Folder `codec:"ignoredFolders" json:"ignoredFolders"`
	NewFolders      []Folder `codec:"newFolders" json:"newFolders"`
}

func (o FavoritesResult) DeepCopy() FavoritesResult {
	return FavoritesResult{
		FavoriteFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.FavoriteFolders),
		IgnoredFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.IgnoredFolders),
		NewFolders: (func(x []Folder) []Folder {
			if x == nil {
				return nil
			}
			ret := make([]Folder, len(x))
			for i, v := range x {
				vCopy := v.DeepCopy()
				ret[i] = vCopy
			}
			return ret
		})(o.NewFolders),
	}
}

type FavoriteAddArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type FavoriteIgnoreArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Folder    Folder `codec:"folder" json:"folder"`
}

type GetFavoritesArg struct {
	SessionID int `codec:"sessionID" json:"sessionID"`
}

type FavoriteInterface interface {
	// Adds a folder to a user's list of favorite folders.
	FavoriteAdd(context.Context, FavoriteAddArg) error
	// Removes a folder from a user's list of favorite folders.
	FavoriteIgnore(context.Context, FavoriteIgnoreArg) error
	// Returns all of a user's favorite folders.
	GetFavorites(context.Context, int) (FavoritesResult, error)
}

func FavoriteProtocol(i FavoriteInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.favorite",
		Methods: map[string]rpc.ServeHandlerDescription{
			"favoriteAdd": {
				MakeArg: func() interface{} {
					var ret [1]FavoriteAddArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]FavoriteAddArg)
					if !ok {
						err = rpc.NewTypeError((*[1]FavoriteAddArg)(nil), args)
						return
					}
					err = i.FavoriteAdd(ctx, typedArgs[0])
					return
				},
			},
			"favoriteIgnore": {
				MakeArg: func() interface{} {
					var ret [1]FavoriteIgnoreArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]FavoriteIgnoreArg)
					if !ok {
						err = rpc.NewTypeError((*[1]FavoriteIgnoreArg)(nil), args)
						return
					}
					err = i.FavoriteIgnore(ctx, typedArgs[0])
					return
				},
			},
			"getFavorites": {
				MakeArg: func() interface{} {
					var ret [1]GetFavoritesArg
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[1]GetFavoritesArg)
					if !ok {
						err = rpc.NewTypeError((*[1]GetFavoritesArg)(nil), args)
						return
					}
					ret, err = i.GetFavorites(ctx, typedArgs[0].SessionID)
					return
				},
			},
		},
	}
}

type FavoriteClient struct {
	Cli rpc.GenericClient
}

// Adds a folder to a user's list of favorite folders.
func (c FavoriteClient) FavoriteAdd(ctx context.Context, __arg FavoriteAddArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteAdd", []interface{}{__arg}, nil)
	return
}

// Removes a folder from a user's list of favorite folders.
func (c FavoriteClient) FavoriteIgnore(ctx context.Context, __arg FavoriteIgnoreArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.favorite.favoriteIgnore", []interface{}{__arg}, nil)
	return
}

// Returns all of a user's favorite folders.
func (c FavoriteClient) GetFavorites(ctx context.Context, sessionID int) (res FavoritesResult, err error) {
	__arg := GetFavoritesArg{SessionID: sessionID}
	err = c.Cli.Call(ctx, "keybase.1.favorite.getFavorites", []interface{}{__arg}, &res)
	return
}
