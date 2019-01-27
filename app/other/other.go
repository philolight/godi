package other

import "godi/framework/dependency"

func init(){
	dependency.FactoryRegister(Factory)
}

func Factory() interface{}{
	return &Other{}
}

type Other struct{
	a int
}