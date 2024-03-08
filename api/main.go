package main

import (
    _ "github.com/lib/pq"
    "github.com/gin-gonic/gin"
    "url-shortener/src"
    _"fmt"
    "database/sql"
    "net/http"
)

var db *sql.DB


func print_doc(c *gin.Context){
    c.JSON(
        http.StatusOK,
        gin.H {
            "GET /stats":"return statistics of saved urls",
            "GET /url/:short_id":"redirect you to saved url",
            "POST /url/short":"provide arg 'url' with your url, return short id",
        },
    )
}

func statistics(c *gin.Context){
    count := src.Countrows(db)
    c.JSON(
        http.StatusOK, 
        gin.H {
            "Nums of short urls":count,
        },
    )
}

func create_short_url(c *gin.Context){
    rand_str := src.Gen_random_string()
    for src.CheckExist(db, rand_str) == true {
        rand_str = src.Gen_random_string()
    }

    url := c.PostForm("url")

    src.Insertdb(db, url, rand_str)
    c.JSON(
        http.StatusOK,
        gin.H {
            "short_id":rand_str,
        },
    )
}

func go_short_url(c *gin.Context){
    short_id := c.Param("short")
    is_exist := src.CheckExist(db, short_id)

    if is_exist == false{
        c.JSON(
            http.StatusNotFound,
            "no data for this id",
        )
        return
    }
    
    orig_url := src.Get_orig_url(db, short_id)
    c.Redirect(http.StatusMovedPermanently, orig_url)
}





func main(){
    
    db = src.Opendb()
    defer db.Close()


    r := gin.Default()

    r.GET("/", print_doc)
    r.GET("/stats", statistics)
    r.POST("/url/short", create_short_url)
    r.GET("/url/:short", go_short_url)



    r.Run("0.0.0.0:8080")
}
