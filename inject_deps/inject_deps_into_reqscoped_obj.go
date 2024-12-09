package inject_deps

import (
	"errors"
)

// 领域对象是请求作用域内的对象，即生命周期跟请求的生命周期绑定在一起，随请求的到来而创建，随请求处理的结束而终止。

// 第一种将依赖注入领域对象的方式：将依赖对象实现的领域接口声明为数据成员，在构造领域对象时注入依赖对象。
// 缺点：
//   - 领域对象不是在服务启动过程中创建的，而是按处理 API 请求的需要而创建。因此依赖依赖会传递给负责构造它的对象，
//     导致构造它的对象有较多依赖，并且这些依赖并非它自身的直接依赖，导致可读性差一些。
//   - 领域对象的依赖依赖，几乎都是在执行其行为时才会实际使用它们。相比执行行为时注入，构造时注入可读性差，因为构造
//     它们的场景有时候并不执行其行为；构造它们的场景为其注入依赖有些困难，需要额外工作
type accountV1 struct {
	id       int
	userName string
	password string
	// ...

	accountRepo   iAccountRepo
	cryptoService iCryptoService
}

type iAccountRepo interface {
	save1(a *accountV1)
	save2(a *accountV2)
	load(username string) *accountV2
}

type iCryptoService interface {
	encrypt(data string) string
	decrypt(data string) string
}

// 第二种将依赖注入领域对象的方式：对象封装了状态和行为。实现时，通常人们会使用具体编程语言的一个元素（例如，跟对
// 象绑定的函数）来表示对象行为。但是这种对行为这个概念的实现灵活性、扩展性不够好。因此这里将对象行为建模为一个自
// 定义类型，再跟具体对象绑定，既实现了编程语言原生对象的效果，又有更好的灵活性、扩展性。
//
// 跟方式一相比，优点：
// - 构造领域对象时，只需要关注其状态/数据，不需要关注其行为需要的依赖依赖，理解起来自然、直观，构造成本低
// - 构造行为时注入依赖依赖，这样让调用其行为的地方负责注入依赖依赖，注入方便，同时可读性更好
// - 领域对象的数据成员（业务状态）跟它所需的依赖依赖成员完全分离，可读性更好
// - 甚至可以为编译型、强类型语言的对象动态添加行为
type accountV2 struct {
	id       int
	userName string
	password string
	// ...
}

// iBehavior is a tag interface
type iBehavior interface {
	name() string
}

func (a *accountV2) ChangePwd(accountRepo iAccountRepo, cryptoService iCryptoService) changePwd {
	return changePwd{
		this:          a,
		accountRepo:   accountRepo,
		cryptoService: cryptoService,
	}
}

type changePwd struct {
	this *accountV2

	accountRepo   iAccountRepo
	cryptoService iCryptoService
}

func (b *changePwd) name() string {
	return "accountV2.changePwd"
}

func (b *changePwd) Do(username, password, newPwd string) error {
	// new and old password is encrypted
	if b.cryptoService.decrypt(password) != b.this.password {
		return errors.New("wrong password")
	}

	b.this.password = b.cryptoService.decrypt(newPwd)
	b.accountRepo.save2(b.this)
	return nil
}
