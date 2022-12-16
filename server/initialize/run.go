package initialize

func Run()  {
	LoadConfig()
	Mysql()
	Redis()
	Alipay()
	Router()
}