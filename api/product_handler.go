package api

import (
	"hexagonal-architecture-example/api/utility"
	"hexagonal-architecture-example/core/serializer"
	"hexagonal-architecture-example/core/service"
	"hexagonal-architecture-example/dependency"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type ProductHandler interface {
	ServeProduct(http.ResponseWriter, *http.Request)
}
type productHandler struct {
	productService service.ProductService
	serializer serializer.ProductSerializer
}

func (p productHandler) ServeProduct(writer http.ResponseWriter, request *http.Request) {
	if  request.Method==http.MethodPost {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println(err.Error())
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		dto,err:=p.serializer.Decode(body)
		if err != nil {
			log.Println(err.Error())
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		err=p.productService.Store(dto.GetProduct())
		if err!=nil{
			log.Println(err.Error())
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		utility.ShowSuccess(writer, request, http.StatusOK, "Operation Successful.")
		return

	}else if request.Method==http.MethodGet {
		var pathValue string
		pathValue=request.URL.Path
		pathValueList:=strings.Split(pathValue,"/")
			code:=pathValueList[4]
			if code!="" {
				utility.ShowSuccess(writer, request, http.StatusOK, p.productService.FindByCode(code))
				return
			}

		utility.ShowSuccess(writer, request, http.StatusOK, p.productService.FindAll())
		return
	}else if request.Method==http.MethodPut {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Println(err.Error())
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		dto,err:=p.serializer.Decode(body)
		if err != nil {
			log.Println(err.Error())
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		err=p.productService.Update(*dto)
		if err!=nil{
			utility.ShowError(writer, request, http.StatusBadRequest, err.Error())
			return
		}
		utility.ShowSuccess(writer, request, http.StatusOK, "Operation Successful.")
		return
	}else if request.Method==http.MethodDelete {
		var pathValue string
		pathValue=request.URL.Path
		pathValueList:=strings.Split(pathValue,"/")
		code:=pathValueList[4]
		if code!="" {
			err:=p.productService.Delete(code)
			if err!=nil{
				utility.ShowError(writer, request, http.StatusOK, err.Error())
				return
			}
			utility.ShowSuccess(writer, request, http.StatusOK, "Operation Successful!")
			return
		}

		utility.ShowError(writer, request, http.StatusOK, "Please provide a code!")
		return
	}else if request.Method==http.MethodPatch {
		utility.ShowSuccess(writer, request, http.StatusOK, "Yet to implement!")
		return
	}
}

func NewProductHandler() ProductHandler {
	return &productHandler{productService: dependency.GetProductService(),serializer: dependency.GetProductSerializer()}
}
