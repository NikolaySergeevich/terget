package target

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type HTTPTransport struct {
	context         context.Context
	logger          *slog.Logger
	router          *gin.Engine
	service         *Service
	timeoutDuration time.Duration
}

func NewHTTPTransport(context context.Context, logger *slog.Logger, router *gin.Engine, service *Service) *HTTPTransport {
	var timeoutInMilliseconds int
	timeoutInMilliseconds, err := strconv.Atoi(os.Getenv("TIMEOUT_IN_MILLISECONDS"))
	if err != nil {
		timeoutInMilliseconds = TIMEOUT_IN_MILLISECONDS
	}

	return &HTTPTransport{
		context:         context,
		logger:          logger,
		router:          router,
		service:         service,
		timeoutDuration: time.Duration(timeoutInMilliseconds) * time.Millisecond,
	}
}

func (transport *HTTPTransport) ListenAndServe(host, port string) {
	transport.router.Run(host + ":" + port)
}

func (transport *HTTPTransport) RegisterRoutes() {
	routerGroup := transport.router.Group("/target/v1")
	routerGroup.GET("/product", transport.findAllProduct)
}

// @Summary Поиск счета по его продуктов
// @Description Получить все продукты в списке
// @Tags Поиск
// @Router /product [get]
func (transport *HTTPTransport) findAllProduct(c *gin.Context) {

	// ctx, cancel := context.WithTimeout(transport.context, transport.timeoutDuration)
	// defer func() {
	// 	cancel()
	// }()

	// // result, err := transport.service.FindAccountByNumber(ctx, findAccountByNumberpPayload)
	// // if err != nil {
	// // 	transport.logger.Error(ERROR_CALLING_SERVICE_LAYER, "msg", err.Error())
	// // 	c.JSON(http.StatusBadRequest, FindAccountByNumberErrorResponse{
	// // 		Error: err.Error(),
	// // 	})
	// // 	return
	// // }

	// transport.logger.Info(INFO_RESULT_RECEIVED, "result", result)

	// response := FindAccountByNumberSuccessResponse{
	// 	Number:  result.AccountData.Number,
	// 	Balance: result.AccountData.Balance,
	// 	Status:  result.AccountData.Status,
	// }

	c.JSON(http.StatusOK, nil)
}