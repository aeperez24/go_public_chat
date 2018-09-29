package delivery

import (
	"log"
	usecase "message/useCase"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/spf13/viper"
)

//handler definition
type HttpMessageHandler struct {
	UseCase    usecase.MessageUseCase
	imgUsecase usecase.ImageUseCase
}

func (a *HttpMessageHandler) getMessagesByTime(c echo.Context) error {

	m := echo.Map{}
	c.Bind(&m)
	// text := c.Param("text")
	from := m["from"].(string)
	to := m["to"].(string)

	dayFormat := viper.GetString("timeFormat")
	fromTime, error := time.Parse(dayFormat, from)
	if error != nil {
		log.Printf("%v", error)
		panic("error al formatear fecha de consulta")
	}
	toTime, error2 := time.Parse(dayFormat, to)
	if error2 != nil {
		panic("error al formatear fecha de consulta")
	}

	result := a.UseCase.GetMessages(fromTime, toTime)
	toTime.String()
	return c.JSON(http.StatusOK, result)
}

func (a *HttpMessageHandler) saveMessage(c echo.Context) error {
	m := echo.Map{}
	c.Bind(&m)
	// text := c.Param("text")
	user := m["user"].(string)
	text := m["text"].(string)
	//log.Printf("user %v", user)
	a.UseCase.SaveMessage(user, text)
	return c.JSON(http.StatusOK, true)
}

func (a *HttpMessageHandler) upLoadFile(c echo.Context) error {
	//TODO: VALIDATE FILZE SIZE
	//TODO: VALIDATE FILE IS IMAGE
	c.FormFile("file")
	file, err := c.FormFile("file")
	if err != nil {
		panic(err)
	}
	filename := a.imgUsecase.UpLoadImage(file)

	return c.JSON(http.StatusOK, filename)
}

func (a *HttpMessageHandler) getImageByName(c echo.Context) error {

	fileName := c.QueryParam("name")
	result := a.imgUsecase.RecoverImage(fileName)
	return c.JSON(http.StatusOK, result)
}

//routing
func NewMesajeHttpHandler(e *echo.Echo, us usecase.MessageUseCase, imgcase usecase.ImageUseCase) {
	handler := &HttpMessageHandler{
		UseCase:    us,
		imgUsecase: imgcase,
	}

	public := e.Group("/chatMessage")
	public.POST("/getByDate", handler.getMessagesByTime)
	public.POST("/postMessage", handler.saveMessage)
	public.POST("/upLoadFile", handler.upLoadFile)
	public.GET("/getImage", handler.getImageByName)

}
