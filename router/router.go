package router

import (
    "fmt"
    "github.com/uas/controllers"
    "net/http"
    "reflect"
    //"strconv"
    "strings"
)

type FastCGIServer struct{}

func (this FastCGIServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    url := r.URL
    parts := strings.Split(url.Path, "/")
    controllerName := "index"
    methodName := "index"
    if len(parts) < 2 {
        //index
    } else if len(parts) < 3 {
        if parts[1] != "" {
            controllerName = parts[1]
        }
    } else {
        controllerName = parts[1]
        if parts[2] != "" {
            methodName = parts[2]
        }
    }
    controller := FindController(controllerName)
    if controller != nil {
        controller.Elem().FieldByName("Request").Set(reflect.ValueOf(r))
        controller.Elem().FieldByName("Response").Set(reflect.ValueOf(w))
        cType := controller.Type()
        cMethod := FindMethod(cType, methodName)
        if cMethod != nil {
            params := PopulateParams(*cMethod, parts)
            allParams := make([]reflect.Value, len(params)+1)
            allParams[0] = *controller
            for i := 0; i < len(params); i++ {
                allParams[i+1] = params[i]
            }
            cMethod.Func.Call(allParams)
        } else {
            http.Error(w, "Unable to locate index method in controller.", http.StatusMethodNotAllowed)
        }
    } else {
        http.Error(w, "Unable to locate default controller.", http.StatusMethodNotAllowed)
    }
}

func FindController(controllerName string) *reflect.Value {
    baseController := new(controllers.BaseController)
    cmt := reflect.TypeOf(baseController)
    count := cmt.NumMethod()
    for i := 0; i < count; i++ {
        cmt_method := cmt.Method(i)
        if strings.ToLower(cmt_method.Name) == strings.ToLower(controllerName) {
            params := make([]reflect.Value, 1)
            params[0] = reflect.ValueOf(baseController)
            result := cmt_method.Func.Call(params)
            return &result[0]
        }
    }
    return nil
}

func FindMethod(cType reflect.Type, methodName string) *reflect.Method {
    count := cType.NumMethod()
    for i := 0; i < count; i++ {
        method := cType.Method(i)
        if strings.ToLower(method.Name) == strings.ToLower(methodName) {
            return &method
        }
    }
    return nil
}

func PopulateParams(method reflect.Method, parts []string) []reflect.Value {
    numParams := method.Type.NumIn() - 1
    //fmt.Println("name: ", method.Name)
    params := make([]reflect.Value, numParams)
    //fmt.Println("num params: ", numParams)
    for x := 0; x < numParams; x++ {
        //it := method.Type.In(x)
        //itk := it.Kind()
        if len(parts) > (x + 3) {
            //fmt.Println("p: ", parts[x+3])
            fmt.Printf("\n%s: type= %s\n", parts[x+3], reflect.TypeOf(parts[x+3]))
            //if itk == reflect.String {
            params[x] = reflect.ValueOf(parts[x+3])
            //} else if itk == reflect.Int {
            //intval, err := strconv.Atoi(parts[x+3])
            //if err != nil {
            //    intval = -1
            //}
            //params[x] = reflect.ValueOf(intval)
            //}
        } else {
            //if itk == reflect.String {
                params[x] = reflect.ValueOf("")
            //} else if itk == reflect.Int {
            //    params[x] = reflect.ValueOf(-1)
            //}
        }
    }
    fmt.Println("answer params: ", params)
    return params
}
