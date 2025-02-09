module euclidean/gosandboxmain

go 1.23.5

require euclidean/trymux v0.0.0-unpublished

replace euclidean/trymux => ./mux

require euclidean/trybasicrouter v0.0.0-unpublished

replace euclidean/trybasicrouter => ./basic

require github.com/gorilla/mux v1.8.1 // indirect

