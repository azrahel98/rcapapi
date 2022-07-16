package main

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/aws/aws-lambda-go/events"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"villas.com/graph/generated"
	"villas.com/graph/resolvers"
	"villas.com/middleware"
	"villas.com/src/data"
	"villas.com/src/impl/mysql"
	"villas.com/src/service"
)

var ginLambda *ginadapter.GinLambda

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}

func init() {
	godotenv.Load()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"*"},
	}))
	r.Use(middleware.MiddlewareGin())
	r.POST("/query", graphqlHandler())
	r.GET("/graph", playgroundHanlder())
	r.GET("/asist/:token", func(ctx *gin.Context) {
		doc := data.NewAsisteData(mysql.AsistenciaImpl{})
		token := ctx.Param("token")
		deco, err := middleware.DecodeTokenJWT(token)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		resutl, err := doc.BuscarAsistencia(deco.Dni, deco.Mes)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, resutl)

	})

	r.GET("/employ/:token", func(ctx *gin.Context) {
		doc := data.NewEmployData(mysql.EmpleadoImpl{})
		token := ctx.Param("token")
		deco, err := middleware.DecodeTokenJWT(token)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		resutl, err := doc.FindByID(deco.Dni)
		if err != nil {
			ctx.JSON(404, err.Error())
			return
		}
		ctx.JSON(200, resutl)

	})
	// r.GET("/papeletas/:token", func(ctx *gin.Context) {
	// 	doc := data.NewDocumentsData(mysql.DocumentsImpl{})
	// 	token := ctx.Param("token")
	// 	deco, err := middleware.DecodeTokenJWT(token)
	// 	if err != nil {
	// 		ctx.JSON(404, err.Error())
	// 		return
	// 	}
	// 	resutl, err := doc.BuscarDocumentosPorDNI(deco.Dni)
	// 	if err != nil {
	// 		ctx.JSON(404, err.Error())
	// 		return
	// 	}
	// 	ctx.JSON(200, resutl)

	// })

	_, err := service.GetInstance()
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println(err)
	}

	r.Run()

	//ginLambda = ginadapter.New(r)
}

func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers:  &resolvers.Resolver{},
		Directives: generated.DirectiveRoot{},
		Complexity: generated.ComplexityRoot{},
	}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHanlder() gin.HandlerFunc {
	h := playground.Handler("Playground", "/query")
	return func(ctx *gin.Context) {
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func main() {
	//log.SetFlags(log.LstdFlags | log.Lshortfile)
	//lambda.Start(Handler)
}
