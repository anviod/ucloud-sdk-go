// Code is generated by ucloud-model, DO NOT EDIT IT.

package udisk

/*
RecycleUDiskSet - 回收站列表
*/
type RecycleUDiskSet struct {

	// 销毁倒计时
	CountdownTime int

	// 创建时间
	CreateTime int

	// 过期时间
	ExpiredTime int

	// 磁盘名称
	Name string

	// 磁盘容量
	Size int

	// 业务组
	Tag string

	// 磁盘id
	UDiskId string

	// 可用区
	Zone string
}

/*
UDiskDataSet - DescribeUDisk
*/
type UDiskDataSet struct {

	// 是否支持开启方舟，1支持 ，0不支持
	ArkSwitchEnable int

	// 该盘的备份方式。快照服务："SnapshotService"；数据方舟："UDataArk"；无备份方式：""
	BackupMode string

	// Year,Month,Dynamic,Trial,Postpay
	ChargeType string

	// 是否支持克隆，1支持 ，0不支持
	CloneEnable int

	// 该盘的cmk id
	CmkId string

	// cmk id 别名
	CmkIdAlias string

	// 该盘cmk的状态, Enabled(正常)，Disabled(失效)，Deleted(删除)，NoCmkId(非加密盘)
	CmkIdStatus string

	// 创建时间
	CreateTime int

	// 该盘的密文密钥
	DataKey string

	// 挂载的设备名称
	DeviceName string

	// 请求中的ProtocolVersion字段为1时，需结合IsBoot确定具体磁盘类型:普通数据盘：DiskType:"CLOUD_NORMAL",IsBoot:"False"； 普通系统盘：DiskType:"CLOUD_NORMAL",IsBoot:"True"；SSD数据盘：DiskType:"CLOUD_SSD",IsBoot:"False"；SSD系统盘：DiskType:"CLOUD_SSD",IsBoot:"True"；RSSD数据盘：DiskType:"CLOUD_RSSD",IsBoot:"False"；RSSD系统盘：DiskType:"CLOUD_RSSD",IsBoot:"True"；高效数据盘：DiskType:"CLOUD_EFFICIENCY",IsBoot:"False"；高效系统盘：DiskType:"CLOUD_EFFICIENCY",IsBoot:"True"。请求中的ProtocolVersion字段为0或没有该字段时，云硬盘类型参照如下:普通数据盘：DataDisk；普通系统盘：SystemDisk；SSD数据盘：SSDDataDisk；SSD系统盘：SSDSystemDisk；RSSD数据盘：RSSDDataDisk；RSSD系统盘：RSSDSystemDisk；高效数据盘：EfficiencyDataDisk；高效系统盘：EfficiencySystemDisk。
	DiskType string

	// 过期时间
	ExpiredTime int

	// 挂载的Host的IP
	HostIP string

	// 挂载的Host的Id
	HostId string

	// 挂载的Host的Name
	HostName string

	// 是否是系统盘，是："True", 否："False"
	IsBoot string

	// 资源是否过期，过期:"Yes", 未过期:"No"
	IsExpire string

	// 实例名称
	Name string

	// RDMA集群id，仅RSSD返回该值；其他类型云盘返回""。当云盘的此值与快杰云主机的RdmaClusterId相同时，RSSD可以挂载到这台云主机。
	RdmaClusterId string

	// 容量单位GB
	Size int

	// 是否支持快照，1支持 ，0不支持
	SnapEnable int

	// 该盘快照个数
	SnapshotCount int

	// 该盘快照上限
	SnapshotLimit int

	// 状态:Available(可用),Attaching(挂载中), InUse(已挂载), Detaching(卸载中), Initializating(分配中), Failed(创建失败),Cloning(克隆中),Restoring(恢复中),RestoreFailed(恢复失败)
	Status string

	// 业务组名称
	Tag string

	// 是否开启数据方舟，开启:"Yes", 不支持:"No"
	UDataArkMode string

	// UDisk实例Id
	UDiskId string

	// 挂载的UHost的IP。【即将废弃，建议使用HostIP】
	UHostIP string

	// 挂载的UHost的Id。【即将废弃，建议使用HostId】
	UHostId string

	// 挂载的UHost的Name。【即将废弃，建议使用HostName】
	UHostName string

	// 是否是加密盘，是:"Yes", 否:"No"
	UKmsMode string

	// 是否支持数据方舟，支持:"2.0", 不支持:"1.0"
	Version string

	// 可用区
	Zone string
}

/*
UDiskPriceDataSet - DescribeUDiskPrice
*/
type UDiskPriceDataSet struct {

	// "UDataArk","SnapshotService","UDisk","Total"
	ChargeName string

	// Year， Month， Dynamic，Trial
	ChargeType string

	// 原价(对应计费OriginalPrice)
	ListPrice int

	// 用户折后价(对应计费CustomPrice)
	OriginalPrice int

	// 实际价格 (单位: 分)
	Price int
}

/*
UDiskSnapshotSet - DescribeUDiskSnapshot
*/
type UDiskSnapshotSet struct {

	// 该快照的cmk id
	CmkId string

	// cmk id 别名
	CmkIdAlias string

	// 该快照cmk的状态, Enabled(正常)，Disabled(失效)，Deleted(删除)，NoCmkId(非加密盘)
	CmkIdStatus string

	// 快照描述
	Comment string

	// 创建时间
	CreateTime int

	// 该快照的密文密钥
	DataKey string

	// 磁盘类型，0：普通数据盘；1：普通系统盘；2：SSD数据盘；3：SSD系统盘；4：RSSD数据盘；5：RSSD系统盘。
	DiskType int

	// 【已废弃】过期时间
	ExpiredTime int

	// 对应磁盘是否处于可用状态
	IsUDiskAvailable bool

	// 快照名称
	Name string

	// 容量单位GB
	Size int

	// 快照Id
	SnapshotId string

	// 快照状态，Normal:正常,Failed:失败,Creating:制作中
	Status string

	// 快照的源UDisk的Id
	UDiskId string

	// 快照的源UDisk的Name
	UDiskName string

	// 对应磁盘制作快照时所挂载的主机
	UHostId string

	// 是否是加密盘快照，是:"Yes", 否:"No"
	UKmsMode string

	// 快照版本
	Version string

	// 可用区
	Zone string
}
