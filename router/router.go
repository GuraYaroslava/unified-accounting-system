package router

import (
    "github.com/uas/controllers"
    "net/http"
    "reflect"
    "strconv"
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
            params[0] = *controller
            cMethod.Func.Call(params)
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
    numParams := method.Type.NumIn()
    params := make([]reflect.Value, numParams)
    for x := 3; x < numParams; x++ {
        it := method.Type.In(x)
        itk := it.Kind()
        if len(parts) > (x) {
            if itk == reflect.String {
                params[x] = reflect.ValueOf(parts[x])
            } else if itk == reflect.Int {
                intval, err := strconv.Atoi(parts[x])
                if err != nil {
                    intval = -1
                }
                params[x] = reflect.ValueOf(intval)
            }
        } else {
            if itk == reflect.String {
                params[x] = reflect.ValueOf("")
            } else if itk == reflect.Int {
                params[x] = reflect.ValueOf(-1)
            }
        }
    }
    return params
}
