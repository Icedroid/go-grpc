package dao

import "time"

// Playbook 数据库表
type Playbook struct {
	ID             int64     // playbook项目ID
	UserID         int64     // 用户ID
	PlaybookFileId int64     // playbook文件ID
	Name           string    // 项目名称
	Description    string    // 项目描述
	Created        time.Time // 创建时间
	Updated        time.Time // 更新时间
}

type PlaybookRepository interface {
	// 插入记录
	Insert(playbook *Playbook) (int64, error)
	// 更新记录
	Update(id int64, playbook *Playbook) (int64, error)
	// 更新最新版本的playbook_file
	UpdatePlaybookFileID(id, playbookFileID int64) (int64, error)
	// FindAll 分页查询
	FindAll(where string, offset, limit int) ([]Playbook, error)
	// FindOne 查询单行记录
	FindOne(id int64) (Playbook, error)
	// FindCount 查询记录总数
	FindCount(where string) (int, error)
	// 根据传入的id列表查询Playbook（in, 单个时为等于）
	FindPlaybookByIDs(str string, strIDs []interface{}) ([]Playbook, error)
}
