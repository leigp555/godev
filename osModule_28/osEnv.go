package main

func main()  {
	//获取所有环境变量
	//environ := os.Environ()
	//fmt.Printf("%v",environ)

	//获取某个环境变量
	//environ2:=os.Getenv("APPDATA")
	//fmt.Printf("%v",environ2)

	//查找
	//env, b := os.LookupEnv("APPDATA")
	//fmt.Println(env)
	//fmt.Println(b)

	//设置环境变量
	//err := os.Setenv("env1", "env2")




	//替换
	//os.Setenv("NAME","gopher")
	//os.Setenv("BURROW","/user/gopher")
	//fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}"))    //NAME和BURROW会被上面的gopher 和 /user/gopher替换

	//清空环境变量 危险操作
	//os.Clearenv()
}

