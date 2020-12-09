package modles

type Notifier struct {
}

type NotifierType int

const (
	EntrustNewNotifierType NotifierType = 1 + iota
	EntrustCancelNotifierType
	EntrustCompleteNotifierType
)
