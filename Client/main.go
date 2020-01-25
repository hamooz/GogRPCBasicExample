package main
import (
	"google.golang.org/grpc"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"fmt"
	"idefinitive"
)
func main(){
	conn, err := grpc.Dial("localhost:4040",grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	client := math.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:fnum/:snum", func(ctx *gin.Context){
		fnum, err := strconv.ParseUint(ctx.Param("fnum"),10,64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid first number"})
			return
		}
		snum, err :=strconv.ParseUint(ctx.Param("snum"),10,64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error":"invalid second number"})
			return
		}
		req := &math.Request{Firstnumber:int32(fnum),Secondnumber:int32(snum)}
		if respons, err := client.AddService(ctx, req); err==nil {
			ctx.JSON(http.StatusOK, gin.H{"result:":fmt.Sprint(respons.Result)})
		}else{
			ctx.JSON(http.StatusInternalServerError, gin.H{"error":"Internal server error"})
		}
	})
	if err := g.Run(":8080"); err != nil {
		panic("failed client")	
		} 
}