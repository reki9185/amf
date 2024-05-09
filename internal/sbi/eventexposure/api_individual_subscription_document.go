// /*
//  * Namf_EventExposure
//  *
//  * AMF Event Exposure Service
//  *
//  * API version: 1.0.0
//  * Generated by: OpenAPI Generator (https://openapi-generator.tech)
//  */

package eventexposure

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"

// 	"github.com/free5gc/amf/internal/logger"
// 	"github.com/free5gc/amf/internal/sbi/producer"
// 	"github.com/free5gc/openapi"
// 	"github.com/free5gc/openapi/models"
// 	"github.com/free5gc/util/httpwrapper"
// )

// // DeleteSubscription - Namf_EventExposure Unsubscribe service Operation
// func HTTPDeleteSubscription(c *gin.Context) {
// 	req := httpwrapper.NewRequest(c.Request, nil)
// 	req.Params["subscriptionId"] = c.Param("subscriptionId")

// 	rsp := producer.HandleDeleteAMFEventSubscription(req)

// 	if rsp.Status == http.StatusOK {
// 		c.JSON(http.StatusOK, gin.H{})
// 	} else {
// 		responseBody, err := openapi.Serialize(rsp.Body, "application/json")
// 		if err != nil {
// 			logger.EeLog.Errorln(err)
// 			problemDetails := models.ProblemDetails{
// 				Status: http.StatusInternalServerError,
// 				Cause:  "SYSTEM_FAILURE",
// 				Detail: err.Error(),
// 			}
// 			c.JSON(http.StatusInternalServerError, problemDetails)
// 		} else {
// 			c.Data(rsp.Status, "application/json", responseBody)
// 		}
// 	}
// }

// // ModifySubscription - Namf_EventExposure Subscribe Modify service Operation
// func HTTPModifySubscription(c *gin.Context) {
// 	var modifySubscriptionRequest models.ModifySubscriptionRequest

// 	requestBody, err := c.GetRawData()
// 	if err != nil {
// 		logger.EeLog.Errorf("Get Request Body error: %+v", err)
// 		problemDetail := models.ProblemDetails{
// 			Title:  "System failure",
// 			Status: http.StatusInternalServerError,
// 			Detail: err.Error(),
// 			Cause:  "SYSTEM_FAILURE",
// 		}
// 		c.JSON(http.StatusInternalServerError, problemDetail)
// 		return
// 	}

// 	err = openapi.Deserialize(&modifySubscriptionRequest, requestBody, "application/json")
// 	if err != nil {
// 		problemDetail := "[Request Body] " + err.Error()
// 		rsp := models.ProblemDetails{
// 			Title:  "Malformed request syntax",
// 			Status: http.StatusBadRequest,
// 			Detail: problemDetail,
// 		}
// 		logger.EeLog.Errorln(problemDetail)
// 		c.JSON(http.StatusBadRequest, rsp)
// 		return
// 	}

// 	req := httpwrapper.NewRequest(c.Request, modifySubscriptionRequest)
// 	req.Params["subscriptionId"] = c.Param("subscriptionId")

// 	rsp := producer.HandleModifyAMFEventSubscription(req)

// 	responseBody, err := openapi.Serialize(rsp.Body, "application/json")
// 	if err != nil {
// 		logger.EeLog.Errorln(err)
// 		problemDetails := models.ProblemDetails{
// 			Status: http.StatusInternalServerError,
// 			Cause:  "SYSTEM_FAILURE",
// 			Detail: err.Error(),
// 		}
// 		c.JSON(http.StatusInternalServerError, problemDetails)
// 	} else {
// 		c.Data(rsp.Status, "application/json", responseBody)
// 	}
// }
