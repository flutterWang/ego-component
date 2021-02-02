package edingtalk

type Department struct {
	DeptId            int     `json:"dept_id,omitempty"`
	Name              string  `json:"name,omitempty"`
	ParentId          int     `json:"parent_id,omitempty"`
	HideDept          bool    `json:"hide_dept,omitempty"`
	DeptPermits       Ints    `json:"dept_permits,omitempty"`
	UserPermits       Strings `json:"user_permits,omitempty"`
	OuterDept         bool    `json:"outer_dept,omitempty"`
	OuterDeptOnlySelf string  `json:"outer_dept_only_self,omitempty"`
	OuterPermitUsers  Strings `json:"outer_permit_users,omitempty"`
	OuterPermitDepts  Ints    `json:"outer_permit_depts,omitempty"`
	CreateDeptGroup   bool    `json:"create_dept_group,omitempty"`
	Order             int     `json:"order,omitempty"`
	SourceIdentifier  string  `json:"source_identifier,omitempty"`
}

type departmentGetRes struct {
	OpenAPIResponse
	Result Department `json:"result"`
}

type departmentCreateReq = Department

type DepartmentCreateRes struct {
	OpenAPIResponse
	Result struct {
		DeptId int `json:"dept_id"`
	} `json:"result"`
}

type DepartmentUpdateReq struct {
	DeptId            int     `json:"dept_id,omitempty"`
	ParentId          *int    `json:"parent_id,omitempty"`
	HideDept          *bool   `json:"hide_dept,omitempty"`
	DeptPermits       *string `json:"dept_permits,omitempty"`
	UserPermits       *string `json:"user_permits,omitempty"`
	CreateDeptGroup   *bool   `json:"create_dept_group,omitempty"`
	Order             *int    `json:"order,omitempty"`
	Name              *string `json:"name,omitempty"`
	SourceIdentifier  *string `json:"source_identifier,omitempty"`
	OuterDept         *bool   `json:"outer_dept,omitempty"`
	OuterPermitUsers  *string `json:"outer_permit_users,omitempty"`
	OuterPermitDepts  *string `json:"outer_permit_depts,omitempty"`
	OuterDeptOnlySelf *string `json:"outer_dept_only_self,omitempty"`
}

func NewDepartmentUpdateReq(did int) *DepartmentUpdateReq {
	return &DepartmentUpdateReq{DeptId: did}
}

func (d *DepartmentUpdateReq) SetParentId(did int) *DepartmentUpdateReq {
	d.ParentId = &did
	return d
}

func (d *DepartmentUpdateReq) SetHideDept(hideDept bool) *DepartmentUpdateReq {
	d.HideDept = &hideDept
	return d
}

func (d *DepartmentUpdateReq) SetDeptPermits(deptPermits string) *DepartmentUpdateReq {
	d.DeptPermits = &deptPermits
	return d
}

func (d *DepartmentUpdateReq) SetUserPermits(userPermits string) *DepartmentUpdateReq {
	d.UserPermits = &userPermits
	return d
}

func (d *DepartmentUpdateReq) SetCreateDeptGroup(createDeptGroup bool) *DepartmentUpdateReq {
	d.CreateDeptGroup = &createDeptGroup
	return d
}

func (d *DepartmentUpdateReq) SetOrder(order int) *DepartmentUpdateReq {
	d.Order = &order
	return d
}

func (d *DepartmentUpdateReq) SetName(name string) *DepartmentUpdateReq {
	d.Name = &name
	return d
}

func (d *DepartmentUpdateReq) SetSourceIdentifier(name string) *DepartmentUpdateReq {
	d.SourceIdentifier = &name
	return d
}

func (d *DepartmentUpdateReq) SetOuterDept(outerDept bool) *DepartmentUpdateReq {
	d.OuterDept = &outerDept
	return d
}

func (d *DepartmentUpdateReq) SetOuterPermitUsers(outerPermitUsers string) *DepartmentUpdateReq {
	d.OuterPermitUsers = &outerPermitUsers
	return d
}

func (d *DepartmentUpdateReq) SetOuterPermitDepts(outerPermitDepts string) *DepartmentUpdateReq {
	d.OuterPermitDepts = &outerPermitDepts
	return d
}

func (d *DepartmentUpdateReq) SetOuterDeptOnlySelf(outerDeptOnlySelf string) *DepartmentUpdateReq {
	d.OuterDeptOnlySelf = &outerDeptOnlySelf
	return d
}