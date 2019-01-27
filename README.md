# godi
File based dependency injection framework in golang

Main Feature
- Can set dependencies and parameters with single configuration file
  (So you can change application instances without change source codes)
- No need in application DI like main()
- Can print object diagram to see instance's dependencies
- Can call by method name for all application instances (and guarantee no duplicated method call for a instance)
- Easy configuration file as direct DI code
- No redundant dependency between application instances
  (If your application instance depends on a interface, it doesn't know concrete struct)
- Can set complex type parameters (slice, map and the combinations) with json style string

You can set parameters and set dependencies of application instances with file by godi.

With godi, your application instances no need direct dependencies, so you can easily separate your application if you want to make it as a micro service server.

You don't need setter or constructor, and you don't need dependency injection codes for your application.

1. Register Factory Method

If you have struct App like :

    type App struct {
	    property string
	    Storage  storage.Storage
	    Name     string
	    Value    int
	    Dur      time.Duration
	    T        time.Time
    }

Make factory method like :

    func Factory() interface{} {
        return &App{}
    }

And then register factory method like :

    func init() {
        dependency.FactoryRegister(Factory)
    }

Now godi can create App instance and set dependency.

2. Write config file

App struct's name in godi is "app:app"

First "app" means App struct is in package "app"

Second "app" means App struct is in file "app.go"

2-1. Create Only

If you want to create App struct instance, write confiuration file like :

    app:app{}

2-2. Initial Parameter Setting

If App struct has Name string field, and you want to set it as "Application", write configuration file like :

    app:app{
        Name=Application
    }

You should use uppercase field name like "Name".(not like "name")

2-3. Create and Inject

If you have App and Dep struct, and App has field Dependency Dep(in package dep and file dep.go), and want to set it with Dep instance, write configuration file like :

    app:app{
       Dependency=dep:dep
    }

If you have App and rdb struct, and App has field Storage Storage (Storage is Interface), and want to set Storage field with rdb instance(and rdb implements Storage interface) :

    app:app{
       Storage=rdb:rdb
    }

So you can set dependency concrete instance as a interface.

3. Add Imports

For the last step, you should add imports for main(). In this project, we are using imports/imports.go for all imports.

    import (
        _ "godi/app"
        _ "godi/storage/nosql"
        _ "godi/storage/rdb"
        _ "godi/bapp"
    )

It is needed because godi can make "no dependency" in application, so init() method cannot call without imports.



Compare with code and configuration file

application structs :

    type App struct {
	    property string
	    Storage  storage.Storage
	    Name     string
	    Value    int
	    Dur      time.Duration
	    T        time.Time
    }

    type Storage interface{}
    type rdb struct{
        Connections int
    } (as Storage)
    type nosql struct{} (as Storage)


DI in application :
    rdbInstance := rdb{Connections:100}
    app := App{
        Storage:rdbInstance,
        Name:"Application",
        Value:987654321,
        Dur:time.Duration(time.Second*10),
        T:time.Parse("2006-02-03T15:04:05Z","2006-02-03T15:04:05Z")
    }


With godi's configuration file :
    storage.rdb:rdb{
        Connections=100
    }
    app:app{
        Storage=storage.rdb:rdb
    	Name=Application
    	Value=987654321
    	Dur=10s
    	T=2006-02-03T15:04:05Z
    }