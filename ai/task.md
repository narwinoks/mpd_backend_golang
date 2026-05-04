[88.198ms] [rows:1] SELECT * FROM "personal_access_tokens_m" WHERE token = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbXBsb3llZV9pZCI6MiwiZXhwIjoxNzc3ODcxNTY2LCJwcm9maWxlX2lkIjoxLCJyb2xlX2lkIjoyLCJ0eXBlIjoiYWNjZXNzIiwidXNlcl9pZCI6MiwidXNlcm5hbWUiOiJ3aW4ifQ.tmfW64MSo-Zyb6KhJCtL4qk7oi8mdoE5rLC-OpgRPZo' AND "personal_access_tokens_m"."deleted_at" IS NULL ORDER BY "personal_access_tokens_m"."id" LIMIT 1


2026/05/04 12:04:55 [Recovery] 2026/05/04 - 12:04:55 panic recovered:
GET /api/v1/master/location/cities HTTP/1.1                                                                                                                                         
Host: localhost:8081                                                                                                                                                                
Accept: */*                                                                                                                                                                         
Accept-Encoding: gzip, deflate, br                                                                                                                                                  
Authorization: *
Cache-Control: no-cache                                                                                                                                                             
Connection: keep-alive                                                                                                                                                              
Postman-Token: f8bdca8e-25c8-4f41-862c-754986a7f9bc                                                                                                                                 
User-Agent: PostmanRuntime/7.51.1


runtime error: invalid memory address or nil pointer dereference                                                                                                                    
C:/Program Files/Go/src/runtime/panic.go:262 (0x7ff734f4b2d7)                                                                                                                       
panicmem: panic(memoryError)                                                                                                                                                
C:/Program Files/Go/src/runtime/signal_windows.go:393 (0x7ff734f4b2a7)                                                                                                              
sigpanic: panicmem()                                                                                                                                                        
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/modules/master/controller/location/city_controller.go:28 (0x7ff735d5dd44)                                                          
(*CityController).FindAll: items, meta, err := h.cityService.GetAll(c.Request.Context(), paginateReq)                                                                       
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/modules/auth/middleware/auth_middleware.go:118 (0x7ff735d67f26)                                                                    
main.(*AuthMiddleware).Handle.func7: c.Next()                                                                                                                               
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/core/response/error_handler.go:16 (0x7ff735d68a66)                                                                                 
main.GlobalErrorHandler.func6: c.Next()                                                                                                                                     
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/core/middleware/logger_middleware.go:11 (0x7ff735d685c9)                                                                           
main.LoggerMiddleware.func5: c.Next()                                                                                                                                       
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/core/middleware/response_id_middleware.go:23 (0x7ff735d6857e)                                                                      
main.ResponseIDMiddleware.func4: c.Next()                                                                                                                                   
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/jasamedika/prj/mpd_rs_v1/backend-app/internal/core/middleware/cors_middleware.go:42 (0x7ff735d683d8)                                                                             
main.CORSMiddleware.func3: c.Next()                                                                                                                                         
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/recovery.go:90 (0x7ff735cfa130)                                                                                          
CustomRecoveryWithWriter.func1: c.Next()                                                                                                                                    
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/context.go:192 (0x7ff735ceaefe)                                                                                          
(*Context).Next: c.handlers[c.index](c)                                                                                                                                     
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/gin.go:722 (0x7ff735cf8634)                                                                                              
(*Engine).handleHTTPRequest: c.Next()                                                                                                                                       
C:/Users/narno/go/pkg/mod/github.com/gin-gonic/gin@v1.12.0/gin.go:672 (0x7ff735cf7f3b)                                                                                              
(*Engine).ServeHTTP: c.Next()                                                                                                                                               
C:/Program Files/Go/src/net/http/server.go:3340 (0x7ff7352d9e4d)                                                                                                                    
serverHandler.ServeHTTP: handler.ServeHTTP(rw, req)                                                                                                                         
C:/Program Files/Go/src/net/http/server.go:2109 (0x7ff7352ca9e4)                                                                                                                    
(*conn).serve: handler.ServeHTTP(rw, req)                                                                                                                                   
C:/Program Files/Go/src/runtime/asm_amd64.s:1693 (0x7ff734f704e0)                                                                                                                   
goexit: BYTE    $0x90   // NOP