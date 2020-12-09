package modles

import "github.com/shopspring/decimal"

// Entrust bean is order record
type Entrust struct {
	*BaseEntity

	MarketName         string          `sql:""`
	MarketId           int             `sql:""`
	UserId             int             `sql:""`
	Price              decimal.Decimal `sql:""`
	Numbers            decimal.Decimal `sql:""`
	TotalMoney         decimal.Decimal `sql:""`
	CompleteNumbers    decimal.Decimal `sql:""`
	CompleteTotalMoney decimal.Decimal `sql:""`
	Types              string          `sql:""`
	Status             string          `sql:""`
	AcctType           string          `sql:""`
	CancelAt           string          `sql:""`
	CompleteAt         string          `sql:""`
}

/**
为IdeaVim插件增加自动切换为英文输入法的功能
输入法自动切换功能不会默认启用
编辑器中normal模式下输入输入下面的指令以启用自动切换输入法功能：
:set keep-english-in-normal 开启输入法自动切换功能
:set keep-english-in-normal-and-restore-in-insert 回到insert模式时恢复输入法
:set nokeep-english-in-normal-and-restore-in-insert 保留输入法自动切换功能，但是回到insert模式不恢复输入法
:set nokeep-english-in-normal 关闭输入法自动切换功能
也可以通过将`set keep-english-in-normal[-and-restore-in-insert]`加入到`~/.ideavimrc`文件中并重启IDE来启用插件功能。

注意:支持MacOS和Windows, 通过fcitx-remote支持Linux
MacOS需要开启英语美国键盘或ABC键盘 Windows需要开启英语美国键盘 Linux需要使用fcitx输入法，通过fcitx-remote切换
*/
