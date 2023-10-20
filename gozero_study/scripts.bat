goctl api go -api ./api/gozero_study.api -dir ./ --style=go_zero

goctl model mysql datasource -url="root:101325@tcp(127.0.0.1:3306)/gozero_study" -table="user" -dir="./model/" --style=go_zero --cache=true