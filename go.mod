module gee

go 1.15

require gee_web v0.0.0
require gee_cache v0.0.0

replace gee_web => ./gee-web
replace gee_cache => ./gee-cache
