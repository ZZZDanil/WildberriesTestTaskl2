module dev11/main

go 1.19

replace server => ./server

//exclude ./dev01 v0.0.0-00010101000000-000000000000

require server v0.0.0-00010101000000-000000000000
