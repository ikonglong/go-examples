package design_pattern

type User struct {
	*address
	name   string
	gender string
}

func (u *User) Accept(v IUserVisitor) {
	v.Name(u.name)
	v.Gender(u.gender)

	// 方式一
	v.Address(u.address)

	// 方式二，将内嵌对象平铺
	v.AddrCity(u.address.city)
	v.AddrState(u.address.state)
	v.AddrState(u.address.zip)
	v.AddrZip(u.address.zip)
}

type IUserVisitor interface {
	Name(n string)
	Gender(g string)

	// 方式一

	Address(addr *address)

	// 方式二，将内嵌对象平铺

	AddrStreet(s string)
	AddrCity(s string)
	AddrState(s string)
	AddrZip(s string)
}

type address struct {
	street string
	city   string
	state  string
	zip    string
}

func (a *address) Accept(v IAddrVisitor) {
	v.Street(a.street)
	v.City(a.city)
	v.State(a.state)
	v.Zip(a.zip)
}

type IAddrVisitor interface {
	Street(s string)
	City(c string)
	State(s string)
	Zip(s string)
}

// userMapper maps user/user_record to user_record/user.
type userMapper struct {
	addrMapper addrMapper
	targetR    userRecord
}

func (m *userMapper) toRecord(u *User) *userRecord {
	// do something after visiting
	return &m.targetR
}

func (m *userMapper) Name(n string) {
	m.targetR.Name = n
}

func (m *userMapper) Gender(g string) {
	m.targetR.Gender = g
}

// 方式一，将内嵌对象平铺

func (m *userMapper) AddrStreet(s string) {
	m.targetR.AddrStreet = s
}

func (m *userMapper) AddrCity(c string) {
	m.targetR.AddrCity = c
}

func (m *userMapper) AddrState(s string) {
	m.targetR.AddrState = s
}

func (m *userMapper) AddrZip(z string) {
	m.targetR.AddrZip = z
}

// 方式二

func (m *userMapper) Address(a *address) {
	a.Accept(&m.addrMapper)
}

type addrMapper struct {
	targetR userRecord
}

func (m *addrMapper) Street(s string) {
	m.targetR.AddrStreet = s
}

func (m *addrMapper) City(c string) {
	m.targetR.AddrCity = c
}

func (m *addrMapper) State(s string) {
	m.targetR.AddrState = s
}

func (m *addrMapper) Zip(s string) {
	m.targetR.AddrZip = s
}

// userRecord 表示 db 中的一条用户记录
type userRecord struct {
	Name       string
	Gender     string
	AddrStreet string
	AddrCity   string
	AddrState  string
	AddrZip    string
}
